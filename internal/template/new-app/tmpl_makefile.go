package newapp

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
