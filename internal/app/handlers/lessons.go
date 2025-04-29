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

type lessonHandler struct {
	nodeService   *services.NodeService
	lessonService *services.LessonService
	reverse       web.Reverse
}

func NewLessonHandler(service *services.LessonService, nodeService *services.NodeService, reverse web.Reverse) *lessonHandler {
	return &lessonHandler{
		lessonService: service,
		nodeService:   nodeService,
		reverse:       reverse,
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
		{Method: web.GET, RoutePath: routes.LessonEdit, HandlerName: routes.GetEditLesson, HandlerFunc: h.showEdit},
	}
}

func (h *lessonHandler) showDetails(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	nodeData := h.nodeDetails(path, nodes)
	component := mt.LessonDetailsPage{
		NodeDetailsPage: nodeData,
		AssetsURLFunc:   web.AssetsURLFunc,
		ViewMarkdownURL: web.URLFunc(web.HandlerName(routes.LessonViewFile), h.reverse),
	}.Component()
	return web.Respond(c, "", component, mt.BaseLayout(h.reverse, component, nodes.User.(dto.User)))

}

func (h *lessonHandler) showEdit(c echo.Context) error {
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
	lessonPage := mt.LessonDetailsPage{
		NodeDetailsPage: nodeData,
		AssetsURLFunc:   web.AssetsURLFunc,
		ViewMarkdownURL: web.URLFunc(web.HandlerName(routes.LessonViewFile), h.reverse),
	}
	component := nodeData.DetailsFormComponent(true)
	altComponent := lessonPage.Component()
	return web.Respond(c, "", component, mt.BaseLayout(h.reverse, altComponent, nodes.User.(dto.User)))

}

func (h *lessonHandler) nodeDetails(path routes.NodePath, nodes node.Nodes) mt.NodeDetailsPage {
	nodePage := mt.NodeDetailsPage{
		Node:            nodes.Lesson,
		ParentNode:      nodes.Unit,
		GetEditNodeURL:  h.reverse(routes.GetEditLesson.String(), path.ToSlice()...),
		PostEditNodeURL: h.reverse(routes.PostEditLesson.String(), path.ToSlice()...),
		UpNavURL:        h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		CancelEditURL:   h.reverse(routes.GetLesson.String(), path.ToSlice()...),
		BreadCrumbs:     BreadCrumbs(nodes, path, h.reverse),
	}
	return nodePage
}
