package cli

import (
	"fmt"

	"github.com/hemv/task-manager/internal/app/task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all of your tasks.",
	Run:   list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	tasks, err := db.FetchAll()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	tasks = pendingTasks(tasks)
	if len(tasks) == 0 {
		fmt.Println("you have no tasks to complete.")
		return
	}
	fmt.Println("you have the following tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Data)
	}
}

func pendingTasks(tasks []db.Task) []db.Task {
	var res []db.Task
	for _, task := range tasks {
		if task.Status == db.PendingStatus {
			res = append(res, task)
		}
	}
	return res
}
