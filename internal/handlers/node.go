package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"gh_static_portfolio/internal/util"
	"log"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type NodeRouter interface {
	ListChildren(echo.Context) error
	ShowFiles(echo.Context) error
	PostFile(echo.Context) error
	ViewFile(echo.Context) error
	ShowDetails(echo.Context) error
	ShowNewChild(echo.Context) error // e.g. if node is course, new child would be new unit
	PostNewChild(echo.Context) error
	ShowEdit(echo.Context) error
	PostEdit(echo.Context) error
	Delete(echo.Context) error // i.e. delete node itself (not child)
	GetRouter() Router
}

type EmptyNode domain.CourseNode

type EmptyNodeSet []EmptyNode

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
func ChildNodesRouteName(nodes ...EmptyNode) RouteName {
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
	return RouteName(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id"
func NodeRouteName(nodes ...EmptyNode) RouteName {
	var path string = "/"
	for _, node := range nodes {
		path = filepath.Join(path, util.KebabCase(node.TypeName()+"s"))
		path = filepath.Join(path, fmt.Sprintf("/:%s-id", util.KebabCase(node.TypeName())))
	}
	return RouteName(path)

}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/new
func NewChildRouteName(nodes ...EmptyNode) RouteName {
	var name = ChildNodesRouteName(nodes...)
	var path = string(name)
	path = filepath.Join(path, "/new")
	return RouteName(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id/edit
func EditNodeRouteName(nodes ...EmptyNode) RouteName {
	var path = string(NodeRouteName(nodes...))
	path = filepath.Join(path, "/edit")
	return RouteName(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id/files/*
func NodeFilesRouteName(nodes ...EmptyNode) RouteName {
	var name = NodeRouteName(nodes...)
	path := string(name)
	path = filepath.Join(path, "/files/*")
	return RouteName(path)
}

func ViewNodeFilesRouteName(nodes ...EmptyNode) RouteName {
	var name = NodeRouteName(nodes...)
	path := string(name)
	path = filepath.Join(path, "/view-markdown/files/*")
	return RouteName(path)
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

func ShowNodeFilesRHN(nodes ...EmptyNode) RouteHandlerName {
	name := NodeFilesRouteName(nodes...)
	rhn := RouteHandlerName(GET + name)
	return rhn
}

func ViewNodeFilesRHN(nodes ...EmptyNode) RouteHandlerName {
	name := ViewNodeFilesRouteName(nodes...)
	rhn := RouteHandlerName(GET + name)
	return rhn
}

func PostNodeFilesRHN(nodes ...EmptyNode) RouteHandlerName {
	name := NodeFilesRouteName(nodes...)
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
		RouteName:   EditNodeRouteName(nodes...),
		HandlerName: ShowEditRHN(nodes...),
		Method:      GET,
		HandlerFunc: handlerFunc,
	}

}

func PostEditHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RouteName:   EditNodeRouteName(nodes...),
		HandlerName: PostEditRHN(nodes...),
		Method:      POST,
		HandlerFunc: handlerFunc,
	}
}

func ShowFilesHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := NodeFilesRouteName(nodes...)
	rhn := ShowNodeFilesRHN(nodes...)
	return RouteHandler{
		RouteName:   routeName,
		HandlerName: rhn,
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}

func ViewFilesHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := ViewNodeFilesRouteName(nodes...)
	rhn := ViewNodeFilesRHN(nodes...)
	return RouteHandler{
		RouteName:   routeName,
		HandlerName: rhn,
		Method:      GET,
		HandlerFunc: handlerFunc,
	}

}

func PostFileHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := NodeFilesRouteName(nodes...)
	rhn := PostNodeFilesRHN(nodes...)
	return RouteHandler{
		RouteName:   routeName,
		HandlerName: rhn,
		Method:      POST,
		HandlerFunc: handlerFunc,
	}
}

func ShowNodeDetailsHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RouteName:   NodeRouteName(nodes...),
		HandlerName: ShowNodeDetailsRHN(nodes...),
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}

func ListChildrenHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	routeName := ChildNodesRouteName(nodes...)
	rhn := ListChildrenRHN(nodes...)
	return RouteHandler{
		RouteName:   routeName,
		HandlerName: rhn,
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}

func ShowNewChildHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RouteName:   NewChildRouteName(nodes...),
		HandlerName: ShowNewChildRHN(nodes...),
		Method:      GET,
		HandlerFunc: handlerFunc,
	}
}
func PostNewChildHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RouteName:   NewChildRouteName(nodes...),
		HandlerName: PostNewChildRHN(nodes...),
		Method:      POST,
		HandlerFunc: handlerFunc,
	}
}

func DeleteHandler(handlerFunc echo.HandlerFunc, nodes ...EmptyNode) RouteHandler {
	return RouteHandler{
		RouteName:   NodeRouteName(nodes...),
		HandlerName: DeleteRHN(nodes...),
		Method:      DELETE,
		HandlerFunc: handlerFunc,
	}
}

// to be used only if *all* NodeRouter methods are implemented.
func NodeHandlers(router NodeRouter) []RouteHandler {
	r := router.GetRouter()
	var handlers = []RouteHandler{
		ListChildrenHandler(router.ListChildren, r.nodeSet...),
		ShowNewChildHandler(router.ShowNewChild, r.nodeSet...),
		PostNewChildHandler(router.PostNewChild, r.nodeSet...),
		ShowFilesHandler(router.ShowFiles, r.nodeSet...),
		ViewFilesHandler(router.ViewFile, r.nodeSet...),
		PostFileHandler(router.PostFile, r.nodeSet...),
		ShowNodeDetailsHandler(router.ShowDetails, r.nodeSet...),
		ShowEditHandler(router.ShowEdit, r.nodeSet...),
		PostEditHandler(router.PostEdit, r.nodeSet...),
		DeleteHandler(router.Delete, r.nodeSet...),
	}
	return handlers
}

func ShowDetailsURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ShowNodeDetailsRHN(r.nodeSet...)), r.params.ToSlice()...)
}

func ListChildrenURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ListChildrenRHN(r.nodeSet...)), r.params.ToSlice()...)
}

func ListSiblingsURL(router NodeRouter) string {
	r := router.GetRouter()
	log.Println("router nodeSet", r.nodeSet)
	return r.app.Reverse(string(ListChildrenRHN(EmptyNodeSet(r.nodeSet).ParentNodeSet()...)), r.params.ToSlice()...)
}

func ShowNewChildURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ShowNewChildRHN(r.nodeSet...)), r.params.ToSlice()...)
}

func PostNewChildURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(PostNewChildRHN(r.nodeSet...)), r.params.ToSlice()...)
}

func ShowFilesURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ShowNodeFilesRHN(r.nodeSet...)), r.params.ToSlice("")...)
}

func ShowEditNodeURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(ShowEditRHN(r.nodeSet...)), r.params.ToSlice()...)
}

func PostEditNodeURL(handler NodeRouter) string {
	r := handler.GetRouter()
	return r.app.Reverse(string(PostEditRHN(r.nodeSet...)), r.params.ToSlice()...)
}

func NodeDetailsPage(h NodeRouter, isEdit bool) mt.NodeDetailsPage {
	r := h.GetRouter()
	return mt.NodeDetailsPage{
		Params:          r.params,
		Node:            r.node,
		GetEditNodeURL:  ShowEditNodeURL(h),
		PostEditNodeURL: PostEditNodeURL(h),
		ListChildrenURL: ListChildrenURL(h),
		UpNavURL:        ListSiblingsURL(h),
		ServerFilesURL:  ShowFilesURL(h),
		BreadCrumbsData: BreadCrumbs(r.app, r.params, r.ancestors...),
		IsEdit:          isEdit,
	}
}

func NodeListPage(h NodeRouter) mt.NodeListPage {
	r := h.GetRouter()

	var upNavURL string
	if _, ok := r.node.(domain.User); ok {
		upNavURL = ShowDetailsURL(h)
	} else {
		upNavURL = ListSiblingsURL(h)
	}
	return mt.NodeListPage{
		Params:          r.params,
		ParentNode:      r.node,
		Children:        r.node.Children(),
		ChildDetailsRHN: string(ShowChildDetailsRHN(r.nodeSet...)),
		ShowNewChildURL: ShowNewChildURL(h),
		// ShowNewChildURL:  string(ShowNewChildRHN(h.NodeSet()...)),
		ChildChildrenRHN: string(ListChildChildrenRHN(r.nodeSet...)),
		DeleteChildRHN:   string(DeleteChildRHN(r.nodeSet...)),
		UpNavURL:         upNavURL,
		E:                r.app,
		BreadCrumbsData:  BreadCrumbs(r.app, r.params, r.ancestors...),
	}
}

func NodeCreateChildPage(h NodeRouter) mt.NodeCreatePage {
	r := h.GetRouter()
	log.Println("NodeCreateChildPage():", PostNewChildURL(h))
	// log.Println("NodeCreateChildPage(): ", h.Node().TypeName(), h.Node().GetName())
	return mt.NodeCreatePage{
		ParentNode:        r.node,
		NodeType:          domain.NodeTypeName(r.node.ChildTypeName()),
		Params:            r.params,
		PostCreateNodeURL: PostNewChildURL(h),
		CancelURL:         ListChildrenURL(h),
		BreadCrumbsData:   BreadCrumbs(r.app, r.params, r.ancestors...),
	}
}

func NodeFilesPage(r NodeRouter, path string, files []mt.FilesPageItem) mt.FilesPage {
	return mt.FilesPage{
		Node:            r.GetRouter().node,
		Params:          r.GetRouter().params,
		CurrentPath:     path,
		OpenFileRHN:     ShowUnitFiles.String(),
		ViewMarkdownRHN: GetUnitViewMarkdown.String(),
		Files:           files,
		E:               r.GetRouter().app,
		BreadCrumbsData: BreadCrumbs(r.GetRouter().app, r.GetRouter().params, r.GetRouter().ancestors...),
	}
}
