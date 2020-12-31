package pages

import (
	"encoding/gob"
	"yubbe-server/authentication"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Page struct {
	Session *authentication.Session
}

var Pages *Page

var store *sessions.CookieStore

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(authentication.User{})

}
