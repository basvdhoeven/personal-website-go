package main

import (
	"crypto/tls"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/basvdhoeven/personal-website-go/internal/quotes"
	"github.com/basvdhoeven/personal-website-go/internal/units"
)

type application struct {
	logger         *slog.Logger
	templateCache  map[string]*template.Template
	unitConverter  *units.UnitConverter
	quoteRetriever *quotes.QuoteRetriever
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

	unitConverter := units.NewUnitConverter()
	unitMapData := map[string]string{
		"length": "./config/units/length.yml",
		"mass":   "./config/units/mass.yml",
		"volume": "./config/units/volume.yml",
	}
	unitConverter.LoadConvRatesFromYaml(unitMapData)

	quoteRetriever := quotes.NewQuotesRetriever()
	quoteRetriever.LoadQuotesFromTextFiles("./config/quotes")

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		unitConverter: unitConverter,
		quoteRetriever: quoteRetriever,
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server", "addr", *addr)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
