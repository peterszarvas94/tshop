package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

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
	Short:   "List all products with variant and price",
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			fmt.Println("Error getting the products")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printProducts(products.Data)
	},
}

var describeProductCmd = &cobra.Command{
	Use:     "describe [name / id]",
	Short:   "Get description of a product by name or id",
	Aliases: []string{"d"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			fmt.Println("Error getting the products")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		var found *terminal.Product
		for _, product := range products.Data {
			if product.Name == args[0] {
				found = &product
				break
			}
			if product.ID == args[0] {
				found = &product
			}
		}

		if found == nil {
			fmt.Printf("There is no product with the name or ID \"%s\"\n", args[0])
			os.Exit(1)
		}

		fmt.Println(found.Description)
	},
}

func printProducts(products []terminal.Product) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Product ID\tVariant ID\tName\tVariant\tPrice")
	fmt.Fprintln(w, "----------\t----------\t-----\t------\t-----")
	for _, product := range products {
		for _, variant := range product.Variants {
			price := helpers.PrettyPrice(variant.Price)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", product.ID, variant.ID, product.Name, variant.Name, price)
		}
	}
	w.Flush()
}
