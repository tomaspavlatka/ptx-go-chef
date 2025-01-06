package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var easypayContractsCmd = &cobra.Command{
	Use:   "contracts",
	Short: "Easypay Contracts",
	Run: func(cmd *cobra.Command, args []string) {
		contracts, err := easypay.GetContracts()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

    // VISUAL
		fmt.Printf(
			"Total: %d, Count: %d, Offset: %d\n",
			contracts.Metadata.Total,
			contracts.Metadata.Count,
			contracts.Metadata.Offset,
		)
    fmt.Println()

    for _, contract := range(contracts.Records) {
      decorators.ToContract(contract);
    }

	},
}

func init() {
	easypayCommand.AddCommand(easypayContractsCmd)
}
