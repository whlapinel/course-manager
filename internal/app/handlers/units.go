package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	mt "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/labstack/echo/v4"
)

type unitHandler struct {
	service     *services.UnitService
	nodeService *services.NodeService
	reverse     web.Reverse
}

func NewUnitHandler(service *services.UnitService, nodeService *services.NodeService, reverse web.Reverse) *unitHandler {
	return &unitHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
	}
}

func RegisterUnitRoutes(group *echo.Group, h *unitHandler) error {
	for _, handler := range unitRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func unitRouteHandlers(h *unitHandler) []web.RouteHandler {
	return []web.RouteHandler{
		{Method: web.GET, RoutePath: routes.Lessons, HandlerName: routes.GetLessons, HandlerFunc: h.listLessons},
		{Method: web.GET, RoutePath: routes.Unit, HandlerName: routes.GetUnit, HandlerFunc: h.showDetails},
		{Method: web.GET, RoutePath: routes.UnitEdit, HandlerName: routes.GetEditUnit, HandlerFunc: h.showEdit},
	}
}

func (h *unitHandler) showDetails(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	component := h.nodeDetails(path, nodes).Component()
	layout := mt.BaseLayout(h.reverse, component, nodes.User.(dto.User))
	return web.Respond(c, "", component, layout)

}

func (h *unitHandler) listLessons(c echo.Context) error {
	log.Println("UnitHandler.listLessons running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	unitDTO, err := h.service.ListLessons(path.UnitID)
	if err != nil {
		return err
	}
	component := mt.NodeListPage{
		ParentNode:      unitDTO,
		Children:        unitDTO.Children(),
		ChildDetailsURL: web.URLFunc(routes.GetLesson, h.reverse, path.ToSlice()...),
		DeleteChildURL:  web.URLFunc(routes.DeleteLesson, h.reverse, path.ToSlice()...),
		ShowNewChildURL: h.reverse(routes.GetNewLesson.String(), path.ToSlice()...),
		UpNavURL:        h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		BreadCrumbsData: BreadCrumbs(nodes, path, h.reverse),
	}.Component()
	layout := mt.BaseLayout(h.reverse, component, dto.User{})
	return web.Respond(c, "", component, layout)
}

func (h *unitHandler) showEdit(c echo.Context) error {
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
	alt := nodeData.Component()
	return web.Respond(c, "", component, mt.BaseLayout(h.reverse, alt, nodes.User.(dto.User)))

}

func (h *unitHandler) nodeDetails(path routes.NodePath, nodes node.Nodes) mt.NodeDetailsPage {
	nodePage := mt.NodeDetailsPage{
		Node:              nodes.Unit,
		ParentNode:        nodes.Course,
		GetEditNodeURL:    h.reverse(routes.GetEditUnit.String(), path.ToSlice()...),
		PostEditNodeURL:   h.reverse(routes.PostEditUnit.String(), path.ToSlice()...),
		ListChildrenURL:   h.reverse(routes.GetUnits.String(), path.ToSlice()...),
		UpNavURL:          h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		CancelEditURL:     h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		CourseCalendarURL: h.reverse(routes.GetCourseCalendar.String(), path.ToSlice()...),
		ServerFilesURL:    h.reverse(routes.GetUnitFiles.String(), path.ToSlice()...),
		BreadCrumbs:       BreadCrumbs(nodes, path, h.reverse),
	}
	return nodePage
}
