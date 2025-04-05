package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/peterszarvas94/tshop/helpers"
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

		fmt.Println("Items in cart:")
		printCartItems(cart.Data.Items)
		fmt.Println()

		address, err := Client.Address.Get(cmd.Context(), cart.Data.AddressID)
		if err != nil {
			fmt.Println("Could not get address info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Selected address:")
		printAddresses([]terminal.Address{address.Data})
		fmt.Println()

		card, err := Client.Card.Get(cmd.Context(), cart.Data.CardID)
		if err != nil {
			fmt.Println("Could not get card info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Selected card:")
		printCards([]terminal.Card{card.Data})

	},
}

var updateItemInCartCmd = &cobra.Command{
	Use:     "update",
	Short:   "Add / Update item in cart",
	Long:    "Add / Update item in cart.\nIf the item does not exists in the cart yet, this command will add it.\nIf the item is already in the cart, it will overwrite the quantity.\nIf you set it to zero, it will delete the item from the cart.",
	Aliases: []string{"u"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cart, err := Client.Cart.SetItem(cmd.Context(), terminal.CartSetItemParams{
			ProductVariantID: terminal.F(Item.ProductVariantID),
			Quantity:         terminal.F(Item.Quantity),
		})
		if err != nil {
			fmt.Println("Could not add item to the cart")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if Item.Quantity == 0 {
			fmt.Println("Item removed from cart")
			os.Exit(0)
		}

		fmt.Println("Item in cart updated")
		printCartItems(cart.Data.Items)
	},
}

var selectAddressForCartCmd = &cobra.Command{
	Use:     "address",
	Short:   "Select address",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Cart.SetAddress(cmd.Context(), terminal.CartSetAddressParams{
			AddressID: terminal.F(args[0]),
		})
		if err != nil {
			fmt.Println("Could not select address")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		address, err := Client.Address.Get(cmd.Context(), args[0])
		if err != nil {
			fmt.Println("Address selected, but could not get address info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printAddresses([]terminal.Address{address.Data})
		fmt.Println("Address selected")
	},
}

var selectCardForCartCmd = &cobra.Command{
	Use:     "card",
	Short:   "Select card",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Cart.SetCard(cmd.Context(), terminal.CartSetCardParams{
			CardID: terminal.F(args[0]),
		})
		if err != nil {
			fmt.Println("Could not select card")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		card, err := Client.Card.Get(cmd.Context(), args[0])
		if err != nil {
			fmt.Println("Card selected, but could not get card info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Card selected")
		printCards([]terminal.Card{card.Data})
	},
}

func printCartItems(items []terminal.CartItem) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Variant ID\tQuantity\tSubtotal")
	fmt.Fprintln(w, "----------\t--------\t--------")
	for _, item := range items {
		total := helpers.PrettyPrice(item.Subtotal)
		fmt.Fprintf(w, "%s\t%d\t%s\n", item.ProductVariantID, item.Quantity, total)
	}
	w.Flush()
}
