// Package template contains templates for infrastructure
package template

// InfrastructureRepository defines the repository template for infrastructure
const InfrastructureRepository = `// Package {{ ToLower .Entity }} contains the {{ ToLower .Entity }} repository
package {{ ToLower .Entity }}

import (
	"{{ .Module }}/internal/infrastructure/{{ ToLower .Folder }}/{{ ToLower .Entity }}/postgres"
	"{{ .Module }}/pkg/database"
)

type repository struct {
	pg *postgres.Postgres
}

// NewRepository creates a new repository for customer entity
func NewRepository(tx database.Transaction) *repository {
	return &repository{pg: postgres.NewPostgres(tx)}
}
`

// InfrastructurePostgresData defines the postgres data template for infrastructure
const InfrastructurePostgresData = `// Package postgres contains the implementation of the {{ ToLower .Entity }} repository
package postgres

import (
	"context"

	"{{ .Module }}/pkg/database"

	"github.com/Masterminds/squirrel"
)

// Postgres is the implementation of repository
type Postgres struct {
	db      database.Transaction
	builder squirrel.StatementBuilderType
}

// NewPostgres returns a new instance of postgres
func NewPostgres(db database.Transaction) *Postgres {
	return &Postgres{db: db, builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}
`

// InfrastructurePostgresModel defines the postgres model template for infrastructure
const InfrastructurePostgresModel = `// Package postgres contains the implementation of the {{ ToLower .Entity }} repository
package postgres
`
