package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/zoc"
)

var zocSavingsCmd = &cobra.Command{
	Use:     "savings [string]",
	Short:   "View savings",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		data := args[0]

		_, err := zoc.GetSavings(data)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("DONE")
	},
}

func init() {
	zocCommand.AddCommand(zocSavingsCmd)
}
