package cmd

import (
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
		fmt.Printf("Resource: %s ===>>> Id: %d\n", resourceName, resourceId)
	},
}
