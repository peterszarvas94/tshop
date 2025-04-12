package cmd

import (
	"fmt"

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
			helpers.HandleError("Error getting the cart", err, 1)
		}

		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			helpers.HandleError("Error getting the products", err, 1)
		}

		helpers.Section("Items in cart:", func() {
			helpers.PrintCartItems(cart.Data.Items, products.Data)
		})

		helpers.Section("Subtotal:", func() {
			fmt.Println(helpers.FormatPrice(cart.Data.Subtotal))
		})

		address, err := Client.Address.Get(cmd.Context(), cart.Data.AddressID)
		if err != nil {
			helpers.HandleError("Error getting the address", err, 1)
		}

		helpers.Section("Selected address:", func() {
			helpers.PrintAddresses([]terminal.Address{address.Data})
		})

		card, err := Client.Card.Get(cmd.Context(), cart.Data.CardID)
		if err != nil {
			helpers.HandleError("Error getting the card", err, 1)
		}

		helpers.Section("Selected card:", func() {
			helpers.PrintCards([]terminal.Card{card.Data})
		})
	},
}

var addToCartCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add items to the cart",
	Aliases: []string{"k"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		oldCart, err := Client.Cart.Get(cmd.Context())
		if err != nil {
			helpers.HandleError("Error getting the cart", err, 1)
		}

		var count int64 = 0
		for _, item := range oldCart.Data.Items {
			if item.ProductVariantID == Item.ProductVariantID {
				count += item.Quantity
			}
		}

		quantity := Item.Quantity
		if quantity == 0 {
			quantity = 1
		}

		cart, err := Client.Cart.SetItem(cmd.Context(), terminal.CartSetItemParams{
			ProductVariantID: terminal.F(Item.ProductVariantID),
			Quantity:         terminal.F(quantity + count),
		})

		helpers.Section("Item in cart updated", func() {
			products, err := Client.Product.List(cmd.Context())
			if err != nil {
				helpers.HandleError("Error getting products", err, 1)
			}

			helpers.PrintCartItems(cart.Data.Items, products.Data)
		})
	},
}

var removeFromCartCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove items from cart",
	Aliases: []string{"j"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		oldCart, err := Client.Cart.Get(cmd.Context())
		if err != nil {
			helpers.HandleError("Error getting the cart", err, 1)
		}

		var count int64 = 0
		for _, item := range oldCart.Data.Items {
			if item.ProductVariantID == Item.ProductVariantID {
				count += item.Quantity
			}
		}

		quantity := Item.Quantity
		if quantity == 0 {
			quantity = 1
		}

		cart, err := Client.Cart.SetItem(cmd.Context(), terminal.CartSetItemParams{
			ProductVariantID: terminal.F(Item.ProductVariantID),
			Quantity:         terminal.F(count - quantity),
		})

		helpers.Section("Item in cart updated", func() {
			products, err := Client.Product.List(cmd.Context())
			if err != nil {
				helpers.HandleError("Error getting the products", err, 1)
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
			helpers.HandleError("Error selecting the address", err, 1)
		}

		helpers.Section("Address selected", func() {
			address, err := Client.Address.Get(cmd.Context(), args[0])
			if err != nil {
				helpers.HandleError("Error getting the address", err, 1)
			}

			helpers.PrintAddresses([]terminal.Address{address.Data})
		})
	},
}

var selectCardForCartCmd = &cobra.Command{
	Use:     "card",
	Short:   "Select card for shopping cart",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Cart.SetCard(cmd.Context(), terminal.CartSetCardParams{
			CardID: terminal.F(args[0]),
		})
		if err != nil {
			helpers.HandleError("Error selecting the card", err, 1)
		}

		helpers.Section("Card selected", func() {
			card, err := Client.Card.Get(cmd.Context(), args[0])
			if err != nil {
				helpers.HandleError("Error getting the card", err, 1)
			}

			helpers.PrintCards([]terminal.Card{card.Data})
		})
	},
}

var clearCartCmd = &cobra.Command{
	Use:     "clear",
	Short:   "Clear shopping cart",
	Aliases: []string{"delete", "x"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Cart.Clear(cmd.Context())
		if err != nil {
			helpers.HandleError("Error clearing the cart", err, 1)
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
		order, err := Client.Cart.Convert(cmd.Context())
		if err != nil {
			helpers.HandleError("Error placing order", err, 1)
		}

		helpers.Section("Order placed", func() {
			helpers.PrintOrders([]terminal.Order{order.Data})
		})
	},
}

func init() {
	cartCmd.AddCommand(cartInfoCmd)

	addToCartCmd.Flags().StringVarP(&Item.ProductVariantID, "variant", "v", "", "Variant ID")
	addToCartCmd.Flags().Int64VarP(&Item.Quantity, "quantity", "q", 0, "Quantity")
	addToCartCmd.MarkFlagRequired("variant")
	cartCmd.AddCommand(addToCartCmd)

	removeFromCartCmd.Flags().StringVarP(&Item.ProductVariantID, "variant", "v", "", "Variant ID")
	removeFromCartCmd.Flags().Int64VarP(&Item.Quantity, "quantity", "q", 0, "Quantity")
	removeFromCartCmd.MarkFlagRequired("variant")
	cartCmd.AddCommand(removeFromCartCmd)

	cartCmd.AddCommand(selectAddressForCartCmd)
	cartCmd.AddCommand(selectCardForCartCmd)
	cartCmd.AddCommand(clearCartCmd)
	cartCmd.AddCommand(placeOrderCmd)
	rootCmd.AddCommand(cartCmd)
}
