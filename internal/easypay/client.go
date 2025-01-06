package easypay

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func url(part string) string {
	return os.Getenv("EASYPAY_BASE_URL") + "/" + part
}

func Get(part string, statusCode int) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url(part), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-SOURCE", "ptx-go-chef")
  req.Header.Add("Authorization", getBearer())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != statusCode {
		return nil, errors.New(resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func getBearer() string {
  return "Bearer " + os.Getenv("EASYPAY_BEARER_TOKEN")
}
