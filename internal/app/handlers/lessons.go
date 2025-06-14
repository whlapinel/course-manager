package handlers

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	lessonviews "gh_static_portfolio/internal/app/views/lesson"
	markdownviews "gh_static_portfolio/internal/app/views/markdown"
	"gh_static_portfolio/internal/features/lesson"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"strconv"
	"strings"

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
	*baseHandler[dto.Lesson, int, int]
}

func NewLessonHandler(
	service *services.LessonService,
	nodeService *services.NodeService,
	reverse web.Reverse,
	paths ports.PathingService,
	slides *services.SlidesService,
	files *services.FileService,
	markdown *services.MarkdownService,
	previewer ports.SiteGenerator,

) *lessonHandler {
	return &lessonHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
		paths:       paths,
		slides:      slides,
		files:       files,
		markdown:    markdown,
		baseHandler: &baseHandler[dto.Lesson, int, int]{
			previewer:        previewer,
			service:          service,
			files:            files,
			markdown:         markdown,
			nodes:            nodeService,
			reverse:          reverse,
			getNode:          routes.GetLesson,
			viewNodeFile:     routes.ViewLessonFile,
			deleteNodeFile:   routes.DeleteLessonFile,
			getNodeFile:      routes.GetLessonFile,
			getNodeFiles:     routes.GetLessonFiles,
			getNodeEditFile:  routes.GetLessonEditFile,
			postNodeFile:     routes.PostLessonFile,
			postNodeEditFile: routes.PostLessonEditFile,
		},
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
		// base handler
		web.NewRouteHandler(web.GET, routes.LessonFiles, routes.GetLessonFiles, h.showFiles),
		web.NewRouteHandler(web.POST, routes.LessonFiles, routes.PostLessonFile, h.postUploadedFile),
		web.NewRouteHandler(web.GET, routes.LessonEditFile, routes.GetLessonEditFile, h.showEditFile),
		web.NewRouteHandler(web.POST, routes.LessonEditFile, routes.PostLessonEditFile, h.postEditMarkdown),
		web.NewRouteHandler(web.GET, routes.LessonViewFile, routes.ViewLessonFile, h.viewMarkdown),
		web.NewRouteHandler(web.DELETE, routes.LessonFile, routes.DeleteLessonFile, h.deleteFile),

		// overrides
		web.NewRouteHandler(web.GET, routes.Lessons, routes.GetLessons, h.listByUnit),
		web.NewRouteHandler(web.GET, routes.Lesson, routes.GetLesson, h.showDetails),
		web.NewRouteHandler(web.GET, routes.LessonEdit, routes.GetEditLesson, h.showEdit),
		web.NewRouteHandler(web.POST, routes.LessonEdit, routes.PostEditLesson, h.postEdit),
		web.NewRouteHandler(web.GET, routes.LessonSlides, routes.GetLessonSlides, h.showSlides),
		web.NewRouteHandler(web.GET, routes.LessonEditSlides, routes.GetEditLessonSlides, h.showEditSlides),
		web.NewRouteHandler(web.POST, routes.LessonEditSlides, routes.PostEditLessonSlides, h.postEditSlides),
		web.NewRouteHandler(web.GET, routes.NewLesson, routes.GetNewLesson, h.showCreateNew),
		web.NewRouteHandler(web.POST, routes.NewLesson, routes.PostNewLesson, h.postNew),
		web.NewRouteHandler(web.DELETE, routes.Lesson, routes.DeleteLesson, h.delete),
	}
}

