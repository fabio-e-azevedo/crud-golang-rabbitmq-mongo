# crud-golang-rabbitmq-mongo

![concept map](.img/concept-map.jpg)

## Command Line
```
$ ./crud --help
CRUD, command created to apply knowledge in learning the Go language

Usage:
  crud [flags]
  crud [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create resources by reading JSON file
  delete      Deleting a resource by ID
  help        Help about any command
  load        Get Resources JSON and Send API HTTP
  read        Getting by ID or Listing all resources
  version     Print the version number of CRUD

Flags:
  -h, --help   help for crud

Use "crud [command] --help" for more information about a command.
```

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











