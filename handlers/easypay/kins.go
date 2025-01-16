package easypay

import (
	"encoding/json"
	"time"

	"github.com/tomaspavlatka/ptx-go-chef/internal/easypay"
)

type KinAudit struct {
	Id                   string
	KinId                string
	AuditType            string
	ResourceId           string
	Value                string
	Type                 string
	Version              int
	Gate                 string
	Seat                 string
	Txid                 string
	CreatedBy            string
	CreatedAt            *time.Time
}

type KinsAudit struct {
	Records  []KinAudit
	Metadata Metadata
}

func GetResourceKinsAudit(resourceId string) (*KinsAudit, error) {
	resp, err := easypay.Get("audits/kins?q=eq(resourceId,"+resourceId+")sort(createdAt)", 200)
	if err != nil {
		return nil, err
	}

	var audits KinsAudit
	if err := json.Unmarshal(resp, &audits); err != nil {
		return nil, err
	}

	return &audits, nil
}
