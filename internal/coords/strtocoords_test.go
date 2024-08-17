package coords

import (
	"testing"
)

func TestGetCoordsFromString(t *testing.T) {
	var tests = []struct {
		s        string
		error    bool
		lat, lng float64
	}{
		{"lat: 50.25234 long: 0.4354", false, 50.25234, 0.4354},
		{"(50.25234, 0.4354)", false, 50.25234, 0.4354},
		{"(50.25234 0.4354)", false, 50.25234, 0.4354},
		{"lat: 50 long: 7", false, 50, 7},
		{"50,7", false, 50, 7},
		{"lat: 50,25234 long: 0,4354", false, 50.25234, 0.4354},
		{"50.25234,0.4354", false, 50.25234, 0.4354},
		{"50", true, 0, 0},
		{"", true, 0, 0},
		{"test", true, 0, 0},
	}

	for _, tt := range tests {
		lat, lng, err := GetCoordsFromString(tt.s)
		if (lat != tt.lat) || (lng != tt.lng) {
			t.Errorf("got coords (%f, %f), wanted (%f, %f) from '%s'", lat, lng, tt.lat, tt.lng, tt.s)
		}
		if tt.error && err == nil {
			t.Errorf("expected error for '%s'", tt.s)
		}
	}
}
