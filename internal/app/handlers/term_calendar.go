package handlers

import (
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type termCalendarHandler struct {
	service     *services.TermCalendarService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewTermCalHandler(svc *services.TermCalendarService, nodeService *services.NodeService, reverse web.Reverse) *termCalendarHandler {
	return &termCalendarHandler{
		service:     svc,
		nodeService: nodeService,
		reverse:     reverse,
	}
}

func RegisterTermCalRoutes(group *echo.Group, h *termCalendarHandler) error {
	for _, handler := range termCalRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func termCalRouteHandlers(h *termCalendarHandler) []web.RouteHandler {
	return []web.RouteHandler{}
}