func (h *lessonHandler) delete(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	return h.service.Delete(nodePath.LessonID)
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
			BaseNode: ports.BaseNode[int, int]{
				ParentID: nodePath.UnitID,
			},
		},
	}
	lesson.Name = c.FormValue("name")
	lesson.Description = c.FormValue("description")
	numParam := c.FormValue("number")
	lesson.Number, err = strconv.Atoi(numParam)
	if err != nil {
		return err
	}
	err = h.service.Save(lesson)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetLessons.String(), nodePath.ToSlice()...))
}
func (h *lessonHandler) showCreateNew(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	page := ac.NodeCreatePage{
		ParentNode:          info.Unit,
		NodeType:            dto.LessonTypeName,
		Params:              info.NodePath,
		PostCreateNodeURL:   h.reverse(routes.PostNewLesson.String(), info.NodePath.ToSlice()...),
		CancelURL:           h.reverse(routes.GetLessons.String(), info.NodePath.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(info.Nodes, info.NodePath, h.reverse),
		CourseManagerLayout: BaseLayout3(h.reverse, info.User.(dto.User)),
	}
	return Respond(c, page)
}

func (h *lessonHandler) postEditSlides(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	content := c.FormValue("code-editor")
	log.Println(content)
	err = h.slides.UpdateSlides(info.Nodes, []byte(content))
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
	var newTab bool
	newTabQueryParam := c.QueryParam("new-tab")
	log.Println("newTabQueryParam", newTabQueryParam)
	if newTabQueryParam != "" {
		newTab, err = strconv.ParseBool(newTabQueryParam)
		if err != nil {
			return err
		}
	}
	if newTab {
		return c.File(h.paths.NodeSlidesHTMLPath(nodes.ToSlice()...))
	}
	html, err := h.slides.GetSlides(nodes.ToSlice()...)
	if err != nil {
		return err
	}
	slides := lessonviews.Slides{
		HTML:          string(html),
		EditSlidesURL: h.reverse(routes.GetEditLessonSlides.String(), path.ToSlice()...),
		LessonDetailsPage: lessonviews.LessonDetailsPage{
			GetSlidesURL: h.reverse(routes.GetLessonSlides.String(), path.ToSlice()...) + "?new-tab=true",
		},
	}
	return web.Respond(c, routes.GetLesson.String(), slides.Component(), nil)

}

func (h *lessonHandler) showEditSlides(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	content, err := h.slides.SlidesContent(info.Nodes)
	if err != nil {
		return err
	}
	editor := markdownviews.MarkdownEditor{
		Name:            "Slides",
		Contents:        string(content),
		PostEditFileURL: h.reverse(routes.PostEditLessonSlides.String(), info.NodePath.ToSlice()...),
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
	lessons, err := h.service.ByParentID(path.UnitID)
	if err != nil {
		return err
	}
	unit.Lessons = lessons
	page := ac.NodeListPage{
		ParentNode:          unit,
		Children:            unit.GetChildren(),
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
		FileURL:         web.URLFunc(routes.GetLessonFiles, h.reverse, path.ToSlice()...),
		AssetsURLFunc:   web.AssetsURLFunc,
		ViewMarkdownURL: web.URLFunc(routes.ViewLessonFile, h.reverse),
		GetSlidesURL:    h.reverse(routes.GetLessonSlides.String(), path.ToSlice()...),
		EditSlidesURL:   h.reverse(routes.GetEditLessonSlides.String(), path.ToSlice()...),
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
	lastName := strings.ToLower(nodes.User.(dto.User).LastName)

	nodePage := ac.NodeDetailsPage{
		GenerateSiteURL:     h.reverse(routes.PostGenerateSite.String(), path.ToSlice()...),
		StaticSiteURL:       h.previewer.StaticSiteURL(lastName, path.CourseID),
		Node:                nodes.Lesson,
		ParentNode:          nodes.Unit,
		ServerFilesURL:      h.reverse(routes.GetLessonFiles.String(), path.ToSlice()...),
		GetEditNodeURL:      h.reverse(routes.GetEditLesson.String(), path.ToSlice()...),
		PostEditNodeURL:     h.reverse(routes.PostEditLesson.String(), path.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetLessons.String(), path.ToSlice()...),
		CancelEditURL:       h.reverse(routes.GetLesson.String(), path.ToSlice()...),
		BreadCrumbs:         BreadCrumbs(nodes, path, h.reverse),
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return nodePage
}
