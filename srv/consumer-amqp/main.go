package main

import (
	"log"
	"time"

	"crud-golang-rabbitmq-mongo/pkg/config"
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
			log.Printf("- failed to connect to rabbitmq: %s\n", err)
			time.Sleep(20 * time.Second)
			continue
		}
		break
	}

	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		cfg.RabbitQueueConsumer, // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	utils.FailOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "failed to register a consumer")

	var db = mongo.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: cfg.MongoCollection,
	}

	var document jph.IResource
	var resultInsert string
	var acknowledge bool

	var forever chan struct{}
	go func() {
		for d := range msgs {
			acknowledge = true
			resourceType := cfg.MongoCollection

			document, _ = jph.GetResource(resourceType, d.Body)

			resultInsert, err = db.DbInsert(document)
			if err != nil {
				log.Printf("- error insert document in mongodb: %s\n", err)
				acknowledge = false
			}

			err = d.Ack(acknowledge)
			if err != nil {
				log.Printf("- error to confirm message: %s\n", err)
				continue
			}

			log.Printf("- success in inserting document with _id: %s\n", resultInsert)
		}
	}()

	log.Printf(" [*] Waiting for messages from the queue \"%s\".", cfg.RabbitQueueConsumer)
	<-forever
}
