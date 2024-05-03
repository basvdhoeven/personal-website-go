package units

import (
	"fmt"
	"regexp"
	"strconv"
)

// Retrieve the amount and unit from a string.
// Expect a integer or decimal number followed by an unit
func ParseUnitsFromString(s string) (amount float64, unit string, err error) {
	// Get decimal coordinated from the string.
	r := regexp.MustCompile(`^(\d+(?:\.\d+)?)\s*([a-zA-Z0-9]+)$`)
	matches := r.FindStringSubmatch(s)
	fmt.Println(matches)
	fmt.Println(len(matches))
	// we should find ONLY a latitude and longitude
	if len(matches) != 3 {
		return 0, "", fmt.Errorf("could not retrieve amount + unit from string '%s'", s)
	}

	amount, err = strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, "", fmt.Errorf("could not convert amount to float 65: %w", err)
	}

	return amount, matches[2], nil
}
