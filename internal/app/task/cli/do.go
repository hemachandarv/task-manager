package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks tasks as complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse the argument: %s\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Printf("Done: %v.\n", ids)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
