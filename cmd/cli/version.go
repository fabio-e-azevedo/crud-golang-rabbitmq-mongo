package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of crud",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("“My favorite things in life don’t cost any money. It’s really clear that the most precious resource we all have is time.” - Steve Jobs")
		fmt.Println("CRUD v0.9")
	},
}
