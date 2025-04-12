package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/helpers"
	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
)

var productCmd = &cobra.Command{
	Use:     "product",
	Short:   "Manage products",
	Aliases: []string{"p"},
	Args:    cobra.ExactArgs(0),
}

var listProductsCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all products",
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			helpers.HandleError("Error getting the products", err, 1)
		}

		helpers.PrintProducts(products.Data)
	},
}

var describeProductCmd = &cobra.Command{
	Use:     "info [id]",
	Short:   "Get description of a product",
	Aliases: []string{"describe", "i"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			helpers.HandleError("Error getting the product", err, 1)
		}

		var found *terminal.Product
		for _, product := range products.Data {
			if product.ID == args[0] {
				found = &product
			}
		}

		if found == nil {
			fmt.Printf("Could not find product")
			os.Exit(1)
		}

		helpers.Section("Description:", func() {
			fmt.Println(found.Description)
		})

		helpers.Section("Details:", func() {
			helpers.PrintProducts(products.Data)
		})
	},
}

func init() {
	productCmd.AddCommand(listProductsCmd)
	productCmd.AddCommand(describeProductCmd)
	rootCmd.AddCommand(productCmd)
}
