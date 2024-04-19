# Print Help
help:
	mmake help
.PHONY: help

# Run applications by docker-compose.
start:
	docker-compose up -d
.PHONY: start

# Stop applications by docker-compose.
stop:
	docker-compose down
.PHONY: stop
