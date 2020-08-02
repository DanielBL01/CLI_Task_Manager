package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/DanielBL01/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Printf("Added '%s'\n", task)
	},
}

// The init() function runs before main.go for initial setup
func init() {
	RootCmd.AddCommand(addCmd)
}
