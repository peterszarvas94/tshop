package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all products",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		product, err := Client.Product.List(context.Background())
		if err != nil {
			panic(err.Error())
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintln(w, "Name\tVariant\tPrice")
		fmt.Fprintln(w, "----\t-------\t-----")
		for _, p := range product.Data {
			for _, v := range p.Variants {
				price := fmt.Sprintf("$%.2f", float64(v.Price)/100)
				fmt.Fprintf(w, "%s\t%s\t%s\n", p.Name, v.Name, price)
			}
		}
		w.Flush()
	},
}
