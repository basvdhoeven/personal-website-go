package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
	}

	app.logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}
