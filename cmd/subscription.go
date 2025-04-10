package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/helpers"
	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
)

var subscriptionCmd = &cobra.Command{
	Use:     "subscription",
	Short:   "Manage subscription",
	Aliases: []string{"b"},
}

var listSubsCmd = &cobra.Command{
	Use:     "list",
	Short:   "List subscriptions",
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		subs, err := Client.Subscription.List(cmd.Context())
		if err != nil {
			fmt.Println("Error getting the subscriptions")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			fmt.Println("Could not get products")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.PrintSubs(products.Data, subs.Data)
	},
}

var createSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create subscription",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := Client.Subscription.New(cmd.Context(), terminal.SubscriptionNewParams{
			Subscription: terminal.SubscriptionParam{
				AddressID:        terminal.F(Sub.AddressID),
				CardID:           terminal.F(Sub.CardID),
				ProductVariantID: terminal.F(Sub.ProductVariantID),
				Quantity:         terminal.Int(Sub.Quantity),
				// TODO: schedule, next?
			},
		})
		if err != nil {
			fmt.Println("Error creating the subscriptions")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.Section("Subscription created", func() {
			Sub.ID = string(res.Data)

			products, err := Client.Product.List(cmd.Context())
			if err != nil {
				fmt.Println("Error getting the products")
				fmt.Println(err.Error())
				os.Exit(1)
			}
			helpers.PrintSubs(products.Data, []terminal.Subscription{*Sub})
		})
	},
}

var subInfoCmd = &cobra.Command{
	Use:     "info [id]",
	Short:   "More info about a subscription",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sub, err := Client.Subscription.Get(cmd.Context(), args[0])
		if err != nil {
			fmt.Println("Error getting the subscription")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.Section("Subscription details:", func() {
			products, err := Client.Product.List(cmd.Context())
			if err != nil {
				fmt.Println("Error getting the products")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			helpers.PrintSubs(products.Data, []terminal.Subscription{sub.Data})
		})

		helpers.Section("Subscription address:", func() {
			address, err := Client.Address.Get(cmd.Context(), sub.Data.AddressID)
			if err != nil {
				fmt.Println("Error getting the address info")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			helpers.PrintAddresses([]terminal.Address{address.Data})
		})

		helpers.Section("Subscription card:", func() {
			card, err := Client.Card.Get(cmd.Context(), sub.Data.CardID)
			if err != nil {
				fmt.Println("Error getting the card info")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			helpers.PrintCards([]terminal.Card{card.Data})
		})
	},
}

var cancelSubCmd = &cobra.Command{
	Use:     "cancel [id]",
	Short:   "Cancel a subscription",
	Aliases: []string{"x"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Subscription.Delete(cmd.Context(), args[0])
		if err != nil {
			fmt.Println("Error canceling the subscription")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Subscription cancelled")
	},
}

func init() {
	subscriptionCmd.AddCommand(listSubsCmd)
	createSubCmd.Flags().StringVarP(&Sub.AddressID, "address", "a", "", "Address")
	createSubCmd.Flags().StringVarP(&Sub.CardID, "card", "c", "", "Card")
	createSubCmd.Flags().StringVarP(&Sub.ProductVariantID, "variant", "v", "", "Variant")
	createSubCmd.Flags().Int64VarP(&Sub.Quantity, "quantity", "q", 0, "Quantity")
	createSubCmd.MarkFlagRequired("address")
	createSubCmd.MarkFlagRequired("card")
	createSubCmd.MarkFlagRequired("variant")
	createSubCmd.MarkFlagRequired("quantity")
	subscriptionCmd.AddCommand(createSubCmd)
	subscriptionCmd.AddCommand(subInfoCmd)
	subscriptionCmd.AddCommand(cancelSubCmd)
	rootCmd.AddCommand(subscriptionCmd)
}
