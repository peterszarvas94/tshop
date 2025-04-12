package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/env"
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

// these hold the flags
var (
	User    = &terminal.ProfileUser{}
	Address = &terminal.Address{}
	Item    = &terminal.CartItem{}
	Sub     = &terminal.Subscription{}
	SubType = "" // can't cast Sub.Schedule.Type into string in flags
)

func init() {

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		var urlOption option.RequestOption
		switch env.Config.TERMINAL_ENV {
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
			option.WithBearerToken(env.Config.TERMINAL_TOKEN),
		)
	}
}
