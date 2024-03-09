package main

import (
    "net/http"

    "github.com/basvdhoeven/personal-website-go/routes"
)

func main() {
    // Register routes
    routes.RegisterRoutes()

    // Serve static files
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Start the server
    http.ListenAndServe(":8080", nil)
}