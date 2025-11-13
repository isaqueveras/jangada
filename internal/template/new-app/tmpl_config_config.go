package newapp

const tmplConfigConfig = `// Package config provides configuration loading and management
package config

import (
	"fmt"
	"net/url"
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
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
func (c *Config) GetDatabases() []database {
	return c.databases
}

// GetPrometheusPushgateway returns the prometheus pushgateway url
func (a Application) GetPrometheusPushgateway() string {
	return a.prometheusPushgateway
}

// LoadDatabase loads the databases configuration for the current environment
func (c *Config) LoadDatabase(databases ...string) error {
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
	Nick               string      ` + "env:\"DB_NICK\"`" + `
	Name               string      ` + "env:\"DB_NAME\"`" + `
	Username           string      ` + "env:\"DB_USER\"`" + `
	Password           string      ` + "env:\"DB_PASS\"`" + `
	Host               string      ` + "env:\"DB_HOST\"`" + `
	Port               string      ` + "env:\"DB_PORT\"`" + `
	MaxConn            int         ` + "env:\"DB_MAX_CONN\"`" + `
	MaxIdle            int         ` + "env:\"DB_MAX_IDLE\"`" + `
	ReadOnly           bool        ` + "env:\"DB_READ_ONLY\"`" + `
	Main               bool        ` + "env:\"DB_MAIN\"`" + `
	TransactionTimeout int         ` + "env:\"DB_TIMEOUT\"`" + `
	SSLMode            string      ` + "env:\"DB_SSL_MODE\"`" + `
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
	Certificate          string ` + "`env:\"DB_SSL_CERT\"`" + `
	PrivateKey           string ` + "`env:\"DB_SSL_KEY\"`" + `
	CertificateAuthority string ` + "`env:\"DB_SSL_CA\"`" + `
}
`
