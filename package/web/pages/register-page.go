package pages

import (
	"YubbeServer/yubbe-server/package/authentication"
	"YubbeServer/yubbe-server/package/database/DB"
	"YubbeServer/yubbe-server/package/database/register"
	"fmt"
	"log"
	"net/http"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./package/web/pages/register-page/static/form.html")
	case "POST":
		session := &authentication.Session{store}
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		user, _ := DB.NewRegisterUser(r.FormValue("name"), r.FormValue("email"), r.FormValue("password"))
		dbr := register.NewDBRegister()

		err := dbr.Write(user)
		if err != nil {
			log.Println("Email inválido")
			http.Redirect(w, r, "/register", 301)
			return
		}

		err = session.AuthenticateSession(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/welcome", 301)
	}
}
