package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/zoc"
)

var zocYearlySavingsCmd = &cobra.Command{
	Use:     "yearly [json] ",
	Short:   "View yearly savings",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		json := args[0]
		zoc.GetYearlySavings(json)
	},
}

func init() {
	zocSavingsCmd.AddCommand(zocYearlySavingsCmd)
}
