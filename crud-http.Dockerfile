FROM golang:1.22.3-alpine AS builder
WORKDIR /crud-golang-rabbitmq-mongo
COPY ./go.mod ./go.sum ./srv/crud-http/main.go ./
COPY ./pkg/ ./pkg
RUN go build -o http main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /crud-golang-rabbitmq-mongo/http .

CMD ["/app/http"]
