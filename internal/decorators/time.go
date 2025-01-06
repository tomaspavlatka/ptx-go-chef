package decorators

import (
	"fmt"
	"time"
)

func ToDateWithAge(t *time.Time) string {
	if t == nil {
		return "---"
	}

	return toDetails(t, "02.01. 15:04")
}

func ToDateWithAgeDetailed(t *time.Time) string {
	if t == nil {
		return "---"
	}

	return toDetails(t, "02.01. 15:04:05")
}

func toDetails(t *time.Time, format string) string {
	now := time.Now()
	sub := now.Sub(*t)

	days := int(sub.Hours() / 24)
	hours := int(int(sub.Hours()) - days*24)

	return t.Format(format) + fmt.Sprintf(" ~ %d days %d hours", days, hours)
}

func ToDuration(months int) string {
	years := int(months / 12)
	return fmt.Sprintf("%d months, %d years", months, years)
}
