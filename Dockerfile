FROM golang:latest AS build
WORKDIR /app
ADD . .
RUN go build -ldflags="-s -w" -o /app
EXPOSE 8080
ENTRYPOINT [ "/app" ]
