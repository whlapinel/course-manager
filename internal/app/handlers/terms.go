package handlers

import (
	"fmt"
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	termviews "gh_static_portfolio/internal/app/views/term"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

type termHandler struct {
	nodeService *services.NodeService
	service     *services.TermService
	fileService *services.FileService
	markdown    *services.MarkdownService
	reverse     web.Reverse
	*baseHandler[dto.Term, int, string]
}

func NewTermHandler(
	service *services.TermService,
	nodeService *services.NodeService,
	fileService *services.FileService,
	markdownService *services.MarkdownService,
	reverse web.Reverse,
) *termHandler {
	return &termHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
		fileService: fileService,
		markdown:    markdownService,
		baseHandler: &baseHandler[dto.Term, int, string]{
			service:          service,
			files:            fileService,
			markdown:         markdownService,
			nodes:            nodeService,
			reverse:          reverse,
			getNode:          routes.GetTerm,
			viewNodeFile:     routes.ViewTermFile,
			getNodeFile:      routes.GetTermFile,
			getNodeFiles:     routes.GetTermFiles,
			getNodeEditFile:  routes.GetTermEditFile,
			postNodeFile:     routes.PostTermFile,
			postNodeEditFile: routes.PostTermEditFile,
		},
	}
}

func RegisterTermRoutes(group *echo.Group, h *termHandler) error {
	for _, handler := range termRouteHandlers(h) {
		err := web.RegisterRoute(group, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func termRouteHandlers(h *termHandler) []web.RouteHandler {
	return []web.RouteHandler{
		// base handler methods
		web.NewRouteHandler(web.GET, routes.TermFiles, routes.GetTermFiles, h.showFiles),
		web.NewRouteHandler(web.POST, routes.TermFiles, routes.PostTermFile, h.postFile),
		web.NewRouteHandler(web.GET, routes.TermEditFile, routes.GetTermEditFile, h.showEditFile),
		web.NewRouteHandler(web.POST, routes.TermEditFile, routes.PostTermEditFile, h.postEditFile),
		web.NewRouteHandler(web.GET, routes.TermViewFile, routes.ViewTermFile, h.viewMarkdown),

		web.NewRouteHandler(web.GET, routes.Term, routes.GetTerm, h.showDetails),
		web.NewRouteHandler(web.GET, routes.Terms, routes.GetTerms, h.listByUser),
		web.NewRouteHandler(web.GET, routes.NewTerm, routes.GetNewTerm, h.showCreateNew),
		web.NewRouteHandler(web.POST, routes.NewTerm, routes.PostTerm, h.postNew),
		web.NewRouteHandler(web.DELETE, routes.Term, routes.DeleteTerm, h.delete),
		web.NewRouteHandler(web.GET, routes.TermEdit, routes.GetEditTerm, h.showEdit),
		web.NewRouteHandler(web.POST, routes.TermEdit, routes.PostEditTerm, h.postEdit),
		web.NewRouteHandler(web.GET, routes.TermDates, routes.GetTermDates, h.showEditTermDates),
		web.NewRouteHandler(web.POST, routes.TermDates, routes.PostTermDate, h.postTermDate),
	}
}

func (h *termHandler) postTermDate(c echo.Context) error {
	info, err := parseNodeInfo(c, h.nodeService)
	if err != nil {
		return err
	}
	dateParam := c.FormValue("date")
	parsed, err := time.Parse(time.DateOnly, dateParam)
	if err != nil {
		return err
	}
	err = h.service.RemoveInstructionalDay(parsed, info.NodePath.TermID)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetTermDates.String(), info.NodePath.ToSlice()...))
}

func (h *termHandler) showEditTermDates(c echo.Context) error {
	info, err := parseNodeInfo(c, h.nodeService)
	if err != nil {
		return err
	}
	page := termviews.AddNonInstructDayPage{
		Term:                info.Nodes.Term.(dto.Term),
		GetAddDayURL:        "testing, this is to show add a day",
		PostAddDayURL:       h.reverse(routes.PostTermDate.String(), info.NodePath.ToSlice()...),
		TermDetailsURL:      h.reverse(routes.GetTerm.String(), info.NodePath.ToSlice()...),
		BreadCrumbs:         BreadCrumbs(info.Nodes, info.NodePath, h.reverse),
		CourseManagerLayout: BaseLayout3(h.reverse, info.User.(dto.User)),
	}
	return Respond(c, page)
}

// creates new term (previously this would be create new course)
func (h *termHandler) showCreateNew(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(nodePath)
	if err != nil {
		return err
	}
	page := appcomponents.NodeCreatePage{
		ParentNode:          nodes.User,
		NodeType:            dto.TermTypeName,
		Params:              nodePath,
		PostCreateNodeURL:   h.reverse(routes.PostTerm.String(), nodePath.ToSlice()...),
		CancelURL:           h.reverse(routes.GetTerms.String(), nodePath.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(nodes, nodePath, h.reverse),
		CourseManagerLayout: BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)
}

func (h *termHandler) postNew(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	term := dto.Term{
		Term: term.Term{
			UserID: nodePath.UserID,
		},
	}
	form := c.Request().Form
	for k, v := range form {
		switch k {
		case "name":
			term.Name = v[0]
		case "description":
			term.Description = v[0]
		case "start-date":
			date, err := time.Parse(time.DateOnly, v[0])
			if err != nil {
				return err
			}
			term.Start = date
		case "end-date":
			date, err := time.Parse(time.DateOnly, v[0])
			if err != nil {
				return err
			}
			term.End = date
		default:
			return fmt.Errorf("unexpected value: %s: %s", k, v)
		}
	}
	log.Println(term)
	err = h.service.Save(term)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetTerms.String(), nodePath.ToSlice()...))
}

