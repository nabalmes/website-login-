package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	session := uadmin.IsAuthenticated(r)
	user := uadmin.User{}

	if session != nil {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		uadmin.Get(&user, "username = ?", username)
		session := user.Login(password, "")

		if session != nil {
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "username",
				Value: username,
			})
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "session",
				Value: session.Key,
			})
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			if username == "" && password != "" {
				data["err_username"] = "Enter your username"
			} else if password == "" && username != "" {
				data["err_password"] = "Enter your password"
			} else if username == "" && password == "" {
				data["err_username"] = "Enter your username"
				data["err_password"] = "Enter your password"
			} else {
				data["err_username"] = "Invalid username"
				data["err_password"] = "Invalid password"
			}
		}
	}

	data["Title"] = "Log in | Pudding"
	uadmin.RenderHTML(w, r, "./templates/login.html", data)
}
