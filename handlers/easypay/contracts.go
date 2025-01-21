package easypay

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tomaspavlatka/ptx-go-chef/internal/easypay"
)

type Contract struct {
	Id                   string
	ExternalId           string
	PartnerId            string
	MonthlyInstallment   Money
	Investment           Money
	DownPayment          Money
	DurationMonths       int
	NominalInterestRate  InterestRate
	TotalCreditAmount    Money
	Status               string
	Name                 string
	Version              int
	AccessToken          string
	AccessTokenExpiresAt time.Time
	ReviewedBy           string
	ReviewedAt           *time.Time
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
}

type ContractAudit struct {
	Id                   string
	ContractId           string
	ExternalId           string
	PartnerId            string
	AuditType            string
	MonthlyInstallment   int
	CompanyId            string
	Name                 string
	Currency             string
	ApplicantId          *string
	Investment           int
	DownPayment          int
	DurationMonths       int
	Country              string
	InterestRate         *int
	Status               string
	Version              int
	AccessToken          string
	AccessTokenExpiresAt time.Time
	ReviewedBy           string
	ReviewedAt           *time.Time
	Gate                 string
	Seat                 string
	Txid                 string
	CreatedBy            string
	CreatedAt            *time.Time
}

type ContractsAudit struct {
	Records  []ContractAudit
	Metadata Metadata
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

type ContractsOpts struct {
	Limit     int8
	Offset    int8
	SortBy    string
	CompanyId string
	Status    string
}

func GetContractAudits(contractId string) (*ContractsAudit, error) {
	resp, err := easypay.Get("audits/contracts?q=eq(contractId,"+contractId+")sort(createdAt)limit(100,0)", 200)
	if err != nil {
		return nil, err
	}

	var audits ContractsAudit
	if err := json.Unmarshal(resp, &audits); err != nil {
		return nil, err
	}

	return &audits, nil
}

func GetContract(contractId string) (*Contract, error) {
	resp, err := easypay.Get("contracts/"+contractId, 200)
	if err != nil {
		return nil, err
	}

	var contract Contract
	if err := json.Unmarshal(resp, &contract); err != nil {
		return nil, err
	}

	return &contract, nil
}

func GetContracts(ops ContractsOpts) (*Contracts, error) {
	var query = ""
	query += fmt.Sprintf("limit(%d,%d)sort(%s)", ops.Limit, ops.Offset, ops.SortBy)
	if ops.CompanyId != "" {
		query += fmt.Sprintf("eq(companyId,%s)", ops.CompanyId)
	}
	if ops.Status != "" {
		query += fmt.Sprintf("eq(status,%s)", ops.Status)
	}

	resp, err := easypay.Get("contracts?q="+query, 200)
	if err != nil {
		return nil, err
	}

	var contracts Contracts
	if err := json.Unmarshal(resp, &contracts); err != nil {
		return nil, err
	}

	return &contracts, nil
}
