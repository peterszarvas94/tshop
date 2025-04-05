package helpers

import "fmt"

func PrettyPrice(price int64) string {
	return fmt.Sprintf("$%.2f", float64(price)/100)
}
