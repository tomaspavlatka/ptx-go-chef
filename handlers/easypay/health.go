package easypay

import (
	"encoding/json"

	"github.com/tomaspavlatka/ptx-go-chef/internal/easypay"
)

type Health struct {
}

func GetHealth() (*Health, error) {
  resp, err := easypay.Get("health", 200)
	if err != nil {
		return nil, err
	}

	var health Health
	if err := json.Unmarshal(resp, &health); err != nil {
		return nil, err
	}

	return &health, nil

}
