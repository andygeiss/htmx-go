FROM golang:latest AS build
WORKDIR /go-build/
ADD . .
RUN go build -ldflags="-s -w" -o server main.go

FROM scratch
COPY --from=build /go-build/server /server
ENTRYPOINT [ "/server" ]
