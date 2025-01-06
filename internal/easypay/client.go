package easypay

import "os"

type ErrorResponse struct {
	ErrorMsg  string
	Timestamp string
	ErrorCode int
}

func url(part string) string {
  return os.Getenv("EASYPAY_BASE_URL") + "/" + part
}
