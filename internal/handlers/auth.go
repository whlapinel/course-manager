package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	auth "gh_static_portfolio/internal/authentication"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	mt "gh_static_portfolio/internal/templates/app"
	components "gh_static_portfolio/internal/templates/components/base"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func AuthHandlers(svc service.CourseService, echo *echo.Echo) []RouteHandler {
	ar := newAuthRouter(svc, echo)
	var routeHandlers []RouteHandler
	authRouteHandlers := []RouteHandler{
		{Signin, GetSignin, GET, ar.GetSignin},
		{Signin, PostSignin, POST, ar.PostSignin},
		{Signup, GetSignup, GET, ar.GetSignup},
		{Signup, PostSignup, POST, ar.PostSignup},
		{Signout, PostSignout, POST, ar.PostSignout},
	}
	routeHandlers = append(routeHandlers, authRouteHandlers...)
	return routeHandlers
}

func newAuthRouter(svc service.CourseService, app *echo.Echo) authRouter {
	return authRouter{
		router: router{
			svc: svc,
			app: app,
		},
	}
}

const (
	Signin  RoutePath = "/signin"
	Signup  RoutePath = "/signup"
	Signout RoutePath = "/signout"
)

const (
	GetSignin   RouteHandlerName = RouteHandlerName(GET + Signin)
	PostSignin  RouteHandlerName = RouteHandlerName(POST + Signin)
	GetSignup   RouteHandlerName = RouteHandlerName(GET + Signup)
	PostSignup  RouteHandlerName = RouteHandlerName(POST + Signup)
	PostSignout RouteHandlerName = RouteHandlerName(POST + Signout)
)

type authRouter struct {
	router
}

// GetRouter implements Router.
func (a *authRouter) GetRouter() router {
	panic("unimplemented")
}

// SetRouter implements Router.
func (a *authRouter) SetRouter(router router) {
	panic("unimplemented")
}

func (h authRouter) GetSignin(c echo.Context) error {
	page := mt.SigninPage{
		GoogleClientID:  os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleSigninURL: h.app.Reverse(PostSignin.String()),
	}
	c.Request().Header.Add("HX-Retarget", "#page")
	component := page.Component()
	layout := CourseManagerLayout(h.app, page.Component(), domain.User{})
	return Respond(c, "", component, layout)
}

func (h authRouter) PostSignin(c echo.Context) error {
	payload, err := auth.GoogleAuth(c)
	if err != nil {
		log.Println(err)
		return c.String(500, "Failed to authenticate")
	}
	sub := payload.Claims["sub"].(string)
	user, err := h.svc.GetUser(sub)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return h.RespondUnauthorized(c, NotRegistered)
		}
		log.Println(err)
		return c.String(500, err.Error())
	}
	log.Println("userID", user.ID)
	t, err := auth.IssueToken(auth.TokenParams{User: domain.User{
		ID:        sub,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Picture:   user.Picture,
	}})
	if err != nil {
		return c.String(500, "Failed to issue token")
	}
	auth.WriteToken(c, t)
	return c.Redirect(303, "/users")
}

func (h authRouter) GetSignup(c echo.Context) error {

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	page := mt.SignUpPage{
		GoogleClientID:  clientID,
		GoogleSignupURL: h.app.Reverse(GetSignup.String()),
	}
	component := page.Component()
	layout := CourseManagerLayout(h.app, page.Component(), domain.User{})
	return Respond(c, "", component, layout)
}

func (h authRouter) PostSignup(c echo.Context) error {
	payload, err := auth.GoogleAuth(c)
	if err != nil {
		log.Println("Failed to authenticate user: ", err)
		return c.String(401, "Failed to authenticate user")
	}
	id := payload.Claims["sub"].(string)
	email := payload.Claims["email"].(string)
	first := payload.Claims["given_name"].(string)
	last := payload.Claims["family_name"].(string)
	picture := payload.Claims["picture"].(string)
	// create user
	if !authorized(email) {
		return h.RespondUnauthorized(c, WrongDomain)
	}
	dupe, err := h.isDuplicate(id)
	if err != nil {
		log.Println(err)
		return c.String(500, fmt.Sprintf("error checking for duplicate user: %v", err))
	}
	if dupe {
		return h.RespondUnauthorized(c, DuplicateUser)
	}
	user, err := h.svc.SaveUser(service.SaveUserParams{
		User: domain.User{
			ID:        id,
			Email:     email,
			FirstName: first,
			LastName:  last,
			Picture:   picture,
		},
	})
	if err != nil {
		log.Println("Failed to create user: ", err)
		return c.String(500, "Failed to create user")
	}
	t, err := auth.IssueToken(auth.TokenParams{User: user})
	if err != nil {
		return c.String(500, "Failed to issue token")
	}
	auth.WriteToken(c, t)
	return c.Redirect(303, h.app.Reverse(UserHome.String(), user.ID))
}

func (r authRouter) isDuplicate(id string) (bool, error) {
	_, err := r.svc.GetUser(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("error verifying user is not already created: %v", err)
	}
	return true, nil
}
func (h authRouter) PostSignout(c echo.Context) error {
	auth.WriteToken(c, "")
	page := mt.SignoutPage{}
	return Respond(c, "", page.Component(), CourseManagerLayout(h.app, page.Component(), domain.User{}))
}

func authorized(email string) bool {
	return isValidEmail(email)
}

func isValidEmail(email string) bool {
	return verifyCMSDomain(email) || email == "whlapinel@gmail.com"
}

func verifyCMSDomain(email string) bool {
	// Check if the email is from a CMS domain
	cmsDomain := "cms.k12.nc.us"
	return email[len(email)-len(cmsDomain):] == cmsDomain
}

type UnauthorizedAction struct {
	Message  string // message to display
	LinkName string // text for link
	URL      string // URL to provide link to
}

type UnauthReason int

const (
	DuplicateUser UnauthReason = iota
	WrongDomain
	NotRegistered
)

func UnauthAction(reason UnauthReason) UnauthorizedAction {
	return [...]UnauthorizedAction{
		{
			Message:  "User already registered. Please sign in.",
			LinkName: "Sign in",
			URL:      "/signin",
		},
		{
			Message:  "Must use a CMS account. If you have a CMS google account, please use it to sign up.",
			LinkName: "Sign up",
			URL:      "/signup",
		},
		{
			Message:  "User not registered. Please create an account.",
			LinkName: "Sign up",
			URL:      "/signup",
		},
	}[reason]
}

func (r authRouter) RespondUnauthorized(c echo.Context, reason UnauthReason) error {
	action := UnauthAction(reason)
	component := mt.UnauthorizedPage{
		Message: action.Message,
		Link: components.Link{
			Text:   action.LinkName,
			URL:    action.URL,
			Target: "#page",
		},
	}.Component()
	return Respond(c, "", component, CourseManagerLayout(r.app, component, domain.User{}))

}
