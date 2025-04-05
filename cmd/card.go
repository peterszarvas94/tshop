package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

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
			fmt.Println("Could not list cards")
			fmt.Println(err)
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "ID\tBrand\tExipration\tNumber")
		fmt.Fprintln(w, "--\t-----\t----------\t------")
		for _, card := range cards.Data {
			expiration := fmt.Sprintf("%d/%d", card.Expiration.Month, card.Expiration.Year)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", card.ID, card.Brand, expiration, card.Last4)
		}

		w.Flush()
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
			fmt.Println("Could not create card")
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Enter card information in the following URL:")
		fmt.Println(res.Data.URL)
		helpers.OpenBrowser(res.Data.URL)
	},
}

var deleteCardCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete card",
	Aliases: []string{"d"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Card.Delete(cmd.Context(), args[0])
		if err != nil {
			fmt.Println("Could not delete card")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Card deleted succeccfully")
	},
}
