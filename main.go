package main

import (
	"net/http"

	"github.com/mbdeguzman/pudding_delivery/models"
	"github.com/mbdeguzman/pudding_delivery/views"
	"github.com/uadmin/uadmin"
)

func main() {

	DatabaseConfig()
	RegisterModel()
	RegisterHandlers()
	ServerStart()
}

func DatabaseConfig() {
	uadmin.Trail(uadmin.INFO, "Database Config")
}
func RegisterModel() {
	uadmin.Trail(uadmin.INFO, "Register Models")
	uadmin.Register(
		models.Tags{},
		models.Location{},
		models.Deliveries{},
		models.Employee{},
	)
}

func RegisterHandlers() {
	uadmin.Trail(uadmin.INFO, "Register Handlers")
	http.HandleFunc("/", uadmin.Handler(views.LandingPageHandler))
	http.HandleFunc("/login", uadmin.Handler(views.LogInHandler))
	http.HandleFunc("/logout", uadmin.Handler(views.LogOutHandler))
	http.HandleFunc("/pudding", uadmin.Handler(views.PuddingHandler))
}

func ServerStart() {
	uadmin.Port = 2024
	uadmin.RootURL = "/admin/"
	uadmin.StartServer()
}
