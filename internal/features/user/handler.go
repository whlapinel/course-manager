package user

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	reverse web.Reverse
	service *Service
}

func NewHandler(service *Service, e *echo.Echo) *Handler {
	return &Handler{service: service, reverse: e.Reverse}
}

func RegisterRoutes(group *echo.Group, h *Handler) {
	for _, handler := range RouteHandlers(h) {
		web.RegisterRoute(group, handler)
	}
}

func RouteHandlers(h *Handler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.GET, routes.Users, routes.GetUsers, h.redirectToDashboard),
		web.NewRouteHandler(web.GET, routes.User, routes.GetUser, h.showDashboard),
	}
}

// this reads the user's id from the context and redirects to the user's dashboard
func (h *Handler) redirectToDashboard(c echo.Context) error {
	userID := c.Get("id")
	if userID == nil {
		return fmt.Errorf("userID is nil")
	}
	dashboardURL := h.reverse(routes.GetUser.String(), userID)
	return c.Redirect(302, dashboardURL)
}

func (h *Handler) showDashboard(c echo.Context) error {
	log.Println("showDashboard running")
	params, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	log.Println("params.UserID", params.UserID)
	user, err := h.service.ByID(params.UserID)
	if err != nil {
		return err
	}
	userDTO := dto.User{
		User: user,
	}
	page := managertemplates.UserHomePage{
		GenerateSiteURL: h.reverse(routes.PostGenerateSite.String(), params.ToSlice()...),
		ListTermsURL:    h.reverse(routes.GetTerms.String(), params.ToSlice()...),
		User:            userDTO,
	}
	component := page.Component()
	layout := managertemplates.BaseLayout(h.reverse, component, user)
	return web.Respond(c, "", component, layout)
}
