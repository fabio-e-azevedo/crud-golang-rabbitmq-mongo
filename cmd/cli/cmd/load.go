package cmd

import (
	"encoding/json"

	"crud-golang-rabbitmq-mongo/internal"

	jph "crud-golang-rabbitmq-mongo/jsonplaceholder"
	rt "crud-golang-rabbitmq-mongo/rabbitmq"

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
			SendRabbitMQ[jph.User](resourceType, resultGet)
		case "photos":
			SendRabbitMQ[jph.Photo](resourceType, resultGet)
		}
	},
}

func SendRabbitMQ[T jph.ResourceGeneric](resourceType string, resourceBody []byte) {
	var resource jph.Resource[T]
	err := json.Unmarshal(resourceBody, &resource)
	internal.FailOnError(err, "Failed to Unmarshal message")

	for i := range resource.Data {
		body, err := json.Marshal(resource.Data[i])
		internal.FailOnError(err, "Failed to marshal message")
		rt.Publisher(body, resource.ResourceType)
	}
}
