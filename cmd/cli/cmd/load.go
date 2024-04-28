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

		cfg := config.NewConfigRabbit()
		rabbit := rabbitmq.RabbitMQ{
			URI:       cfg.RabbitURI,
			QueueName: resourceType,
		}

		for i := range resultGet {
			body, err := json.Marshal(resultGet[i])
			internal.FailOnError(err, "Failed to marshal message")
			rabbit.Publisher(body)
		}
	},
}
