package decorators

import (
	"fmt"
	"math"
	"strings"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

func ToMoney(m easypay.Money, abs bool) string {
	return ToMoneyFromCentAmount(m.CentAmount, m.Currency, abs)
}

func ToMoneyFromCentAmount(centAmount int, currency string, abs bool) string {
	value := float64(centAmount) / 100.0
	if abs {
		value = math.Abs(value)
	}

	money := fmt.Sprintf("%.2f", value)

	parts := strings.Split(money, ".")

	intPart := parts[0]
	formattedIntPart := addUnderscores(intPart)

	if len(parts) > 1 {
		return formattedIntPart + "." + parts[1] + " " + currency
	}

	return formattedIntPart + " " + currency
}

func addUnderscores(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var sb strings.Builder
	mod := n % 3
	if mod > 0 {
		sb.WriteString(s[:mod])
		sb.WriteString("_")
	}

	for i := mod; i < n; i += 3 {
		sb.WriteString(s[i : i+3])
		if i+3 < n {
			sb.WriteString("_")
		}
	}

	return sb.String()
}
