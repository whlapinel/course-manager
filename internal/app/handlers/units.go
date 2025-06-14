package handlers

import (
	"fmt"
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/features/unit"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type unitHandler struct {
	service     *services.UnitService
	nodeService *services.NodeService
	files       *services.FileService
	markdown    *services.MarkdownService
	reverse     web.Reverse
	*baseHandler[dto.Unit, int, int]
}

func NewUnitHandler(
	service *services.UnitService,
	nodeService *services.NodeService,
	files *services.FileService,
	markdown *services.MarkdownService,
	previewer ports.SiteGenerator,
	reverse web.Reverse,
) *unitHandler {
	return &unitHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
		files:       files,
		markdown:    markdown,
		baseHandler: &baseHandler[dto.Unit, int, int]{
			service:          service,
			files:            files,
			markdown:         markdown,
			nodes:            nodeService,
			reverse:          reverse,
			previewer:        previewer,
			getNode:          routes.GetUnit,
			viewNodeFile:     routes.ViewUnitFile,
			deleteNodeFile:   routes.DeleteUnitFile,
			getNodeFile:      routes.GetUnitFile,
			getNodeFiles:     routes.GetUnitFiles,
			getNodeEditFile:  routes.GetUnitEditFile,
			postNodeFile:     routes.PostUnitFile,
			postNodeEditFile: routes.PostUnitEditFile,
		},
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
		// base handler
		web.NewRouteHandler(web.GET, routes.UnitFiles, routes.GetUnitFiles, h.showFiles),
		web.NewRouteHandler(web.POST, routes.UnitFiles, routes.PostUnitFile, h.postUploadedFile),
		web.NewRouteHandler(web.GET, routes.UnitEditFile, routes.GetUnitEditFile, h.showEditFile),
		web.NewRouteHandler(web.POST, routes.UnitEditFile, routes.PostUnitEditFile, h.postEditMarkdown),
		web.NewRouteHandler(web.GET, routes.UnitViewFile, routes.ViewUnitFile, h.viewMarkdown),
		web.NewRouteHandler(web.DELETE, routes.UnitFile, routes.DeleteUnitFile, h.deleteFile),

		// overrides
		web.NewRouteHandler(web.GET, routes.Units, routes.GetUnits, h.listByCourse),
		web.NewRouteHandler(web.GET, routes.Unit, routes.GetUnit, h.showDetails),
		web.NewRouteHandler(web.GET, routes.UnitEdit, routes.GetEditUnit, h.showEdit),
		web.NewRouteHandler(web.POST, routes.UnitEdit, routes.PostEditUnit, h.postEdit),
		web.NewRouteHandler(web.GET, routes.NewUnit, routes.GetNewUnit, h.showCreateNew),
		web.NewRouteHandler(web.POST, routes.NewUnit, routes.PostNewUnit, h.postNew),
		web.NewRouteHandler(web.DELETE, routes.Unit, routes.DeleteUnit, h.delete),
	}
}

func (h *unitHandler) delete(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	return h.service.Delete(nodePath.UnitID)

}
func (h *unitHandler) postNew(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	unit := dto.Unit{
		Unit: unit.Unit{
			BaseNode: ports.BaseNode[int, int]{
				ParentID: nodePath.CourseID,
			},
		},
	}
	unit.Name = c.FormValue("name")
	unit.Description = c.FormValue("description")
	numParam := c.FormValue("number")
	unit.Number, err = strconv.Atoi(numParam)
	if err != nil {
		return err
	}
	err = h.service.Save(unit)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetUnits.String(), nodePath.ToSlice()...))
}

func (h *unitHandler) showCreateNew(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	log.Println("unitHandler.showCreateNew running")
	if info.NodePath.CourseID == 0 {
		return fmt.Errorf("course ID is 0")
	}
	if info.Nodes.Course == nil {
		return fmt.Errorf("course is nil")
	}
	page := ac.NodeCreatePage{
		ParentNode:          info.Course,
		NodeType:            dto.UnitTypeName,
		Params:              info.NodePath,
		PostCreateNodeURL:   h.reverse(routes.PostNewUnit.String(), info.NodePath.ToSlice()...),
		CancelURL:           h.reverse(routes.GetUnits.String(), info.NodePath.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(info.Nodes, info.NodePath, h.reverse),
		CourseManagerLayout: BaseLayout3(h.reverse, info.User.(dto.User)),
	}
	return Respond(c, page)
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
	units, err := h.service.ByParentID(path.CourseID)
	if err != nil {
		return err
	}
	course.Units = units
	page := ac.NodeListPage{
		ParentNode:          course,
		Children:            course.GetChildren(),
		ChildDetailsURL:     web.URLFunc(routes.GetUnit, h.reverse, path.ToSlice()...),
		DeleteChildURL:      web.URLFunc(routes.DeleteUnit, h.reverse, path.ToSlice()...),
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
			unit.Sequence = num
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

func (h *unitHandler) nodeDetails(path routes.NodePath, nodes ports.Nodes) ac.NodeDetailsPage {
	lastName := strings.ToLower(nodes.User.(dto.User).LastName)
	nodePage := ac.NodeDetailsPage{
		GenerateSiteURL:     h.reverse(routes.PostGenerateSite.String(), path.ToSlice()...),
		StaticSiteURL:       h.previewer.StaticSiteURL(lastName, path.CourseID),
		Node:                nodes.Unit,
		ParentNode:          nodes.Course,
		GetEditNodeURL:      h.reverse(routes.GetEditUnit.String(), path.ToSlice()...),
		PostEditNodeURL:     h.reverse(routes.PostEditUnit.String(), path.ToSlice()...),
		ListChildrenURL:     h.reverse(routes.GetLessons.String(), path.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetUnits.String(), path.ToSlice()...),
		CancelEditURL:       h.reverse(routes.GetUnit.String(), path.ToSlice()...),
		CourseCalendarURL:   h.reverse(routes.GetCourseCalendar.String(), path.ToSlice()...),
		ServerFilesURL:      h.reverse(routes.GetUnitFiles.String(), path.ToSlice()...),
		BreadCrumbs:         BreadCrumbs(nodes, path, h.reverse),
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return nodePage
}
