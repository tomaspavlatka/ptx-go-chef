package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var zocSavingsCmd = &cobra.Command{
	Use:     "savings",
	Short:   "View savings",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DONE")
	},
}

func init() {
	zocCommand.AddCommand(zocSavingsCmd)
}
