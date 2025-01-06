package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/internal/easypay"
)

var easypayApplicantsCmd = &cobra.Command{
	Use:   "applicants",
	Short: "Applicants",
  Long: "Search among applicants",
	Run: func(cmd *cobra.Command, args []string) {
    applicants, err := easypay.GetApplicants();
    if err != nil {
			fmt.Println("ERROR:", err)
      return
    }

    fmt.Println(applicants)

	},
}

func init() {
	easypayCommand.AddCommand(easypayApplicantsCmd)
}
