package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var (
	limit     int8
	offset    int8
	sortBy    string
	companyId string
	status    string
)

var easypayContractSearchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Searches among contracts",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := easypay.GetContracts(easypay.ContractsOpts{
      Limit: limit,
      Offset: offset,
      SortBy: sortBy,
      CompanyId: companyId,
      Status: status,
		})
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		for _, contract := range resp.Records {
			decorators.ToContract(contract)
			fmt.Println()
		}

    if resp.Metadata.Count == 0 {
      fmt.Println("NO RESULTS")
    }
	},
}

func init() {
	easypayContractsCmd.AddCommand(easypayContractSearchCmd)
	easypayContractSearchCmd.Flags().Int8VarP(&limit, "limit", "l", 20, "max number of results")
	easypayContractSearchCmd.Flags().Int8VarP(&offset, "offset", "o", 0, "offset")
	easypayContractSearchCmd.Flags().StringVarP(&companyId, "company", "c", "", "restrict to a company only")
	easypayContractSearchCmd.Flags().StringVarP(&status, "status", "t", "", "restrict to a status only")
	easypayContractSearchCmd.Flags().StringVarP(&sortBy, "sort", "s", "createdAt", "how to sort the data, use - (eg. -createdAt) for descending order")
}
