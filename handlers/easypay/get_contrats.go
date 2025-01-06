package easypay

import (
	"encoding/json"
	"time"

	"github.com/tomaspavlatka/ptx-go-chef/internal/easypay"
)

type Contract struct {
	Id                 string
	MonthlyInstallment Money
	Investment         Money
	DownPayment        Money
	DurationMonths     int
	InterestRate       InterestRate
	Status             string
	Version            int
	ReviewedBy         string
	ReviewedAt         *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

type InterestRate struct {
	Rate     int
	Metadata InterestRateMeta
}

type InterestRateMeta struct {
	Decimal    float64
	Percentage float64
}

type Contracts struct {
	Records  []Contract
	Metadata Metadata
}

func GetContracts() (*Contracts, error) {
	resp, err := easypay.Get("contracts", 200)
	if err != nil {
		return nil, err
	}

	var contracts Contracts
	if err := json.Unmarshal(resp, &contracts); err != nil {
		return nil, err
	}

	return &contracts, nil
}
