package user

import (
	"gh_static_portfolio/internal/shared/routes"
	"log"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	reverse Reverse
	service Service
}

type Reverse func(name string, params ...any) string

func NewHandler(service Service, e *echo.Echo) *Handler {
	return &Handler{service: service, reverse: e.Reverse}
}

func RegisterRoutes(group *echo.Group, h *Handler) {
	for _, handler := range RouteHandlers(h) {
		routes.RegisterRoute(group, handler)
	}
}

func RouteHandlers(h *Handler) []routes.RouteHandler {
	return []routes.RouteHandler{
		routes.NewRouteHandler(routes.GET, routes.Users, routes.GetUsers, h.showDashboard),
	}
}

func (h *Handler) showDashboard(c echo.Context) error {
	log.Println("showDashboard running")
	return nil
}
