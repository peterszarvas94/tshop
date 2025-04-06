package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

var rootCmd = &cobra.Command{
	Use:   "tshop",
	Short: "tshop - CLI for terminal.shop",
	Args:  cobra.ExactArgs(0),
}

var Client *terminal.Client

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error running the command")
		fmt.Println(err.Error())
		os.Exit(1)
		os.Exit(1)
	}
}

// these stucts hold the flags
var User = &terminal.ProfileUser{}
var Address = &terminal.Address{}
var Item = &terminal.CartItem{}

func init() {
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

		var urlOption option.RequestOption
		switch env {
		case ("dev"):
			urlOption = option.WithEnvironmentDev()
			break
		case ("prod"):
			urlOption = option.WithEnvironmentProduction()
			break
		default:
			fmt.Println("Invalid environment variable \"TERMINAL_ENV\", should be \"dev\" or \"prod\"")
			os.Exit(1)
		}

		Client = terminal.NewClient(
			urlOption,
			option.WithBearerToken(token),
		)
	}
}
