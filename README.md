# crud-golang-rabbitmq-mongo

## Stack

- [Golang](https://go.dev/)
- [RabbitMQ](https://www.rabbitmq.com/)
- [MongoDB](https://www.mongodb.com/)


To run and stop the applications by Makefile:

```
make start

make stop
```

## API

| Method |          Resource           |
|:------:|:---------------------------:|
|  GET   |         /users              |
|  GET   |         /users/{id}         |
|  GET   |         /photos             |
|  GET   |         /photos/{id}        |
|  GET   |         /posts              |
|  GET   |         /posts/{id}         |