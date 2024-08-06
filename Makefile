all: test build run

APP=`basename $(PWD)`
IMAGE=$(APP):latest

build:
	@go build -ldflags="-s -w" -pgo=default.pgo -o bin/app

build-docker:
	@docker build -t $(IMAGE) .

profile:
	@curl -o default.pgo http://localhost:8080/debug/pprof/profile?seconds=30

register:
	@curl -X POST http://localhost:8080/api/v1/account \
	-H "Content-Type: application/x-www-form-urlencoded" \
	-d "email=foo&password=bar"

run:
	@go run main.go

run-docker:
	@docker run -it -p 8080:8080 $(IMAGE)

test:
	@go test -v ./...
