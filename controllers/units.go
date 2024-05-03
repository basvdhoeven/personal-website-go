package controllers

import (
	"html/template"
	"net/http"

	"github.com/basvdhoeven/personal-website-go/units"
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
		Title      string
		ParseError string
		Unit       string
	}{
		Title: "Unit Converter",
	}

	data.Unit = r.URL.Query().Get("unit")

	if data.Unit != "" {
		amount, unit, err := units.ParseUnitsFromString(data.Unit)
	}

	if err != nil {
		data.ParseError = "Could not retrieve amount and unit"
	}

	// Execute the template and write the output to the response
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
