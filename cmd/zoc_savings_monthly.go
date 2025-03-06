package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/zoc"
)

var zocMonthlySavingsCmd = &cobra.Command{
	Use:     "monthly [json] ",
	Short:   "View monthly savings",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		json := args[0]
		zoc.GetMonthlySavings(json)
	},
}

func init() {
	zocSavingsCmd.AddCommand(zocMonthlySavingsCmd)
}
