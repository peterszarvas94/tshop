package cmd

import (
	"fmt"

	"github.com/peterszarvas94/tshop/constants"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print tshop version",
	Aliases: []string{"v"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tshop - CLI for terminal.shop")
		fmt.Println(constants.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
