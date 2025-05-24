package handlers

import (
	"bytes"
	"context"
	"fmt"
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	fileviews "gh_static_portfolio/internal/app/views/files"
	lessonviews "gh_static_portfolio/internal/app/views/lesson"
	markdownviews "gh_static_portfolio/internal/app/views/markdown"
	"gh_static_portfolio/internal/core/lesson"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

type lessonHandler struct {
	nodeService *services.NodeService
	service     *services.LessonService
	paths       ports.PathingService
	reverse     web.Reverse
	slides      *services.SlidesService
	files       *services.FileService
	markdown    *services.MarkdownService
}

func NewLessonHandler(
	service *services.LessonService,
	nodeService *services.NodeService,
	reverse web.Reverse,
	paths ports.PathingService,
	slides *services.SlidesService,
	files *services.FileService,
	markdown *services.MarkdownService,

) *lessonHandler {
	return &lessonHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
		paths:       paths,
		slides:      slides,
		files:       files,
		markdown:    markdown,
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
		web.NewRouteHandler(web.GET, routes.LessonSlides, routes.GetLessonSlides, h.showSlides),
		web.NewRouteHandler(web.GET, routes.LessonEditSlides, routes.GetEditLessonSlides, h.showEditSlides),
		web.NewRouteHandler(web.POST, routes.LessonEditSlides, routes.PostEditLessonSlides, h.postEditSlides),
		web.NewRouteHandler(web.GET, routes.LessonFiles, routes.GetLessonFiles, h.showFiles),
		web.NewRouteHandler(web.GET, routes.NewLesson, routes.GetNewLesson, h.showCreateNew),
		web.NewRouteHandler(web.POST, routes.NewLesson, routes.PostLesson, h.postNew),
		web.NewRouteHandler(web.DELETE, routes.Lesson, routes.DeleteLesson, h.delete),
		web.NewRouteHandler(web.POST, routes.LessonFiles, routes.PostLessonFile, h.postFile),
		web.NewRouteHandler(web.GET, routes.LessonEditFile, routes.GetLessonEditFile, h.showEditFile),
		web.NewRouteHandler(web.POST, routes.LessonEditFile, routes.PostLessonEditFile, h.postEditFile),
		web.NewRouteHandler(web.GET, routes.LessonViewFile, routes.ViewUnitFile, h.viewMarkdown),
	}
}

