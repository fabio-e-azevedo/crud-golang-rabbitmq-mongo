package main

import (
	"encoding/json"

	"crud-golang-rabbitmq-mongo/internal"

	rt "crud-golang-rabbitmq-mongo/rabbitmq"
	data "crud-golang-rabbitmq-mongo/users"
)

func main() {
	users, err := data.GetUsers()
	internal.FailOnError(err, "Failed to get users")

	for i := range *users {
		body, err := json.Marshal((*users)[i])
		internal.FailOnError(err, "Failed to marshal message")
		rt.Publisher(body)
	}
}
