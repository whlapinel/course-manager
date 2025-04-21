package term

import (
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	reverse Reverse
	service *Service
}

type Reverse func(name string, params ...any) string

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func RegisterRoutes(group *echo.Group, h *Handler) {
	for _, handler := range routeHandlers(h) {
		web.RegisterRoute(group, handler)
	}
}

func routeHandlers(h *Handler) []web.RouteHandler {
	return []web.RouteHandler{}
}
