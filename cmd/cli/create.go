package cmd

import (
	"bytes"
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var filePath string

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().VarP(&resourceName, "resource", "r", fmt.Sprintf("resource name. options: %s", strings.Join(validResourceOptions, ", ")))
	createCmd.Flags().StringVarP(&filePath, "file", "f", "", "JSON file path with resource array")
	createCmd.MarkFlagRequired("resource")
	createCmd.MarkFlagRequired("file")
}

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create resources by reading JSON file",
	Example: "crud create --resource name --file /path/to/file.json",
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := resourceName.String()

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(fileContent))

		resources, err := jph.GetResources(resourceType, 1, fileContent)
		utils.FailOnError(err, fmt.Sprintf("failed to read file %s", filePath))

		urlPost := fmt.Sprintf("http://localhost/api/v1/%s", resourceType)

		for i := range resources {
			msgByte, err := json.Marshal(resources[i])
			utils.FailOnError(err, "failed to marshal message")

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
