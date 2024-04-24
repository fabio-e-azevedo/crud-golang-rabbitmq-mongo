package rabbitmq

import (
	"context"
	"log"
	"time"

	"crud-golang-rabbitmq-mongo/internal"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	URI       string
	QueueName string
}

func (r RabbitMQ) Publisher(body []byte) {
	conn, err := amqp.Dial(r.URI)
	internal.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	internal.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		r.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
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
