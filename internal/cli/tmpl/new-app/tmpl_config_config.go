package newapp

import "strings"

var tmplConfigConfig = strings.ReplaceAll(tmplConfigConfigTemp, "'", "`")

const tmplConfigConfigTemp = `package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	configAppFile      = "config/app.yaml"
	configJobFile      = "config/job.yaml"
	configDatabaseFile = "config/database.yaml"
)

type Config struct {
	environment string       'env:"ENVIRONMENT,required"'
	app         *application 'yaml:"app"'
	database    *database    'yaml:"database"'
	job         *job         'yaml:"job"'
}

type job struct{}

type application struct {
	Name    string 'yaml:"name"'
	Address string 'yaml:"address"'
	Port    string 'yaml:"port"'
	Version string 'yaml:"version"'
	Debug   bool   'yaml:"debug"'
}

type database struct {
	Driver string 'yaml:"driver" env:"DATABASE_DRIVER,required"'
	URL    string 'yaml:"url" env:"DATABASE_URL,required"'
}

func NewConfig() (cfg *Config, err error) {
	cfg = &Config{environment: "development"}

	if cfg.app, err = load[application](configAppFile, cfg.environment); err != nil {
		return nil, err
	}

	if cfg.database, err = load[database](configDatabaseFile, cfg.environment); err != nil {
		return nil, err
	}

	if cfg.job, err = load[job](configJobFile, cfg.environment); err != nil {
		return nil, err
	}

	c, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v\n", string(c))
	return cfg, nil
}

func load[T any](path string, env string) (*T, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	type block[T any] map[string]T

	var data block[T]
	if err := yaml.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	cfg, ok := data[env]
	if !ok {
		return nil, fmt.Errorf("configuração para ambiente '%s' não encontrada no arquivo %s", env, path)
	}

	return &cfg, nil
}
`
