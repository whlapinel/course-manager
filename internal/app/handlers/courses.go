package handlers

import (
	"fmt"
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	mt "gh_static_portfolio/internal/app/views/course"
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

type courseHandler struct {
	service     *services.CourseService
	nodeService *services.NodeService
	fileService *services.FileService
	markdown    *services.MarkdownService
	sitegen     *services.SiteGeneratorService
	reverse     web.Reverse
	*baseHandler[dto.Course, int, int]
}

func NewCourseHandler(
	sitegen *services.SiteGeneratorService,
	service *services.CourseService,
	nodeService *services.NodeService,
	fileService *services.FileService,
	markdown *services.MarkdownService,
	reverse web.Reverse,
) *courseHandler {
	return &courseHandler{
		sitegen:     sitegen,
		service:     service,
		nodeService: nodeService,
		fileService: fileService,
		reverse:     reverse,
		markdown:    markdown,
		baseHandler: &baseHandler[dto.Course, int, int]{
			hasStatic:          true,
			previewer:          sitegen,
			service:            service,
			files:              fileService,
			markdown:           markdown,
			nodes:              nodeService,
			reverse:            reverse,
			getNode:            routes.GetCourse,
			viewNodeFile:       routes.ViewCourseFile,
			deleteNodeFile:     routes.DeleteCourseFile,
			getNodeFile:        routes.GetCourseFile,
			getNodeFiles:       routes.GetCourseFiles,
			getNodeEditFile:    routes.GetCourseEditFile,
			getCreateMarkdown:  routes.GetCourseCreateMarkdown,
			postCreateMarkdown: routes.PostCourseCreateMarkdown,
			postNodeFile:       routes.PostCourseFile,
			postNodeEditFile:   routes.PostCourseEditFile,
		},
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
		// base handler
		web.NewRouteHandler(web.GET, routes.CourseFiles, routes.GetCourseFiles, h.showFiles),
		web.NewRouteHandler(web.GET, routes.CourseFile, routes.GetCourseFile, h.showFiles),
		web.NewRouteHandler(web.POST, routes.CourseFiles, routes.PostCourseFile, h.postUploadedFile),
		web.NewRouteHandler(web.GET, routes.CourseEditFile, routes.GetCourseEditFile, h.showEditFile),
		web.NewRouteHandler(web.POST, routes.CourseEditFile, routes.PostCourseEditFile, h.postEditMarkdown),
		web.NewRouteHandler(web.GET, routes.CourseViewFile, routes.ViewCourseFile, h.viewMarkdown),
		web.NewRouteHandler(web.DELETE, routes.CourseFile, routes.DeleteCourseFile, h.deleteFile),
		web.NewRouteHandler(web.GET, routes.CourseCreateMarkdown, routes.GetCourseCreateMarkdown, h.getCreateMarkdownFile),
		web.NewRouteHandler(web.POST, routes.CourseCreateMarkdown, routes.PostCourseCreateMarkdown, h.postCreateMarkdownFile),

		// course handler overrides
		web.NewRouteHandler(web.POST, routes.GenerateSite, routes.PostGenerateSite, h.generateSite),
		web.NewRouteHandler(web.GET, routes.Courses, routes.GetCourses, h.listByTerm),
		web.NewRouteHandler(web.GET, routes.Course, routes.GetCourse, h.showDetails),
		web.NewRouteHandler(web.GET, routes.CourseEdit, routes.GetEditCourse, h.showEdit),
		web.NewRouteHandler(web.POST, routes.CourseEdit, routes.PostEditCourse, h.postEdit),
		web.NewRouteHandler(web.GET, routes.NewCourse, routes.GetNewCourse, h.showCreateNew),
		web.NewRouteHandler(web.POST, routes.NewCourse, routes.PostNewCourse, h.postNew),
		web.NewRouteHandler(web.DELETE, routes.Course, routes.DeleteCourse, h.delete),
	}
}

func (h *courseHandler) generateSite(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	err = h.sitegen.Build(info.User, info.Term, info.Course)
	if err != nil {
		return err
	}
	return c.String(200, "Site generated!")
}

func (h *courseHandler) delete(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	return h.service.Delete(nodePath.CourseID)

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
			BaseNode: ports.BaseNode[int, int]{
				ParentID: nodePath.TermID,
			},
		},
	}
	course.Name = c.FormValue("name")
	course.Description = c.FormValue("description")
	err = h.service.Save(course)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetCourses.String(), nodePath.ToSlice()...))
}

