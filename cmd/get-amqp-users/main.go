package main

import (
	"fmt"
	"log"

	"crud-golang-rabbitmq-mongo/config"
	"crud-golang-rabbitmq-mongo/database"
	"crud-golang-rabbitmq-mongo/internal"
	mongo "crud-golang-rabbitmq-mongo/mongodb"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	resourceType := "photos"
	cfg := config.NewConfig()

	conn, err := amqp.Dial(cfg.RabbitURI)
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

	var db database.Database = mongo.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			fmt.Println(db.DbInsert(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages from the queue \"%s\".", resourceType)
	<-forever
}
