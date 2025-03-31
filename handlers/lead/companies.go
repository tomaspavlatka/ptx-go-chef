package lead

import (
	"encoding/json"
	"os"

	"github.com/tomaspavlatka/ptx-go-chef/internal/lead"
)

type Company struct {
	Id string
}

type Billable struct {
	PeriodStart string        `json:"periodStart"`
	PeriodEnd   string        `json:"periodEnd"`
	Orgs        []OrgBillable `json:"orgs"`
}

type OrgBillable struct {
	Id      string    `json:"id"`
	Auth0Id string    `json:"auth0Id"`
	Items   []Payable `json:"items"`
}

type Payable struct {
	Id        string  `json:"id"`
	Price     float32 `json:"price"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"createdAt"`
}

func ConvertCompanies(year, month string) ([]Company, error) {
	billable, err := retrieveData(year, month)
	if err != nil {
		return nil, err
	}

	companies := make([]Company, 0, len(billable.Orgs))
	for _, org := range billable.Orgs {
		company := Company{
			Id: org.Auth0Id,
		}

    companies = append(companies, company);
	}

	return companies, nil
}

func retrieveData(year, month string) (*Billable, error) {
	partner := os.Getenv("LEAD_ENGINE_PARTNER_ID")
	url := "billings/" + partner + "/leads?year=" + year + "&month=" + month

	resp, err := lead.Get(url, 200)
	if err != nil {
		return nil, err
	}

	var billable Billable
	if err := json.Unmarshal(resp, &billable); err != nil {
		return nil, err
	}

	return &billable, nil
}
