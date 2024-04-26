package coords

import (
	"math"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	var tests = []struct {
		a        Coords
		b        Coords
		distance float64
		err      bool
	}{
		{Coords{lat: 0, lng: 0}, Coords{lat: 1, lng: 1}, 136.182823, false},
		{Coords{lat: 500, lng: 0}, Coords{lat: 1, lng: 1}, 0, true},
		{Coords{lat: 0.000000000001, lng: 0}, Coords{lat: 0, lng: 0}, 0, false},
	}

	for _, tt := range tests {
		dist, err := CalculateDistance(tt.a.lat, tt.a.lng, tt.b.lat, tt.b.lng)
		if math.Abs(dist-tt.distance) > 0.001 {
			t.Errorf("expected distance %f but got %f", tt.distance, dist)
		}
		if tt.err && err == nil {
			t.Errorf("expected distance error for (%f,%f) , (%f,%f)", tt.a.lat, tt.a.lng, tt.b.lat, tt.b.lng)
		}
	}
}
