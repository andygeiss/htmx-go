all: test build
	@go run main.go

FLAGS=-ldflags="-s -w"
NAME=`basename $(PWD)`

build:
	@go build -o bin/${NAME} ${FLAGS}
	@upx bin/${NAME}

test:
	@go test -v ./...
