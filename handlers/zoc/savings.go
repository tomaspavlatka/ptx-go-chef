package zoc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
	"github.com/tomaspavlatka/ptx-go-chef/internal/decorators"
	"github.com/tomaspavlatka/ptx-go-chef/internal/savings"
)

type EnergyData struct {
	ID                                     int     `json:"id"`
	ElectricityFixGridPrice                float64 `json:"electricity_fix_grid_price"`
	PriceSaleToGrid                        float64 `json:"price_sale_to_grid"`
	PriceEEGAppointment                    float64 `json:"price_eeg_appointment"`
	PricePostEEGPayment                    float64 `json:"price_post_eeg_payment"`
	ElectricityKwhPrice                    float64 `json:"electricity_kwh_price"`
	PriceRebuyStorageByKwh                 float64 `json:"price_rebuy_storage_by_kwh"`
	FactorIncreaseTotalConsumptionPerAnnum float64 `json:"factor_increase_total_consumption_per_annum"`
	FactorInflationRate                    float64 `json:"factor_inflation_rate"`
	FactorInflationElectricityRate         float64 `json:"factor_inflation_electricity_rate"`
	FactorModuleDegradation                float64 `json:"factor_module_degradation"`
	FactorOwnConsumptionCorrection         float64 `json:"factor_own_consumption_correction"`
	StorageCapacity                        float64 `json:"storage_capacity"`
	StorageMaxLoadPower                    float64 `json:"storage_max_load_power"`
	Roof                                   string  `json:"roof"`
	SystemCapacity                         float64 `json:"system_capacity"`
	EnergyConsumption                      float64 `json:"energy_consumption"`
	SubTotal                               float64 `json:"sub_total"`
	State                                  string  `json:"state"`
}

type Metric struct {
	Investment float32 `json:"investment_amount"`
	Bought     float32 `json:"cost_bought_from_grid"`
	Income     float32 `json:"income_sold_to_grid"`
	EegCharge  float32 `json:"cost_eeg_charging_fee"`
	Value      uint    `json:"value"`
}

type Metrics struct {
	Metrics []Metric `json:"metrics"`
}

type Group struct {
	Months  Metrics `json:"months"`
	Seasons Metrics `json:"seasons"`
	Years   Metrics `json:"years"`
}

type Simulation struct {
	Origin  Group `json:"origin"`
	Planned Group `json:"planned"`
}

type Saving struct {
	Savings easypay.Money
	Origin  easypay.Money
	Planned easypay.Money
	Value   uint
}

func GetYearlySavings(raw string) (string, error) {
	resp, err := getSimulation(raw)
	if err != nil {
		return "", err
	}

	var simulation Simulation
	if err := json.Unmarshal(resp, &simulation); err != nil {
		return "", err
	}

	var planned = simulation.Planned.Years.Metrics

	var savings []Saving

	for _, month := range planned {
		peer, err := getPeer(month, simulation.Origin.Years.Metrics)
		if err != nil {
			return "", err
		}

		var plannedCost = getCost(month)
		var originCost = getCost(*peer)

		savings = append(savings, Saving{
			Savings: easypay.Money{
				CentAmount: int((plannedCost - originCost) * 100),
				Currency:   "EUR",
			},
			Origin: easypay.Money{
				CentAmount: int(originCost * 100),
				Currency:   "EUR",
			},
			Planned: easypay.Money{
				CentAmount: int(plannedCost * 100),
				Currency:   "EUR",
			},
			Value: month.Value,
		})
	}

	toSaving(savings)

	return "", nil
}

func GetMonthlySavings(raw string) (string, error) {
	resp, err := getSimulation(raw)
	if err != nil {
		return "", err
	}

	var simulation Simulation
	if err := json.Unmarshal(resp, &simulation); err != nil {
		return "", err
	}

	var planned = simulation.Planned.Months.Metrics

	var savings []Saving

	for _, month := range planned {
		peer, err := getPeer(month, simulation.Origin.Months.Metrics)
		if err != nil {
			return "", err
		}

		var plannedCost = getCost(month)
		var originCost = getCost(*peer)

		savings = append(savings, Saving{
			Savings: easypay.Money{
				CentAmount: int((plannedCost - originCost) * 100),
				Currency:   "EUR",
			},
			Origin: easypay.Money{
				CentAmount: int(originCost * 100),
				Currency:   "EUR",
			},
			Planned: easypay.Money{
				CentAmount: int(plannedCost * 100),
				Currency:   "EUR",
			},
			Value: month.Value,
		})
	}

	toSaving(savings)

	return "", nil
}

