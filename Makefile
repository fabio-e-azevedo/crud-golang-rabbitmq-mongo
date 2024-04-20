# Print Help
help:
	mmake help
.PHONY: help

# Run all containers by docker-compose.
start:
	docker-compose up -d
.PHONY: start

# Stop all containers by docker-compose.
stop:
	docker-compose down
.PHONY: stop

# Run container mongodb by docker-compose.
mongodb:
	docker-compose up -d mongo
	docker-compose up -d mongo mongo-express
.PHONY: mongodb

# Run container rabbitmq by docker-compose.
rabbitmq:
	docker-compose up -d rabbitmq
.PHONY: rabbitmq
