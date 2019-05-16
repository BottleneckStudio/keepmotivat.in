package controllers

import (
	"fmt"
	"net/http"
)

// PrivacyController is the Privacy screen
func PrivacyController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Privacy Controller")
	}
}
