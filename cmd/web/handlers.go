package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/basvdhoeven/personal-website-go/internal/coords"
	"github.com/basvdhoeven/personal-website-go/internal/units"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.tmpl", templateData{})
}

func (app *application) toolsHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "tools.tmpl", templateData{})
}

func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "about.tmpl", templateData{})
}

func (app *application) ipHandler(w http.ResponseWriter, r *http.Request) {
	var ipAddress string
	parts := strings.Split(r.RemoteAddr, ":")
	if len(parts) > 0 {
		ipAddress = parts[0]
	}

	app.render(w, r, http.StatusOK, "ip.tmpl", templateData{Ip: ipAddress})
}

func (app *application) coordinatesHandler(w http.ResponseWriter, r *http.Request) {
	data := CoordinatesData{
		ParseError:  make(map[string]string),
		CoordsOrder: "latlng",
	}

	data.PointA = r.URL.Query().Get("pointa")
	data.PointB = r.URL.Query().Get("pointb")

	coordsOrder := r.URL.Query().Get("coords_order")
	if coordsOrder != "" {
		data.CoordsOrder = coordsOrder
	}

	if data.PointA == "" && data.PointB == "" {
		app.render(w, r, http.StatusOK, "coords.tmpl", templateData{CoordinatesData: data})
		return
	}

	var erra error
	data.LatA, data.LngA, erra = coords.GetCoordsFromString(data.PointA)
	if erra != nil {
		data.ParseError["pointa"] = "Could not retrieve coordinates from Point A"
	}

	var errb error
	data.LatB, data.LngB, errb = coords.GetCoordsFromString(data.PointB)
	if errb != nil && data.PointB != "" {
		data.ParseError["pointb"] = "Could not retrieve coordinates from Point B"
	}

	if data.CoordsOrder == "lnglat" {
		data.LatA, data.LngA = data.LngA, data.LatA
		data.LatB, data.LngB = data.LngB, data.LatB
	}

	if erra == nil && errb == nil {
		distance, err := coords.CalculateDistance(data.LatA, data.LngA, data.LatB, data.LngB)
		data.Distance = coords.FloatToDistanceString(distance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	app.render(w, r, http.StatusOK, "coords.tmpl", templateData{CoordinatesData: data})
}

func (app *application) unitHandler(w http.ResponseWriter, r *http.Request) {
	data := UnitData{
		Input: r.URL.Query().Get("input"),
	}

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

	app.render(w, r, http.StatusOK, "unit.tmpl", templateData{UnitData: data})
}

func (app *application) validateJson(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "json.tmpl", templateData{})
}

func (app *application) validateJsonPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData := r.PostForm.Get("json")

	valid := json.Valid([]byte(jsonData))
	if valid {
		var prettyJson bytes.Buffer
		json.Indent(&prettyJson, []byte(jsonData), "", " ")
		jsonData = prettyJson.String()
	}

	app.render(w, r, http.StatusOK, "json.tmpl", templateData{JsonValidation: JsonValidation{Data: jsonData, Valid: valid}})
}
