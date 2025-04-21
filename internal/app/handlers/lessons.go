package handlers

import (
	"gh_static_portfolio/internal/app/services"
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

type lessonHandler struct {
	nodeService   *services.NodeService
	lessonService *services.LessonService
	reverse       web.Reverse
}

func NewLessonHandler(service *services.LessonService, nodes *services.NodeService, reverse web.Reverse) *lessonHandler {
	return &lessonHandler{
		lessonService: service,
		nodeService:   nodes,
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
	lesson, err := h.lessonService.ByID(path.LessonID)
	if err != nil {
		return err
	}
	nodeData := managertemplates.NodeDetailsPage{
		Node:            lesson,
		ParentNode:      nodes.Unit,
		GetEditNodeURL:  h.reverse(routes.GetEditLesson.String(), path.ToSlice()...),
		PostEditNodeURL: h.reverse(routes.PostEditLesson.String(), path.ToSlice()...),
		UpNavURL:        h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		CancelEditURL:   h.reverse(routes.GetLesson.String(), path.ToSlice()...),
		IsEdit:          false,
		BreadCrumbsData: managertemplates.BreadCrumbs{
			Nodes: nodes,
		},
	}
	component := managertemplates.LessonDetailsPage{}
	panic("not implemented")
}
