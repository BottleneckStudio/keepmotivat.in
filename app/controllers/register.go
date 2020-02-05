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

		// hash the password
		hashedPassword := hashPassword(password)

		userRepo := models.NewDBUserRepository(db)

		// check if we have the existing email address already
		getUserUsecase := usecase.NewGetUserUsecase(userRepo)
		user, err := getUserUsecase.GetUser(models.Query{})
		if err != nil {
			log.Println(err.Error())
			session.SetFlash(w, r, "error", "Something went wrong")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		if user.EmailAddress == emailAddress {
			session.SetFlash(w, r, "error", "Email Address is already taken.")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		user.EmailAddress = emailAddress
		user.Password = hashedPassword
		user.Ctime = time.Now().Unix()

		createUserUsecase := usecase.NewCreateUserUsecase(userRepo)
		if err := createUserUsecase.CreateUser(*user); err != nil {
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
