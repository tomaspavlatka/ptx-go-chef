package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

// easypayMeCmd represents the easypayMe command
var easypayMeCmd = &cobra.Command{
	Use:     "health",
	Short:   "Health",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := easypay.GetHealth()
		if err != nil {
			fmt.Printf(`ERROR: %v`, err)
		} else {
			fmt.Println("EASYPAY is up and running")
		}

	},
}

func init() {
	easypayCommand.AddCommand(easypayMeCmd)
}
