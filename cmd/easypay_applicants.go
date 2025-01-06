package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

var easypayApplicantsCmd = &cobra.Command{
	Use:   "applicants",
	Short: "Applicants",
	Long:  ".",
	Run: func(cmd *cobra.Command, args []string) {
		applicants, err := easypay.GetApplicants()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Printf("Total: %d, Count: %d, Offset: %d\n", applicants.Metadata.Total, applicants.Metadata.Count, applicants.Metadata.Offset)
	},
}

func init() {
	easypayCommand.AddCommand(easypayApplicantsCmd)
}
