package easybill

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func url(part string) string {
	return os.Getenv("EASYBILL_BASE_URL") + "/" + part
}

func Post(part string, statusCode int, body io.Reader) ([]byte, error) {
	client := &http.Client{}

  u := url(part);
  fmt.Println("url", u)

	req, err := http.NewRequest(http.MethodPost, url(part), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-SOURCE", "ptx-go-chef")
  req.Header.Add("Content-Type", "application/json")
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
  return "Bearer " + os.Getenv("EASYBILL_AUTH_TOKEN")
}
