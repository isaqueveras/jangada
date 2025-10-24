help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build: ## Build jangada binary
	go build -o ./bin/jangada main.go
	cp ./bin/jangada ~/go/bin/jangada-dev

deps: ### deps tidy + verify
	go mod tidy && go mod verify
.PHONY: deps

.PHONY: build help deps
