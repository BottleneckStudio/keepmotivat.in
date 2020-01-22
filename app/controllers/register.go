package controllers

import (
	"log"
	"net/http"
	"time"

	"BottleneckStudio/keepmotivat.in/app/session"
	"BottleneckStudio/keepmotivat.in/models"
	tmpl "BottleneckStudio/keepmotivat.in/template"
	"BottleneckStudio/keepmotivat.in/usecase"
	"golang.org/x/crypto/bcrypt"
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
		if err := r.ParseForm(); err != nil {
			session.SetFlash(w, r, "error", err.Error())
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		emailAddress := r.Form.Get("email_address")
		password := r.Form.Get("password")
		confirmPassword := r.Form.Get("confirm_password")

		if password != confirmPassword {
			session.SetFlash(w, r, "error", "Passwords dont match")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		// check if we have the existing email address already
		log.Println(emailAddress)

		// hash the password
		hashedPassword := hashPassword(password)

		userRepo := models.NewDBUserRepository(db)
		userUsecase := usecase.NewCreateUserUsecase(userRepo)

		user := models.User{
			EmailAddress: emailAddress,
			Password:     hashedPassword,
			Ctime:        time.Now().Unix(),
		}

		if err := userUsecase.CreateUser(user); err != nil {
			log.Println("Create User Error: " + err.Error())
			session.SetFlash(w, r, "error", "Something went wrong registering your account.")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func hashPassword(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords() bool {
	return false
}
