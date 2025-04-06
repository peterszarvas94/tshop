package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/helpers"
	"github.com/spf13/cobra"
)

var orderCmd = &cobra.Command{
	Use:     "order",
	Short:   "Manage orders",
	Aliases: []string{"o"},
	Args:    cobra.ExactArgs(0),
}

var listOrdersCmd = &cobra.Command{
	Use:     "list",
	Short:   "List orders",
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		orders, err := Client.Order.List(cmd.Context())
		if err != nil {
			fmt.Println("Could not get orders")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.PrintOrders(orders.Data)
	},
}

var orderInfoCmd = &cobra.Command{
	Use:     "info [id]",
	Short:   "Get more info about an order by ID",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		order, err := Client.Order.Get(cmd.Context(), args[0])
		if err != nil {
			fmt.Println("Could not get order info")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			fmt.Println("Could not get products")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.Section("Items in order:", func() {
			helpers.PrintOrderItems(order.Data.Items, products.Data)
		})

		helpers.Section("Address:", func() {
			helpers.PrintOrderAddress(order.Data.Shipping)
		})

		helpers.Section("Amount:", func() {
			helpers.PrintOrderAmount(order.Data.Amount)
		})

		helpers.Section("Tracking:", func() {
			helpers.PrintOrderTracking(order.Data.Tracking)
		})
	},
}

func init() {
	orderCmd.AddCommand(listOrdersCmd)
	orderCmd.AddCommand(orderInfoCmd)
	rootCmd.AddCommand(orderCmd)
}
