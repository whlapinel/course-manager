package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	mt "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type courseHandler struct {
	service     *services.CourseService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewCourseHandler(service *services.CourseService, nodeService *services.NodeService, reverse web.Reverse) *courseHandler {
	return &courseHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
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
		{Method: web.GET, RoutePath: routes.Course, HandlerName: routes.GetCourse, HandlerFunc: h.showDetails},
		{Method: web.GET, RoutePath: routes.Units, HandlerName: routes.GetUnits, HandlerFunc: h.listUnits},
		{Method: web.GET, RoutePath: routes.CourseEdit, HandlerName: routes.GetEditCourse, HandlerFunc: h.showEdit},
	}
}

func (h *courseHandler) showDetails(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	nodePage := h.nodeDetails(path, nodes)
	component := mt.CourseDetailsPage{
		NodeDetailsPage:          nodePage,
		GetCopyCourseURL:         "please-create-me",
		PostSelectStandardSetURL: "please-create-me",
	}.Component()
	layout := mt.BaseLayout(h.reverse, component, nodes.User.(dto.User))
	return web.Respond(c, "", component, layout)
}

func (h *courseHandler) listUnits(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	courseDTO, err := h.service.ListUnits(path.CourseID)
	if err != nil {
		return err
	}
	component := mt.NodeListPage{
		ParentNode:       courseDTO,
		Children:         courseDTO.Children(),
		ChildDetailsURL:  web.URLFunc(routes.GetUnit, h.reverse, path.ToSlice()...),
		ChildChildrenURL: web.URLFunc(routes.GetLessons, h.reverse, path.ToSlice()...),
		DeleteChildURL:   web.URLFunc(routes.DeleteCourse, h.reverse, path.ToSlice()...),
		ShowNewChildURL:  h.reverse(routes.GetNewUnit.String(), path.ToSlice()...),
		UpNavURL:         h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		BreadCrumbsData:  BreadCrumbs(nodes, path, h.reverse),
	}.Component()
	layout := mt.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}

func (h *courseHandler) showEdit(c echo.Context) error {
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
	component := nodeData.DetailsFormComponent(true)
	alt := mt.CourseDetailsPage{
		NodeDetailsPage:          nodeData,
		GetCopyCourseURL:         "please-create-me",
		PostSelectStandardSetURL: "please-create-me",
	}.Component()
	return web.Respond(c, "", component, mt.BaseLayout(h.reverse, alt, nodes.User.(dto.User)))

}

func (h *courseHandler) nodeDetails(path routes.NodePath, nodes node.Nodes) mt.NodeDetailsPage {
	nodeData := mt.NodeDetailsPage{
		Node:            nodes.Course,
		ParentNode:      nodes.Term,
		GetEditNodeURL:  h.reverse(routes.GetEditCourse.String(), path.ToSlice()...),
		PostEditNodeURL: h.reverse(routes.PostEditCourse.String(), path.ToSlice()...),
		UpNavURL:        h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		CancelEditURL:   h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		BreadCrumbs:     BreadCrumbs(nodes, path, h.reverse),
		IsEdit:          true,
	}
	return nodeData
}
