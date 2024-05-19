package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/basvdhoeven/personal-website-go/cmd/web/config"
	"github.com/basvdhoeven/personal-website-go/cmd/web/controllers"
)

// Define an application struct to hold the application-wide dependencies for the
// web application.
type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	app := &config.Application{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", controllers.HomeHandler(app))
	mux.HandleFunc("/about", controllers.AboutHandler(app))
	mux.HandleFunc("/projects", controllers.ProjectsHandler(&config.Application{}))
	mux.HandleFunc("/projects/ip", controllers.IpHandler(app))
	mux.HandleFunc("/projects/coordinates", controllers.CoordinatesHandler(app))
	mux.HandleFunc("/projects/unit", controllers.UnitHandler(app))

	app.Logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	app.Logger.Error(err.Error())
	os.Exit(1)
}
