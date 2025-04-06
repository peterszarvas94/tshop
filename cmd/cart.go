package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/helpers"
	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
)

var cartCmd = &cobra.Command{
	Use:     "cart",
	Short:   "Manage shopping cart",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(0),
}

var cartInfoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Get shopping cart info",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cart, err := Client.Cart.Get(cmd.Context())
		if err != nil {
			fmt.Println("Could not get cart info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			fmt.Println("Could not get products")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.Section("Items in cart:", func() {
			helpers.PrintCartItems(cart.Data.Items, products.Data)
		})

		helpers.Section("Subtotal:", func() {
			fmt.Println(helpers.FormatPrice(cart.Data.Subtotal))
		})

		address, err := Client.Address.Get(cmd.Context(), cart.Data.AddressID)
		if err != nil {
			fmt.Println("Could not get address info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.Section("Selected address:", func() {
			helpers.PrintAddresses([]terminal.Address{address.Data})
		})

		card, err := Client.Card.Get(cmd.Context(), cart.Data.CardID)
		if err != nil {
			fmt.Println("Could not get card info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.Section("Selected card:", func() {
			helpers.PrintCards([]terminal.Card{card.Data})
		})
	},
}

var updateItemInCartCmd = &cobra.Command{
	Use:     "update",
	Short:   "Add / Update item in shopping cart",
	Long:    "Add / Update item in shopping cart.\nIf the item does not exists in the cart yet, this command will add it.\nIf the item is already in the cart, it will overwrite the quantity.\nIf you set it to zero, it will delete the item from the cart.",
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

		helpers.Section("Item in cart updated", func() {
			products, err := Client.Product.List(cmd.Context())
			if err != nil {
				fmt.Println("Could not get products")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			helpers.PrintCartItems(cart.Data.Items, products.Data)
		})
	},
}

var selectAddressForCartCmd = &cobra.Command{
	Use:     "address",
	Short:   "Select address for shopping cart",
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

		helpers.Section("Address selected", func() {
			address, err := Client.Address.Get(cmd.Context(), args[0])
			if err != nil {
				fmt.Println("Could not get address info")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			helpers.PrintAddresses([]terminal.Address{address.Data})
		})
	},
}

var selectCardForCartCmd = &cobra.Command{
	Use:     "card",
	Short:   "Select card for shopping cart",
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

		helpers.Section("Card selected", func() {
			card, err := Client.Card.Get(cmd.Context(), args[0])
			if err != nil {
				fmt.Println("Could not get card info")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			helpers.PrintCards([]terminal.Card{card.Data})
		})
	},
}

var clearCartCmd = &cobra.Command{
	Use:     "clear",
	Short:   "Clear shopping cart",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Cart.Clear(cmd.Context())
		if err != nil {
			fmt.Println("Could not clear shopping cart")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Shopping cart cleared")
	},
}

var placeOrderCmd = &cobra.Command{
	Use:     "order",
	Short:   "Place order with current shopping cart",
	Aliases: []string{"o"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := Client.Cart.Convert(cmd.Context())
		if err != nil {
			fmt.Println("Could place order")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.Section("Order placed", func() {
			helpers.PrintOrders([]terminal.Order{res.Data})
		})
	},
}

func init() {
	cartCmd.AddCommand(cartInfoCmd)
	updateItemInCartCmd.Flags().StringVarP(&Item.ProductVariantID, "variant", "v", "", "Variant ID")
	updateItemInCartCmd.Flags().Int64VarP(&Item.Quantity, "quantity", "q", 0, "Quantity")
	updateItemInCartCmd.MarkFlagRequired("variant")
	updateItemInCartCmd.MarkFlagRequired("quantity")
	cartCmd.AddCommand(updateItemInCartCmd)
	cartCmd.AddCommand(selectAddressForCartCmd)
	cartCmd.AddCommand(selectCardForCartCmd)
	cartCmd.AddCommand(clearCartCmd)
	cartCmd.AddCommand(placeOrderCmd)
	rootCmd.AddCommand(cartCmd)
}
