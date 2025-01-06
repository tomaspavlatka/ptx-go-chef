package easypay

import (
	"encoding/json"
)

type Health struct {
}

func GetHealth() (*Health, error) {
	resp, err := get("health", 200)
	if err != nil {
		return nil, err
	}

	var health Health
	if err := json.Unmarshal(resp, &health); err != nil {
		return nil, err
	}

	return &health, nil

}
