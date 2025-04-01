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
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Welcome to terminal.shop!\nTo get started, run \"tshop --help\"")
	// },
}

var Client *terminal.Client

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// version
	rootCmd.AddCommand(versionCmd)

	// product
	productCmd.AddCommand(listProductsCmd)
	productCmd.AddCommand(getProductCmd)
	rootCmd.AddCommand(productCmd)

	// profil
	profilCmd.AddCommand(profilInfoCmd)
	rootCmd.AddCommand(profilCmd)

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		token, ok := os.LookupEnv("TERMINAL_TOKEN")
		if !ok || token == "" {
			fmt.Printf("Environment variable \"TERMINAL_TOKEN\" is missing\n")
			os.Exit(1)
		}

		env, ok := os.LookupEnv("TERMINAL_ENV")
		if !ok || env == "" {
			fmt.Printf("Environment variable \"TERMINAL_ENV\" is missing\n")
			os.Exit(1)
		}

		var url string
		switch env {
		case ("dev"):
			url = constants.DevUrl
			break
		case ("prod"):
			url = constants.ProdUrl
			break
		default:
			fmt.Println("Invalid environment variable \"TERMINAL_ENV\", should be \"dev\" or \"prod\"")
			os.Exit(1)
		}

		Client = terminal.NewClient(
			option.WithBaseURL(url),
			option.WithBearerToken(token),
		)
	}
}
