package views

import (
	"net/http"
)

func DeliveriesHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}
	context["Title"] = "Deliveries | Pudding"

	return context
}
