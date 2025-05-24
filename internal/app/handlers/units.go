package handlers

import (
	"bytes"
	"context"
	"fmt"
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	fileviews "gh_static_portfolio/internal/app/views/files"
	markdownviews "gh_static_portfolio/internal/app/views/markdown"
	"gh_static_portfolio/internal/core/unit"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

type unitHandler struct {
	service     *services.UnitService
	nodeService *services.NodeService
	files       *services.FileService
	markdown    *services.MarkdownService

	reverse web.Reverse
}

func NewUnitHandler(
	service *services.UnitService,
	nodeService *services.NodeService,
	files *services.FileService,
	markdown *services.MarkdownService,
	reverse web.Reverse,
) *unitHandler {
	return &unitHandler{
		service:     service,
		nodeService: nodeService,
		reverse:     reverse,
		files:       files,
		markdown:    markdown,
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
		web.NewRouteHandler(web.GET, routes.Units, routes.GetUnits, h.listByCourse),
		web.NewRouteHandler(web.GET, routes.Unit, routes.GetUnit, h.showDetails),
		web.NewRouteHandler(web.GET, routes.UnitEdit, routes.GetEditUnit, h.showEdit),
		web.NewRouteHandler(web.GET, routes.UnitEdit, routes.PostEditUnit, h.postEdit),
		web.NewRouteHandler(web.GET, routes.UnitFiles, routes.GetUnitFiles, h.showFiles),
		web.NewRouteHandler(web.GET, routes.NewUnit, routes.GetNewUnit, h.showCreateNew),
		web.NewRouteHandler(web.POST, routes.NewUnit, routes.PostUnit, h.postNew),
		web.NewRouteHandler(web.DELETE, routes.Unit, routes.DeleteUnit, h.delete),
		web.NewRouteHandler(web.POST, routes.UnitFiles, routes.PostUnitFile, h.postFile),
		web.NewRouteHandler(web.GET, routes.UnitEditFile, routes.GetUnitEditFile, h.showEditFile),
		web.NewRouteHandler(web.POST, routes.UnitEditFile, routes.PostUnitEditFile, h.postEditFile),
		web.NewRouteHandler(web.GET, routes.UnitViewFile, routes.ViewUnitFile, h.viewMarkdown),
	}
}

func (h *unitHandler) viewMarkdown(c echo.Context) error {
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
func (h *unitHandler) postEditFile(c echo.Context) error {
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
			routes.ViewUnitFile,
			h.reverse,
			path.ToSlice()...,
		)(filePath),
	)

}
func (h *unitHandler) showEditFile(c echo.Context) error {
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
		PostEditFileURL:     web.URLFunc(routes.PostUnitEditFile, h.reverse, path.ToSlice()...)(filePath),
		CourseManagerLayout: BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}
	return web.Respond(
		c,
		h.reverse(
			routes.GetUnit.String(),
			path.ToSlice()...,
		),
		page.Component(),
		nil,
	)

}
func (h *unitHandler) postFile(c echo.Context) error {
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
			CourseID: nodePath.CourseID,
		},
	}
	unit.Name = c.FormValue("name")
	unit.Description = c.FormValue("description")
	err = h.service.Save(unit)
	if err != nil {
		return err
	}
	return c.Redirect(303, h.reverse(routes.GetUnits.String(), nodePath.ToSlice()...))
}

func (h *unitHandler) showCreateNew(c echo.Context) error {
	info, err := parseNodeInfo(c, h.nodeService)
	if err != nil {
		return err
	}
	page := ac.NodeCreatePage{
		ParentNode:          info.User,
		NodeType:            dto.UnitTypeName,
		Params:              info.NodePath,
		PostCreateNodeURL:   h.reverse(routes.PostUnit.String(), info.NodePath.ToSlice()...),
		CancelURL:           h.reverse(routes.GetUnits.String(), info.NodePath.ToSlice()...),
		BreadCrumbsData:     BreadCrumbs(info.Nodes, info.NodePath, h.reverse),
		CourseManagerLayout: BaseLayout3(h.reverse, info.User.(dto.User)),
	}
	return Respond(c, page)
}

func (h *unitHandler) showFiles(c echo.Context) error {
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
			URL:   h.reverse(routes.GetUnitFile.String(), nodes.ToSlice(filePath)),
			Path:  parentPath,
			IsDir: true,
		},
		CurrentDirectory: fileviews.FilesPageItem{
			Name:  filePath,
			URL:   h.reverse(routes.GetUnitFile.String(), nodes.ToSlice(filePath)),
			Path:  filePath,
			IsDir: filepath.Ext(filePath) == "",
		},
		OpenDirURL:          web.URLFunc(routes.GetUnitFiles, h.reverse, nodePath.ToSlice()...),
		ViewMarkdownURL:     web.URLFunc(routes.ViewUnitFile, h.reverse, nodePath.ToSlice()...),
		EditMarkdownFileURL: web.URLFunc(routes.GetUnitEditFile, h.reverse, nodePath.ToSlice()...),
		OpenFileURL:         web.URLFunc(routes.GetUnitFile, h.reverse, nodePath.ToSlice()...),
		UploadFileURL:       web.URLFunc(routes.PostUnitFile, h.reverse, nodePath.ToSlice()...),
		Node:                nodes.Unit,
		Files:               files,
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
		BreadCrumbsData:     BreadCrumbs(nodes, nodePath, h.reverse),
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
	units, err := h.service.ByCourseID(path.CourseID)
	if err != nil {
		return err
	}
	course.Units = units
	page := ac.NodeListPage{
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

func (h *unitHandler) nodeDetails(path routes.NodePath, nodes ports.Nodes) ac.NodeDetailsPage {
	nodePage := ac.NodeDetailsPage{
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
