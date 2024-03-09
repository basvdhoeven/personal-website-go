package controllers

import (
    "html/template"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // Parse and execute the template
    tmpl, err := template.ParseFiles("views/home.html", "views/layouts/base.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Create a data structure to pass to the template
    data := struct {
        Title string
    }{
        Title: "My Personal Website",
    }

    // Execute the template and write the output to the response
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}