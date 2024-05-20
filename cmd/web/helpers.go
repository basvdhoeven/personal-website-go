package main

import (
	"fmt"
	"net/http"
)

type CoordinatesData struct {
	Distance    string
	PointA      string
	PointB      string
	LatA        float64
	LngA        float64
	LatB        float64
	LngB        float64
	ParseError  map[string]string
	CoordsOrder string
}

type UnitData struct {
	Input          string
	DetectedUnit   string
	ConvertedInput string
	ParseError     string
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}