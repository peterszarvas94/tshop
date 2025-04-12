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
			helpers.HandleError("Error getting the subscriptions", err, 1)
		}

		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			helpers.HandleError("Error getting the products", err, 1)
		}

		helpers.PrintSubs(products.Data, subs.Data)
	},
}

var createSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create subscription.",
	Example: "--type weekly --interval 3\n--type fixed",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if SubType != string(terminal.SubscriptionScheduleTypeFixed) && SubType != string(terminal.SubscriptionScheduleTypeWeekly) {
			fmt.Printf("Schedule type can only be \"%s\" or \"%s\"\n", terminal.SubscriptionScheduleTypeFixed, terminal.SubscriptionScheduleTypeWeekly)
			os.Exit(1)
		}

		if SubType == string(terminal.SubscriptionScheduleTypeWeekly) && Sub.Schedule.Interval < 1 {
			fmt.Printf("If scheduel type is \"%s\", interval must be set to a value greater than 0\n", terminal.SubscriptionScheduleTypeWeekly)
			os.Exit(1)
		}

		res, err := Client.Subscription.New(cmd.Context(), terminal.SubscriptionNewParams{
			Subscription: terminal.SubscriptionParam{
				AddressID:        terminal.F(Sub.AddressID),
				CardID:           terminal.F(Sub.CardID),
				ProductVariantID: terminal.F(Sub.ProductVariantID),
				Quantity:         terminal.Int(Sub.Quantity),
				Schedule: terminal.Raw[terminal.SubscriptionScheduleUnionParam](terminal.SubscriptionScheduleParam{
					Type:     terminal.F(terminal.SubscriptionScheduleType(SubType)),
					Interval: terminal.Int(Sub.Schedule.Interval),
				}),
			},
		})
		if err != nil {
			helpers.HandleError("Error creating the subscription", err, 1)
		}

		helpers.Section("Subscription created", func() {
			Sub.ID = string(res.Data)

			products, err := Client.Product.List(cmd.Context())
			if err != nil {
				helpers.HandleError("Error getting the products", err, 1)
			}

			subs, err := Client.Subscription.List(cmd.Context())
			if err != nil {
				helpers.HandleError("Error getting the subscription", err, 1)
			}

			var found *terminal.Subscription
			for _, sub := range subs.Data {
				if sub.ProductVariantID == Sub.ProductVariantID {
					found = &sub
				}
			}

			if found == nil {
				fmt.Println("Could not find subscription")
				os.Exit(1)
			}

			Sub.Schedule.Type = terminal.SubscriptionScheduleType(SubType)
			Sub.ID = found.ID
			Sub.Next = found.Next
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
			helpers.HandleError("Error getting the subscription", err, 1)
		}

		helpers.Section("Subscription details:", func() {
			products, err := Client.Product.List(cmd.Context())
			if err != nil {
				helpers.HandleError("Error getting the products", err, 1)
			}

			helpers.PrintSubs(products.Data, []terminal.Subscription{sub.Data})
		})

		helpers.Section("Subscription address:", func() {
			address, err := Client.Address.Get(cmd.Context(), sub.Data.AddressID)
			if err != nil {
				helpers.HandleError("Error getting the address", err, 1)
			}

			helpers.PrintAddresses([]terminal.Address{address.Data})
		})

		helpers.Section("Subscription card:", func() {
			card, err := Client.Card.Get(cmd.Context(), sub.Data.CardID)
			if err != nil {
				helpers.HandleError("Error getting the card", err, 1)
			}

			helpers.PrintCards([]terminal.Card{card.Data})
		})
	},
}

var cancelSubCmd = &cobra.Command{
	Use:     "delete [id]",
	Short:   "Cancel subscription",
	Aliases: []string{"cancel", "x"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Subscription.Delete(cmd.Context(), args[0])
		if err != nil {
			helpers.HandleError("Error deleting the subscription", err, 1)
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
	createSubCmd.Flags().StringVarP(&SubType, "type", "t", "", fmt.Sprintf("Type (%s or %s)", terminal.SubscriptionScheduleTypeFixed, terminal.SubscriptionScheduleTypeWeekly))
	createSubCmd.Flags().Int64VarP(&Sub.Schedule.Interval, "interval", "i", 0, "Interval")

	createSubCmd.MarkFlagRequired("address")
	createSubCmd.MarkFlagRequired("card")
	createSubCmd.MarkFlagRequired("variant")
	createSubCmd.MarkFlagRequired("quantity")
	createSubCmd.MarkFlagRequired("type")
	subscriptionCmd.AddCommand(createSubCmd)
	subscriptionCmd.AddCommand(subInfoCmd)
	subscriptionCmd.AddCommand(cancelSubCmd)
	rootCmd.AddCommand(subscriptionCmd)
}
