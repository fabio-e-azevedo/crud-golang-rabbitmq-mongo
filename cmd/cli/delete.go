package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Resource by ID",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("enter the name of a resource: albums, comments, photos, posts, todos, users")
			os.Exit(1)
		}

		if len(args) == 1 {
			fmt.Println("enter the resource id")
			os.Exit(1)
		}

		resourceType := args[0]

		valueID, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		resources := map[string]bool{
			"albums":   true,
			"comments": true,
			"photos":   true,
			"posts":    true,
			"todos":    true,
			"users":    true,
		}

		if resources[resourceType] {
			url := fmt.Sprintf("http://localhost:5000/api/v1/%s/%d", resourceType, valueID)
			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				fmt.Println("error request delete:", err)
				os.Exit(1)
			}

			client := http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("error send request delete:", err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			// body, err := io.ReadAll(resp.Body)
			// if err != nil {
			// 	fmt.Println("error reading response body:", err)
			// 	return
			// }

			if resp.StatusCode == 404 {
				fmt.Printf("not found id %d\n", valueID)
				os.Exit(1)
			}
		} else {
			fmt.Printf("not found resource \"%s\"\n", resourceType)
			fmt.Println("enter the name of a resource: albums, comments, photos, posts, todos, users")
			os.Exit(1)
		}
	},
}
