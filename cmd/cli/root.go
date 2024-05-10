package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var resourceName validResourceValue
var resourceId int

var rootCmd = &cobra.Command{
	Use:   "crud",
	Short: "CRUD, command created to apply knowledge in learning the Go language",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
