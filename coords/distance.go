package coords

import (
	"fmt"
	"math"
)

type coords struct {
	lat float64
	lng float64
}

const earthRadius = 6371 // average Earth radius in km

func CalculateDistance(aLat, aLng, bLat, bLng float64) (float64, error) {
	a := coords{aLat, aLng}
	err := a.validCoordinates()
	if err != nil {
		return 0, fmt.Errorf("coordinates of Point A are invalid; %w", err)
	}

	b := coords{bLat, bLng}
	err = b.validCoordinates()
	if err != nil {
		return 0, fmt.Errorf("coordinates of Point B are invalid; %w", err)
	}

	return haversineDistance(a, b), nil
}

func (point coords) validCoordinates() error {
	if math.Abs(point.lat) > 90 {
		return fmt.Errorf("latitude must be between -90 and 90 degrees: (%f)", point.lat)
	}

	if (math.Abs(point.lng) > 180) || (math.Abs(point.lng) < 0) {
		return fmt.Errorf("longitude must be between 0 and 180 degrees: (%f, %f)", point.lat, point.lng)
	}

	return nil
}

// See https://en.wikipedia.org/wiki/Haversine_formula
func haversineDistance(a, b coords) (distance float64) {
	// Convert degrees to radians
	dLat := degToRad(b.lat - a.lat)
	dLng := degToRad(b.lng - a.lng)

	aLat := degToRad(a.lat)
	bLat := degToRad(b.lat)

	fraction := haversine(dLat) + math.Cos(aLat)*math.Cos(bLat)*(haversine(dLng))/2

	return 2 * earthRadius * math.Asin(math.Sqrt(fraction))
}

func haversine(alpha float64) float64 {
	return (1 - math.Cos(alpha)) / 2
}

func degToRad(degrees float64) float64 {
	return degrees * (math.Pi) / 180
}
