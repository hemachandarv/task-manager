package cli

import (
	"fmt"
	"strconv"

	"github.com/hemv/task-manager/internal/app/task/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
}

// Execute is the entry point for cli to start engaging with the user.
// It will initialise all sub-commands.
func Execute() error {
	err := rootCmd.Execute()
	return err
}

func parseArgs(args []string) []int {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("failed to parse the argument: %s, hence skipping.\n", arg)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
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
