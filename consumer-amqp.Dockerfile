FROM golang:1.22.2-alpine AS builder
WORKDIR /crud-golang-rabbitmq-mongo
COPY ./go.mod ./go.sum ./srv/consumer-amqp/main.go .
COPY ./pkg/ ./pkg
RUN go build -o consumer main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /crud-golang-rabbitmq-mongo/consumer .

CMD ["/app/consumer"]
