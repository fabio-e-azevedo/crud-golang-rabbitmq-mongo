package rabbitmq

import (
	"context"
	"log"
	"time"

	"crud-golang-rabbitmq-mongo/pkg/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMQ) Publisher(body []byte) {
	log.SetPrefix("[RBT] ")

	err := GetConnection(r)
	utils.FailOnError(err, "failed to connect to rabbitmq")

	ch, err := r.Connection.Channel()
	utils.FailOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		r.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "failed to declare a queue")

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

	utils.FailOnError(err, "failed to publish a message")
	log.Printf("| [x] Sent: %s\n", body)
	log.SetPrefix("")
}
