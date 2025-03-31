package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/constants"
	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

var rootCmd = &cobra.Command{
	Use:   "tshop",
	Short: "CLI for terminal.shop",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to terminal.shop!\nTo get started, run \"tshop --help\"")
	},
}

var Client *terminal.Client

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(listCmd)

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		token, ok := os.LookupEnv("TERMINAL_TOKEN")
		if !ok || token == "" {
			fmt.Printf("Environment variable \"TERMINAL_TOKEN\" is missing\n")
			os.Exit(1)
		}

		Client = terminal.NewClient(
			option.WithBaseURL(constants.BaseUrl),
			option.WithBearerToken(token),
		)
	}
}
