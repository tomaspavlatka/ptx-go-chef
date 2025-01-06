package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var easypayContractSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches among contracts",
  Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := easypay.GetContracts()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

    for _, contract := range(resp.Records) {
      decorators.ToContract(contract);
    }
	},
}

func init() {
	easypayContractsCmd.AddCommand(easypayContractSearchCmd)
}
