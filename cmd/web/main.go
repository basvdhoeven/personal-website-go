package main

import (
	"log"
	"net/http"

	"github.com/basvdhoeven/personal-website-go/cmd/web/controllers"
)

func test(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
}

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", controllers.HomeHandler)	
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/about", controllers.AboutHandler)
	mux.HandleFunc("/projects", controllers.ProjectsHandler)
	mux.HandleFunc("/projects/ip", controllers.IpHandler)
	mux.HandleFunc("/projects/coordinates", controllers.CoordinatesHandler)
	mux.HandleFunc("/projects/unit", controllers.UnitHandler)

	log.Print("starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
