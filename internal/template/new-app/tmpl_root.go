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

const tmplMakefile = `templ:
	templ generate

dev: templ
	@echo "> Running project on http://{{ .DefaultHost }}..."
	@go run cmd/app/main.go 

clean:
	@echo "> Cleaning project..."
	go clean
	rm -rf ./bin

test:
	@echo "> Running tests..."
	go test ./... -coverprofile=coverage.out -covermode=count
	go tool cover -func=coverage.out

build: clean test
	@echo "> Building project..."
	@go run github.com/a-h/templ/cmd/templ@latest generate
	go build -o ./bin/jangada-app ./cmd/app/main.go

run-build: build
	@echo "> Running project on http://{{ .DefaultHost }}..."
	./bin/jangada-server

lint:
	@echo "> Linting project..."
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

.PHONY: dev clean test lint build run-build
`

const tmplGitIgnore string = `db/*.db
db/*.db-*

bin/
tmp/

*_templ.go
*_templ.txt`
