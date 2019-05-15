package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// PostController is the main screen
func PostController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "postID")
		fmt.Fprintln(w, "Post # is "+postID)
	}
}
