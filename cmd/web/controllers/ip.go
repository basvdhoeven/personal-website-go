package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func IpHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/ip.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
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
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
