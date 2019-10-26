package controllers

import (
	"log"
	"net/http"

	"BottleneckStudio/keepmotivat.in/app/session"
	"BottleneckStudio/keepmotivat.in/models"
	tmpl "BottleneckStudio/keepmotivat.in/template"
	"BottleneckStudio/keepmotivat.in/usecase"
)

// FeedController is the main screen
func FeedController(db *models.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session.SetFlash(w, r, "error", "Some error message")

		tpl := tmpl.New("./app/views/")

		data := map[string]interface{}{}
		data["message"] = "Another World"

		// Get all post ...
		postRepo := models.NewDBPostRepository(db)
		postUsecase := usecase.NewGetPostUsecase(postRepo)
		posts, err := postUsecase.GetAllPosts(&models.Query{})
		log.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@")
		log.Println(posts)
		if err != nil {
			log.Println(err)
		}

		data["posts"] = posts

		if err := tpl.RenderHTML(w, "feed.html", data); err != nil {
			return
		}
	}
}
