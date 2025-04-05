package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
)

var cartCmd = &cobra.Command{
	Use:     "cart",
	Short:   "Manage cart",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(0),
}

var cartInfoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Get cart info",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cart, err := Client.Cart.Get(cmd.Context())
		if err != nil {
			fmt.Println("Could not get cart info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Items in cart")
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "ID\tVariant ID\tQuantity\tSubtotal")
		fmt.Fprintln(w, "--\t----------\t--------\t--------")
		for _, item := range cart.Data.Items {
			fmt.Fprintf(w, "%s\t%s\t%d\t%d\n", item.ID, item.ProductVariantID, item.Quantity, item.Subtotal)
		}
		w.Flush()
	},
}

var addItemToCartCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add item to cart",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Cart.SetItem(cmd.Context(), terminal.CartSetItemParams{
			ProductVariantID: terminal.F(Item.ProductVariantID),
			Quantity:         terminal.F(Item.Quantity),
		})
		if err != nil {
			fmt.Println("Could not add item to the cart")
			fmt.Println(err.Error())
		}

		fmt.Println("Item added to cart")
	},
}
