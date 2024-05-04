package cmd

import (
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().VarP(&resourceName, "resource", "r", fmt.Sprintf("resource name. options: %s", strings.Join(validResourceOptions, ", ")))
	readCmd.Flags().IntVar(&resourceId, "id", 0, "resource id")
	readCmd.MarkFlagRequired("resource")
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read one or all values of a resource",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("%s/%s", "http://localhost:5000/api/v1", resourceName.String())

		if resourceId > 0 {
			urlId := fmt.Sprintf("%s/%d", url, resourceId)
			resources, err := jph.Get(urlId, resourceName.String(), false)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("%s\n", resources[0].Echo())
		} else {
			resources, err := jph.Get(url, resourceName.String(), true)
			if err != nil {
				fmt.Println(err)
			}

			for i := range resources {
				fmt.Printf("%s\n", resources[i].Echo())
			}
		}
	},
}
