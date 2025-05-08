package home

import (
	"gh_static_portfolio/internal/app/dto"
	appcomponents "gh_static_portfolio/internal/newtemplates/components/app"
	"gh_static_portfolio/internal/shared/routes"
	web "gh_static_portfolio/internal/shared/web"
	"log"

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
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func routeHandlers(h *Handler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.GET, routes.Home, routes.GetHome, h.showHome),
	}
}

func (h Handler) showHome(c echo.Context) error {
	log.Println("showHome running")
	pageData := HomePage{
		UsersURL:  h.reverse(routes.GetUsers.String()),
		SigninURL: h.reverse(routes.GetSignin.String()),
		SignupURL: h.reverse(routes.GetSignup.String()),
	}
	template := HomePageComponent(pageData)
	layout := appcomponents.CourseManagerLayout{
		User:       dto.User{},
		HomeURL:    "/",
		SigninURL:  h.reverse(routes.GetSignin.String()),
		SignupURL:  h.reverse(routes.GetSignup.String()),
		SignoutURL: h.reverse(routes.PostSignout.String()),
		Page:       pageData.Component(),
	}.Component()
	return web.Respond(c, "", template, layout)
}
