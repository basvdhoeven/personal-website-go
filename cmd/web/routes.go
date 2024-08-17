package main

import (
	"net/http"

	"github.com/basvdhoeven/personal-website-go/ui"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileserver := http.FileServerFS(ui.Files)
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/about", app.aboutHandler)
	mux.HandleFunc("/projects", app.projectHandler)
	mux.HandleFunc("/projects/ip", app.ipHandler)
	mux.HandleFunc("/projects/coordinates", app.coordinatesHandler)
	mux.HandleFunc("/projects/unit", app.unitHandler)

	return app.recoverPanic(app.logRequest(mux))
}
