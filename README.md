# crud-golang-rabbitmq-mongo

## Stack

- [Golang](https://go.dev/)
- [RabbitMQ](https://www.rabbitmq.com/)
- [MongoDB](https://www.mongodb.com/)
- [Redis](https://redis.io/)


To run and stop the applications by Makefile:

```
make up

make down
```

## API

| Method      |          Resource           |
|:-----------:|:---------------------------:|
|  GET        |     /api/v1/albums          |
|  GET        |     /api/v1/comments        |
|  GET        |     /api/v1/photos          |
|  GET        |     /api/v1/posts           |
|  GET        |     /api/v1/todos           |
|  GET        |     /api/v1/users           |
| GET, DELETE |     /api/v1/albums/{id}     |
| GET, DELETE |     /api/v1/comments/{id}   |
| GET, DELETE |     /api/v1/photos/{id}     |
| GET, DELETE |     /api/v1/posts/{id}      |
| GET, DELETE |     /api/v1/todos/{id}      |
| GET, DELETE |     /api/v1/users/{id}      |











