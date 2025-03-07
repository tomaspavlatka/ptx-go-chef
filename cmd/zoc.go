package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var zocCommand = &cobra.Command{
	Use:     "zoc",
	Short:   "ZOC",
	Long:    "Welcome to the ZOC world",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zoc called")
	},
}

func init() {
	rootCmd.AddCommand(zocCommand)
}
