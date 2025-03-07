package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/zoc"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var zocSavingsCmd = &cobra.Command{
	Use:     "savings [json]",
	Short:   "View savings",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		json := args[0]
    savings, err := zoc.GetSavings(json);
    if (err != nil) {
      fmt.Println(err)
    }

    decorators.ToSavings(savings);
	},
}

func init() {
	zocCommand.AddCommand(zocSavingsCmd)
}
