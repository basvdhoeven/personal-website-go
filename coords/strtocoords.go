package coords

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetCoordsFromString(s string) (lat, lng float64, err error) {
	// Get decimal coordinated from the string.
	matchesDecimal, ok := matchCoordsInString(s, `[-+]?\d+([.,]\d*)?`)
	if ok {
		return convertToFloat(matchesDecimal)
	}

	// If not succesful, attempt to get non-decimal coordinates.
	matchesNonDecimal, ok := matchCoordsInString(s, `[-+]?\d+`)
	if ok {
		return convertToFloat(matchesNonDecimal)
	}

	return 0, 0, fmt.Errorf("could not retrieve a latitude/longitude pair from string %s", s)
}

func matchCoordsInString(s, regex string) (matches []string, ok bool) {
	r := regexp.MustCompile(regex)
	matches = r.FindAllString(s, -1)

	// we should find ONLY a latitude and longitude
	if len(matches) == 2 {
		return matches, true
	}

	return nil, false
}

func convertToFloat(coords []string) (lat, lng float64, err error) {
	latString := strings.Replace(coords[0], ",", ".", -1)
	lat, err = strconv.ParseFloat(latString, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("could not convert latitude to float: %s; %w", coords[0], err)
	}

	lngString := strings.Replace(coords[1], ",", ".", -1)
	lng, err = strconv.ParseFloat(lngString, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("could not convert longitude to float: %s; %w", coords[1], err)
	}

	return lat, lng, nil
}
