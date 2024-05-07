package rabbitmq

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	URI        string
	QueueName  string
	Connection *amqp.Connection
}

var singletonConnection *amqp.Connection
var connectionError error
var once sync.Once

func GetConnection(cfg *RabbitMQ) error {
	once.Do(func() {
		conn, err := amqp.Dial(cfg.URI)
		if err != nil {
			connectionError = err
		}
		singletonConnection = conn
	})
	cfg.Connection = singletonConnection
	return connectionError
}
