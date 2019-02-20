package cli

import (
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
