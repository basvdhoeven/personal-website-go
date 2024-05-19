package units

import (
	"fmt"
	"regexp"
	"strconv"
)

type Measure struct {
	Amount float64
	Unit   string
}

// Retrieve the amount and unit from a string.
// Expect a integer or decimal number followed by an unit
func ParseUnitsFromString(s string) (measure Measure, err error) {
	// Get decimal coordinated from the string.
	r := regexp.MustCompile(`^(\d+(?:\.\d+)?)\s*([a-zA-Z0-9\s]+)$`)
	matches := r.FindStringSubmatch(s)

	// we should find ONLY a latitude and longitude
	if len(matches) != 3 {
		return Measure{}, fmt.Errorf("could not retrieve amount + unit from string '%s'", s)
	}

	amount, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return Measure{}, fmt.Errorf("could not convert amount to float 65: %w", err)
	}

	return Measure{amount, matches[2]}, nil
}
