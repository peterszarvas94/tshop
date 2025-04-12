package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/env"
	"github.com/peterszarvas94/tshop/helpers"
	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
)

var tokenCmd = &cobra.Command{
	Use:     "token",
	Short:   "Manage access tokens",
	Aliases: []string{"t"},
}

var listTokensCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all access tokens",
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tokens, err := Client.Token.List(cmd.Context())
		if err != nil {
			helpers.HandleError("Error getting the tokens", err, 1)
		}

		helpers.PrintTokens(tokens.Data)
	},
}

var createTokenCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create access token",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		token, err := Client.Token.New(cmd.Context())
		if err != nil {
			helpers.HandleError("Error creating the token", err, 1)
		}

		// "created" is missing from response, so we make a new request
		newToken, err := Client.Token.Get(cmd.Context(), token.Data.ID)
		if err != nil {
			helpers.HandleError("Error getting the token", err, 1)
		}

		helpers.Section("Token value will not be shown again", func() {
			helpers.PrintTokens([]terminal.Token{{
				ID:      newToken.Data.ID,
				Token:   token.Data.Token,
				Created: newToken.Data.Created,
			}})
		})
	},
}

var deleteTokenCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Revoke access token",
	Aliases: []string{"revoke", "x"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if env.Config.TERMINAL_TOKEN_ID == args[0] {
			if !helpers.Confirm("This token is used by your environment, do you want to delete it?") {
				fmt.Println("Cancelled")
				os.Exit(0)
			}
		}

		fmt.Println("Confirmed")

		_, err := Client.Token.Delete(cmd.Context(), args[0])
		if err != nil {
			helpers.HandleError("Error deleting the token", err, 1)
		}

		fmt.Println("Token deleted successfully")
	},
}

func init() {
	tokenCmd.AddCommand(listTokensCmd)
	tokenCmd.AddCommand(createTokenCmd)
	tokenCmd.AddCommand(deleteTokenCmd)
	rootCmd.AddCommand(tokenCmd)
}
