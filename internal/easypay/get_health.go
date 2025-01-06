package easypay

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type Health struct {
}

func GetHealth() (*Health, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url("health") , nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-SOURCE", "ptx-go-chef")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var error ErrorResponse
		if err := json.Unmarshal(body, &error); err != nil {
			return nil, errors.New(strconv.Itoa(error.ErrorCode) + ": " + error.ErrorMsg)
		}

		return nil, err
	}

	var health Health
	if err := json.Unmarshal(body, &health); err != nil {
		return nil, err
	}

	return &health, nil

}
