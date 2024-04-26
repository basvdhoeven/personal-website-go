package controllers

import (
	"fmt"
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

	// Create a data structure to pass to the template
	data := struct {
		Title      string
		Distance   string
		PointA     string
		PointB     string
		LatA       float64
		LngA       float64
		LatB       float64
		LngB       float64
		ParseError map[string]string
	}{
		Title:      "Coordinates",
		ParseError: make(map[string]string),
	}

	data.PointA = r.URL.Query().Get("pointa")
	data.PointB = r.URL.Query().Get("pointb")

	if data.PointA == "" && data.PointB == "" {
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var erra error
	data.LatA, data.LngA, erra = coords.GetCoordsFromString(data.PointA)
	if erra != nil {
		data.ParseError["pointa"] = "Could not retrieve coordinates from Point A"
	}

	var errb error
	data.LatB, data.LngB, errb = coords.GetCoordsFromString(data.PointB)
	if errb != nil {
		data.ParseError["pointb"] = "Could not retrieve coordinates from Point B"
	}

	var distance float64
	if erra == nil && errb == nil {
		distance, err = coords.CalculateDistance(data.LatA, data.LngA, data.LatB, data.LngB)
		data.Distance = floatToDistanceString(distance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// Execute the template and write the output to the response
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func floatToDistanceString(distance float64) string {
	// more than 1 km
	if distance > 1 {
		return fmt.Sprintf("%.2f km", distance)
	}
	// less than 1 km
	distanceMeters := distance * 1000
	return fmt.Sprintf("%.2f m", distanceMeters)
}
