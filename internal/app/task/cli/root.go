package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
	Long: `Usage:
	task [command]

Available Commands:
	add         Add a new task to your TODO list
	do          Mark a task on your TODO list as complete
	list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.`,
}

// Execute is the entry point for cli
// to start engaging with the user.
// It will initialise all sub-commands.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
