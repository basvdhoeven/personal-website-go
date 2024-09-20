package units

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

type UnitConverter struct {
	convRates map[string]map[string]float64
}

func NewUnitConverter() *UnitConverter {
	return &UnitConverter{
		convRates: make(map[string]map[string]float64),
	}
}

func (uc *UnitConverter) LoadConvRatesFromYaml(paths map[string]string) error {
	for quantity, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading conv rates YAML file: %v", err)
		}

		// Unmarshal YAML into a map
		var units map[string]float64
		err = yaml.Unmarshal(data, &units)
		if err != nil {
			return fmt.Errorf("error unmarshaling conv rates YAML: %v", err)
		}

		uc.convRates[quantity] = units
	}

	return nil
}

func (uc *UnitConverter) GetAllUnits(quantity string) ([]string, error) {
	quantityData, ok := uc.convRates[quantity]
	if !ok {
		return nil, errors.New("could not find quantity data")
	}

	units := make([]string, 0, len(quantityData))
	for u := range quantityData {
		units = append(units, u)
	}

	sort.Strings(units)

	return units, nil
}

func (uc *UnitConverter) Convert(quantity, inputUnit, outputUnit string, amount float64) (float64, error) {
	quantityData, ok := uc.convRates[quantity]
	if !ok {
		return 0, errors.New("could not find quantity data")
	}
	convRateInput, ok := quantityData[inputUnit]
	if !ok {
		return 0, errors.New("could not retrieve conversion rate of input unit")
	}

	convRateOutput, ok := quantityData[outputUnit]
	if !ok {
		return 0, errors.New("could not retrieve conversion rate of output unit")
	}

	conversionRate := convRateInput / convRateOutput

	convertedAmount := amount * conversionRate

	return convertedAmount, nil
}
