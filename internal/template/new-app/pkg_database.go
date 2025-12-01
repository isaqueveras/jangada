package newapp

const serviceDatabaseInterfaceTemp = `// Package database provides interfaces and implementations for database connections and transactions.
package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"{{ .ModuleName }}/config"
)

// Database is the interface for database connections
type Database interface {
	/**
	 * Open opens the database connection
	 *
	 * @ctx represents the context for the operation
	 * @db represents the database configuration
	 *
	 * @err returns an error
	 */
	Open(ctx context.Context, db *config.Database) (err error)

	/**
	 * NewTx returns a new transaction from the database connection
	 *
	 * @ctx represents the context for the operation
	 * @readonly indicates if the transaction should be read-only
	 *
	 * @tx returns a new transaction interface
	 * @err returns an error
	 */
	NewTx(ctx context.Context, readonly bool) (tx pgx.Tx, err error)

	/**
	 * Close closes the database connection
	 */
	Close()
}

// ConnectionPool is the interface for database connection pools
type ConnectionPool interface {
	/**
	 * NewTransaction returns a new transaction from the database connection
	 *
	 * @ctx represents the context for the operation
	 * @readonly indicates if the transaction should be read-only
	 * @database is the name of the database to use

	 * @tx returns a new transaction interface
	 * @err returns an error
	 */
	NewTransaction(ctx context.Context, readonly bool, database ...string) (tx Transaction, err error)

	/**
	 * CloseConnections closes all database connections in the pool.
	 */
	CloseConnections()
}

// Transaction is the interface for database transactions
type Transaction interface {
	/**
	 * Exec executes a query without returning any rows.
	 *
	 * @ctx represents the context for the operation
	 * @sql is the query to execute
	 * @args is the arguments to the query
	 *
	 * @cmd returns a pgconn.CommandTag object
	 * @err returns an error
	 */
	Exec(ctx context.Context, sql string, args ...any) (cmd pgconn.CommandTag, err error)

	/**
	 * QueryRow execute a query that is expected to return at most one row.
	 *
	 * @ctx represents the context for the operation
	 * @sql is the query to execute
	 * @args is the arguments to the query
	 *
	 * @row returns a pgx.Row object
	 */
	QueryRow(ctx context.Context, sql string, args ...any) (row pgx.Row)

	/**
	 * Query executes a query that returns rows, typically a SELECT.
	 *
	 * @ctx represents the context for the operation
	 * @sql is the query to execute
	 * @args is the arguments to the query

	 * @row returns a pgx.Rows object
	 * @err returns an error
	 */
	Query(ctx context.Context, sql string, args ...any) (row pgx.Rows, err error)

	/**
	 * Prepare prepares the given query for later execution.
	 *
	 * if the name is empty, the anonymous prepared statement will be used. This
	 * allows Prepare to also to describe statements without creating a server-side prepared statement.
	 *
	 * @ctx represents the context for the operation
	 * @name is the name of the prepared statement
	 * @sql is the query to prepare
	 *
	 * @statement returns a pgconn.StatementDescription object
	 * @err returns an error
	 */
	Prepare(ctx context.Context, name string, sql string) (statement *pgconn.StatementDescription, err error)

	/**
	 * CopyFrom copys data from the reader to the database.
	 *
	 * requires all values use the binary format. A pgtype.Type that supports the binary format must be registered
	 * for the type of each column. Almost all types implemented by pgx support the binary format.
	 *
	 * @ctx represents the context for the operation
	 * @tableName is the name of the table to copy data to
	 * @columnNames is the names of the columns to copy data to
	 * @rowSrc is the source of the data to copy
	 *
	 * @rowsAffected returns the number of rows affected
	 * @err returns an error
	 */
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (rowsAffected int64, err error)

	/**
	 * SendBatch send a batch of queries to the database.
	 *
	 * @ctx represents the context for the operation
	 * @batch is the batch of queries
	 *
	 * @results returns a pgx.BatchResults object
	 */
	SendBatch(ctx context.Context, batch *pgx.Batch) (results pgx.BatchResults)

	/**
	 * Commit commits the transaction to the database
	 *
	 * @err returns an error
	 */
	Commit() (err error)

	/**
	 * Rollback rollbacks the transaction from the database.
	 *
	 * rolls back the transaction if this is a real transaction or rolls back to the savepoint
	 * if this is a pseudo nested transaction.
	 */
	Rollback()
}
`

