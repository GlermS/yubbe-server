package pages

import (
	"YubbeServer/yubbe-server/package/authentication"
	"YubbeServer/yubbe-server/package/database/DB"
	"YubbeServer/yubbe-server/package/database/login"
	"fmt"
	"log"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	session := &authentication.Session{store}
	auth, err := session.CheckSession(r)
	if err != nil {
		log.Println("Limpou o cache")
		session.CleanSession(w, r)
		http.Redirect(w, r, "/login", 301)
	}

	if auth {
		log.Println("Autenticado")
		http.Redirect(w, r, "/welcome", 301)
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./web/pages/login-page/static/form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			return
		}
		user := DB.NewLoginUser(r.FormValue("email"), r.FormValue("password"))
		dblv := login.NewDBLoginVerifier()

		err := dblv.Check(user)
		if err != nil {
			log.Println("Erro no login-page")
			http.Redirect(w, r, "/login", 301)
		} else {

			err = session.AuthenticateSession(r, w)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println("Sem erro")
			http.Redirect(w, r, "/welcome", 301)

		}

	}
}
