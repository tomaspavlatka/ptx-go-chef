package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var easypayContractViewCmd = &cobra.Command{
	Use:   "view [contract_id]",
	Short: "Views a contract",
	Args:  cobra.MinimumNArgs(1),
  Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		contract, err := easypay.GetContract(args[0])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		decorators.ToContract(*contract)
	},
}

func init() {
	easypayContractsCmd.AddCommand(easypayContractViewCmd)
}
