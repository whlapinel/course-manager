package handlers

import (
	"bytes"
	"context"
	"fmt"
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	mt "gh_static_portfolio/internal/app/views/course"
	fileviews "gh_static_portfolio/internal/app/views/files"
	markdownviews "gh_static_portfolio/internal/app/views/markdown"
	"gh_static_portfolio/internal/core/course"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type courseHandler struct {
	service     *services.CourseService
	nodeService *services.NodeService
	fileService *services.FileService
	markdown    *services.MarkdownService
	reverse     web.Reverse
}

func NewCourseHandler(
	service *services.CourseService,
	nodeService *services.NodeService,
	fileService *services.FileService,
	markdown *services.MarkdownService,

	reverse web.Reverse,
) *courseHandler {
	return &courseHandler{
		service:     service,
		nodeService: nodeService,
		fileService: fileService,
		reverse:     reverse,
		markdown:    markdown,
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
		web.NewRouteHandler(web.GET, routes.Course, routes.GetCourse, h.showDetails),
		web.NewRouteHandler(web.GET, routes.CourseEdit, routes.GetEditCourse, h.showEdit),
		web.NewRouteHandler(web.GET, routes.CourseFiles, routes.GetCourseFiles, h.showFiles),
		web.NewRouteHandler(web.POST, routes.CourseEdit, routes.PostEditCourse, h.postEdit),
		web.NewRouteHandler(web.GET, routes.NewCourse, routes.GetNewCourse, h.showCreateNew),
		web.NewRouteHandler(web.POST, routes.NewCourse, routes.PostCourse, h.postNew),
		web.NewRouteHandler(web.DELETE, routes.Course, routes.DeleteCourse, h.delete),
		web.NewRouteHandler(web.POST, routes.CourseFiles, routes.PostCourseFile, h.postFile),
		web.NewRouteHandler(web.GET, routes.CourseEditFile, routes.GetCourseEditFile, h.showEditFile),
		web.NewRouteHandler(web.POST, routes.CourseEditFile, routes.PostCourseEditFile, h.postEditFile),
		web.NewRouteHandler(web.GET, routes.CourseViewFile, routes.ViewCourseFile, h.viewMarkdown),
	}
}

func (h *courseHandler) viewMarkdown(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	if filePath == "" {
		return fmt.Errorf("path param is empty")
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	html, err := h.markdown.ViewMarkdown(filePath, nodes.ToSlice()...)
	if err != nil {
		return err
	}
	doc := markdownviews.MarkdownDocument{
		Title:   filepath.Base(filePath),
		Content: string(html),
		Static:  false,
	}
	var buf bytes.Buffer
	err = markdownviews.DocLayout(doc).Render(context.Background(), &buf)
	if err != nil {
		return err
	}
	doc.Content = buf.String()
	component := markdownviews.MarkdownIFrame(doc)
	layout := BaseLayout3(h.reverse, nodes.User.(dto.User))
	return web.Respond(c, "", component, layout.Component2(component))
}

func (h *courseHandler) postEditFile(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	if filePath == "" {
		return fmt.Errorf("path param is empty")
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	fileInfo, err := h.fileService.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if fileInfo.IsDir {
		return fmt.Errorf("%s is a directory", filePath)
	}
	content := c.FormValue("code-editor")
	log.Println("content", content)
	err = h.fileService.Update([]byte(content), filePath, nodes)
	if err != nil {
		return err
	}
	return c.Redirect(
		303,
		web.URLFunc(
			routes.ViewCourseFile,
			h.reverse,
			path.ToSlice()...,
		)(filePath),
	)
}
func (h *courseHandler) showEditFile(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	if filePath == "" {
		return fmt.Errorf("path param is empty")
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	fileInfo, err := h.fileService.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if fileInfo.IsDir {
		return fmt.Errorf("%s is a directory", filePath)
	}
	content, err := h.fileService.FileContent(filePath, nodes)
	if err != nil {
		return err
	}
	page := markdownviews.MarkdownEditor{
		Contents:            string(content),
		PostEditFileURL:     web.URLFunc(routes.PostCourseEditFile, h.reverse, path.ToSlice()...)(filePath),
		CourseManagerLayout: BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}
	return web.Respond(
		c,
		h.reverse(
			routes.GetCourse.String(),
			path.ToSlice()...,
		),
		page.Component(),
		nil,
	)
}
func (h *courseHandler) postFile(c echo.Context) error {
	filePath := c.Param("*")
	if filePath == "*" {
		filePath = "."
	}
	params, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(params)
	if err != nil {
		return err
	}
	// Parse the form to retrieve the file
	err = c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		return err
	}
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	filePath = filepath.Join(filePath, file.Filename)
	err = h.fileService.Save(file, filePath, nodes)
	if err != nil {
		return err
	}
	// Respond to the client
	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", file.Filename))
}
func (h *courseHandler) delete(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	return h.service.Delete(nodePath.TermID)

}
func (h *courseHandler) postNew(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	course := dto.Course{
		Course: course.Course{
			ParentID: nodePath.TermID,
		},
	}
	course.CourseName = c.FormValue("name")
	course.Description = c.FormValue("description")
	err = h.service.Save(course)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetCourses.String(), nodePath.ToSlice()...))
}

func (h *courseHandler) showCreateNew(c echo.Context) error {
	info, err := parseNodeInfo(c, h.nodeService)
	if err != nil {
		return err
	}
	page := appcomponents.NodeCreatePage{
		ParentNode:          info.User,
		NodeType:            dto.CourseTypeName,
		Params:              info.NodePath,
		PostCreateNodeURL:   h.reverse(routes.PostCourse.String(), info.NodePath.ToSlice()...),
		CancelURL:           h.reverse(routes.GetCourses.String(), info.NodePath.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(info.Nodes, info.NodePath, h.reverse),
		CourseManagerLayout: BaseLayout3(h.reverse, info.User.(dto.User)),
	}
	return Respond(c, page)
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

func (h *courseHandler) nodeDetails(path routes.NodePath, nodes ports.Nodes) appcomponents.NodeDetailsPage {
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
