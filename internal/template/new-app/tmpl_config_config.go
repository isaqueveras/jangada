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

const (
	configAppFile      = "config/app.yaml"
	configDatabaseFile = "config/database.yaml"
)

// Config represents the application configuration
type Config struct {
	Environment string      ` + "`env:\"ENVIRONMENT,required\" envDefault:\"development\"`" + `
	App         Application ` + "`yaml:\"app\"`" + `
	Databases   []Database  ` + "`yaml:\"databases\"`" + `
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

// IsDevelopment returns true if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsTesting returns true if the environment is testing
func (c *Config) IsTesting() bool {
	return c.Environment == "testing"
}

// GetDatabases returns the databases configuration for the current environment
func (c *Config) GetDatabases() []Database {
	return c.Databases
}

// Application represents the application settings
type Application struct {
	Name        string ` + "`yaml:\"name\"`" + `
	Description string ` + "`yaml:\"description\"`" + `
	Address     string ` + "`yaml:\"address\" env:\"APP_ADDRESS\"`" + `
	Version     string ` + "`yaml:\"version\" env:\"APP_VERSION\"`" + `
	Debug       bool   ` + "`yaml:\"debug\" env:\"APP_DEBUG\"`" + `
}

// Database represents the database connection settings
type Database struct {
	Nick               string      ` + "`yaml:\"nick\" env:\"DB_NICK\"`" + `
	Name               string      ` + "`yaml:\"name\" env:\"DB_NAME\"`" + `
	Username           string      ` + "`yaml:\"username\" env:\"DB_USER\"`" + `
	Password           string      ` + "`yaml:\"password\" env:\"DB_PASS\"`" + `
	Host               string      ` + "`yaml:\"hostname\" env:\"DB_HOST\"`" + `
	Port               string      ` + "`yaml:\"port\" env:\"DB_PORT\"`" + `
	MaxConn            int         ` + "`yaml:\"max_conn\" env:\"DB_MAX_CONN\"`" + `
	MaxIdle            int         ` + "`yaml:\"max_idle\" env:\"DB_MAX_IDLE\"`" + `
	ReadOnly           bool        ` + "`yaml:\"read_only\" env:\"DB_READ_ONLY\"`" + `
	Main               bool        ` + "`yaml:\"main\" env:\"DB_MAIN\"`" + `
	TransactionTimeout int         ` + "`yaml:\"transaction_timeout\" env:\"DB_TIMEOUT\"`" + `
	SSLMode            string      ` + "`yaml:\"ssl_mode\" env:\"DB_SSL_MODE\"`" + `
	SSLClient          Certificate ` + "`yaml:\"ssl_client\"`" + `
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
	Certificate          string ` + "`yaml:\"path_cert\" env:\"DB_SSL_CERT\"`" + `
	PrivateKey           string ` + "`yaml:\"path_key\" env:\"DB_SSL_KEY\"`" + `
	CertificateAuthority string ` + "`yaml:\"path_ca\" env:\"DB_SSL_CA\"`" + `
}

// NewConfig returns a new config
func NewConfig() (cfg *Config, err error) {
	cfg = &Config{Environment: "development"}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("error loading env: %w", err)
	}

	if cfg.App, err = load[Application](configAppFile, cfg.Environment); err != nil {
		return nil, err
	}

	if cfg.Databases, err = load[[]Database](configDatabaseFile, cfg.Environment); err != nil {
		return nil, err
	}

	for i := range cfg.Databases {
		if err := env.Parse(&cfg.Databases[i]); err != nil {
			return nil, fmt.Errorf("error loading env for db %d: %w", i, err)
		}
	}

	return cfg, nil
}

// load loads the configuration from a file
func load[T any](path string, env string) (T, error) {
	type block[T any] map[string]T

	var (
		cfg  T
		data block[T]
	)

	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	if err := yaml.Unmarshal(b, &data); err != nil {
		return cfg, err
	}

	if cfg, ok := data[env]; ok {
		return cfg, nil
	}

	return cfg, fmt.Errorf("configuração para ambiente %s não encontrada no arquivo %s", env, path)
}
`
