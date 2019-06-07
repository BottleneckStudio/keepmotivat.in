package controllers

import (
	"fmt"
	"net/http"

	"github.com/BottleneckStudio/keepmotivat.in/app/session"
)

// TermsOfServiceController is the TermsOfService screen
func TermsOfServiceController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TermsOfService Controller")

		flash := session.GetFlash(w, r)

		// horray!
		fmt.Fprintln(w, flash.Value)
	}
}
