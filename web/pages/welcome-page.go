package pages

import (
	"YubbeServer/yubbe-server/authentication"
	"log"
	"net/http"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	session := &authentication.Session{store}
	auth, err := session.CheckSession(r)
	if err != nil {
		session.CleanSession(w, r)
		http.Redirect(w, r, "/login", 301)
	}
	_, err = session.GetUser(r)
	if err != nil {
		log.Println("Get user")
		log.Println(err)
	}

	if !auth {
		http.Redirect(w, r, "/login", 301)
	}

	http.ServeFile(w, r, "./web/pages/welcome-page/static/welcome.html")
}
