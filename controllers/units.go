package controllers

import (
	"html/template"
	"net/http"

	"github.com/basvdhoeven/personal-website-go/projects/units"
)

func UnitHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and execute the template
	tmpl, err := template.ParseFiles("views/layouts/base.html", "views/unit.html")
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a data structure to pass to the template
	data := struct {
		Title          string
		ParseError     string
		Input          string
		DetectedUnit   string
		ConvertedInput string
	}{
		Title: "Unit Converter",
	}

	data.Input = r.URL.Query().Get("input")

	if data.Input != "" {
		parsedMeasure, err := units.ParseUnitsFromString(data.Input)
		if err != nil {
			data.ParseError = "Could not retrieve amount and unit from input."
		} else {
			data.ConvertedInput, data.DetectedUnit = units.ConvertUnits(parsedMeasure)
			if data.ConvertedInput == "" {
				data.ParseError = "Could not convert the input."
			}
		}
	}

	// Execute the template and write the output to the response
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
