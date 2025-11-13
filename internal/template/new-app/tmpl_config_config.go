package newapp

const tmplConfigConfig = `// Package config provides configuration loading and management
package config

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/caarlos0/env/v11"
)

// NewConfig returns a new config
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Config represents the application configuration
type Config struct {
	environment string ` + "`env:\"ENVIRONMENT,required\" envDefault:\"development\"`" + `
	app         Application
	databases   []Database
}

// GetApplication returns the application configuration
func (c *Config) GetApplication() Application {
	return c.app
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	return c.environment == "production"
}

// IsDevelopment returns true if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.environment == "development"
}

// IsTesting returns true if the environment is testing
func (c *Config) IsTesting() bool {
	return c.environment == "testing"
}

// GetDatabases returns the databases configuration for the current environment
func (c *Config) GetDatabases() []Database {
	return c.databases
}

// GetPrometheusPushgateway returns the prometheus pushgateway url
func (a Application) GetPrometheusPushgateway() string {
	return a.prometheusPushgateway
}

// LoadDatabase loads the databases configuration for the current environment
func (c *Config) LoadDatabase(databases ...string) error {
	if len(databases) == 0 {
		return fmt.Errorf("no databases provided")
	}

	for _, name := range databases {
		db := &Database{}
		opt := env.Options{Prefix: fmt.Sprintf("%s_", strings.ToUpper(name))}
		if err := env.ParseWithOptions(db, opt); err != nil {
			return fmt.Errorf("error loading database configuration. %s: %w", name, err)
		}
		c.databases = append(c.databases, *db)
	}

	return nil
}

// Application represents the application settings
type Application struct {
	Name        string ` + "`env:\"APP_NAME\"`" + `
	Description string ` + "`env:\"APP_DESCRIPTION\"`" + `
	Address     string ` + "`env:\"APP_ADDRESS\"`" + `
	Version     string ` + "`env:\"APP_VERSION\"`" + `
	Debug       bool   ` + "`env:\"APP_DEBUG\"`" + `

	prometheusPushgateway string ` + "`env:\"PROMETHEUS_PUSHGATEWAY\"`" + `
}

// Database represents the database connection settings
type Database struct {
	Nick               string      ` + "`env:\"DATABASE_NICK\"`" + `
	Name               string      ` + "`env:\"DATABASE_NAME\"`" + `
	Username           string      ` + "`env:\"DATABASE_USER\"`" + `
	Password           string      ` + "`env:\"DATABASE_PASS\"`" + `
	Host               string      ` + "`env:\"DATABASE_HOST\"`" + `
	Port               string      ` + "`env:\"DATABASE_PORT\"`" + `
	MaxConn            int         ` + "`env:\"DATABASE_MAX_CONN\"`" + `
	MaxIdle            int         ` + "`env:\"DATABASE_MAX_IDLE\"`" + `
	ReadOnly           bool        ` + "`env:\"DATABASE_READ_ONLY\"`" + `
	Main               bool        ` + "`env:\"DATABASE_MAIN\"`" + `
	TransactionTimeout int         ` + "`env:\"DATABASE_TIMEOUT\"`" + `
	SSLMode            string      ` + "`env:\"DATABASE_SSL_MODE\"`" + `
	SSLClient          Certificate
}

// String returns the database connection string
func (d Database) String() string {
	baseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", url.QueryEscape(d.Username), url.QueryEscape(d.Password), d.Host, d.Port, d.Name)

	params := url.Values{}
	if d.SSLMode != "" {
		params.Set("sslmode", d.SSLMode)
	}

	if d.SSLMode != "disable" {
		if d.SSLClient.CertificateAuthority != "" {
			params.Set("sslrootcert", d.SSLClient.CertificateAuthority)
		}
		if d.SSLClient.Certificate != "" {
			params.Set("sslcert", d.SSLClient.Certificate)
		}
		if d.SSLClient.PrivateKey != "" {
			params.Set("sslkey", d.SSLClient.PrivateKey)
		}
	}

	if len(params) > 0 {
		baseURL += "?" + params.Encode()
	}

	return baseURL
}

// Certificate represents the SSL certificate settings
type Certificate struct {
	Certificate          string ` + "`env:\"DATABASE_SSL_CERT\"`" + `
	PrivateKey           string ` + "`env:\"DATABASE_SSL_KEY\"`" + `
	CertificateAuthority string ` + "`env:\"DATABASE_SSL_CA\"`" + `
}
`
