package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/basvdhoeven/personal-website-go/internal/units"
)

func UnitHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/unit.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
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
			data.ConvertedInput, data.DetectedUnit = units.ConvertUnits(parsedMeasure)
			if data.ConvertedInput == "" {
				data.ParseError = "Could not convert the input."
			}
		}
	}

	// Execute the template and write the output to the response
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
