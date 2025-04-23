package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	mt "gh_static_portfolio/internal/newtemplates/app"
	templates "gh_static_portfolio/internal/newtemplates/shared"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type termHandler struct {
	nodeService services.NodeService
	service     *services.TermService
	reverse     web.Reverse
}

func NewTermHandler(service *services.TermService, nodeService *services.NodeService, reverse web.Reverse) *termHandler {
	return &termHandler{
		service:     service,
		nodeService: *nodeService,
		reverse:     reverse,
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
		{Method: web.GET, RoutePath: routes.Term, HandlerName: routes.GetTerm, HandlerFunc: h.showDetails},
		{Method: web.GET, RoutePath: routes.TermEdit, HandlerName: routes.GetEditTerm, HandlerFunc: h.showEdit},
	}
}

func (h *termHandler) showDetails(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	nodePage := h.nodeDetails(path, nodes)
	component := mt.TermDetailsPage{
		NodeDetailsPage:      nodePage,
		ShowEditTermDatesURL: h.reverse(routes.GetTermDates.String(), path.ToSlice()...),
	}.Component()
	layout := mt.BaseLayout(h.reverse, component, nodes.User.(dto.User))
	return web.Respond(c, "", component, layout)

}

func (h *termHandler) listCourses(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	termDTO, err := h.service.ListCourses(path.TermID)
	if err != nil {
		return err
	}
	nodePage := mt.NodeListPage{
		ParentNode:       termDTO,
		Children:         termDTO.Children(),
		ChildDetailsURL:  web.URLFunc(routes.GetCourse, h.reverse, path.ToSlice()...),
		ChildChildrenURL: web.URLFunc(routes.GetUnits, h.reverse, path.ToSlice()...),
		DeleteChildURL:   web.URLFunc(routes.DeleteCourse, h.reverse, path.ToSlice()...),
		ShowNewChildURL:  h.reverse(routes.GetNewCourse.String(), path.ToSlice()...),
		UpNavURL:         h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		BreadCrumbsData:  BreadCrumbs(nodes, path, h.reverse),
	}
	component := mt.CoursesListPage{
		ShowCourseCalendarURL: web.URLFunc(routes.GetTermCalendar, h.reverse, path.ToSlice()...),
		ShowAssessmentsURL:    web.URLFunc(routes.GetCourseAssessments, h.reverse, path.ToSlice()...),
		NodeListPage:          nodePage,
	}.Component()
	layout := mt.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}

func (h *termHandler) showEdit(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	nodeData := h.nodeDetails(path, nodes)
	nodeData.IsEdit = true
	component := mt.TermDetailsPage{
		NodeDetailsPage:      nodeData,
		ShowEditTermDatesURL: h.reverse(routes.GetTermDates.String(), path.ToSlice()...),
	}.DetailsFormComponent(true)
	alt := nodeData.Component()
	return web.Respond(c, "", component, mt.BaseLayout(h.reverse, alt, nodes.User.(dto.User)))

}

func (h *termHandler) nodeDetails(path routes.NodePath, nodes templates.Nodes) mt.NodeDetailsPage {
	nodePage := mt.NodeDetailsPage{
		Node:              nodes.Term,
		ParentNode:        nodes.User,
		GetEditNodeURL:    h.reverse(routes.GetEditTerm.String(), path.ToSlice()...),
		PostEditNodeURL:   h.reverse(routes.PostEditTerm.String(), path.ToSlice()...),
		ListChildrenURL:   h.reverse(routes.GetCourses.String(), path.ToSlice()...),
		UpNavURL:          h.reverse(routes.GetUser.String(), path.ToSlice()...),
		CancelEditURL:     h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		CourseCalendarURL: h.reverse(routes.GetCourseCalendar.String(), path.ToSlice()...),
		ServerFilesURL:    h.reverse(routes.GetTermFiles.String(), path.ToSlice()...),
		BreadCrumbs:       BreadCrumbs(nodes, path, h.reverse),
	}
	return nodePage
}
