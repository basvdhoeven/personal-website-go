package units

import (
	"testing"
)

func TestParseUnitsFromString(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		amount float64
		unit   string
		err    bool
	}{
		{
			name:   "valid input",
			input:  "12.5 kg",
			amount: 12.5,
			unit:   "kg",
			err:    false,
		},
		{
			name:   "valid input with decimal",
			input:  "3.14 m",
			amount: 3.14,
			unit:   "m",
			err:    false,
		},
		{
			name:   "valid input with integer",
			input:  "19 L",
			amount: 19,
			unit:   "L",
			err:    false,
		},
		{
			name:   "valid input with unit containing number",
			input:  "19 dm3",
			amount: 19,
			unit:   "dm3",
			err:    false,
		},
		{
			name:   "valid input with space in unit",
			input:  "19 fl oz",
			amount: 19,
			unit:   "fl oz",
			err:    false,
		},
		{
			name:   "invalid input",
			input:  "invalid",
			amount: 0,
			unit:   "",
			err:    true,
		},
		{
			name:   "empty input",
			input:  "",
			amount: 0,
			unit:   "",
			err:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsedMeasure, err := ParseUnitsFromString(tc.input)
			if tc.err && err == nil {
				t.Errorf("expected error but got nil")
			} else if !tc.err && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if parsedMeasure.Amount != tc.amount {
				t.Errorf("expected amount %f but got %f", tc.amount, parsedMeasure.Amount)
			}

			if parsedMeasure.Unit != tc.unit {
				t.Errorf("expected unit %s but got %s", tc.unit, parsedMeasure.Unit)
			}
		})
	}
}
