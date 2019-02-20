package cli

import (
	"fmt"

	"github.com/hemv/task-manager/internal/app/task/db"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a task from your task list.",
	Run:   remove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func remove(cmd *cobra.Command, args []string) {
	ids := parseArgs(args)
	tasks, err := db.FetchAll()
	if err != nil {
		fmt.Printf("unable to fetch tasks: %v\n", err)
		return
	}
	tasks = pendingTasks(tasks)
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Printf("invalid task number: %d\n", id)
			continue
		}
		task := tasks[id-1]
		err := db.Delete(task.ID)
		if err != nil {
			fmt.Printf("failed to delete task \"%d\". error: %v\n", id, err)
		} else {
			fmt.Printf("deleted task \"%d\".\n", id)
		}
	}
}
