package decorators

import (
	"fmt"
	"strings"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

func ToMoney(m easypay.Money) string {
	money := fmt.Sprintf("%.2f", float64(m.CentAmount)/100.0)

	parts := strings.Split(money, ".")

	intPart := parts[0]
	formattedIntPart := addUnderscores(intPart)

  if len(parts) > 1 { 
    return formattedIntPart + "." + parts[1] + " " + m.Currency
  }

  return formattedIntPart + " " + m.Currency
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
