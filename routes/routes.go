package routes

import (
	"net/http"

	"github.com/basvdhoeven/personal-website-go/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/ip", controllers.IpHandler)
	http.HandleFunc("/projects", controllers.ProjectsHandler)
	http.HandleFunc("/about", controllers.AboutHandler)
}
