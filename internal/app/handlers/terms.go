package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	templates "gh_static_portfolio/internal/newtemplates/shared"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/labstack/echo/v4"
)

type termHandler struct {
	service *services.TermService
	reverse web.Reverse
}

func NewTermHandler(service *services.TermService, reverse web.Reverse) *termHandler {
	return &termHandler{
		service: service,
		reverse: reverse,
	}
}

func RegisterTermRoutes(group *echo.Group, h *termHandler) error {
	for _, handler := range termRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func termRouteHandlers(h *termHandler) []web.RouteHandler {
	return []web.RouteHandler{
		{Method: web.GET, RoutePath: routes.Courses, HandlerName: routes.GetCourses, HandlerFunc: h.listCourses},
	}
}

func (h *termHandler) listCourses(c echo.Context) error {
	log.Println("TermHandler.listCourses running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	termDTO, err := h.service.ListCourses(path.TermID)
	if err != nil {
		return err
	}
	nodePage := managertemplates.NodeListPage{
		ParentNode:       termDTO,
		Children:         termDTO.Children(),
		ChildDetailsURL:  web.URLFunc(routes.GetCourse, h.reverse, path.ToSlice()...),
		ChildChildrenURL: web.URLFunc(routes.GetUnits, h.reverse, path.ToSlice()...),
		DeleteChildURL:   web.URLFunc(routes.DeleteCourse, h.reverse, path.ToSlice()...),
		ShowNewChildURL:  h.reverse(routes.GetNewCourse.String(), path.ToSlice()...),
		UpNavURL:         h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		BreadCrumbsData: managertemplates.BreadCrumbs{
			Nodes: templates.Nodes{
				Term: termDTO,
			},
			TermDetailsURL: h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		},
	}
	component := managertemplates.CoursesListPage{
		ShowCourseCalendarURL: web.URLFunc(routes.GetTermCalendar, h.reverse, path.ToSlice()...),
		ShowAssessmentsURL:    web.URLFunc(routes.GetCourseAssessments, h.reverse, path.ToSlice()...),
		NodeListPage:          nodePage,
	}.Component()
	layout := managertemplates.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}
