# Print Help
help:
	mmake help
.PHONY: help

# Up all containers by docker-compose.
start:
	docker-compose up -d --build
.PHONY: start

# Down all containers by docker-compose.
stop:
	docker-compose down
.PHONY: stop

# Run container mongodb by docker-compose.
mongodb:
	docker-compose up -d mongodb
	docker-compose up -d mongo-express
.PHONY: mongodb

# Run container rabbitmq by docker-compose.
rabbitmq:
	docker-compose up -d rabbitmq
.PHONY: rabbitmq

# Go build command line "crud".
build:
	go build -o ./crud cmd/main.go
	./crud -h
.PHONY: build