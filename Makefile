# info
VERSION=v0.0.1

OS_ARCH=$(shell go version  | awk '{print $$4}')
GO_BIN=$(GOROOT)/bin

## help: Help for this project
help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: Compile the binary.
build:
	@go build -o bin/$(OS_ARCH)/passwd-cli cmd/client/main.go
	@go build -o bin/$(OS_ARCH)/passwd-server cmd/server/main.go


## install: build and install.
install:
	@go build -o bin/$(OS_ARCH)/passwd-cli cmd/client/main.go
	@go build -o bin/$(OS_ARCH)/passwd-server cmd/server/main.go
	@mv bin/$(OS_ARCH)/* $(GO_BIN)

## clean: Clean build files.
clean:
	rm -rf bin/darwin
	rm -rf bin/linux
	rm -rf bin/windows