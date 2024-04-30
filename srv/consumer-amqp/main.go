package main

import (
	"fmt"
	"log"
	"time"

	"crud-golang-rabbitmq-mongo/pkg/config"
	"crud-golang-rabbitmq-mongo/pkg/database"
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	mongo "crud-golang-rabbitmq-mongo/pkg/mongodb"
	"crud-golang-rabbitmq-mongo/pkg/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := config.NewConfig()

	var conn *amqp.Connection
	var err error

	for {
		conn, err = amqp.Dial(cfg.RabbitURI)
		if err != nil {
			log.Printf("- Failed to connect to RabbitMQ: %s\n", err)
			time.Sleep(30 * time.Second)
			continue
		}
		break
	}

	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		cfg.RabbitQueueConsumer, // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var db database.Database = mongo.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: cfg.MongoCollection,
	}

	var document jph.IResource

	var forever chan struct{}
	go func() {
		for d := range msgs {
			document, _ = jph.GetResource(cfg.MongoCollection, d.Body)
			fmt.Println(db.DbInsert(document))
		}
	}()

	log.Printf(" [*] Waiting for messages from the queue \"%s\".", cfg.RabbitQueueConsumer)
	<-forever
}
