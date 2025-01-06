package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var easypayApplicantsCmd = &cobra.Command{
	Use:   "applicants",
	Short: "Applicants",
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Easypay Applicants")
	},
}

func init() {
	easypayCommand.AddCommand(easypayApplicantsCmd)
}
