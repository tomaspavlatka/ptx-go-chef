package lead

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/tomaspavlatka/ptx-go-chef/internal/easybill"
	"github.com/tomaspavlatka/ptx-go-chef/internal/lead"
)

type Company struct {
	Id                string
	Name              string
	Code              string
	EasybilCustomerId uint32
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

type CompanyProfile struct {
	Id   string `csv:"id"`
	Code string `csv:"code"`
	Name string `csv:"name"`
}

type Customer struct {
	Id   uint32 `json:"id"`
	Code string `json:"number"`
}

type Customers struct {
	Customers []Customer `json:"items"`
}

func CompleteCompanies(file string) ([]Company, error) {
	profiles, err := retrieveProfiles(file)
	if err != nil {
		return nil, err
	}

	companies := make([]Company, 0, len(profiles))

	for _, profile := range profiles {
		customer, err := findByCode(profile.Code)
		if err != nil {
			customer, err := createCustomer(profile)
			if err != nil {
				return nil, err
			}

			companies = append(companies, Company{
				Id:                profile.Id,
				Name:              profile.Name,
				Code:              profile.Code,
				EasybilCustomerId: customer.Id,
			})

		} else {
			companies = append(companies, Company{
				Id:                profile.Id,
				Name:              profile.Name,
				Code:              profile.Code,
				EasybilCustomerId: customer.Id,
			})
		}
	}

	return companies, nil
}

func createCustomer(profile *CompanyProfile) (*Customer, error) {
	fmt.Println("Create")
	data := map[string]any{
		"number":       profile.Code,
		"company_name": profile.Name,
		"last_name":    profile.Name,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := "customers"
	resp, err := easybill.Post(url, 201, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	var customer Customer
	if err := json.Unmarshal(resp, &customer); err != nil {
		return nil, err
	}

	return &customer, nil
}

func findByCode(code string) (*Customer, error) {
	fmt.Println("Find by code")
	url := "customers?number=" + code

	resp, err := easybill.Get(url, 200)
	if err != nil {
		return nil, err
	}

	var customers Customers
	if err := json.Unmarshal(resp, &customers); err != nil {
		return nil, err
	}

	if len(customers.Customers) > 0 {
		return &customers.Customers[0], nil
	}

	return nil, errors.New("Not found")
}

func retrieveProfiles(file string) ([]*CompanyProfile, error) {
	in, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	profiles := []*CompanyProfile{}

	if err := gocsv.UnmarshalFile(in, &profiles); err != nil {
		panic(err)
	}

	return profiles, nil
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
				Id:                org.Auth0Id,
				Name:              "",
				Code:              org.Auth0Id,
				EasybilCustomerId: 0,
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
  fmt.Println("lead:url", url)

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
