package rabbitmq

import (
	"context"
	"log"
	"os"
	"time"

	"crud-golang-rabbitmq-mongo/internal"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publisher(body []byte) {
	err := godotenv.Load()
	internal.FailOnError(err, "No .env file found")

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
		"users", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	internal.FailOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	internal.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