func (h *courseHandler) showCreateNew(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodeService)
	if err != nil {
		return err
	}
	page := appcomponents.NodeCreatePage{
		ParentNode:          info.Term,
		NodeType:            dto.CourseTypeName,
		Params:              info.NodePath,
		PostCreateNodeURL:   h.reverse(routes.PostNewCourse.String(), info.NodePath.ToSlice()...),
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

	courses, err := h.service.ByParentID(path.TermID)
	if err != nil {
		return err
	}
	startMonth := int(term.Start.Month())
	startYear := term.Start.Year()

	term.Courses = courses
	nodePage := appcomponents.NodeListPage{
		ParentNode: term,
		Children:   term.GetChildren(),

		ChildDetailsURL: web.URLFunc(routes.GetCourse, h.reverse, path.ToSlice()...),
		DeleteChildURL:  web.URLFunc(routes.DeleteCourse, h.reverse, path.ToSlice()...),
		ShowNewChildURL: h.reverse(routes.GetNewCourse.String(), path.ToSlice()...),
		UpNavURL:        h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		BreadCrumbsData: BreadCrumbs(nodes, path, h.reverse),
	}
	page := mt.CoursesListPage{
		ShowCourseCalendarURL: func(courseID int) string {
			return h.reverse(routes.GetCourseMonthCalendar.String(), path.ToSlice(courseID, startMonth, startYear)...)
		},
		ShowAssessmentsURL:  web.URLFunc(routes.GetCourseAssessments, h.reverse, path.ToSlice()...),
		NodeListPage:        nodePage,
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
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
	if nodes.Course == nil {
		return fmt.Errorf("nodes.Course is nil")
	}
	nodePage := h.nodeDetails(path, nodes)
	nodePage.IsEdit = false
	page := mt.CourseDetailsPage{
		StaticSiteURL:            h.sitegen.StaticSiteURL(nodes.User.(dto.User).LastName, nodes.Course.GetID().(int)),
		GenerateSiteURL:          web.URLFunc(routes.PostGenerateSite, h.reverse, path.ToSlice()...)(),
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

func (h *courseHandler) nodeDetails(path routes.NodePath, nodes ports.Nodes) appcomponents.NodeDetailsPage {
	lastName := strings.ToLower(nodes.User.(dto.User).LastName)
	firstDay := nodes.Term.(dto.Term).Start
	month := firstDay.Month()
	year := firstDay.Year()

	nodeData := appcomponents.NodeDetailsPage{
		GenerateSiteURL: h.reverse(routes.PostGenerateSite.String(), path.ToSlice()...),
		StaticSiteURL:   h.previewer.StaticSiteURL(lastName, path.CourseID),

		Node:                nodes.Course,
		CourseCalendarURL:   h.reverse(routes.GetCourseMonthCalendar.String(), path.UserID, path.TermID, path.CourseID, int(month), year),
		ParentNode:          nodes.Term,
		ListChildrenURL:     h.reverse(routes.GetUnits.String(), path.ToSlice()...),
		GetEditNodeURL:      h.reverse(routes.GetEditCourse.String(), path.ToSlice()...),
		PostEditNodeURL:     h.reverse(routes.PostEditCourse.String(), path.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetCourses.String(), path.ToSlice()...),
		CancelEditURL:       h.reverse(routes.GetCourse.String(), path.ToSlice()...),
		BreadCrumbs:         BreadCrumbs(nodes, path, h.reverse),
		IsEdit:              true,
		ServerFilesURL:      h.reverse(routes.GetCourseFiles.String(), path.ToSlice()...),
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return nodeData
}
