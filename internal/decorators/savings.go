package decorators

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/zoc"
)

func ToSavings(savings []zoc.Saving) {
	if len(savings) == 0 {
		return
	}

	legend := []string{"id", "investment"}
	for i := 1; i < len(savings[0].Savings); i++ {
		legend = append(legend, strconv.Itoa(i))
	}

	fmt.Println(strings.Join(legend, ","))

	for _, saving := range savings {
		row := []string{strconv.Itoa(saving.Id), strconv.FormatFloat(saving.Investment, 'f', 1, 64)}

    for _, p := range saving.Savings {
      row = append(row, strconv.FormatFloat(p, 'f', 1, 64))
    }

		fmt.Println(strings.Join(row, ","))
	}
}
