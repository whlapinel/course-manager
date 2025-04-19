package auth

import (
	mt "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	web "gh_static_portfolio/internal/shared/web"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	reverse web.Reverse
	service Service
}

func NewHandler(service Service, e *echo.Echo) *Handler {
	return &Handler{service: service, reverse: e.Reverse}
}

func RegisterRoutes(group *echo.Group, h *Handler) error {
	for _, handler := range routeHandlers(h) {
		err := routes.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func routeHandlers(h *Handler) []routes.RouteHandler {
	return []routes.RouteHandler{
		routes.NewRouteHandler(routes.GET, routes.Signin, routes.GetSignin, h.showSignin),
		routes.NewRouteHandler(routes.POST, routes.Signin, routes.PostSignin, h.postSignin),
		routes.NewRouteHandler(routes.GET, routes.Signup, routes.GetSignup, h.showSignup),
		routes.NewRouteHandler(routes.POST, routes.Signup, routes.PostSignup, h.postSignup),
		routes.NewRouteHandler(routes.POST, routes.Signout, routes.PostSignout, h.postSignout),
	}
}

func (h Handler) showSignin(c echo.Context) error {
	page := mt.SigninPage{
		GoogleClientID:  os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleSigninURL: h.reverse(routes.PostSignin.String()),
	}
	c.Request().Header.Add("HX-Retarget", "#page")
	component := page.Component()
	layout := web.CourseManagerLayout(h.reverse, page.Component(), web.EmptyUser())
	return web.Respond(c, "", component, layout)
}

func (h Handler) postSignin(c echo.Context) error {
	log.Println("postSignin running")
	return nil
}

func (h Handler) showSignup(c echo.Context) error {
	log.Println("showSignup running")
	return nil
}
func (h Handler) postSignup(c echo.Context) error {
	log.Println("postSignup running")
	return nil
}
func (h Handler) postSignout(c echo.Context) error {
	log.Println("postSignout running")
	return nil
}
