package routes

import (
    "net/http"

    "github.com/basvdhoeven/personal-website-go/controllers"
)

func RegisterRoutes() {
    http.HandleFunc("/", controllers.HomeHandler)
//     http.HandleFunc("/about", controllers.AboutHandler)
//     http.HandleFunc("/portfolio", controllers.PortfolioHandler)
//     http.HandleFunc("/contact", controllers.ContactHandler)
}