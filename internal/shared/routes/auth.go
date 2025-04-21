package routes

import "gh_static_portfolio/internal/shared/web"

const (
	Home    web.RoutePath = "/"
	Signin  web.RoutePath = "/signin"
	Signup  web.RoutePath = "/signup"
	Signout web.RoutePath = "/signout"
)

var (
	GetHome     = web.NewHandlerName(web.GET, Home)
	GetSignin   = web.NewHandlerName(web.GET, Signin)
	PostSignin  = web.NewHandlerName(web.POST, Signin)
	GetSignup   = web.NewHandlerName(web.GET, Signup)
	PostSignup  = web.NewHandlerName(web.POST, Signup)
	PostSignout = web.NewHandlerName(web.POST, Signout)
)