func GetSeasonalSavings(raw string) (string, error) {
	resp, err := getSimulation(raw)
	if err != nil {
		return "", err
	}

	var simulation Simulation
	if err := json.Unmarshal(resp, &simulation); err != nil {
		return "", err
	}

	var planned = simulation.Planned.Seasons.Metrics

	var savings []Saving

	for _, month := range planned {
		peer, err := getPeer(month, simulation.Origin.Seasons.Metrics)
		if err != nil {
			return "", err
		}

		var plannedCost = getCost(month)
		var originCost = getCost(*peer)

		savings = append(savings, Saving{
			Savings: easypay.Money{
				CentAmount: int((plannedCost - originCost) * 100),
				Currency:   "EUR",
			},
			Origin: easypay.Money{
				CentAmount: int(originCost * 100),
				Currency:   "EUR",
			},
			Planned: easypay.Money{
				CentAmount: int(plannedCost * 100),
				Currency:   "EUR",
			},
			Value: month.Value,
		})
	}

	toSaving(savings)

	return "", nil
}

func toSaving(savings []Saving) {
	for _, saving := range savings {

		fmt.Println("Month    : ", saving.Value)
		fmt.Println("Saving   : ", decorators.ToMoney(saving.Savings, true))
		fmt.Println("Planned  : ", decorators.ToMoney(saving.Planned, true))
		fmt.Println("Origin   : ", decorators.ToMoney(saving.Origin, true))
		fmt.Println("========= ")
	}
}

func getPeer(metric Metric, metrics []Metric) (*Metric, error) {
	for _, peer := range metrics {
		if peer.Value == metric.Value {
			return &peer, nil
		}
	}

	return nil, errors.New("peer not found")
}

func getCost(metric Metric) float32 {
	return metric.Investment + metric.Bought + metric.EegCharge - metric.Income;
}

func getSimulation(raw string) ([]byte, error) {
	var data EnergyData
	err := json.Unmarshal([]byte(raw), &data)
	if err != nil {
		return nil, err
	}

	roofs := strings.Split(data.Roof, "|")
	lat, err := strconv.ParseFloat(roofs[0], 32)
	if err != nil {
		return nil, err
	}

	lng, err := strconv.ParseFloat(roofs[1], 32)
	if err != nil {
		return nil, err
	}

	tilt, err := strconv.Atoi(roofs[2])
	if err != nil {
		return nil, err
	}

	request := savings.Request{
		Simulation: savings.Simulation{
			Roofs: []savings.Roof{
				{
					Lat:            float32(lat),
					Lng:            float32(lng),
					Tilt:           uint(tilt),
					Losses:         0,
					Orientation:    roofs[3],
					SystemCapacity: uint(data.SystemCapacity),
				},
			},
			Consumers: []savings.Consumer{
				{
					Profile:     "H1",
					Consumption: uint(data.EnergyConsumption),
				},
			},
			ConsumptionCorrection: uint(data.FactorOwnConsumptionCorrection),
			StorageCapacity:       float32(data.StorageCapacity),
			StorageMaxLoad:        float32(data.StorageMaxLoadPower),
		},
		Economic: savings.Economic{
			ElectricityContract: savings.ElectricityContract{
				FixPrice:        uint(data.ElectricityFixGridPrice),
				SaleToGridPrice: float32(data.PriceSaleToGrid),
				EegChargePrice:  float32(data.PriceEEGAppointment),
				TaxRate:         0,
				PostEegPayment:  float32(data.PricePostEEGPayment),
				KwhPriceRanges: []savings.Price{
					{
						Default: true,
						Amount:  0,
						Price:   float32(data.ElectricityKwhPrice),
					},
				},
			},
			StoragePrice: float32(data.PriceRebuyStorageByKwh),
			Factors: savings.Factors{
				IncreaseElectricityUsage:  float32(data.FactorIncreaseTotalConsumptionPerAnnum),
				InflationRate:             float32(data.FactorInflationRate),
				ElectricityInflationRate:  float32(data.FactorInflationElectricityRate),
				DegradationModulesPerYear: float32(data.FactorModuleDegradation),
			},
			Investment: float32(data.SubTotal),
		},
	}

	return savings.Get(request, 200)
}
