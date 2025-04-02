package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

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
			panic(err.Error())
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "ID\tName\tVariant\tPrice")
		fmt.Fprintln(w, "--\t----\t-------\t-----")
		for _, product := range products.Data {
			for _, variant := range product.Variants {
				price := fmt.Sprintf("$%.2f", float64(variant.Price)/100)
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", variant.ID, product.Name, variant.Name, price)
			}
		}
		w.Flush()
	},
}

var getProductCmd = &cobra.Command{
	Use:     "info [name]",
	Short:   "Get description of a product by name",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		products, err := Client.Product.List(cmd.Context())
		if err != nil {
			panic(err.Error())
		}

		var product *terminal.Product
		for _, p := range products.Data {
			if p.Name == args[0] {
				product = &p
			}
		}

		fmt.Println(product.Description)
	},
}
