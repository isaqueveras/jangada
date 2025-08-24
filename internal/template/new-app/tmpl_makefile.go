package newapp

const tmplMakefile = `# Run templ generation in watch mode
templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false

# Run air for Go hot reload
server:
	air \
	--build.cmd "go build -o bin/jangada ./cmd/app/main.go" \
	--build.bin "bin/jangada" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# Watch Tailwind CSS changes
tailwind:
	tailwindcss -i ./web/assets/css/input.css -o ./web/assets/css/output.css --watch

# Start development server with all watchers
dev:
	make -j3 tailwind templ server

mocks:
	@go tool mockery --output internal/blog/tests/mocks/services --dir internal/blog/application/interfaces --all

clean:
	@echo "> Cleaning project..."
	go clean
	rm -rf ./bin

test:
	@echo "> Running tests..."
	go test ./internal/blog/tests/... -coverprofile=coverage.out -covermode=count
	go tool cover -func=coverage.out

build: clean test
	@echo "> Building project..."
	go run github.com/a-h/templ/cmd/templ@latest generate
	go build -o ./bin/jangada-server ./cmd/server/main.go
	go build -o ./bin/jangada-worker ./cmd/worker/main.go

run: build
	@echo "> Running project on http://localhost:8080..."
	./bin/jangada-server

lint:
	@echo "> Linting project..."
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

.PHONY: dev clean test lint
`
