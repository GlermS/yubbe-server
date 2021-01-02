package router

import (
	"YubbeServer/yubbe-server/package/web/pages"
	"net/http"
)

type Router struct {
}

func (router *Router) ApplyRoutes() {

	http.HandleFunc("/login", pages.LoginPage)
	http.HandleFunc("/logout", pages.Logout)
	http.HandleFunc("/register", pages.RegisterPage)
	http.HandleFunc("/welcome", pages.WelcomePage)
}
