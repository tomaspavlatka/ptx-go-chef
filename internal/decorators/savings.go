package decorators

import (
	"fmt"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/zoc"
)

func ToSavings(savings []zoc.Saving) {
	fmt.Println("cart_id,value,investment,saving");

	for _, saving := range savings {
		fmt.Printf("%d,%d,%.2f,%.2f\n", saving.Id, saving.Value, saving.Investment, saving.Savings);
	}
}
