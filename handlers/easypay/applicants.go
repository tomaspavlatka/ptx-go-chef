package easypay

import (
	"encoding/json"
	"time"

	"github.com/tomaspavlatka/ptx-go-chef/internal/easypay"
)

type Applicant struct {
	Id          string
	DateOfBirth time.Time
	Email       string
	FirstName   string
	LastName    string
	Status      string
	Phone       string
	Version     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Applicants struct {
	Records  []Applicant
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

