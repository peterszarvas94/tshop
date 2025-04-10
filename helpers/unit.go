package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

// since they forgot that EU use metric, this function finds substrings like "12oz" and puts "(340g)" after it
func convertToGramms(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if strings.HasSuffix(word, "oz") {
			numStr := strings.TrimSuffix(word, "oz")
			numStr = strings.TrimRightFunc(numStr, func(r rune) bool {
				return r < '0' || r > '9'
			})
			oz, err := strconv.ParseFloat(numStr, 64)
			if err == nil {
				g := oz * 28.3495
				words[i] = fmt.Sprintf("%soz (%.0fg)", numStr, g)
			}
		}
	}
	return strings.Join(words, " ")
}
