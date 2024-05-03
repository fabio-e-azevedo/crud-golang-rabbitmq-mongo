package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().VarP(&resourceName, "resource", "r", fmt.Sprintf("resource name. options: %s", strings.Join(validResourceOptions, ", ")))
	deleteCmd.Flags().IntVar(&resourceId, "id", 0, "resource id")
	deleteCmd.MarkFlagRequired("resource")
	deleteCmd.MarkFlagRequired("id")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Resource by ID",
	Run: func(cmd *cobra.Command, args []string) {
		deleteResourceById(resourceName.String(), resourceId)
		//fmt.Printf("delete in %s id %d", resourceName.String(), resourceId)
	},
}

func deleteResourceById(resourceType string, resourceId int) {
	url := fmt.Sprintf("http://localhost:5000/api/v1/%s/%d", resourceType, resourceId)
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
		fmt.Printf("not found id %d\n", resourceId)
		return
	}
}
