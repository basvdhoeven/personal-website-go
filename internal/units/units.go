package units

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	Mass   = "mass"
	Length = "length"
	Volume = "volume"
)

func GetUnits(quantity string) ([]string, error) {
	units, err := loadUnitData(quantity)
	if err != nil {
		return nil, fmt.Errorf("loading unit data; %w", err)
	}

	unitSlice := make([]string, 0, len(units))
	for unit := range units {
		unitSlice = append(unitSlice, unit)
	}

	return unitSlice, nil
}

func Convert(quantity, inputUnit, outputUnit string, amount float64) (float64, error) {
	unitData, err := loadUnitData(quantity)
	if err != nil {
		return 0, fmt.Errorf("loading unit data; %w", err)
	}

	conversionRate := unitData[inputUnit] / unitData[outputUnit]

	convertedAmount := amount * conversionRate

	return convertedAmount, nil
}

func loadUnitData(quantity string) (map[string]float64, error) {
	data, err := os.ReadFile("./internal/units/" + quantity + ".yml")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Unmarshal YAML into a map
	var units map[string]float64
	err = yaml.Unmarshal(data, &units)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return units, nil
}
