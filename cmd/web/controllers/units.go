package controllers

import (
	"html/template"
	"net/http"

	"github.com/basvdhoeven/personal-website-go/cmd/web/config"
	"github.com/basvdhoeven/personal-website-go/internal/units"
)

func UnitHandler(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/unit.tmpl",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			ParseError     string
			Input          string
			DetectedUnit   string
			ConvertedInput string
		}{}

		data.Input = r.URL.Query().Get("input")

		if data.Input != "" {
			parsedMeasure, err := units.ParseUnitsFromString(data.Input)
			if err != nil {
				data.ParseError = "Could not retrieve amount and unit from input."
			} else {
				var baseUnits = units.BaseUnits{
					{YamlFile: "./internal/units/volume.yml", Unit: "liter"},
					{YamlFile: "./internal/units/length.yml", Unit: "meter"},
					{YamlFile: "./internal/units/mass.yml", Unit: "kilogram"},
				}

				data.ConvertedInput, data.DetectedUnit = units.ConvertUnits(parsedMeasure, baseUnits)
				if data.ConvertedInput == "" {
					data.ParseError = "Could not convert the input."
				}
			}
		}

		// Execute the template and write the output to the response
		if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
