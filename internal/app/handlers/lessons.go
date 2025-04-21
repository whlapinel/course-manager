package handlers

import (
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type lessonHandler struct {
	service *services.LessonService
	reverse web.Reverse
}

func NewLessonHandler(service *services.LessonService, reverse web.Reverse) *lessonHandler {
	return &lessonHandler{
		service: service,
		reverse: reverse,
	}
}

func RegisterLessonRoutes(group *echo.Group, h *lessonHandler) error {
	for _, handler := range lessonRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func lessonRouteHandlers(h *lessonHandler) []web.RouteHandler {
	return []web.RouteHandler{
		{Method: web.GET, RoutePath: routes.Lesson, HandlerName: routes.GetLesson, HandlerFunc: h.showDetails},
	}
}

func (h *lessonHandler) showDetails(c echo.Context) error {
	panic("not implemented")
}
