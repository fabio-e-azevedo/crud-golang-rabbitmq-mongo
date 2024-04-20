package cmd

import (
	"encoding/json"

	"crud-golang-rabbitmq-mongo/internal"

	ph "crud-golang-rabbitmq-mongo/jsonplaceholder"
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
		resultGet, err := ph.Get("users")
		internal.FailOnError(err, "Failed to get users")

		var users []ph.UserFromJson

		err = json.Unmarshal(*resultGet, &users)
		if err != nil {
			internal.FailOnError(err, "Failed to Unmarshal message")
		}

		for i := range users {
			msg := ph.Resource{
				ResourceType: "users",
				Data:         (users)[i],
			}

			body, err := json.Marshal(msg)
			internal.FailOnError(err, "Failed to marshal message")
			rt.Publisher(body, "users")
		}
	},
}
