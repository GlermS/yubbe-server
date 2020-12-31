package authentication

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

// User holds a users account information
type User struct {
	Username      string
	Authenticated bool
}

// store will hold all session data
type Session struct {
	Store *sessions.CookieStore
}

// getUser returns a user from session s
// on error returns an empty user
func (ss *Session) GetUser(r *http.Request) (User, error) {
	session, err := ss.Store.Get(r, "cookie-name")

	val := session.Values["user"]
	var user = User{}

	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}, err
	}
	return user, err
}

func (ss *Session) CheckSession(r *http.Request) (bool, error) {

	user, err := ss.GetUser(r)

	auth := user.Authenticated

	if err != nil {
		log.Println(err)
	}

	return auth, err
}

func (ss *Session) AuthenticateSession(r *http.Request, w http.ResponseWriter) error {
	session, err := ss.Store.Get(r, "cookie-name")

	if err != nil {
		log.Println(err)
	}

	user := &User{
		Username:      r.FormValue("email"),
		Authenticated: true,
	}

	session.Values["user"] = user

	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (ss *Session) CleanSession(w http.ResponseWriter, r *http.Request) error {
	session, err := ss.Store.Get(r, "cookie-name")
	if err != nil {
		log.Println(err)
	}

	session.Values["user"] = User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
	}
	return err
}
