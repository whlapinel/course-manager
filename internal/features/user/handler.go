package user

import (
	"gh_static_portfolio/internal/core"
	"net/http"

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
		core.RegisterRoute(group, handler)
	}
}

const (
	users core.RoutePath = "/users"
)

const (
	listUsers core.HandlerName = core.HandlerName(("GET: " + users))
)

var RouteHandlers = func(h *Handler) []core.RouteHandler {
	return []core.RouteHandler{
		core.NewRouteHandler(http.MethodGet, users, listUsers, h.listUsers),
	}
}

func (h Handler) listUsers(c echo.Context) error {
	return nil
}
