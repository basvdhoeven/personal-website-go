package main

import (
	"net/http"

	"github.com/basvdhoeven/personal-website-go/ui"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/about", app.aboutHandler)
	mux.HandleFunc("/tools", app.toolsHandler)
	mux.HandleFunc("/tools/ip", app.ipHandler)
	mux.HandleFunc("/tools/coordinates", app.coordinatesHandler)
	mux.HandleFunc("/tools/unit", app.unitHandler)
	mux.HandleFunc("/tools/json", app.jsonHandler)

	return app.recoverPanic(app.logRequest(mux))
}
