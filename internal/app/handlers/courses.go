package handlers

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	mt "gh_static_portfolio/internal/app/views/course"
	fileviews "gh_static_portfolio/internal/app/views/files"
	"gh_static_portfolio/internal/features/files"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type courseHandler struct {
	service     *services.CourseService
	nodeService *services.NodeService
	fileService *files.Service
	reverse     web.Reverse
}

func NewCourseHandler(
	service *services.CourseService,
	nodeService *services.NodeService,
	fileService *files.Service,
	reverse web.Reverse,
) *courseHandler {
	return &courseHandler{
		service:     service,
		nodeService: nodeService,
		fileService: fileService,
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
		web.NewRouteHandler(web.GET, routes.Courses, routes.GetCourses, h.listByTerm),
		{Method: web.GET, RoutePath: routes.Course, HandlerName: routes.GetCourse, HandlerFunc: h.showDetails},
		{Method: web.GET, RoutePath: routes.CourseEdit, HandlerName: routes.GetEditCourse, HandlerFunc: h.showEdit},
		{Method: web.GET, RoutePath: routes.CourseFiles, HandlerName: routes.GetCourseFiles, HandlerFunc: h.showFiles},
		web.NewRouteHandler(web.POST, routes.CourseEdit, routes.PostEditCourse, h.postEdit),
	}
}

func (h *courseHandler) listByTerm(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	term := nodes.Term.(dto.Term)

	courses, err := h.service.ListByTerm(path.TermID)
	if err != nil {
		return err
	}
	term.Courses = courses
	nodePage := appcomponents.NodeListPage{
		ParentNode:       term,
		Children:         term.Children(),
		ChildDetailsURL:  web.URLFunc(routes.GetCourse, h.reverse, path.ToSlice()...),
		ChildChildrenURL: web.URLFunc(routes.GetUnits, h.reverse, path.ToSlice()...),
		DeleteChildURL:   web.URLFunc(routes.DeleteCourse, h.reverse, path.ToSlice()...),
		ShowNewChildURL:  h.reverse(routes.GetNewCourse.String(), path.ToSlice()...),
		UpNavURL:         h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		BreadCrumbsData:  BreadCrumbs(nodes, path, h.reverse),
	}
	page := mt.CoursesListPage{
		ShowCourseCalendarURL: web.URLFunc(routes.GetCourseCalendar, h.reverse, path.ToSlice()...),
		ShowAssessmentsURL:    web.URLFunc(routes.GetCourseAssessments, h.reverse, path.ToSlice()...),
		NodeListPage:          nodePage,
		CourseManagerLayout:   BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)

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
	nodePage.IsEdit = false
	page := mt.CourseDetailsPage{
		NodeDetailsPage:          nodePage,
		GetCopyCourseURL:         "please-create-me",
		PostSelectStandardSetURL: "please-create-me",
		CourseManagerLayout:      BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)
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
	page := nodeData.DetailsEdit()
	return Respond(c, page)
}

func (h *courseHandler) postEdit(c echo.Context) error {
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
	course := nodes.Course.(dto.Course)
	form := c.Request().Form
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "name":
			course.Name = val[0]
		case "description":
			course.Description = val[0]
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}
	}
	err = h.service.Update(course)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetCourse.String(), nodePath.ToSlice()...))
}

func (h *courseHandler) showFiles(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	log.Println("filePath", filePath)
	nodes, err := h.nodeService.Nodes(nodePath)
	if err != nil {
		return err
	}
	if filePath == "" || filePath == "*" {
		filePath = "."
	}
	fileInfo, err := h.fileService.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir {
		c.Attachment(fileInfo.AbsPath, filepath.Base(fileInfo.AbsPath))
	}
	parentPath := filepath.Dir(filePath)
	files, err := h.fileService.NodeFiles(filePath, nodes.ToSlice()...)
	if err != nil {
		return err
	}
	page := fileviews.FilesPage{
		Root: filePath == ".",
		ParentDirectory: fileviews.FilesPageItem{
			Name:  parentPath,
			URL:   h.reverse(routes.GetCourseFile.String(), nodes.ToSlice(filePath)),
			Path:  parentPath,
			IsDir: true,
		},
		CurrentDirectory: fileviews.FilesPageItem{
			Name:  filePath,
			URL:   h.reverse(routes.GetCourseFile.String(), nodes.ToSlice(filePath)),
			Path:  filePath,
			IsDir: filepath.Ext(filePath) == "",
		},
		OpenDirURL:          web.URLFunc(routes.GetCourseFiles, h.reverse, nodePath.ToSlice()...),
		ViewMarkdownURL:     web.URLFunc(routes.ViewCourseFile, h.reverse, nodePath.ToSlice()...),
		EditMarkdownFileURL: web.URLFunc(routes.GetCourseEditFile, h.reverse, nodePath.ToSlice()...),
		OpenFileURL:         web.URLFunc(routes.GetCourseFile, h.reverse, nodePath.ToSlice()...),
		UploadFileURL:       web.URLFunc(routes.PostCourseFile, h.reverse, nodePath.ToSlice()...),
		Node:                nodes.Course,
		Files:               files,
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)
}

func (h *courseHandler) nodeDetails(path routes.NodePath, nodes node.Nodes) appcomponents.NodeDetailsPage {
	nodeData := appcomponents.NodeDetailsPage{
		Node:                nodes.Course,
		ParentNode:          nodes.Term,
		ListChildrenURL:     h.reverse(routes.GetUnits.String(), path.ToSlice()...),
		GetEditNodeURL:      h.reverse(routes.GetEditCourse.String(), path.ToSlice()...),
		PostEditNodeURL:     h.reverse(routes.PostEditCourse.String(), path.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		CancelEditURL:       h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		BreadCrumbs:         BreadCrumbs(nodes, path, h.reverse),
		IsEdit:              true,
		ServerFilesURL:      h.reverse(routes.GetCourseFiles.String(), path.ToSlice()...),
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return nodeData
}
