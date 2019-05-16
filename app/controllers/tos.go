package controllers

import (
	"fmt"
	"net/http"
)

// TermsOfServiceController is the TermsOfService screen
func TermsOfServiceController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TermsOfService Controller")
	}
}
