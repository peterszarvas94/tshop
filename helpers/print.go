package helpers

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func PrintUser(user terminal.ProfileUser) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Name\tEmail")
	fmt.Fprintln(w, "----\t-----")
	fmt.Fprintf(w, "%s\t%s\n", user.Name, user.Email)
	w.Flush()
}

func PrintProducts(products []terminal.Product) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Product ID\tVariant ID\tName\tVariant\tPrice")
	fmt.Fprintln(w, "----------\t----------\t-----\t------\t-----")
	for _, product := range products {
		for _, variant := range product.Variants {
			price := PrettyPrice(variant.Price)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", product.ID, variant.ID, product.Name, variant.Name, price)
		}
	}
	w.Flush()
}

func PrintCards(cards []terminal.Card) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "ID\tBrand\tExipration\tNumber")
	fmt.Fprintln(w, "--\t-----\t----------\t------")
	for _, card := range cards {
		expiration := fmt.Sprintf("%d/%d", card.Expiration.Month, card.Expiration.Year)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", card.ID, card.Brand, expiration, card.Last4)
	}

	w.Flush()
}

func PrintCartItems(items []terminal.CartItem, products []terminal.Product) {
	variantNameMap := map[string]string{}
	for _, p := range products {
		for _, v := range p.Variants {
			variantNameMap[v.ID] = p.Name
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Variant ID\tName\tQuantity\tSubtotal")
	fmt.Fprintln(w, "----------\t----\t--------\t--------")

	for _, item := range items {
		name := variantNameMap[item.ProductVariantID]
		total := PrettyPrice(item.Subtotal)
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", item.ProductVariantID, name, item.Quantity, total)
	}
	w.Flush()
}

func PrintAddresses(addresses []terminal.Address) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "ID\tName\tCountry\tProvince\tCity\tZip\tStreet1\tStreet2")
	fmt.Fprintln(w, "--\t----\t-------\t--------\t----\t---\t-------\t-------")
	for _, address := range addresses {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", address.ID, address.Name, address.Country, address.Province, address.City, address.Zip, address.Street1, address.Street2)
	}
	w.Flush()
}
