package handlers

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	mt "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"strconv"

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
		{Method: web.GET, RoutePath: routes.Units, HandlerName: routes.GetUnits, HandlerFunc: h.listByCourse},
		{Method: web.GET, RoutePath: routes.Unit, HandlerName: routes.GetUnit, HandlerFunc: h.showDetails},
		{Method: web.GET, RoutePath: routes.UnitEdit, HandlerName: routes.GetEditUnit, HandlerFunc: h.showEdit},
		web.NewRouteHandler(web.GET, routes.UnitEdit, routes.PostEditUnit, h.postEdit),
	}
}

func (h *unitHandler) listByCourse(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	course := nodes.Course.(dto.Course)
	units, err := h.service.ByCourseID(path.CourseID)
	if err != nil {
		return err
	}
	course.Units = units
	page := appcomponents.NodeListPage{
		ParentNode:          course,
		Children:            course.Children(),
		ChildDetailsURL:     web.URLFunc(routes.GetUnit, h.reverse, path.ToSlice()...),
		ChildChildrenURL:    web.URLFunc(routes.GetLessons, h.reverse, path.ToSlice()...),
		DeleteChildURL:      web.URLFunc(routes.DeleteCourse, h.reverse, path.ToSlice()...),
		ShowNewChildURL:     h.reverse(routes.GetNewUnit.String(), path.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(nodes, path, h.reverse),
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)

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
	page := h.nodeDetails(path, nodes)
	return Respond(c, page)
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
	page := nodeData.DetailsEdit()
	return Respond(c, page)
}

func (h *unitHandler) postEdit(c echo.Context) error {
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
	unit := nodes.Unit.(dto.Unit)
	form := c.Request().Form
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "sequence":
			num, err := strconv.Atoi(val[0])
			if err != nil {
				return err
			}
			unit.SequenceNum = num
		case "number":
			num, err := strconv.Atoi(val[0])
			if err != nil {
				return err
			}
			unit.Number = num
		case "name":
			unit.Name = val[0]
		case "description":
			unit.Description = val[0]
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}
	}
	err = h.service.Update(unit)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetUnit.String(), nodePath.ToSlice()...))
}

func (h *unitHandler) nodeDetails(path routes.NodePath, nodes ports.Nodes) mt.NodeDetailsPage {
	nodePage := mt.NodeDetailsPage{
		Node:                nodes.Unit,
		ParentNode:          nodes.Course,
		GetEditNodeURL:      h.reverse(routes.GetEditUnit.String(), path.ToSlice()...),
		PostEditNodeURL:     h.reverse(routes.PostEditUnit.String(), path.ToSlice()...),
		ListChildrenURL:     h.reverse(routes.GetUnits.String(), path.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		CancelEditURL:       h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		CourseCalendarURL:   h.reverse(routes.GetCourseCalendar.String(), path.ToSlice()...),
		ServerFilesURL:      h.reverse(routes.GetUnitFiles.String(), path.ToSlice()...),
		BreadCrumbs:         BreadCrumbs(nodes, path, h.reverse),
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return nodePage
}
