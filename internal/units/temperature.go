package units

import (
	"errors"
	"fmt"
)

const (
	kelvin     = "kelvin (K)"
	celsius    = "Celsius (°C)"
	fahrenheit = "Fahrenheit (°F)"
)

type temp struct {
	kelvinTemp float64
}

func ConvertTemp(amount float64, inputUnit, outputUnit string) (float64, error) {
	temp, err := newTemp(amount, inputUnit)
	if err != nil {
		return 0, err
	}

	return temp.toUnit(outputUnit)
}

func GetTempUnits() []string {
	return []string{kelvin, celsius, fahrenheit}
}

func newTemp(val float64, unit string) (temp, error) {
	var kelvinTemp float64
	switch unit {
	case kelvin:
		kelvinTemp = val
	case celsius:
		kelvinTemp = val + 273.15
	case fahrenheit:
		kelvinTemp = (val + 459.67) * (float64(5) / 9)
	default:
		return temp{}, fmt.Errorf("could not create temp for '%s'", unit)
	}

	if kelvinTemp < 0 {
		return temp{}, errors.New("cannot create temp below absolute zero (0K)")
	}

	return temp{kelvinTemp}, nil
}

func (t temp) toUnit(unit string) (float64, error) {
	switch unit {
	case kelvin:
		return t.kelvinTemp, nil
	case celsius:
		return t.kelvinTemp + 273.15, nil
	case fahrenheit:
		return t.kelvinTemp*(float64(9)/5) - 459.67, nil
	default:
		return 0, fmt.Errorf("could not convert to unit '%s'", unit)
	}
}

func (t temp) ToCelsius() float64 {
	return t.kelvinTemp
}