func (h *termHandler) delete(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	return h.service.Delete(nodePath.TermID)
}

func (h *termHandler) listByUser(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(nodePath)
	if err != nil {
		return err
	}
	user := nodes.User.(dto.User)
	userID := user.ID
	terms, err := h.service.ByParentID(userID)
	if err != nil {
		return err
	}
	user.Terms = terms
	nodePage := appcomponents.NodeListPage{
		ParentNode:          user,
		Children:            user.Children(),
		ChildDetailsURL:     web.URLFunc(routes.GetTerm, h.reverse, nodePath.ToSlice()...),
		ChildChildrenURL:    web.URLFunc(routes.GetCourses, h.reverse, nodePath.ToSlice()...),
		DeleteChildURL:      web.URLFunc(routes.DeleteTerm, h.reverse, nodePath.ToSlice()...),
		ShowNewChildURL:     h.reverse(routes.GetNewTerm.String(), nodePath.ToSlice()...),
		UpNavURL:            h.reverse(routes.GetUser.String(), nodePath.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(nodes, nodePath, h.reverse),
		CourseManagerLayout: BaseLayout3(h.reverse, user),
	}
	page := termviews.TermsListPage{
		ShowTermCalendarURL: web.URLFunc(routes.GetTermCalendar, h.reverse, nodePath.ToSlice()...),
		NodeListPage:        nodePage,
	}
	return Respond(c, page)
}

func (h *termHandler) postEdit(c echo.Context) error {
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
	term := nodes.Term.(dto.Term)
	form := c.Request().Form
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "name":
			term.Name = val[0]
		case "description":
			term.Description = val[0]
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}
	}
	err = h.service.Update(term)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetTerm.String(), nodePath.ToSlice()...))
}

func (h *termHandler) showDetails(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodeService.Nodes(path)
	if err != nil {
		return err
	}

	nodePage := h.nodeDetails(path, nodes)
	termPage := termviews.TermDetailsPage{
		NodeDetailsPage:      nodePage,
		ShowEditTermDatesURL: h.reverse(routes.GetTermDates.String(), path.ToSlice()...),
		CourseManagerLayout:  BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, termPage)
}

func (h *termHandler) showEdit(c echo.Context) error {
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
	page := termviews.TermDetailsPage{
		NodeDetailsPage:      nodeData,
		ShowEditTermDatesURL: h.reverse(routes.GetTermDates.String(), path.ToSlice()...),
		CourseManagerLayout:  BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}.DetailsEdit()
	return Respond(c, page)
}

func (h *termHandler) nodeDetails(path routes.NodePath, nodes ports.Nodes) appcomponents.NodeDetailsPage {
	nodePage := appcomponents.NodeDetailsPage{
		Node:            nodes.Term,
		ParentNode:      nodes.User,
		GetEditNodeURL:  h.reverse(routes.GetEditTerm.String(), path.ToSlice()...),
		PostEditNodeURL: h.reverse(routes.PostEditTerm.String(), path.ToSlice()...),
		ListChildrenURL: h.reverse(routes.GetCourses.String(), path.ToSlice()...),
		UpNavURL:        h.reverse(routes.GetUser.String(), path.ToSlice()...),
		CancelEditURL:   h.reverse(routes.GetTerm.String(), path.ToSlice()...),
		ServerFilesURL:  h.reverse(routes.GetTermFiles.String(), path.ToSlice()...),
		BreadCrumbs:     BreadCrumbs(nodes, path, h.reverse),
	}
	return nodePage
}
