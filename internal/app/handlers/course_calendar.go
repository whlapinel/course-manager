package handlers

import (
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type courseCalendarHandler struct {
	service     *services.CourseCalendarService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewCourseCalHandler(svc *services.CourseCalendarService, nodeService *services.NodeService, reverse web.Reverse) *courseCalendarHandler {
	return &courseCalendarHandler{
		service:     svc,
		nodeService: nodeService,
		reverse:     reverse,
	}
}

func RegisterCourseCalRoutes(group *echo.Group, h *courseCalendarHandler) error {
	for _, handler := range courseCalRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func courseCalRouteHandlers(h *courseCalendarHandler) []web.RouteHandler {
	return []web.RouteHandler{}
}
