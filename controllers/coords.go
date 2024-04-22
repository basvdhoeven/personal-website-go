package controllers

import (
	"html/template"
	"net/http"

	"github.com/basvdhoeven/personal-website-go/coords"
)

func CoordinatesHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and execute the template
	tmpl, err := template.ParseFiles("views/layouts/base.html", "views/coords.html")
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alat, alng, err := coords.GetCoordsFromString(r.URL.Query().Get("a"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	blat, blng, err := coords.GetCoordsFromString(r.URL.Query().Get("b"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	distance, err := coords.CalculateDistance(alat, alng, blat, blng)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Create a data structure to pass to the template
	data := struct {
		Title    string
		Distance float64
	}{
		Title:    "Coordinates",
		Distance: distance,
	}

	// Execute the template and write the output to the response
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
