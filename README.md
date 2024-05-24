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
  version     Show version of crud

Flags:
  -h, --help   help for crud

Use "crud [command] --help" for more information about a command.
```

## Stack

- [Golang](https://go.dev/)
- [RabbitMQ](https://www.rabbitmq.com/)
- [MongoDB](https://www.mongodb.com/)
- [Redis](https://redis.io/)
- [Traefik](https://traefik.io/traefik/)


To run and stop the applications by Makefile:

```
make up

make down
```

## API REST

| Method |  Resource  |
|:------:|:----------:|
|  GET   |  /swagger  |
