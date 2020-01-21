package controllers

import (
	"log"
	"net/http"

	"BottleneckStudio/keepmotivat.in/app/session"
	"BottleneckStudio/keepmotivat.in/models"
	tmpl "BottleneckStudio/keepmotivat.in/template"
)

// RegisterViewHandler shows the form.
func RegisterViewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl := tmpl.New("./app/views/")

		flash := session.GetFlash(w, r)

		data := map[string]interface{}{}
		data["flash"] = flash

		log.Println(flash)

		if err := tpl.RenderHTML(w, "register.html", data); err != nil {
			return
		}
	}
}

// RegisterPostController ...
func RegisterPostController(db *models.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// parse form
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		confirmPassword := r.Form.Get("confirm_password")

		if password != confirmPassword {
			session.SetFlash(w, r, "error", "Passwords dont match")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		// hash the password
		log.Println(username)
		log.Println(password)
		log.Println(confirmPassword)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
