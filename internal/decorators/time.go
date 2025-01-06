package decorators

import (
	"fmt"
	"time"
)

func ToDateWithAge(t time.Time) string {
	now := time.Now()
	sub := now.Sub(t)

	days := int(sub.Hours() / 24)
	hours := int(int(sub.Hours()) - days*24)

	return t.Format("02.01. 15:04") + fmt.Sprintf(" ~ %d days,%d hours", days, hours)
}

func ToDuration(months int) string {
	years := int(months / 12)
	return fmt.Sprintf("%d months, %d years", months, years)
}
