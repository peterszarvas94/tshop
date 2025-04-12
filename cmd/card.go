package cmd

import (
	"fmt"

	"github.com/peterszarvas94/tshop/helpers"
	"github.com/spf13/cobra"
)

var cardCmd = &cobra.Command{
	Use:     "card",
	Short:   "Manage credit cards",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(0),
}

var listCardsCmd = &cobra.Command{
	Use:     "list",
	Short:   "List saved cards",
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cards, err := Client.Card.List(cmd.Context())
		if err != nil {
			helpers.HandleError("Error getting the cards", err, 1)
		}

		helpers.PrintCards(cards.Data)
	},
}

var addCardCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create card",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := Client.Card.Collect(cmd.Context())
		if err != nil {
			helpers.HandleError("Error creating card", err, 1)
		}

		fmt.Println("Enter card information in the following URL:")
		fmt.Println(res.Data.URL)
		helpers.OpenBrowser(res.Data.URL)
	},
}

var deleteCardCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete card",
	Aliases: []string{"x"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Card.Delete(cmd.Context(), args[0])
		if err != nil {
			helpers.HandleError("Error deleting the card", err, 1)
		}

		fmt.Println("Card deleted succeccfully")
	},
}

func init() {
	cardCmd.AddCommand(listCardsCmd)
	cardCmd.AddCommand(addCardCmd)
	cardCmd.AddCommand(deleteCardCmd)
	rootCmd.AddCommand(cardCmd)
}
