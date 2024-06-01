# Print Help
help:
	mmake help
.PHONY: help

# Up all containers by docker-compose.
up:
	@docker-compose up -d --build traefik rabbitmq1 rabbitmq2 rabbitmq3 mongodb haproxy mongo-express redis crud-http
	@sleep 15 # wait for the rabbitmq cluster to become healthy
	@docker-compose up -d --build consumer-users consumer-albums consumer-posts consumer-photos consumer-comments consumer-todos
.PHONY: up

# Down all containers by docker-compose.
down:
	@docker-compose down crud-http consumer-users consumer-albums consumer-posts consumer-photos consumer-comments consumer-todos
	@docker-compose down
	@sudo chown $$USER containers/.erlang.cookie
.PHONY: down

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

# Remove images from build.
clean:
	@docker images | fgrep '<none>' | fgrep -v 'kindest/node' | awk '{ print $$3 }' | while read img; do docker rmi $$img; done
.PHONY: clean
