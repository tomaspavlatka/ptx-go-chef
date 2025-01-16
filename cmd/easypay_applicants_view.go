package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
)

var easypayApplicantViewCmd = &cobra.Command{
	Use:     "view [applicant_id]",
	Short:   "Views a applicant",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		applicantId := args[0]
		applicant, err := easypay.GetApplicant(applicantId)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		decorators.ToApplicant(*applicant)

		if includeAudit {
			audits, err := easypay.GetApplicantAudits(applicantId)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

      fmt.Println()
			decorators.ToApplicantAudits(audits.Records)
		}

		if includeKins {
			kins, err := easypay.GetResourceKinsAudit(applicantId)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

      fmt.Println()
			decorators.ToKins(kins.Records)
		}
	},
}

func init() {
	easypayApplicantsCmd.AddCommand(easypayApplicantViewCmd)
	easypayApplicantViewCmd.Flags().BoolVarP(&includeAudit, "audit", "a", false, "include audit")
	easypayApplicantViewCmd.Flags().BoolVarP(&includeKins, "kins", "k", false, "include kins")
}
