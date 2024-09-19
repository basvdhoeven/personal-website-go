package main

import (
	"net/http"

	"github.com/basvdhoeven/personal-website-go/ui"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /", app.homeHandler)
	mux.HandleFunc("GET /about", app.aboutHandler)
	mux.HandleFunc("GET /tools", app.toolsHandler)
	mux.HandleFunc("GET /tools/ip", app.ipHandler)
	mux.HandleFunc("GET /tools/coordinates", app.coordinatesHandler)
	mux.HandleFunc("GET /tools/unit/", app.unitHandler)
	mux.HandleFunc("POST /tools/unit/", app.unitHandlerPost)
	mux.HandleFunc("GET /tools/json", app.validateJson)
	mux.HandleFunc("POST /tools/json", app.validateJsonPost)

	return app.recoverPanic(app.logRequest(mux))
}
