package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"crud-golang-rabbitmq-mongo/pkg/utils"

	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loadCmd)
	loadCmd.Flags().VarP(&resourceName, "resource", "r", fmt.Sprintf("resource name. options: %s", strings.Join(validResourceOptions, ", ")))
	loadCmd.MarkFlagRequired("resource")
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Get Resources JSON and Send API HTTP",

	Run: func(cmd *cobra.Command, args []string) {
		resourceType := resourceName.String()
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/%s", resourceType)

		resultGet, err := jph.Get(url, resourceType, true)
		utils.FailOnError(err, fmt.Sprintf("Failed to GET %s", resourceType))

		urlPost := fmt.Sprintf("http://localhost:5000/api/v1/%s", resourceType)

		for i := range resultGet {
			msgByte, err := json.Marshal(resultGet[i])
			utils.FailOnError(err, "Failed to marshal message")

			msgReader := bytes.NewReader(msgByte)

			resp, err := http.Post(urlPost, "application/json", msgReader)
			if err != nil {
				log.Println(err)
			}

			resultPost, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			resp.Body.Close()

			log.Println(string(resultPost))
		}
	},
}
