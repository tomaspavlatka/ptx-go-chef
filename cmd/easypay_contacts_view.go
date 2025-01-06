package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var includeAudit = false

var easypayContractViewCmd = &cobra.Command{
	Use:     "view [contract_id]",
	Short:   "Views a contract",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		contactId := args[0]
		contract, err := easypay.GetContract(contactId)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		decorators.ToContract(*contract)

		if includeAudit {
			audits, err := easypay.GetContractAudits(contactId)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			decorators.ToAudits(audits.Records)
		}
	},
}

func init() {
	easypayContractsCmd.AddCommand(easypayContractViewCmd)
	easypayContractViewCmd.Flags().BoolVarP(&includeAudit, "audit", "a", false, "include audit")
}
