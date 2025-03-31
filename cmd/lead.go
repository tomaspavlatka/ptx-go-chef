package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var leadCommand = &cobra.Command{
	Use:     "lead",
	Short:   "Lead engine",
	Long:    "Welcome to the Lead world",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lead called")
	},
}

func init() {
	rootCmd.AddCommand(leadCommand)
}
