package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/zoc"
)

var zocSeasonalSavingsCmd = &cobra.Command{
	Use:     "seasonal [json] ",
	Short:   "View seasonal savings",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		json := args[0]
		zoc.GetSeasonalSavings(json)
	},
}

func init() {
	zocSavingsCmd.AddCommand(zocSeasonalSavingsCmd)
}
