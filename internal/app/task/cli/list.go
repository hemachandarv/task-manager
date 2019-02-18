package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("list invoked.")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
