all: test build
	@go run main.go

build:
	@go build -ldflags="-s -w" -o bin/`basename $(PWD)`

test:
	@go test -v ./...
