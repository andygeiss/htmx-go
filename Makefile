all: test build run

build:
	@go build -ldflags="-s -w" -pgo=default.pgo -o bin/`basename $(PWD)`

profile:
	@curl -o default.pgo http://localhost:8080/debug/pprof/profile?seconds=30

run:
	@go run main.go

test:
	@go test -v ./...
