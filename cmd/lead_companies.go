package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/lead"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var leadComaniesCmd = &cobra.Command{
	Use:     "companies [csv]",
	Short:   "Completes companies from CSV file",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		raw := args[0]
		companies, err := lead.CompleteCompanies(raw)
		if err != nil {
			fmt.Println(err)
		}

		decorators.ToDynamoDb(companies)
	},
}

func init() {
	leadCommand.AddCommand(leadComaniesCmd)
}
