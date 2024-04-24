package cmd

import (
	"encoding/json"

	"crud-golang-rabbitmq-mongo/internal"

	"crud-golang-rabbitmq-mongo/config"
	jph "crud-golang-rabbitmq-mongo/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/rabbitmq"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Get Resource JSON and Send RabbitMQ",

	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]

		resultGet, err := jph.Get(resourceType)
		internal.FailOnError(err, "Failed to get users")

		switch resourceType {
		case "users":
			SendRabbitMQ[jph.User](resultGet, resourceType)
		case "photos":
			SendRabbitMQ[jph.Photo](resultGet, resourceType)
		case "posts":
			SendRabbitMQ[jph.Posts](resultGet, resourceType)
		case "comments":
			SendRabbitMQ[jph.Comments](resultGet, resourceType)
		}
	},
}

func SendRabbitMQ[T jph.ResourceGeneric](resourceBody []byte, resourceType string) {
	var resource jph.Resource[T]
	resource.New(resourceBody)

	cfg := config.NewConfigRabbit()
	rabbit := rabbitmq.RabbitMQ{
		URI:       cfg.RabbitURI,
		QueueName: resource.ResourceType,
	}

	for i := range resource.Data {
		body, err := json.Marshal(resource.Data[i])
		internal.FailOnError(err, "Failed to marshal message")
		rabbit.Publisher(body)
	}
}
