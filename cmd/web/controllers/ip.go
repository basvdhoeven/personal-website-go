package controllers

import (
	"html/template"
	"net/http"

	"github.com/basvdhoeven/personal-website-go/cmd/web/config"
)

func IpHandler(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/ip.tmpl",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Ip string
		}{
			Ip: r.RemoteAddr,
		}

		// Execute the template and write the output to the response
		if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
