package handlers

import (
	"database/sql"
	"errors"
	ac "gh_static_portfolio/internal/app/components"
	appcomponents "gh_static_portfolio/internal/app/components"
	av "gh_static_portfolio/internal/app/views/authentication"
	"gh_static_portfolio/internal/core/user"
	"gh_static_portfolio/internal/features/auth"
	authentication "gh_static_portfolio/internal/newauthentication"
	components "gh_static_portfolio/internal/newtemplates/components/base"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	reverse web.Reverse
	service *auth.Service
}

func NewAuthHandler(service *auth.Service, reverse web.Reverse) *authHandler {
	return &authHandler{service: service, reverse: reverse}
}

func RegisterAuthRoutes(group *echo.Group, h *authHandler) error {
	for _, handler := range routeHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func routeHandlers(h *authHandler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.GET, routes.Signin, routes.GetSignin, h.showSignin),
		web.NewRouteHandler(web.POST, routes.Signin, routes.PostSignin, h.postSignin),
		web.NewRouteHandler(web.GET, routes.Signup, routes.GetSignup, h.showSignup),
		web.NewRouteHandler(web.POST, routes.Signup, routes.PostSignup, h.postSignup),
		web.NewRouteHandler(web.POST, routes.Signout, routes.PostSignout, h.postSignout),
	}
}

func (h *authHandler) showSignin(c echo.Context) error {
	signInPage := av.SigninPage{
		GoogleClientID:      os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleSigninURL:     h.reverse(routes.PostSignin.String()),
		CourseManagerLayout: BaseLayoutWithoutUser(h.reverse),
	}
	hxPage := ac.BasicHTMXPage[av.SigninPage]{
		Page: signInPage,
	}
	c.Request().Header.Add("HX-Retarget", "#page")
	return Respond(c, hxPage)
}

func (h *authHandler) postSignin(c echo.Context) error {
	payload, err := authentication.GoogleAuth(c)
	if err != nil {
		log.Println(err)
		return c.String(500, "Failed to authenticate")
	}
	sub := payload.Claims["sub"].(string)
	currUser, err := h.service.GetUser(sub)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return h.RespondUnauthorized(c, NotRegistered)
		}
		log.Println(err)
		return c.String(500, err.Error())
	}
	log.Println("userID", currUser.ID)
	t, err := authentication.IssueToken(authentication.TokenParams{User: user.User{
		ID:        sub,
		Email:     currUser.Email,
		FirstName: currUser.FirstName,
		LastName:  currUser.LastName,
		Picture:   currUser.Picture,
	}})
	if err != nil {
		return c.String(500, "Failed to issue token")
	}
	authentication.WriteToken(c, t)
	return c.Redirect(303, "/users")
}

func (h *authHandler) showSignup(c echo.Context) error {
	log.Println("showSignup running")
	return nil
}
func (h *authHandler) postSignup(c echo.Context) error {
	log.Println("postSignup running")
	return nil
}
func (h *authHandler) postSignout(c echo.Context) error {
	log.Println("postSignout running")
	return nil
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

func (h *authHandler) RespondUnauthorized(c echo.Context, reason UnauthReason) error {
	action := UnauthAction(reason)
	page := av.UnauthorizedPage{
		Message: action.Message,
		Link: components.Link{
			Text:   action.LinkName,
			URL:    action.URL,
			Target: "#page",
		},
	}
	htmxPage := appcomponents.BasicHTMXPage[av.UnauthorizedPage]{
		Page: page,
	}
	return Respond(c, htmxPage)
}
