FROM golang:1.22.3-alpine AS builder
WORKDIR /crud-golang-rabbitmq-mongo
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY ./go.mod ./go.sum ./
COPY crud-http ./crud-http
COPY ./pkg/ ./pkg
RUN swag init -d crud-http --parseDependency --parseInternal --parseDepth 2 -o pkg/docs
RUN go build -o ./goapp ./crud-http/main.go

FROM alpine:latest
WORKDIR /api
COPY --from=builder /crud-golang-rabbitmq-mongo/goapp .

CMD ["/api/goapp"]
