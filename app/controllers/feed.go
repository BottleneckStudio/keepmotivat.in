package controllers

import (
	"net/http"

	tmpl "github.com/BottleneckStudio/keepmotivat.in/template"
)

// FeedController is the main screen
func FeedController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tpl := tmpl.New("./app/views/")

		if err := tpl.RenderHTML(w, "feed.html", "HELLO WORLD ALL CAPS!"); err != nil {
			return
		}
	}
}
