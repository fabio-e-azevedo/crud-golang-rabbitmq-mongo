package cmd

import (
	"encoding/json"

	"crud-golang-rabbitmq-mongo/pkg/utils"

	"crud-golang-rabbitmq-mongo/pkg/config"
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/rabbitmq"

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
		utils.FailOnError(err, "Failed to get users")

		cfg := config.NewConfigRabbit()
		rabbit := rabbitmq.RabbitMQ{
			URI:       cfg.RabbitURI,
			QueueName: resourceType,
		}

		for i := range resultGet {
			body, err := json.Marshal(resultGet[i])
			utils.FailOnError(err, "Failed to marshal message")
			rabbit.Publisher(body)
		}
	},
}
