package routes

const (
	Home    RoutePath = "/"
	Signin  RoutePath = "/signin"
	Signup  RoutePath = "/signup"
	Signout RoutePath = "/signout"
)

var (
	GetHome     = NewHandlerName(GET, Home)
	GetSignin   = NewHandlerName(GET, Signin)
	PostSignin  = NewHandlerName(POST, Signin)
	GetSignup   = NewHandlerName(GET, Signup)
	PostSignup  = NewHandlerName(POST, Signup)
	PostSignout = NewHandlerName(POST, Signout)
)
