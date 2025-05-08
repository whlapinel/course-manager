package handlers

import (
	mt "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	lessonviews "gh_static_portfolio/internal/app/views/lesson"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type lessonHandler struct {
	nodeService *services.NodeService
	service     *services.LessonService
	reverse     web.Reverse
}

func NewLessonHandler(service *services.LessonService, nodeService *services.NodeService, reverse web.Reverse) *lessonHandler {
	return &lessonHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
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
		{Method: web.GET, RoutePath: routes.Lessons, HandlerName: routes.GetLessons, HandlerFunc: h.listByUnit},
		{Method: web.GET, RoutePath: routes.Lesson, HandlerName: routes.GetLesson, HandlerFunc: h.showDetails},
		{Method: web.GET, RoutePath: routes.LessonEdit, HandlerName: routes.GetEditLesson, HandlerFunc: h.showEdit},
		web.NewRouteHandler(web.GET, routes.LessonEdit, routes.PostEditLesson, h.postEdit),
	}
}

func (h *lessonHandler) listByUnit(c echo.Context) error {
	log.Println("UnitHandler.listLessons running...")
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	unit := nodes.Unit.(dto.Unit)
	lessons, err := h.service.ByUnitID(path.UnitID)
	if err != nil {
		return err
	}
	unit.Lessons = lessons
	page := mt.NodeListPage{
		ParentNode:          unit,
		Children:            unit.Children(),
		ChildDetailsURL:     web.URLFunc(routes.GetLesson, h.reverse, path.ToSlice()...),
		DeleteChildURL:      web.URLFunc(routes.DeleteLesson, h.reverse, path.ToSlice()...),
		ShowNewChildURL:     h.reverse(routes.GetNewLesson.String(), path.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(nodes, path, h.reverse),
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)

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
	page := lessonviews.LessonDetailsPage{
		NodeDetailsPage: nodeData,
		AssetsURLFunc:   web.AssetsURLFunc,
		ViewMarkdownURL: web.URLFunc(web.HandlerName(routes.LessonViewFile), h.reverse),
	}
	return Respond(c, page)
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
	lessonData := lessonviews.LessonDetailsPage{
		NodeDetailsPage: nodeData,
		AssetsURLFunc:   web.AssetsURLFunc,
		ViewMarkdownURL: web.URLFunc(web.HandlerName(routes.LessonViewFile), h.reverse),
	}
	page := lessonData.EditDetails()
	return Respond(c, page)
}

func (h *lessonHandler) postEdit(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(nodePath)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	lesson := nodes.Lesson.(dto.Lesson)
	form := c.Request().Form
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "number":
			num, err := strconv.Atoi(val[0])
			if err != nil {
				return err
			}
			lesson.Number = num
		case "name":
			lesson.Name = val[0]
		case "description":
			lesson.Description = val[0]
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}
	}
	err = h.service.Update(lesson)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetUnit.String(), nodePath.ToSlice()...))
}

func (h *lessonHandler) nodeDetails(path routes.NodePath, nodes node.Nodes) mt.NodeDetailsPage {
	nodePage := mt.NodeDetailsPage{
		Node:                nodes.Lesson,
		ParentNode:          nodes.Unit,
		GetEditNodeURL:      h.reverse(routes.GetEditLesson.String(), path.ToSlice()...),
		PostEditNodeURL:     h.reverse(routes.PostEditLesson.String(), path.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		CancelEditURL:       h.reverse(routes.GetLesson.String(), path.ToSlice()...),
		BreadCrumbs:         BreadCrumbs(nodes, path, h.reverse),
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return nodePage
}
