package units

import (
	"testing"
)

func TestConvertUnits(t *testing.T) {
	testCases := []struct {
		name                 string
		input                Measure
		expectedConverted    string
		expectedDetectedUnit string
	}{
		{
			name:                 "valid input mass",
			input:                Measure{Amount: 12.5, Unit: "lb"},
			expectedConverted:    "5.670 kilogram",
			expectedDetectedUnit: "pound",
		},
		{
			name:                 "valid input length",
			input:                Measure{Amount: 13.2, Unit: "ft"},
			expectedConverted:    "4.023 meter",
			expectedDetectedUnit: "foot",
		},
		{
			name:                 "valid input volume",
			input:                Measure{Amount: 2, Unit: "tbsp"},
			expectedConverted:    "30.0 milliliter",
			expectedDetectedUnit: "tablespoon",
		},
		{
			name:                 "invalid input",
			input:                Measure{Amount: 2, Unit: "firetruck"},
			expectedConverted:    "",
			expectedDetectedUnit: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			converted, detectedUnit := ConvertUnits(tc.input)
			if converted != tc.expectedConverted {
				t.Errorf("expected converted amount %s but got %s", tc.expectedConverted, converted)
			}

			if detectedUnit != tc.expectedDetectedUnit {
				t.Errorf("expected detected unit %s but got %s", tc.expectedDetectedUnit, detectedUnit)
			}
		})
	}
}
