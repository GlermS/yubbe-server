package router

import (
	"net/http"
	"yubbe-server/web/pages"
)

type Router struct {
}

func (router *Router) ApplyRoutes() {

	http.HandleFunc("/login", pages.LoginPage)
	http.HandleFunc("/logout", pages.Logout)
	http.HandleFunc("/register", pages.RegisterPage)
	http.HandleFunc("/welcome", pages.WelcomePage)
}
