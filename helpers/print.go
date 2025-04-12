package helpers

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/terminaldotshop/terminal-sdk-go"
)

func Section(text string, next func()) {
	fmt.Println(text)
	fmt.Println()
	next()
	fmt.Println()
}

// TODO: apply to headers
func FormatPrice(price int64) string {
	return fmt.Sprintf("$%.2f", float64(price)/100)
}

func PadPrice(price int64) string {
	formatted := FormatPrice(price)
	return fmt.Sprintf("%*s", 10, formatted)
}

func PrintUser(user terminal.ProfileUser) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Name\tEmail")
	fmt.Fprintln(w, "----\t-----")
	fmt.Fprintf(w, "%s\t%s\n", user.Name, user.Email)
	w.Flush()
}

func PrintProducts(products []terminal.Product) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintf(w, "Product ID\tVariant ID\tName\tVariant\t%*s\n", 10, "Price")
	fmt.Fprintf(w, "----------\t----------\t-----\t------\t%*s\n", 10, "-----")
	for _, product := range products {
		for _, variant := range product.Variants {
			price := PadPrice(variant.Price)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", product.ID, variant.ID, product.Name, convertToGramms(variant.Name), price)
		}
	}
	w.Flush()
}

func PrintCards(cards []terminal.Card) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "ID\tBrand\tExpiration\tLast 4 digits")
	fmt.Fprintln(w, "--\t-----\t----------\t-------------")
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
	fmt.Fprintf(w, "Variant ID\tName\tQuantity\t%*s\n", 10, "Subtotal")
	fmt.Fprintf(w, "----------\t----\t--------\t%*s\n", 10, "--------")

	for _, item := range items {
		name := variantNameMap[item.ProductVariantID]
		total := PadPrice(item.Subtotal)
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", item.ProductVariantID, name, item.Quantity, total)
	}
	w.Flush()
}

func PrintAddresses(addresses []terminal.Address) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "ID\tName\tCountry\tProvince\tCity\tZip\tStreet1\tStreet2\tPhone")
	fmt.Fprintln(w, "--\t----\t-------\t--------\t----\t---\t-------\t-------\t-----")
	for _, address := range addresses {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", address.ID, address.Name, address.Country, address.Province, address.City, address.Zip, address.Street1, address.Street2, address.Phone)
	}
	w.Flush()
}

func PrintOrders(orders []terminal.Order) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintf(w, "ID\tNumber\tItems\tTracking\tURL\t%*s\n", 10, "Amount")
	fmt.Fprintf(w, "--\t------\t-----\t--------\t---\t%*s\n", 10, "------")
	for _, order := range orders {
		subtotal := FormatPrice(order.Amount.Subtotal)
		fmt.Fprintf(w, "%s\t%d\t%d\t%s\t%s\t%s\n", order.ID, order.Index, len(order.Items), subtotal, order.Tracking.Number, order.Tracking.URL)
	}
	w.Flush()
}

func PrintOrderItems(items []terminal.OrderItem, products []terminal.Product) {
	variantNameMap := map[string]string{}
	for _, p := range products {
		for _, v := range p.Variants {
			variantNameMap[v.ID] = p.Name
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintf(w, "Variant ID\tName\tQuantity\t%*s\n", 10, "Amount")
	fmt.Fprintf(w, "----------\t----\t--------\t%*s\n", 10, "------")
	for _, item := range items {
		name := variantNameMap[item.ProductVariantID]
		amount := FormatPrice(item.Amount)
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", item.ProductVariantID, name, item.Quantity, amount)
	}
	w.Flush()
}

func PrintOrderAddress(address terminal.OrderShipping) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Name\tCountry\tProvince\tCity\tZip\tStreet1\tStreet2")
	fmt.Fprintln(w, "----\t-------\t--------\t----\t---\t-------\t-------")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", address.Name, address.Country, address.Province, address.City, address.Zip, address.Street1, address.Street2)
	w.Flush()
}

func PrintOrderAmount(amount terminal.OrderAmount) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	subtotal := PadPrice(amount.Subtotal)
	shipping := PadPrice(amount.Shipping)
	total := PadPrice(amount.Subtotal + amount.Shipping)

	fmt.Fprintf(w, "%s\t%s\n", "Items total", subtotal)
	fmt.Fprintf(w, "%s\t%s\n", "Shipping cost", shipping)
	fmt.Fprintf(w, "%s\t%s\n", "-------------", "----------")
	fmt.Fprintf(w, "%s\t%s\n", "Total", total)
	w.Flush()
}

func PrintOrderTracking(tracking terminal.OrderTracking) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Tracking\tURL")
	fmt.Fprintln(w, "--------\t---")
	fmt.Fprintf(w, "%s\t%s", tracking.Number, tracking.URL)
	w.Flush()
}

func PrintSubs(products []terminal.Product, subscriptions []terminal.Subscription) {
	variantNameMap := map[string]string{}
	for _, p := range products {
		for _, v := range p.Variants {
			variantNameMap[v.ID] = p.Name
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "ID\tQuantity\tVariant ID\tName\tType\tInterval\tNext")
	fmt.Fprintln(w, "--\t--------\t----------\t----\t----\t--------\t----")
	for _, sub := range subscriptions {
		name := variantNameMap[sub.ProductVariantID]
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%d\t%s\n", sub.ID, sub.Quantity, sub.ProductVariantID, name, sub.Schedule.Type, sub.Schedule.Interval, sub.Next)
	}
	w.Flush()
}

func PrintTokens(tokens []terminal.Token) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "ID\tToken\tCreated")
	fmt.Fprintln(w, "--\t-----\t-------")
	for _, token := range tokens {
		fmt.Fprintf(w, "%s\t%s\t%s\n", token.ID, token.Token, token.Created)
	}
	w.Flush()
}
