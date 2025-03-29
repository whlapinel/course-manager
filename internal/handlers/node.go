package handlers

import (
	"bytes"
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	mt "gh_static_portfolio/internal/templates/app"
	"gh_static_portfolio/internal/util"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Router interface {
	GetRouter() router
	SetRouter(router router)
}

type NodeRouter interface {
	ListChildren(echo.Context) error
	ShowFiles(echo.Context) error
	PostFile(echo.Context) error
	ViewFile(echo.Context) error
	ShowDetails(echo.Context) error
	ShowNewChild(echo.Context) error // e.g. if node is course, new child would be new unit
	PostNewChild(echo.Context) error
	ShowEdit(echo.Context) error
	ShowEditFile(echo.Context) error
	PostEditFile(echo.Context) error
	PostEdit(echo.Context) error
	Delete(echo.Context) error // i.e. delete node itself (not child)
	Router
}

type EmptyNode domain.CourseNode

type EmptyNodeSet []EmptyNode

func URL(r NodeRouter, rhn RouteHandlerName, additionalParams ...any) string {
	echo := r.GetRouter().app
	nodes := r.GetRouter().params
	return echo.Reverse(rhn.String(), AddParams(nodes, additionalParams...)...)
}
func (set EmptyNodeSet) ChildNodeSet() EmptyNodeSet {
	return emptyNodes[:len(set)+1]
}

func (set EmptyNodeSet) ParentNodeSet() EmptyNodeSet {
	return emptyNodes[:len(set)-1]
}

var emptyNodes = []EmptyNode{
	domain.User{},
	domain.Term{},
	domain.Course{},
	domain.Unit{},
	domain.Lesson{},
}

var EmptyNodesUser = emptyNodes[:1]
var EmptyNodesTerm = emptyNodes[:2]
var EmptyNodesCourse = emptyNodes[:3]
var EmptyNodesUnit = emptyNodes[:4]
var EmptyNodesLesson = emptyNodes

// e.g. "/users/:user-id/terms/:term-id/courses/:course-id/units"
func ChildNodesRouteName(nodes ...EmptyNode) RoutePath {
	if len(nodes) == 0 {
		log.Fatalf("ChildNodesRouteName(): nodes cannot be empty")
	}
	lastNode := nodes[len(nodes)-1]
	childTypeName := lastNode.ChildTypeName()
	if childTypeName == "" {
		log.Fatalf("node type '%s' has no child type", lastNode.TypeName())
	}
	var path = string(NodeRouteName(nodes...))
	childrenSegment := util.KebabCase(fmt.Sprintf("%ss", childTypeName))
	path = filepath.Join(path, childrenSegment)
	return RoutePath(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id"
func NodeRouteName(nodes ...EmptyNode) RoutePath {
	var path string = "/"
	for _, node := range nodes {
		path = filepath.Join(path, util.KebabCase(node.TypeName()+"s"))
		path = filepath.Join(path, fmt.Sprintf("/:%s-id", util.KebabCase(node.TypeName())))
	}
	return RoutePath(path)

}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/new
func NewChildRouteName(nodes ...EmptyNode) RoutePath {
	var name = ChildNodesRouteName(nodes...)
	var path = string(name)
	path = filepath.Join(path, "/new")
	return RoutePath(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id/edit
func EditNodeRouteName(nodes ...EmptyNode) RoutePath {
	var path = string(NodeRouteName(nodes...))
	path = filepath.Join(path, "/edit")
	return RoutePath(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id/files/*
func NodeFilesRoutePath(nodes ...EmptyNode) RoutePath {
	var name = NodeRouteName(nodes...)
	path := string(name)
	path = filepath.Join(path, "/files/*")
	return RoutePath(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id/files/edit/*
func NodeEditFileRoutePath(nodes ...EmptyNode) RoutePath {
	var name = NodeRouteName(nodes...)
	path := string(name)
	path = filepath.Join(path, "/files/edit/*")
	return RoutePath(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id/view-markdown/files/*
func ViewNodeFilesRoutePath(nodes ...EmptyNode) RoutePath {
	var name = NodeRouteName(nodes...)
	path := string(name)
	path = filepath.Join(path, "/view-markdown/files/*")
	return RoutePath(path)
}

func ViewNodeFilesRHN(nodes ...EmptyNode) RouteHandlerName {
	routeName := ViewNodeFilesRoutePath(nodes...)
	return RouteHandlerName(GET + routeName)
}

func ListChildrenRHN(nodes ...EmptyNode) RouteHandlerName {
	name := ChildNodesRouteName(nodes...)
	rhn := RouteHandlerName(GET + name)
	return rhn
}

func ListChildChildrenRHN(nodes ...EmptyNode) RouteHandlerName {
	return ListChildrenRHN(EmptyNodeSet(nodes).ChildNodeSet()...)
}

func ShowNewChildRHN(nodes ...EmptyNode) RouteHandlerName {
	routeName := NewChildRouteName(nodes...)
	return RouteHandlerName(GET + routeName)
}

func DeleteChildRHN(nodes ...EmptyNode) RouteHandlerName {
	return DeleteRHN(EmptyNodeSet(nodes).ChildNodeSet()...)
}

func PostNewChildRHN(nodes ...EmptyNode) RouteHandlerName {
	routeName := NewChildRouteName(nodes...)
	return RouteHandlerName(POST + routeName)
}

func ShowNodeDetailsRHN(nodes ...EmptyNode) RouteHandlerName {
	return RouteHandlerName(GET + NodeRouteName(nodes...))
}

func ShowParentDetailsRHN(nodes ...EmptyNode) RouteHandlerName {
	return ShowNodeDetailsRHN(EmptyNodeSet(nodes).ParentNodeSet()...)
}

func ShowChildDetailsRHN(nodes ...EmptyNode) RouteHandlerName {
	return ShowNodeDetailsRHN(EmptyNodeSet(nodes).ChildNodeSet()...)
}

func ShowEditNodeFileRHN(nodes ...EmptyNode) RouteHandlerName {
	name := NodeEditFileRoutePath(nodes...)
	rhn := RouteHandlerName(GET + name)
	return rhn
}

func PostEditNodeFileRHN(nodes ...EmptyNode) RouteHandlerName {
	name := NodeEditFileRoutePath(nodes...)
	rhn := RouteHandlerName(POST + name)
	return rhn
}

func ShowNodeFilesRHN(nodes ...EmptyNode) RouteHandlerName {
	name := NodeFilesRoutePath(nodes...)
	rhn := RouteHandlerName(GET + name)
	return rhn
}

func PostNodeFilesRHN(nodes ...EmptyNode) RouteHandlerName {
	name := NodeFilesRoutePath(nodes...)
	rhn := RouteHandlerName(POST + name)
	return rhn
}

func ShowEditRHN(nodes ...EmptyNode) RouteHandlerName {
	routeName := EditNodeRouteName(nodes...)
	rhn := RouteHandlerName(GET + routeName)
	return rhn
}

func PostEditRHN(nodes ...EmptyNode) RouteHandlerName {
	routeName := EditNodeRouteName(nodes...)
	rhn := RouteHandlerName(POST + routeName)
	return rhn
}

func DeleteRHN(nodes ...EmptyNode) RouteHandlerName {
	routeName := NodeRouteName(nodes...)
	rhn := RouteHandlerName(DELETE + routeName)
	return rhn
}

func ShowEditHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RoutePath:   EditNodeRouteName(nodes...),
		HandlerName: ShowEditRHN(nodes...),
		Method:      GET,
		HandlerFunc: handlerFunc,
	}

}

func PostEditHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RoutePath:   EditNodeRouteName(nodes...),
		HandlerName: PostEditRHN(nodes...),
		Method:      POST,
		HandlerFunc: handlerFunc,
	}
}

func ShowFilesHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := NodeFilesRoutePath(nodes...)
	rhn := ShowNodeFilesRHN(nodes...)
	return RouteHandler{
		RoutePath:   routeName,
		HandlerName: rhn,
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}

func ShowEditFileHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := NodeEditFileRoutePath(nodes...)
	rhn := ShowEditNodeFileRHN(nodes...)
	return RouteHandler{
		RoutePath:   routeName,
		HandlerName: rhn,
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}

func PostEditFileHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	path := NodeEditFileRoutePath(nodes...)
	rhn := PostEditNodeFileRHN(nodes...)
	return RouteHandler{
		RoutePath:   path,
		HandlerName: rhn,
		Method:      POST,
		HandlerFunc: handlerFunc,
	}
}

func ViewFilesHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := ViewNodeFilesRoutePath(nodes...)
	rhn := ViewNodeFilesRHN(nodes...)
	return RouteHandler{
		RoutePath:   routeName,
		HandlerName: rhn,
		Method:      GET,
		HandlerFunc: handlerFunc,
	}

}

func PostFileHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := NodeFilesRoutePath(nodes...)
	rhn := PostNodeFilesRHN(nodes...)
	return RouteHandler{
		RoutePath:   routeName,
		HandlerName: rhn,
		Method:      POST,
		HandlerFunc: handlerFunc,
	}
}

func ShowNodeDetailsHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RoutePath:   NodeRouteName(nodes...),
		HandlerName: ShowNodeDetailsRHN(nodes...),
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}

func ListChildrenHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := ChildNodesRouteName(nodes...)
	rhn := ListChildrenRHN(nodes...)
	return RouteHandler{
		RoutePath:   routeName,
		HandlerName: rhn,
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}

func ShowNewChildHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RoutePath:   NewChildRouteName(nodes...),
		HandlerName: ShowNewChildRHN(nodes...),
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}
func PostNewChildHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RoutePath:   NewChildRouteName(nodes...),
		HandlerName: PostNewChildRHN(nodes...),
		Method:      POST,
		HandlerFunc: handlerFunc,
	}
}

func DeleteHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RoutePath:   NodeRouteName(nodes...),
		HandlerName: DeleteRHN(nodes...),
		Method:      DELETE,
		HandlerFunc: handlerFunc,
	}
}

// to be used only if *all* NodeRouter methods are implemented.
func NodeHandlers(router NodeRouter) []RouteHandler {
	r := router.GetRouter()
	var handlers = []RouteHandler{
		ListChildrenHandler(router.ListChildren, r.emptyNodeSet...),
		ShowNewChildHandler(router.ShowNewChild, r.emptyNodeSet...),
		PostNewChildHandler(router.PostNewChild, r.emptyNodeSet...),
		ShowFilesHandler(router.ShowFiles, r.emptyNodeSet...),
		ViewFilesHandler(router.ViewFile, r.emptyNodeSet...),
		PostFileHandler(router.PostFile, r.emptyNodeSet...),
		ShowNodeDetailsHandler(router.ShowDetails, r.emptyNodeSet...),
		ShowEditHandler(router.ShowEdit, r.emptyNodeSet...),
		ShowEditFileHandler(router.ShowEditFile, r.emptyNodeSet...),
		PostEditFileHandler(router.PostEditFile, r.emptyNodeSet...),
		PostEditHandler(router.PostEdit, r.emptyNodeSet...),
		DeleteHandler(router.Delete, r.emptyNodeSet...),
	}
	return handlers
}

func ShowDetailsURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ShowNodeDetailsRHN(r.emptyNodeSet...)), r.params.ToSlice()...)
}

func ListChildrenURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ListChildrenRHN(r.emptyNodeSet...)), r.params.ToSlice()...)
}

func ListSiblingsURL(router NodeRouter) string {
	r := router.GetRouter()
	log.Println("router nodeSet", r.emptyNodeSet)
	return r.app.Reverse(string(ListChildrenRHN(EmptyNodeSet(r.emptyNodeSet).ParentNodeSet()...)), r.params.ToSlice()...)
}

func ShowNewChildURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ShowNewChildRHN(r.emptyNodeSet...)), r.params.ToSlice()...)
}

func PostNewChildURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(PostNewChildRHN(r.emptyNodeSet...)), r.params.ToSlice()...)
}

func ShowFilesURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ShowNodeFilesRHN(r.emptyNodeSet...)), AddParams(r.params, "")...)
}

func ShowEditFileURL(handler NodeRouter) func(relPath string) string {
	r := handler.GetRouter()
	return func(relPath string) string {
		return r.app.Reverse(string(ShowEditNodeFileRHN(r.emptyNodeSet...)), AddParams(r.params, relPath)...)
	}
}
func PostEditFileURL(handler NodeRouter) func(relPath string) string {
	r := handler.GetRouter()
	return func(relPath string) string {
		return r.app.Reverse(string(PostEditNodeFileRHN(r.emptyNodeSet...)), AddParams(r.params, relPath)...)
	}
}

func ShowEditNodeURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ShowEditRHN(r.emptyNodeSet...)), r.params.ToSlice()...)
}

func PostEditNodeURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(PostEditRHN(r.emptyNodeSet...)), r.params.ToSlice()...)
}

func NodeDetailsPage(router NodeRouter, isEdit bool) mt.NodeDetailsPage {
	r := router.GetRouter()
	var listChildrenURL string
	if _, ok := r.nodes.CurrentNode().(domain.Lesson); !ok {
		listChildrenURL = ListChildrenURL(router)
	}
	node := r.nodes.CurrentNode()
	var calURL string
	_, isCourse := node.(domain.Course)
	_, isUnit := node.(domain.Unit)
	_, isLesson := node.(domain.Lesson)
	if isCourse || isUnit || isLesson {
		calURL = r.app.Reverse(string(ShowCourseCalendar), r.params.ToSlice()...)
	}
	return mt.NodeDetailsPage{
		Params:            r.params,
		Node:              r.nodes.CurrentNode(),
		CourseCalendarURL: calURL,
		GetEditNodeURL:    ShowEditNodeURL(router),
		PostEditNodeURL:   PostEditNodeURL(router),
		ListChildrenURL:   listChildrenURL,
		UpNavURL:          ListSiblingsURL(router),
		CancelEditURL:     ShowDetailsURL(router),
		ServerFilesURL:    ShowFilesURL(router),
		BreadCrumbsData:   BreadCrumbs(r.app, r.params, r.nodes.ToSlice()...),
		IsEdit:            isEdit,
	}
}

func NodeListPage(r NodeRouter) mt.NodeListPage {
	router := r.GetRouter()
	var upNavURL string
	if _, ok := router.nodes.CurrentNode().(domain.User); ok {
		upNavURL = ShowDetailsURL(r)
	} else {
		upNavURL = ListSiblingsURL(r)
	}
	var childChildrenRHN string
	if router.nodes.CurrentNode().ChildTypeName() != domain.LessonTypeName.String() {
		childChildrenRHN = string(ListChildChildrenRHN(router.emptyNodeSet...))
	}
	return mt.NodeListPage{
		Params:           router.params,
		ParentNode:       router.nodes.CurrentNode(),
		Children:         router.nodes.CurrentNode().Children(),
		ChildDetailsRHN:  string(ShowChildDetailsRHN(router.emptyNodeSet...)),
		ShowNewChildURL:  ShowNewChildURL(r),
		ChildChildrenRHN: childChildrenRHN,
		DeleteChildRHN:   string(DeleteChildRHN(router.emptyNodeSet...)),
		UpNavURL:         upNavURL,
		E:                router.app,
		BreadCrumbsData:  BreadCrumbs(router.app, router.params, router.nodes.ToSlice()...),
	}
}

func NodeCreateChildPage(h NodeRouter) mt.NodeCreatePage {
	r := h.GetRouter()
	return mt.NodeCreatePage{
		ParentNode:        r.nodes.CurrentNode(),
		NodeType:          domain.NodeTypeName(r.nodes.CurrentNode().ChildTypeName()),
		Params:            r.params,
		PostCreateNodeURL: PostNewChildURL(h),
		CancelURL:         ListChildrenURL(h),
		BreadCrumbsData:   BreadCrumbs(r.app, r.params, r.nodes.ToSlice()...),
	}
}

func NodeFilesPage(router NodeRouter, path string, files []mt.FilesPageItem) mt.FilesPage {
	r := router.GetRouter()
	return mt.FilesPage{
		Root: path == "",
		ParentDirectory: mt.FilesPageItem{
			Name:  filepath.Base(filepath.Dir(path)),
			URL:   string(ShowNodeFilesRHN(r.emptyNodeSet...).String()),
			IsDir: true,
		},
		Node:                r.nodes.CurrentNode(),
		Params:              r.params,
		CurrentPath:         path,
		OpenFileRHN:         ShowNodeFilesRHN(r.emptyNodeSet...).String(),
		ViewMarkdownRHN:     ViewNodeFilesRHN(r.emptyNodeSet...).String(),
		EditMarkdownFileURL: ShowEditFileURL(router),
		Files:               files,
		E:                   r.app,
		BreadCrumbsData:     BreadCrumbs(r.app, r.params, r.nodes.ToSlice()...),
	}
}

func ShowDetails(c echo.Context, router NodeRouter) error {
	r := router.GetRouter()
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	router.SetRouter(r)
	page := NodeDetailsPage(router, false)
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)

}

func ShowFiles(c echo.Context, r NodeRouter) error {
	var router = r.GetRouter()
	path := c.Param("*")
	log.Println("path: ", path)
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	router.params = params
	nodes, err := router.svc.Nodes(params)
	if err != nil {
		return err
	}
	router.nodes = nodes
	err = router.svc.CreateNodeFilesDir(router.nodes.ToSlice()...)
	if err != nil {
		return err
	}
	isDir, err := router.svc.IsDir(path, router.nodes.ToSlice()...)
	if err != nil {
		return err
	}
	if !isDir {
		c.Attachment(router.svc.NodeFilePath(path, router.nodes.ToSlice()...), filepath.Base(path))
	}
	files, err := router.svc.NodeFiles(path, router.nodes.ToSlice()...)
	for _, file := range files {
		log.Println(file.Path)
	}
	if err != nil {
		return err
	}
	r.SetRouter(router)
	page := NodeFilesPage(r, path, files)
	component := page.Component()
	layout := CourseManagerLayout(router.app, component, router.nodes.User)
	return Respond(c, "", component, layout)
}

func ShowEditFile(c echo.Context, router NodeRouter, redirect string) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	path := c.Param("*")
	log.Println("path: ", path)
	path, err = url.PathUnescape(path)
	if err != nil {
		return err
	}
	log.Println("decoded: ", path)
	r := router.GetRouter()
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	root := data.NodeFilesDirPath(nodes.ToSlice()...)
	path = filepath.Join(root, path)
	r.nodes = nodes
	router.SetRouter(r)
	markdownFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return err
	}
	defer markdownFile.Close()
	bytes, err := io.ReadAll(markdownFile)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(len(bytes), "bytes read")
	relPath, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}
	log.Println(string(bytes))
	component := mt.MarkdownEditor{
		Path:            relPath,
		Params:          r.params,
		Contents:        string(bytes),
		PostEditFileURL: PostEditFileURL(router),
		E:               r.app,
	}.Component()
	return Respond(c, ShowDetailsURL(router), component, nil)

}

func PostEditFile(c echo.Context, router NodeRouter, redirect string) error {
	log.Println("PostEditFile")
	r := router.GetRouter()
	log.Println(r.nodes.CurrentNode())
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	path := c.Param("*")
	log.Println("path: ", path)
	path, err = url.PathUnescape(path)
	if err != nil {
		return err
	}
	log.Println("decoded: ", path)
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	router.SetRouter(r)
	content := c.FormValue(string(mt.EditSlidesTextAreaID))
	log.Println("content", content)
	err = r.svc.WriteToMarkdown(path, content, nodes)
	if err != nil {
		return err
	}
	return c.Redirect(303, ShowDetailsURL(router))
}

// redirect is for non-htmx requests
func ViewFile(c echo.Context, router NodeRouter, redirect string) error {
	r := router.GetRouter()
	path := c.Param("*")
	log.Println("path: ", path)
	path, err := url.PathUnescape(path)
	if err != nil {
		return err
	}
	log.Println("decoded: ", path)
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	err = r.svc.CreateNodeFilesDir(nodes.ToSlice()...)
	if err != nil {
		return err
	}
	pathRoot := data.NodeFilesDirPath(nodes.ToSlice()...)
	path = filepath.Join(pathRoot, path)
	content, err := r.svc.RenderMarkdownFile(path)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	data := mt.MarkdownDocument{
		Title:   filepath.Base(path),
		Content: string(content),
		Static:  false,
	}
	err = mt.DocLayout(data).Render(context.Background(), &buf)
	if err != nil {
		return err
	}
	data.Content = buf.String()
	component := mt.MarkdownIFrame(data)
	return Respond(c, redirect, component, nil)

}

func PostFile(c echo.Context, router NodeRouter) error {
	r := router.GetRouter()
	path := c.Param("*")
	log.Println("path: ", path)
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := r.svc.Nodes(params)
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
	err = r.svc.WriteFile(file, path, nodes)
	if err != nil {
		return err
	}
	// Respond to the client
	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", file.Filename))

}
