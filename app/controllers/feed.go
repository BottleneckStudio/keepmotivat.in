package controllers

import (
	"fmt"
	"net/http"
)

// FeedController is the main screen
func FeedController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Feed")
	}
}
