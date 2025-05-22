package handlers

import (
	"bytes"
	"context"
	"fmt"
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	fileviews "gh_static_portfolio/internal/app/views/files"
	markdownviews "gh_static_portfolio/internal/app/views/markdown"
	termviews "gh_static_portfolio/internal/app/views/term"
	"gh_static_portfolio/internal/core/term"
	"gh_static_portfolio/internal/features/files"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
)

type termHandler struct {
	nodeService *services.NodeService
	service     *services.TermService
	fileService *files.Service
	markdown    *services.MarkdownService
	reverse     web.Reverse
}

func NewTermHandler(
	service *services.TermService,
	nodeService *services.NodeService,
	fileService *files.Service,
	markdownService *services.MarkdownService,
	reverse web.Reverse,
) *termHandler {
	return &termHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
		fileService: fileService,
		markdown:    markdownService,
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

		web.NewRouteHandler(web.GET, routes.Term, routes.GetTerm, h.showDetails),
		web.NewRouteHandler(web.GET, routes.Terms, routes.GetTerms, h.listByUser),
		web.NewRouteHandler(web.GET, routes.NewTerm, routes.GetNewTerm, h.showCreateNew),
		web.NewRouteHandler(web.POST, routes.NewTerm, routes.PostTerm, h.postNew),
		web.NewRouteHandler(web.DELETE, routes.Term, routes.DeleteTerm, h.delete),
		web.NewRouteHandler(web.GET, routes.TermEdit, routes.GetEditTerm, h.showEdit),
		web.NewRouteHandler(web.POST, routes.TermEdit, routes.PostEditTerm, h.postEdit),
		web.NewRouteHandler(web.GET, routes.TermFiles, routes.GetTermFiles, h.showFiles),
		web.NewRouteHandler(web.POST, routes.TermFiles, routes.PostTermFile, h.postFile),
		web.NewRouteHandler(web.GET, routes.TermEditFile, routes.GetTermEditFile, h.showEditFile),
		web.NewRouteHandler(web.POST, routes.TermEditFile, routes.PostTermEditFile, h.postEditFile),
		web.NewRouteHandler(web.GET, routes.TermViewFile, routes.ViewTermFile, h.viewMarkdown),
	}
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
		ParentNode:        nodes.User,
		NodeType:          dto.TermTypeName,
		Params:            nodePath,
		PostCreateNodeURL: h.reverse(routes.PostTerm.String(), nodePath.ToSlice()...),
		CancelURL:         h.reverse(routes.GetTerms.String(), nodePath.ToSlice()...),
		BreadCrumbsData:   BreadCrumbs(nodes, nodePath, h.reverse),
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
	terms, err := h.service.ByUserID(userID)
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
		BreadCrumbsData:     BreadCrumbs2(nodes, nodePath, h.reverse),
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

func (h *termHandler) postEditFile(c echo.Context) error {
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
	content := c.FormValue(string(markdownviews.EditSlidesTextAreaID))
	log.Println("content", content)
	err = h.fileService.Update([]byte(content), filePath, nodes)
	if err != nil {
		return err
	}
	return c.Redirect(303, web.URLFunc(routes.ViewTermFile, h.reverse, path.ToSlice()...)(filePath))
}

func (h *termHandler) showEditFile(c echo.Context) error {
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
		PostEditFileURL:     web.URLFunc(routes.GetTermEditFile, h.reverse, path.ToSlice()...)(filePath),
		CourseManagerLayout: BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)
}

func (h *termHandler) viewMarkdown(c echo.Context) error {
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

func (h *termHandler) showFiles(c echo.Context) error {
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
			URL:   h.reverse(routes.GetTermFile.String(), nodes.User, nodes.Term, filePath),
			Path:  parentPath,
			IsDir: true,
		},
		CurrentDirectory: fileviews.FilesPageItem{
			Name:  filePath,
			URL:   h.reverse(routes.GetTermFile.String(), nodes.User, nodes.Term, filePath),
			Path:  filePath,
			IsDir: filepath.Ext(filePath) == "",
		},
		OpenDirURL:          web.URLFunc(routes.GetTermFiles, h.reverse, nodePath.ToSlice()...),
		ViewMarkdownURL:     web.URLFunc(routes.ViewTermFile, h.reverse, nodePath.ToSlice()...),
		EditMarkdownFileURL: web.URLFunc(routes.GetTermEditFile, h.reverse, nodePath.ToSlice()...),
		OpenFileURL:         web.URLFunc(routes.GetTermFile, h.reverse, nodePath.ToSlice()...),
		UploadFileURL:       web.URLFunc(routes.PostTermFile, h.reverse, nodePath.ToSlice()...),
		Node:                nodes.Term,
		Files:               files,
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)
}

func (h *termHandler) postFile(c echo.Context) error {
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
		BreadCrumbs:     BreadCrumbs2(nodes, path, h.reverse),
	}
	return nodePage
}

// (done) ListChildren(echo.Context) error
// (done) ShowFiles(echo.Context) error
// (done) PostFile(echo.Context) error
// (done) ViewFile(echo.Context) error
// (done) ShowDetails(echo.Context) error
// (done) ShowNewChild(echo.Context) error // e.g. if node is course, new child would be new unit
// PostNewChild(echo.Context) error
// ShowEdit(echo.Context) error
// ShowEditFile(echo.Context) error
// PostEditFile(echo.Context) error
// PostEdit(echo.Context) error
// Delete(echo.Context) error // i.e. delete node itself (not child)
