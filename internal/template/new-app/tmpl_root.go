package newapp

import "strings"

const tmplGoMod string = `module {{ .ModuleName }}

go 1.24.6`

var tmplReadme = strings.ReplaceAll(tmplReadmeTemp, "'", "`")

const tmplReadmeTemp string = `# Jangada Framework

A full-stack web framework in Go for building modern web applications, RESTful APIs, and gRPC with integrated frontend and backend.

## Setup

'''bash
# Install Jangada
go install github.com/isaqueveras/jangada@latest
'''

## Run Jangada

'''bash
# Run development server
jangada serve
'''

## Run Jangada with Docker

'''bash
# Run development server
docker compose up
'''

## Tests and Coverage

'''bash
# Run tests
jangada test
'''

## Documentation

[Docs](https://jangada-framework.com/docs)
`

const tmplMakefile = `dev:
	templ generate -watch -cmd "go run cmd/app/main.go"

air:
	air -c .air.toml

clean:
	@echo "> Cleaning project..."
	@go clean
	@rm -rf ./bin

test:
	@echo "> Running tests..."
	@go test -coverprofile=coverage.out ./...
	@sed -i '/_templ\.go/d' coverage.out
	@go tool cover -html=coverage.out -o coverage.html

build: clean test
	@echo "> Building project..."
	@go run github.com/a-h/templ/cmd/templ@latest generate
	go build -o ./bin/app ./cmd/app/main.go

run-build: build
	./bin/app

lint:
	@echo "> Linting project..."
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

.PHONY: dev clean test lint build run-build

# --------------- Docker Tools ---------------

DOCKER_COMMAND :=	docker compose -p {{ ToLower .AppName }} -f docker-compose.yml

docker-dev: 
	@echo "> 🐳 Run development server"
	@$(DOCKER_COMMAND) up {{ ToLower .AppName }}_postgres {{ ToLower .AppName }}_app

docker-app:
	@echo "> 🐳 Start app service"
	@$(DOCKER_COMMAND) up {{ ToLower .AppName }}_app

docker-database:
	@echo "> 🐳 Start database service"
	@$(DOCKER_COMMAND) up {{ ToLower .AppName }}_postgres

# Docker commands

docker-up:
	@echo "> 🐳 Builds, (re)creates, starts, and attaches to containers for a service..."
	@$(DOCKER_COMMAND) up

docker-down:
	@echo "> 🐳 Stops containers and removes containers, networks, volumes, and images..."
	@$(DOCKER_COMMAND) down

docker-build:
	@echo "> 🐳 Build or rebuild services..."
	@$(DOCKER_COMMAND) build

docker-start:
	@echo "> 🐳 Start services..."
	@$(DOCKER_COMMAND) start

docker-stop:
	@echo "> 🐳 Stop services..."
	@$(DOCKER_COMMAND) stop

docker-restart:
	@echo "> 🐳 Restart service containers..."
	@$(DOCKER_COMMAND) restart

docker-logs:
	@echo "> 🐳 View output from containers..."
	@$(DOCKER_COMMAND) logs

docker-clean:
	@echo "> 🐳 Remove orphaned containers, volumes, and images...."
	@$(DOCKER_COMMAND) down -v
	@docker system prune -f
`

const tmplGitIgnore string = `bin/
tmp/

*_templ.go
.env`

const tmplDockerIgnore string = `.air.toml
.coverage.*
Makefile
README.md
bin/
tmp/
.env
.env.example
.gitignore
database.sql`

const tmplEnv string = `# ================== APPICATION ==================

# ENVIRONMENT -> define the environment used during the application
ENVIRONMENT=development # development|testing|production

# APP_NAME -> defines the system name
APP_NAME={{ ToLower .AppName }}

# APP_DESCRIPTION -> defines the system description
APP_DESCRIPTION="Application built with Jangada Framework"

# APP_ADDRESS -> defines the address where the application will be executed
APP_ADDRESS={{ .DefaultHost }}:{{ .DefaultPort }}

# APP_VERSION -> defines the version of the application that is running
APP_VERSION=v0.0.0

# APP_DEBUG -> defines whether the application will run in debug mode
APP_DEBUG=true

# PROMETHEUS_PUSHGATEWAY -> defines the application address for Prometheus Pushgateway
PROMETHEUS_PUSHGATEWAY=http://localhost:9091

# ================== DATABASE ==================

# {{ ToUpper .AppName }}_DATABASE_NICK -> defines the alias for the database connection
{{ ToUpper .AppName }}_DATABASE_NICK={{ .AppName }}

# {{ ToUpper .AppName }}_DATABASE_NAME -> defines the database name
{{ ToUpper .AppName }}_DATABASE_NAME={{ ToLower .AppName }}

# {{ ToUpper .AppName }}_DATABASE_USER -> defines the database username
{{ ToUpper .AppName }}_DATABASE_USER=postgres

# {{ ToUpper .AppName }}_DATABASE_PASS -> define the database user's password
{{ ToUpper .AppName }}_DATABASE_PASS=postgres

# {{ ToUpper .AppName }}_DATABASE_HOST -> defines the database address
{{ ToUpper .AppName }}_DATABASE_HOST=localhost

# {{ ToUpper .AppName }}_DATABASE_PORT -> defines the database port
{{ ToUpper .AppName }}_DATABASE_PORT=5432

# {{ ToUpper .AppName }}_DATABASE_MAX_CONN -> defines the maximum number of connections the database will have access to
{{ ToUpper .AppName }}_DATABASE_MAX_CONN=20

# {{ ToUpper .AppName }}_DATABASE_MAX_IDLE -> defines the maximum time a connection can wait
{{ ToUpper .AppName }}_DATABASE_MAX_IDLE=5

# {{ ToUpper .AppName }}_DATABASE_READ_ONLY -> defines whether the database is read-only or write-only
{{ ToUpper .AppName }}_DATABASE_READ_ONLY=false

# {{ ToUpper .AppName }}_DATABASE_MAIN -> defines whether the database is the primary one or not
{{ ToUpper .AppName }}_DATABASE_MAIN=true

# {{ ToUpper .AppName }}_DATABASE_TIMEOUT -> defines the timeout that a connection can have in the database
{{ ToUpper .AppName }}_DATABASE_TIMEOUT=30

# {{ ToUpper .AppName }}_DATABASE_SSL_MODE -> defines which security mode the database will be used
{{ ToUpper .AppName }}_DATABASE_SSL_MODE=disable

# {{ ToUpper .AppName }}_DATABASE_SSL_CERT -> defines the address for SSL certification
{{ ToUpper .AppName }}_DATABASE_SSL_CERT=/etc/ssl/certs/client.crt

# {{ ToUpper .AppName }}_DATABASE_SSL_KEY -> defines the address for SSL certification
{{ ToUpper .AppName }}_DATABASE_SSL_KEY=/etc/ssl/private/client.key

# {{ ToUpper .AppName }}_DATABASE_SSL_CA -> defines the address for SSL certification
{{ ToUpper .AppName }}_DATABASE_SSL_CA=/etc/ssl/certs/ca.crt
`

const tmplAirToml = `root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./bin/{{ ToLower .AppName }}-app ./cmd/app"
  bin = "tmp/main"
  full_bin = "./bin/{{ ToLower .AppName }}-app"
  include_ext = ["go", "tmpl"]
  exclude_dir = ["vendor", "tmp"]

[log]
  level = "info"
  color = true`
