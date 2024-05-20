package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	app := &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/about", app.aboutHandler)
	mux.HandleFunc("/projects", app.projectHandler)
	mux.HandleFunc("/projects/ip", app.ipHandler)
	mux.HandleFunc("/projects/coordinates", app.coordinatesHandler)
	mux.HandleFunc("/projects/unit", app.unitHandler)

	app.logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	app.logger.Error(err.Error())
	os.Exit(1)
}
