package auth

import (
	"database/sql"
	"errors"
	"gh_static_portfolio/internal/core/user"
	authentication "gh_static_portfolio/internal/newauthentication"
	mt "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	web "gh_static_portfolio/internal/shared/web"
	components "gh_static_portfolio/internal/templates/components/base"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	reverse web.Reverse
	service *Service
}

func NewHandler(service *Service, e *echo.Echo) *Handler {
	return &Handler{service: service, reverse: e.Reverse}
}

func RegisterRoutes(group *echo.Group, h *Handler) error {
	for _, handler := range routeHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func routeHandlers(h *Handler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.GET, routes.Signin, routes.GetSignin, h.showSignin),
		web.NewRouteHandler(web.POST, routes.Signin, routes.PostSignin, h.postSignin),
		web.NewRouteHandler(web.GET, routes.Signup, routes.GetSignup, h.showSignup),
		web.NewRouteHandler(web.POST, routes.Signup, routes.PostSignup, h.postSignup),
		web.NewRouteHandler(web.POST, routes.Signout, routes.PostSignout, h.postSignout),
	}
}

func (h *Handler) showSignin(c echo.Context) error {
	page := mt.SigninPage{
		GoogleClientID:  os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleSigninURL: h.reverse(routes.PostSignin.String()),
	}
	c.Request().Header.Add("HX-Retarget", "#page")
	component := page.Component()
	layout := mt.BaseLayout(h.reverse, page.Component(), user.User{})
	return web.Respond(c, "", component, layout)
}

func (h *Handler) postSignin(c echo.Context) error {
	payload, err := authentication.GoogleAuth(c)
	if err != nil {
		log.Println(err)
		return c.String(500, "Failed to authenticate")
	}
	sub := payload.Claims["sub"].(string)
	currUser, err := h.service.userQuery.ByID(sub)
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

func (h *Handler) showSignup(c echo.Context) error {
	log.Println("showSignup running")
	return nil
}
func (h *Handler) postSignup(c echo.Context) error {
	log.Println("postSignup running")
	return nil
}
func (h *Handler) postSignout(c echo.Context) error {
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

func (h *Handler) RespondUnauthorized(c echo.Context, reason UnauthReason) error {
	action := UnauthAction(reason)
	component := mt.UnauthorizedPage{
		Message: action.Message,
		Link: components.Link{
			Text:   action.LinkName,
			URL:    action.URL,
			Target: "#page",
		},
	}.Component()
	return web.Respond(c, "", component, mt.BaseLayout(h.reverse, component, user.User{}))
}
