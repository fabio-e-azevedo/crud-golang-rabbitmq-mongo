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
	deleteCmd.AddCommand(albumsCmd)
	deleteCmd.AddCommand(commentsCmd)
	deleteCmd.AddCommand(photosCmd)
	deleteCmd.AddCommand(postsCmd)
	deleteCmd.AddCommand(todosCmd)
	deleteCmd.AddCommand(usersCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Resource by ID",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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

var albumsCmd = &cobra.Command{
	Use:   "albums [id] or",
	Short: "Resource albums",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// if len(args) < 1 {
		// 	fmt.Println("enter the resource id")
		// 	return
		// }

		albumsID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("the albums id must be a number")
			return
		}

		deleteResourceById("albums", albumsID)
	},
}

var commentsCmd = &cobra.Command{
	Use:   "comments [id] or",
	Short: "Resource comments",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commentsID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("the comments id must be a number")
			return
		}

		deleteResourceById("comments", commentsID)
	},
}

var photosCmd = &cobra.Command{
	Use:   "photos [id] or",
	Short: "Resource photos",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		photosID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("the photos id must be a number")
			return
		}

		deleteResourceById("photos", photosID)
	},
}

var postsCmd = &cobra.Command{
	Use:   "posts [id] or",
	Short: "Resource posts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		postsID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("the posts id must be a number")
			return
		}

		deleteResourceById("posts", postsID)
	},
}

var todosCmd = &cobra.Command{
	Use:   "todos [id] or",
	Short: "Resource todos",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todosID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("the todos id must be a number")
			return
		}

		deleteResourceById("todos", todosID)
	},
}

var usersCmd = &cobra.Command{
	Use:   "users [id] or",
	Short: "Resource users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("the users id must be a number")
			return
		}

		deleteResourceById("users", userID)
	},
}
