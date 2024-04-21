package main

import (
	"log"
	"os"

	"crud-golang-rabbitmq-mongo/internal"
	mongo "crud-golang-rabbitmq-mongo/mongodb"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	resourceType := "users"

	uri := os.Getenv("RABBITMQ_URI")
	if uri == "" {
		log.Fatal("You must set your 'RABBITMQ_URI' environment variable.")
	}

	conn, err := amqp.Dial(uri)
	internal.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	internal.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		resourceType, // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	internal.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	internal.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			mongo.Insert(d.Body, resourceType)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
