package units

import (
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Conversions map[string]struct {
	ConversionRate float64  `yaml:"conversion"`
	Units          []string `yaml:"units"`
}

type baseUnit struct {
	YamlFile string
	Unit     string
}

func ConvertUnits(measure Measure) (convertedMeasure string, detectedUnit string) {
	measure.Unit = strings.ToLower(measure.Unit)

	var baseUnits = []baseUnit{
		{YamlFile: "units/volume.yml", Unit: "liter"},
		{YamlFile: "units/length.yml", Unit: "meter"},
		{YamlFile: "units/mass.yml", Unit: "kilogram"},
	}

	for _, baseUnit := range baseUnits {
		convertedMeasure, detectedUnit = convert(measure, baseUnit)
		if convertedMeasure != "" {
			return convertedMeasure, detectedUnit
		}
	}

	return "", ""
}

func convert(measure Measure, baseUnit baseUnit) (convertedMeasure string, detectedUnit string) {
	data, err := os.ReadFile(baseUnit.YamlFile)
	if err != nil {
		panic(err)
	}

	var conversions Conversions
	err = yaml.Unmarshal(data, &conversions)
	if err != nil {
		panic(err)
	}

	for detectedUnit, details := range conversions {
		for _, u := range details.Units {
			if measure.Unit == u {
				return strconv.FormatFloat(measure.Amount*details.ConversionRate, 'f', 3, 64) + " " + baseUnit.Unit, detectedUnit
			}
		}
	}

	return "", ""
}
