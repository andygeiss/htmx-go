FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app/
COPY . .
RUN go test -v ./...
RUN CGO_ENABLED=0 go build -o /app/bin/app

FROM scratch
COPY --from=builder /app/data /data
COPY --from=builder /app/bin/app /app
ENTRYPOINT [ "/app" ]
