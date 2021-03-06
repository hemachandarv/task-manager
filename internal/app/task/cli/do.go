package cli

import (
	"fmt"

	"github.com/hemv/task-manager/internal/app/task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks tasks as complete.",
	Run:   complete,
}

func init() {
	rootCmd.AddCommand(doCmd)
}

func complete(cmd *cobra.Command, args []string) {
	ids := parseArgs(args)
	tasks, err := db.FetchAll()
	if err != nil {
		fmt.Printf("unable to fetch tasks: %v\n", err)
		return
	}
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Printf("invalid task number: %d\n", id)
			continue
		}
		task := tasks[id-1]
		err := db.UpdateStatus(task.ID, db.CompletedStatus.String())
		if err != nil {
			fmt.Printf("failed to mark \"%d\" as complete. error: %v\n", id, err)
		} else {
			fmt.Printf("marked \"%d\" as complete.\n", id)
		}
	}
}
