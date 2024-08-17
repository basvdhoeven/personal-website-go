package main

import (
	"html/template" // New import
	"io/fs"
	"path/filepath" // New import

	"github.com/basvdhoeven/personal-website-go/ui"
)

type templateData struct {
	Ip              string
	CoordinatesData CoordinatesData
	UnitData        UnitData
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	// Return the map.
	return cache, nil
}
