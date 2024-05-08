# crud-golang-rabbitmq-mongo

![concept map](.img/concept-map.jpg)

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
|  GET, POST  |     /api/v1/albums          |
|  GET, POST  |     /api/v1/comments        |
|  GET, POST  |     /api/v1/photos          |
|  GET, POST  |     /api/v1/posts           |
|  GET, POST  |     /api/v1/todos           |
|  GET, POST  |     /api/v1/users           |
| GET, DELETE |     /api/v1/albums/{id}     |
| GET, DELETE |     /api/v1/comments/{id}   |
| GET, DELETE |     /api/v1/photos/{id}     |
| GET, DELETE |     /api/v1/posts/{id}      |
| GET, DELETE |     /api/v1/todos/{id}      |
| GET, DELETE |     /api/v1/users/{id}      |











