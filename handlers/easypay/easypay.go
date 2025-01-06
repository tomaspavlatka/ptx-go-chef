package easypay

type Metadata struct {
	Offset    int
	Limit     int
	Count     int
	Total     int
	HardLimit int
}

type Money struct {
	CentAmount int    `json:"centAmount"`
	Currency   string `json:"currency"`
}

func ToEur(m Money) float64 {
	return float64(m.CentAmount) / 100.0
}
