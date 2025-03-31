package decorators

import (
	"fmt"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/lead"
)

func ToCompanies(companies []lead.Company) {
	if len(companies) == 0 {
		return
	}

  for _, company := range companies {
    fmt.Println("- ", company.Id);
  }
}
