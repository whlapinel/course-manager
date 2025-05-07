package handlers

import (
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	service     *services.UserService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewUserHandler(service *services.UserService, nodeService *services.NodeService, reverse web.Reverse) *userHandler {
	return &userHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
	}
}

func RegisterUserRoutes(group *echo.Group, h *userHandler) error {
	for _, handler := range userRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func userRouteHandlers(h *userHandler) []web.RouteHandler {
	return []web.RouteHandler{
		web.NewRouteHandler(web.POST, routes.GenerateSite, routes.PostGenerateSite, h.generateSite),
	}
}

func (h *userHandler) generateSite(c echo.Context) error {
	panic("not implemented")
}
