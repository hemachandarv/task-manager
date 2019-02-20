package cli

import (
	"fmt"
	"strings"

	"github.com/hemv/task-manager/internal/app/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run:   add,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	task := strings.Join(args, " ")
	err := db.Insert(task)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("added \"%s\" to your task list.\n", task)
}
