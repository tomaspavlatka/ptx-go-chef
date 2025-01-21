package easypay

import (
	"encoding/json"
	"time"

	"github.com/tomaspavlatka/ptx-go-chef/internal/easypay"
)

type Applicant struct {
	Id          string
	DateOfBirth *time.Time
	Email       string
	FirstName   string
	LastName    string
	Status      string
	Phone       string
	Version     int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type Applicants struct {
	Records  []Applicant
	Metadata Metadata
}

type ApplicantAudit struct {
	Id          string
	ApplicantId string
	AuditType   string
	CompanyId   string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Status      string
	DateOfBirth *time.Time
	Gate        string
	Seat        string
	Txid        string
	CreatedBy   string
	CreatedAt   *time.Time
}

type ApplicantsAudit struct {
	Records  []ApplicantAudit
	Metadata Metadata
}

func GetApplicants() (*Applicants, error) {
	resp, err := easypay.Get("applicants", 200)
	if err != nil {
		return nil, err
	}

	var applicants Applicants
	if err := json.Unmarshal(resp, &applicants); err != nil {
		return nil, err
	}

	return &applicants, nil
}

func GetApplicant(applicantId string) (*Applicant, error) {
	resp, err := easypay.Get("applicants/"+applicantId, 200)
	if err != nil {
		return nil, err
	}

	var applicant Applicant
	if err := json.Unmarshal(resp, &applicant); err != nil {
		return nil, err
	}

	return &applicant, nil
}

func GetApplicantAudits(applicantId string) (*ApplicantsAudit, error) {
	resp, err := easypay.Get("audits/applicants?q=eq(applicantId,"+applicantId+")sort(createdAt)limit(100,0)", 200)
	if err != nil {
		return nil, err
	}

	var audits ApplicantsAudit
	if err := json.Unmarshal(resp, &audits); err != nil {
		return nil, err
	}

	return &audits, nil
}
