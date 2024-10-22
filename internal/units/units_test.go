package units

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-6

func TestGetAllUnits(t *testing.T) {
	UnitConverter := NewUnitConverter()
	err := UnitConverter.LoadConvRatesFromYaml(map[string]string{"mass": "../../config/units/mass.yml"})
	if err != nil {
		t.Errorf("Expected no error while loading existing unit data file")
	}

	units, err := UnitConverter.GetAllUnits("mass")
	if err != nil {
		t.Errorf("Expected no error while getting all the units")
	}

	if len(units) != 11 {
		t.Errorf("Got %d units, expected 3", len(units))
	}
}

func TestConvertUnits(t *testing.T) {
	UnitConverter := NewUnitConverter()
	err := UnitConverter.LoadConvRatesFromYaml(map[string]string{"mass": "../../config/units/mass.yml"})
	if err != nil {
		t.Errorf("Expected no error while loading existing unit data file")
	}

	var tests = []struct {
		quantity, inputUnit, outputUnit string
		amount, expectedConvertedAmount float64
		expectError                     bool
	}{
		{"mass", "kilogram", "kilogram", 10, 10, false},
		{"mass", "kilogram", "pound", 10, 22.046226, false},
		{"missing_quantity", "kilogram", "pound", 0, 0, true},
		{"mass", "missing_unit", "pound", 0, 0, true},
		{"mass", "kilogram", "missing_unit", 0, 0, true},
		{"mass", "kilogram", "pound", 0, 0, false},
	}

	for _, tt := range tests {
		convertedAmount, err := UnitConverter.Convert(tt.quantity, tt.inputUnit, tt.outputUnit, tt.amount)
		if math.Abs(tt.expectedConvertedAmount-convertedAmount) > float64EqualityThreshold {
			t.Errorf("got converted amount %f, expected %f", convertedAmount, tt.expectedConvertedAmount)
		}
		if tt.expectError && err == nil {
			t.Errorf("expected error for")
		}
	}
}

func TestLoadUnknownYaml(t *testing.T) {
	UnitConverter := NewUnitConverter()
	err := UnitConverter.LoadConvRatesFromYaml(map[string]string{"test": "not existing file"})
	if err == nil {
		t.Errorf("Expected error while loading non-existing unit data file")
	}
}
