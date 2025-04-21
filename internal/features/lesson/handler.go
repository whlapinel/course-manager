package lesson

import (
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	reverse web.Reverse
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
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
	return []web.RouteHandler{}
}
