package decorators

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/lead"
)

func ToCompanies(companies []lead.Company) {
	if len(companies) == 0 {
		return
	}

	for _, company := range companies {
		fmt.Println("- ID         :", company.Id)
		fmt.Println("- Code       :", company.Id)
		fmt.Println("- Name       :", company.Name)
		fmt.Println("- Easybill ID:", company.EasybilCustomerId)
	}
}

type Attribute struct {
	S string `json:"S"`
}

type Output struct {
	CompanyID          Attribute `json:"company_id"`
	EasybillCustomerID Attribute `json:"easybill_customer_id"`
}

func ToDynamoDb(companies []lead.Company) {

	for _, company := range companies {
		data := Output{
			CompanyID:          Attribute{S: company.Code},
      EasybillCustomerID: Attribute{S: strconv.FormatUint(uint64(company.EasybilCustomerId), 10)},
    }

		// Convert struct to JSON
		jsonData, err := json.MarshalIndent(data, "", "  ") // Pretty-print JSON
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		// Print the JSON output
		fmt.Println(string(jsonData))

	}

}
