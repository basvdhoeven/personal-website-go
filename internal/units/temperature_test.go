package units

import (
	"math"
	"testing"
)

func TestGetTempUnits(t *testing.T) {
	units := GetTempUnits()
	if len(units) != 3 {
		t.Errorf("expected 3 temperature units")
	}
}

func TestConvertTemp(t *testing.T) {
	var tests = []struct {
		inputUnit, outputUnit           string
		amount, expectedConvertedAmount float64
		expectedErrorMessage            string
	}{
		{"kelvin (K)", "kelvin (K)", 10, 10, ""},
		{"Celsius (°C)", "kelvin (K)", 25, 298.15, ""},
		{"Celsius (°C)", "Fahrenheit (°F)", 0, 32, ""},
		{"Celsius (°C)", "kelvin (K)", -300, 0, "cannot create temp below absolute zero (0K)"},
		{"invalid input", "kelvin (K)", 0, 0, "could not create temp for 'invalid input'"},
		{"Celsius (°C)", "invalid output", 0, 0, "could not convert to unit 'invalid output'"},
	}

	for _, tt := range tests {
		convertedAmount, err := ConvertTemp(tt.amount, tt.inputUnit, tt.outputUnit)
		if math.Abs(tt.expectedConvertedAmount-convertedAmount) > float64EqualityThreshold {
			t.Errorf("got converted amount %f, expected %f", convertedAmount, tt.expectedConvertedAmount)
		}
		if tt.expectedErrorMessage != "" && err == nil {
			t.Errorf("expected error message: %s", tt.expectedErrorMessage)
		}
		if err != nil && tt.expectedErrorMessage != err.Error() {
			t.Errorf("expected error message: %s but got %s", tt.expectedErrorMessage, err.Error())

		}
	}
}
