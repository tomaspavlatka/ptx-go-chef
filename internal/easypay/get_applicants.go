package easypay

import (
	"encoding/json"
	"time"
)

type Applicant struct {
	id          string
	dateOfBirth time.Time
	email       string
	firstName   string
	lastName    string
	status      string
	phone       string
	version     string
	createdAt   time.Time
	updatedAt   time.Time
}

type Metadata struct {
	offset    int
	limit     int
	count     int
	total     int
	hardLimit int
}

type Applicants struct {
	records  []*Applicant
	metadata *Metadata
}

func GetApplicants() (*Applicants, error) {
	resp, err := get("applicants", 200)
	if err != nil {
		return nil, err
	}

	var applicants Applicants
	if err := json.Unmarshal(resp, &applicants); err != nil {
		return nil, err
	}

	return &applicants, nil
}
