package main

import (
	"encoding/json"
	"fmt"
	"log"

	rt "get-http-users/rabbitmq"
	data "get-http-users/users"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	users, err := data.GetUsers()
	failOnError(err, "Failed to get users")

	for i := range *users {
		body, err := json.Marshal((*users)[i])
		failOnError(err, "Failed to marshal message")
		rt.Publisher(body)
	}
}
