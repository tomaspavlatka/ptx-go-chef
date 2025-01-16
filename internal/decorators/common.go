package decorators

import "github.com/charmbracelet/lipgloss"

var (
	headerStyle = lipgloss.NewStyle().Bold(true)
)

func translateType(t string) string {
	switch t {
	case "I":
		return "created"
	case "U":
		return "updated"
	case "D":
		return "deleted"
	default:
		return "unknown!"
	}
}

func gotChanged[T comparable](currentValue T, newValue *T) (T, bool) {
	if newValue == nil {
		return currentValue, false
	}

	changed := currentValue != *newValue

	return *newValue, changed
}
