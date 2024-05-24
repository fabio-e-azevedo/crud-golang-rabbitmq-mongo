package cmd

import (
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var outputJson bool

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().VarP(&resourceName, "resource", "r", fmt.Sprintf("resource name. options: %s", strings.Join(validResourceOptions, ", ")))
	readCmd.Flags().IntVar(&resourceId, "id", 0, "resource id")
	readCmd.Flags().BoolVar(&outputJson, "json", false, "output format json")
	readCmd.MarkFlagRequired("resource")
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Getting by ID or Listing all resources",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("%s/%s", "http://localhost/api/v1", resourceName.String())

		if resourceId > 0 {
			urlId := fmt.Sprintf("%s/%d", url, resourceId)
			resources, err := jph.Get(urlId, resourceName.String(), false)
			if err != nil {
				fmt.Printf("not found id %d\n", resourceId)
				return
			}

			if outputJson {
				fmt.Println(string(resources[0].Show()))
			} else {
				fmt.Printf("%s\n", resources[0].Echo())
			}
		} else {
			resources, err := jph.Get(url, resourceName.String(), true)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if outputJson {
				result, _ := json.Marshal(resources)
				fmt.Println(string(result))
			} else {
				for i := range resources {
					fmt.Printf("%s\n", resources[i].Echo())
				}
			}
		}
	},
}