const serviceDatabasePoolTemp = `// Package database provides database connection pooling and transaction management.
package database

import (
	"context"
	"errors"

	"{{ .ModuleName }}/core"
	"{{ .ModuleName }}/pkg/database/postgres"
)

const roSuffix = "_ro"
const defaultDatabase = "{{ ToLower .AppName }}"

type poolDatabase struct {
	pool map[string]Database
}

// NewConnectionPool creates a new connection pool with the given databases
func NewConnectionPool(core *core.Core) ConnectionPool {
	databases := core.Config().GetDatabases()
	if len(databases) == 1 {
		databaseRO := databases[0]
		databaseRO.ReadOnly = true
		databases = append(databases, databaseRO)
	}

	pool := make(map[string]Database)
	for _, db := range databases {
		name := db.Nick
		if db.ReadOnly {
			name += roSuffix
		}

		pool[name] = postgres.NewPostgres()
		if err := pool[name].Open(context.Background(), &db); err != nil {
			core.Log().Error("error opening database", "name", name, "error", err)
			continue
		}
	}

	return &poolDatabase{pool: pool}
}

// CloseConnections closes all connections in the pool
func (p *poolDatabase) CloseConnections() {
	for _, db := range p.pool {
		db.Close()
	}
}

// NewTransaction starts a new transaction in the database
func (p *poolDatabase) NewTransaction(ctx context.Context, readonly bool, database ...string) (Transaction, error) {
	databaseName := defaultDatabase
	if len(database) > 0 {
		databaseName = database[0]
	}

	if readonly {
		databaseName += roSuffix
	}

	conn, ok := p.pool[databaseName]
	if !ok {
		return nil, errors.New("database not found: " + databaseName)
	}

	tx, err := conn.NewTx(ctx, readonly)
	if err != nil {
		return nil, err
	}

	return &transactionDB{tx: tx}, nil
}
`

const serviceDatabaseTransactionTemp = `// Package database provides database connection pooling and transaction management.
package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type transactionDB struct{ 
	tx pgx.Tx
}

// Commit commits the transaction to the database
func (t *transactionDB) Commit() error {
	if t.tx == nil {
		return errors.New("transaction is nil")
	}
	return t.tx.Commit(context.Background())
}

// Rollback rollbacks the transaction from the database
func (t *transactionDB) Rollback() {
	_ = t.tx.Rollback(context.Background())
}

// Exec executes a query without returning any rows.
func (t *transactionDB) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return t.tx.Exec(ctx, sql, args...)
}

// QueryRow execute a query that is expected to return at most one row.
func (t *transactionDB) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return t.tx.QueryRow(ctx, query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (t *transactionDB) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return t.tx.Query(ctx, query, args...)
}

// Prepare prepares the given query for later execution.
func (t *transactionDB) Prepare(ctx context.Context, name string, sql string) (*pgconn.StatementDescription, error) {
	return t.tx.Prepare(ctx, name, sql)
}

// CopyFrom copys data from the reader to the database.
func (t *transactionDB) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return t.tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

// SendBatch send a batch of queries to the database.
func (t *transactionDB) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return t.tx.SendBatch(ctx, b)
}
`

const serviceDatabasePostgresTemp = `// Package postgres provides a connection to a postgres database
package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"{{ .ModuleName }}/config"
)

// postgres is a connection to a postgres database
type postgres struct {
	pool    *pgxpool.Pool
	timeout int
}

// NewPostgres returns a new connection to a postgres database
func NewPostgres() *postgres {
	return &postgres{}
}

// Open opens the connection to the database
func (p *postgres) Open(ctx context.Context, cfg *config.Database) error {
	config, err := pgxpool.ParseConfig(cfg.String())
	if err != nil {
		return err
	}

	config.MaxConns = int32(cfg.MaxConn)
	config.MaxConnLifetime = time.Minute * 1
	config.ConnConfig.RuntimeParams = map[string]string{
		"application_name": "aomercado",
		"DateStyle":        "ISO",
		"IntervalStyle":    "iso_8601",
		"search_path":      "public",
		"TimeZone":         "America/Fortaleza",
	}

	p.timeout = cfg.TransactionTimeout
	if p.pool, err = pgxpool.NewWithConfig(ctx, config); err != nil {
		return err
	}

	return p.pool.Ping(ctx)
}

// Close closes the connection to the database
func (p *postgres) Close() {
	if p.pool != nil {
		p.pool.Close()
	}
}

// NewTx starts a new transaction in the database
func (p *postgres) NewTx(ctx context.Context, readonly bool) (pgx.Tx, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(p.timeout+1)*time.Second)
	defer cancel()

	opts := pgx.TxOptions{AccessMode: pgx.ReadWrite}
	if readonly {
		opts.AccessMode = pgx.ReadOnly
	}

	return p.pool.BeginTx(ctx, opts)
}
`

const serviceDatabaseSeedsTemp = `// Package seeds provides seeds for the database
package seeds

type Seeds struct{}

func NewSeeds() *Seeds {
	return &Seeds{}
}

func (s *Seeds) Seed() error {
	return nil
}
`
