package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var easypayCommand = &cobra.Command{
	Use:     "easypay",
	Short:   "Easypay",
	Long:    "Welcome to the easypay world",
	Aliases: []string{"ep"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("easypay called")
	},
}

var (
	includeAudit bool
	includeKins  bool
)

func init() {
	rootCmd.AddCommand(easypayCommand)
}