func (h *lessonHandler) viewMarkdown(c echo.Context) error {
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

func (h *lessonHandler) postEditFile(c echo.Context) error {
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
	fileInfo, err := h.files.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if fileInfo.IsDir {
		return fmt.Errorf("%s is a directory", filePath)
	}
	content := c.FormValue("code-editor")
	log.Println("content", content)
	err = h.files.Update([]byte(content), filePath, nodes)
	if err != nil {
		return err
	}
	return c.Redirect(
		303,
		web.URLFunc(
			routes.ViewLessonFile,
			h.reverse,
			path.ToSlice()...,
		)(filePath),
	)
}
func (h *lessonHandler) showEditFile(c echo.Context) error {
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
	fileInfo, err := h.files.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if fileInfo.IsDir {
		return fmt.Errorf("%s is a directory", filePath)
	}
	content, err := h.files.FileContent(filePath, nodes)
	if err != nil {
		return err
	}
	page := markdownviews.MarkdownEditor{
		Contents:            string(content),
		PostEditFileURL:     web.URLFunc(routes.PostLessonEditFile, h.reverse, path.ToSlice()...)(filePath),
		CourseManagerLayout: BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}
	return web.Respond(
		c,
		h.reverse(
			routes.GetLesson.String(),
			path.ToSlice()...,
		),
		page.Component(),
		nil,
	)
}
func (h *lessonHandler) postFile(c echo.Context) error {
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
	err = h.files.Save(file, filePath, nodes)
	if err != nil {
		return err
	}
	// Respond to the client
	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", file.Filename))
}

func (h *lessonHandler) delete(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	return h.service.Delete(nodePath.UnitID)
}
func (h *lessonHandler) postNew(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	lesson := dto.Lesson{
		Lesson: lesson.Lesson{
			UnitID: nodePath.UnitID,
		},
	}
	lesson.Name = c.FormValue("name")
	lesson.Description = c.FormValue("description")
	err = h.service.Save(lesson)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetLessons.String(), nodePath.ToSlice()...))
}
func (h *lessonHandler) showCreateNew(c echo.Context) error {
	info, err := parseNodeInfo(c, h.nodeService)
	if err != nil {
		return err
	}
	page := ac.NodeCreatePage{
		ParentNode:          info.User,
		NodeType:            dto.LessonTypeName,
		Params:              info.NodePath,
		PostCreateNodeURL:   h.reverse(routes.PostLesson.String(), info.NodePath.ToSlice()...),
		CancelURL:           h.reverse(routes.GetLessons.String(), info.NodePath.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(info.Nodes, info.NodePath, h.reverse),
		CourseManagerLayout: BaseLayout3(h.reverse, info.User.(dto.User)),
	}
	return Respond(c, page)
}

func (h *lessonHandler) showFiles(c echo.Context) error {
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
	fileInfo, err := h.files.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir {
		c.Attachment(fileInfo.AbsPath, filepath.Base(fileInfo.AbsPath))
	}
	parentPath := filepath.Dir(filePath)
	files, err := h.files.NodeFiles(filePath, nodes.ToSlice()...)
	if err != nil {
		return err
	}
	page := fileviews.FilesPage{
		Root: filePath == ".",
		ParentDirectory: fileviews.FilesPageItem{
			Name:  parentPath,
			URL:   h.reverse(routes.GetLessonFile.String(), nodes.ToSlice(filePath)),
			Path:  parentPath,
			IsDir: true,
		},
		CurrentDirectory: fileviews.FilesPageItem{
			Name:  filePath,
			URL:   h.reverse(routes.GetLessonFile.String(), nodes.ToSlice(filePath)),
			Path:  filePath,
			IsDir: filepath.Ext(filePath) == "",
		},
		OpenDirURL:          web.URLFunc(routes.GetLessonFiles, h.reverse, nodePath.ToSlice()...),
		ViewMarkdownURL:     web.URLFunc(routes.ViewLessonFile, h.reverse, nodePath.ToSlice()...),
		EditMarkdownFileURL: web.URLFunc(routes.GetLessonEditFile, h.reverse, nodePath.ToSlice()...),
		OpenFileURL:         web.URLFunc(routes.GetLessonFile, h.reverse, nodePath.ToSlice()...),
		UploadFileURL:       web.URLFunc(routes.PostLessonFile, h.reverse, nodePath.ToSlice()...),
		Node:                nodes.Lesson,
		Files:               files,
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
		BreadCrumbsData:     BreadCrumbs(nodes, nodePath, h.reverse),
	}
	return Respond(c, page)
}

func (h *lessonHandler) postEditSlides(c echo.Context) error {
	info, err := parseNodeInfo(c, h.nodeService)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	for k, v := range form {
		log.Println(k, v)
	}
	content := form[lessonviews.EditSlidesTextAreaID]
	log.Println(content)
	err = h.files.UpdateSlides(info.Nodes, []byte(content[0]))
	if err != nil {
		return err
	}
	return c.Redirect(
		302,
		h.reverse(
			routes.GetLessonSlides.String(),
			info.NodePath.ToSlice()...,
		),
	)

}

func (h *lessonHandler) showSlides(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}
	html, err := h.slides.GetSlides(nodes.ToSlice()...)
	if err != nil {
		return err
	}
	slides := lessonviews.Slides{
		HTML:          string(html),
		EditSlidesURL: h.reverse(routes.GetEditLessonSlides.String(), path.ToSlice()...),
	}
	return web.Respond(c, routes.GetLesson.String(), slides.Component(), nil)

}

func (h *lessonHandler) showEditSlides(c echo.Context) error {
	info, err := parseNodeInfo(c, h.nodeService)
	if err != nil {
		return err
	}
	content, err := h.files.SlidesContent(info.Nodes)
	if err != nil {
		return err
	}
	editor := lessonviews.SlidesEditor{
		Content:           string(content),
		PostEditSlidesURL: h.reverse(routes.PostEditLessonSlides.String(), info.NodePath.ToSlice()...),
	}
	return web.Respond(
		c,
		h.reverse(
			routes.GetLessonSlides.String(),
			info.NodePath.ToSlice()...,
		),
		editor.Component(),
		nil,
	)

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
	page := ac.NodeListPage{
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
		GetSlidesURL:    h.reverse(routes.GetLessonSlides.String(), path.ToSlice()...),
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

func (h *lessonHandler) nodeDetails(path routes.NodePath, nodes ports.Nodes) ac.NodeDetailsPage {
	nodePage := ac.NodeDetailsPage{
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
