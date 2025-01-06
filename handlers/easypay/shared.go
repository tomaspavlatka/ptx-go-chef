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
