package controllers

import (
	"fmt"
	"net/http"
)

// AboutController is the about us screen
func AboutController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "About Controller")
	}
}
