package routes

import (
	"net/http"

	"github.com/basvdhoeven/personal-website-go/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/about", controllers.AboutHandler)
	http.HandleFunc("/projects", controllers.ProjectsHandler)
	http.HandleFunc("/projects/ip", controllers.IpHandler)
	http.HandleFunc("/projects/coordinates", controllers.CoordinatesHandler)
	http.HandleFunc("/projects/unit", controllers.UnitHandler)
}
