package lead

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
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

type Relation struct {
	CompanyId          string `csv:"company_id"`
	EasybillCustomerId string `csv:"easybill_customer_id"`
}

func GetMissingRelations(year, month string) ([]Company, error) {
	billable, err := retrieveData(year, month)
	if err != nil {
		return nil, err
	}

	relations, err := retrieveRelations()
	if err != nil {
		return nil, err
	}

	data := make(map[string]bool)

	for _, relation := range relations {
		if len(relation.EasybillCustomerId) > 0 {
			data[relation.CompanyId] = true
		}
	}

	missing := make([]Company, 0, len(billable.Orgs))
	for _, org := range billable.Orgs {

		_, ok := data[org.Auth0Id]
		if !ok {
			missing = append(missing, Company{
				Id: org.Auth0Id,
			})
		}
	}

	return missing, nil
}

func retrieveRelations() ([]*Relation, error) {
	in, err := os.Open("data/relations-tbl.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	relations := []*Relation{}

	if err := gocsv.UnmarshalFile(in, &relations); err != nil {
		panic(err)
	}

	return relations, nil
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
