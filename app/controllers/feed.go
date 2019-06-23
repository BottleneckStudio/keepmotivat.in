package controllers

import (
	"net/http"

	"BottleneckStudio/keepmotivat.in/app/session"
	tmpl "BottleneckStudio/keepmotivat.in/template"
)

// FeedController is the main screen
func FeedController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session.SetFlash(w, r, "error", "Some error message")

		tpl := tmpl.New("./app/views/")

		data := map[string]interface{}{}
		data["message"] = "Another World"

		if err := tpl.RenderHTML(w, "feed.html", data); err != nil {
			return
		}
	}
}
