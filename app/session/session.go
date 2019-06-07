package session

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
)

// Flash message
type Flash struct {
	Type    string
	Message string
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// Get wraps the session store Getter
func Get(r *http.Request, name string) (*sessions.Session, error) {
	return store.Get(r, name)
}

// SetFlash sets the flash message with the given
// type and message
func SetFlash(w http.ResponseWriter, r *http.Request, t, m string) {
	session, _ := store.Get(r, "notification")
	session.AddFlash(fmt.Sprintf("%s<>%s", t, m))
	if err := session.Save(r, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetFlash ...
func GetFlash(w http.ResponseWriter, r *http.Request) *Flash {
	session, _ := store.Get(r, "notification")

	if flashes := session.Flashes(); len(flashes) > 0 {
		chunks := strings.Split(flashes[0].(string), "<>")
		f := &Flash{
			Type:    chunks[0],
			Message: chunks[1],
		}
		if err := session.Save(r, w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return nil
		}
		return f
	}

	return nil
}
