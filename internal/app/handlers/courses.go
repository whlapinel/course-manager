package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/labstack/echo/v4"
)

type courseHandler struct {
	service *services.CourseService
	reverse web.Reverse
}

func NewCourseHandler(service *services.CourseService, reverse web.Reverse) *courseHandler {
	return &courseHandler{
		service: service,
		reverse: reverse,
	}
}

func RegisterCourseRoutes(group *echo.Group, h *courseHandler) error {
	for _, handler := range courseRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func courseRouteHandlers(h *courseHandler) []web.RouteHandler {
	return []web.RouteHandler{
		{Method: web.GET, RoutePath: routes.Courses, HandlerName: routes.GetCourses, HandlerFunc: h.listUnits},
	}
}

func (h *courseHandler) listUnits(c echo.Context) error {
	log.Println("CourseHandler.listUnits running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	courseDTO, err := h.service.ListUnits(path.CourseID)
	if err != nil {
		return err
	}
	nodePage := managertemplates.NodeListPage{
		ParentNode:       courseDTO,
		Children:         courseDTO.Children(),
		ChildDetailsURL:  web.URLFunc(routes.GetUnit, h.reverse, path.ToSlice()...),
		ChildChildrenURL: web.URLFunc(routes.GetLessons, h.reverse, path.ToSlice()...),
		DeleteChildURL:   web.URLFunc(routes.DeleteCourse, h.reverse, path.ToSlice()...),
		ShowNewChildURL:  h.reverse(routes.GetNewUnit.String(), path.ToSlice()...),
		UpNavURL:         h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		BreadCrumbsData: managertemplates.BreadCrumbs{
			Course:           courseDTO,
			CourseDetailsURL: h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		},
	}
	component := managertemplates.CoursesListPage{
		ShowCourseCalendarURL: web.URLFunc(routes.GetCourseCalendar, h.reverse, path.ToSlice()...),
		NodeListPage:          nodePage,
	}.Component()
	layout := managertemplates.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}
