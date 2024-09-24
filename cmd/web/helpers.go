package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
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

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func getIp(r *http.Request) string {
	// When deployed in Google Cloud Run, the forwarded client ip can
	// be found in the "X-Forwarded-For" header
	ip := r.Header.Get("X-Forwarded-For")

	// if deployed locally, get ip from RemoteAddr field of request
	if ip == "" {
		parts := strings.Split(r.RemoteAddr, ":")
		if len(parts) > 0 {
			ip = parts[0]
		}
	}

	return ip
}
