build:
	go build -o ./bin/jangada main.go
	cp ./bin/jangada ~/go/bin/jangada-dev

TIME := $(shell date +%S)

newapp: build
	cd example && ../bin/jangada new app$(TIME) && cd ../

sail: build
	cd example/app49/ && ../../bin/jangada sail transport ticket$(TIME)/ownership/ownership && cd ../../

.PHONY: build newapp sail
