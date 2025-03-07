package zoc

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"

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

type Inputs struct {
	Data []EnergyData `json:"data"`
}

type Metric struct {
	Investment float64 `json:"investment_amount"`
	Bought     float64 `json:"cost_bought_from_grid"`
	Income     float64 `json:"income_sold_to_grid"`
	EegCharge  float64 `json:"cost_eeg_charging_fee"`
	Value      int     `json:"value"`
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
	Id         int
	Investment float64
	Savings    []float64
}

func GetSavings(raw string) ([]Saving, error) {
	inputs, err := getInputs(raw)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	results := make(chan Saving)

	for _, input := range inputs.Data {
		wg.Add(1)
		go processInput(input, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var savings []Saving

	for result := range results {
		savings = append(savings, result)
	}

	return savings, nil
}

func processInput(input EnergyData, results chan<- Saving, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := getSimulation(input)
	if err != nil {
		return
	}

	var simulation Simulation
	if err := json.Unmarshal(resp, &simulation); err != nil {
		return
	}

	var savings []float64
	var planned = simulation.Planned.Months.Metrics
	for _, month := range planned {
		peer, _ := getPeer(month, simulation.Origin.Months.Metrics)

		noPv := peer.Bought - peer.Income     // how much we would pay without PV system
		withPv := month.Bought - month.Income // how much we would pay with PV system

		savings = append(savings, noPv - withPv)
	}

	results <- Saving{
		Id:         input.ID,
		Investment: input.SubTotal,
		Savings:    savings,
	}
}

func getPeer(metric Metric, peers []Metric) (*Metric, error) {
	for _, peer := range peers {
		if peer.Value == metric.Value {
			return &peer, nil
		}
	}

	return nil, errors.New("No peer has been found")
}

func getInputs(raw string) (*Inputs, error) {
	var inputs Inputs
	err := json.Unmarshal([]byte(raw), &inputs)
	if err != nil {
		return nil, err
	}

	return &inputs, nil
}

func getSimulation(data EnergyData) ([]byte, error) {
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
					Lat:            lat,
					Lng:            lng,
					Tilt:           uint(tilt),
					Losses:         0,
					Orientation:    roofs[3],
					SystemCapacity: data.SystemCapacity,
				},
			},
			Consumers: []savings.Consumer{
				{
					Profile:     "H1",
					Consumption: uint(data.EnergyConsumption),
				},
			},
			ConsumptionCorrection: uint(data.FactorOwnConsumptionCorrection),
			StorageCapacity:       data.StorageCapacity,
			StorageMaxLoad:        data.StorageMaxLoadPower,
		},
		Economic: savings.Economic{
			ElectricityContract: savings.ElectricityContract{
				FixPrice:        uint(data.ElectricityFixGridPrice),
				SaleToGridPrice: data.PriceSaleToGrid,
				EegChargePrice:  data.PriceEEGAppointment,
				TaxRate:         0,
				PostEegPayment:  data.PricePostEEGPayment,
				KwhPriceRanges: []savings.Price{
					{
						Default: true,
						Amount:  0,
						Price:   data.ElectricityKwhPrice,
					},
				},
			},
			StoragePrice: data.PriceRebuyStorageByKwh,
			Factors: savings.Factors{
				IncreaseElectricityUsage:  data.FactorIncreaseTotalConsumptionPerAnnum,
				InflationRate:             data.FactorInflationRate,
				ElectricityInflationRate:  data.FactorInflationElectricityRate,
				DegradationModulesPerYear: data.FactorModuleDegradation,
			},
			Investment: data.SubTotal,
		},
	}

	return savings.Get(request, 200)
}
