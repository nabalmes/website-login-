package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)



func PuddingHandler(w http.ResponseWriter, r *http.Request){
	session := r.FormValue("session")
	if  session != "" {
		http.SetCookie(w, &http.Cookie{
			Name: "session",
			Value: session,
			Path: "/",
		})
	}


	//check if user is not authenticated
	//and redirect to login page


	sessionn := uadmin.IsAuthenticated(r)
	if sessionn == nil {
		http.Redirect(w, r, "/", 303)
		return
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/pudding")
	page := strings.TrimSuffix(r.URL.Path, "/")
	context := map[string]interface{}{}


	switch page {
	case "deliveries":
		context = DeliveriesHandler(w, r)
	case "employee":
		context = EmployeeHandler(w, r)
	default:
		page = "deliveries"
	}

	RenderInterface(w, r, page, context)
}

func RenderInterface(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}){
	templateList := []string{}
	templateList = append(templateList, "./templates/base.html")

	path := "./templates/"+ page +".html"
	templateList = append(templateList, path)
	uadmin.Trail(uadmin.DEBUG,"PAGE: %v",page)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}