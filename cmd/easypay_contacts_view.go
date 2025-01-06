package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var (
	includeAudit bool
	includeKins  bool
)

var easypayContractViewCmd = &cobra.Command{
	Use:     "view [contract_id]",
	Short:   "Views a contract",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		contractId := args[0]
		contract, err := easypay.GetContract(contractId)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		decorators.ToContract(*contract)

		if includeAudit {
			audits, err := easypay.GetContractAudits(contractId)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			decorators.ToContractAudits(audits.Records)
		}

		if includeKins {
			kins, err := easypay.GetContractKinsAudit(contractId)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			decorators.ToContractKins(kins.Records)
		}
	},
}

func init() {
	easypayContractsCmd.AddCommand(easypayContractViewCmd)
	easypayContractViewCmd.Flags().BoolVarP(&includeAudit, "audit", "a", false, "include audit")
	easypayContractViewCmd.Flags().BoolVarP(&includeKins, "kins", "k", false, "include kins")
}
