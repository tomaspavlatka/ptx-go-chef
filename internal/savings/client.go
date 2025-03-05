package savings

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type Roof struct {
	Lat            float32 `json:"lat"`
	Lng            float32 `json:"lon"`
	Tilt           uint    `json:"tilt"`
	Losses         uint    `json:"losses"`
	Orientation    string  `json:"orientation"`
	SystemCapacity uint    `json:"system_capacity"`
}

type Consumer struct {
	Profile     string `json:"load_profile"`
	Consumption uint   `json:"energy_consumption"`
}

type Simulation struct {
	Roofs                 []Roof     `json:"roofs"`
	Consumers             []Consumer `json:"consumers"`
	ConsumptionCorrection uint       `json:"factor_own_consumption_correction"`
	StorageCapacity       float32    `json:"storage_capacity"`
	StorageMaxLoad        float32    `json:"storage_max_load_power"`
}

type Price struct {
	Default bool    `json:"default"`
	Amount  uint    `json:"kwh_amount"`
	Price   float32 `json:"kwh_price"`
}

type Factors struct {
	IncreaseElectricityUsage  float32 `json:"increase_electricity_usage"`
	InflationRate             float32 `json:"inflation_rate"`
	ElectricityInflationRate  float32 `json:"electricity_inflation_rate"`
	DegradationModulesPerYear float32 `json:"degradation_modules_per_years"`
}

type ElectricityContract struct {
	FixPrice        uint    `json:"fix_price"`
	SaleToGridPrice float32 `json:"sale_to_grid_price"`
	EegChargePrice  float32 `json:"eeg_charge_price"`
	TaxRate         float32 `json:"tax_rate"`
	PostEegPayment  float32 `json:"post_eeg_payment"`
	KwhPriceRanges  []Price `json:"kwh_price_ranges"`
}

type Economic struct {
	ElectricityContract ElectricityContract `json:"current_electricity_contract"`
	StoragePrice        float32             `json:"storage_kwh_price"`
	Factors             Factors             `json:"factors"`
	Investment          float32             `json:"investment_amount"`
}

type Request struct {
	Simulation Simulation `json:"simulation"`
	Economic   Economic   `json:"economic"`
}

func url() string {
	return os.Getenv("CALCULATOR_BASE_URL")
}

func Get(request Request, statusCode int) ([]byte, error) {
	client := &http.Client{}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

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
