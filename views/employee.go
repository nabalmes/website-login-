package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func EmployeeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}

	username_cookie, _ := r.Cookie("username")
	username := username_cookie.Value
	user := uadmin.User{}

	uadmin.Get(&user, "username = ?", username)
	data["user"] = user

	session_cookie, _ := r.Cookie("session")
	if session_cookie != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: session_cookie.Value,
			Path:  "/",
		})
	}

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return data
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/employee")
	data["Title"] = "Employee | Pudding"
	return data
}
