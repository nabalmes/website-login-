package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	data["Title"] = "Log in | Pudding"
	uadmin.RenderHTML(w, r, "./templates/index.html", data)
}
