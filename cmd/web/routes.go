package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/about", app.aboutHandler)
	mux.HandleFunc("/projects", app.projectHandler)
	mux.HandleFunc("/projects/ip", app.ipHandler)
	mux.HandleFunc("/projects/coordinates", app.coordinatesHandler)
	mux.HandleFunc("/projects/unit", app.unitHandler)

	return mux
}
