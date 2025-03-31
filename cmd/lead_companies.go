package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/lead"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var (
	year  string
	month string
)

var leadCompaniesCmd = &cobra.Command{
	Use:     "companies [json]",
	Short:   "Extract companies from JSON",
	Args:    cobra.MinimumNArgs(0),
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		companies, err := lead.ConvertCompanies(year, month)
		if err != nil {
			fmt.Println(err)
		}

		decorators.ToCompanies(companies)
	},
}

func init() {
	leadCommand.AddCommand(leadCompaniesCmd)
	leadCompaniesCmd.Flags().StringVarP(&year, "year", "y", "2025", "year you want data for")
	leadCompaniesCmd.Flags().StringVarP(&month, "month", "m", "1", "month you want data for")
}
