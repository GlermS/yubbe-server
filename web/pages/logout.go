package pages

import (
	"net/http"
	"yubbe-server/authentication"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session := &authentication.Session{store}
	session.CleanSession(w, r)
	http.Redirect(w, r, "/login", 301)
}
