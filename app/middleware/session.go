package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("very-secure-one"))

// Session middleware.
func Session(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// initialize gorilla/session

		log.Println(store)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
