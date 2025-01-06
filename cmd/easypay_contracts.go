package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var easypayContractsCmd = &cobra.Command{
	Use:   "contracts",
	Short: "Contracts",
  Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("easypay contracts")
	},
}

func init() {
	easypayCommand.AddCommand(easypayContractsCmd)
}
