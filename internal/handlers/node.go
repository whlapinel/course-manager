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

type NodeHandler interface {
	ListChildren(echo.Context) error
	ShowFiles(echo.Context) error
	ViewFile(echo.Context) error
	ShowDetails(echo.Context) error
	ShowNewChild(echo.Context) error // e.g. if node is course, new child would be new unit
	PostNewChild(echo.Context) error
	ShowEdit(echo.Context) error
	PostEdit(echo.Context) error
	Delete(echo.Context) error // i.e. delete node itself (not child)
	Router() *echo.Echo
	Node() domain.CourseNode
	AncestorPath() []domain.CourseNode
	NodeSet() []EmptyNode
	Params() mt.CourseIDParams
}

type EmptyNode domain.CourseNode

type EmptyNodeSet []EmptyNode

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
var EmptyNodesLessons = emptyNodes[:5]
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

func ListChildrenRHN(nodes ...EmptyNode) RouteHandlerName {
	name := ChildNodesRouteName(nodes...)
	rhn := RouteHandlerName(GET + name)
	return rhn
}

func ShowNewChildRHN(nodes ...EmptyNode) RouteHandlerName {
	routeName := NewChildRouteName(nodes...)
	return RouteHandlerName(GET + routeName)
}

func PostNewChildRHN(nodes ...EmptyNode) RouteHandlerName {
	routeName := NewChildRouteName(nodes...)
	return RouteHandlerName(POST + routeName)
}

func ShowNodeDetailsRHN(nodes ...EmptyNode) RouteHandlerName {
	return RouteHandlerName(GET + NodeRouteName(nodes...))
}

func ShowNodeFilesRHN(nodes ...EmptyNode) RouteHandlerName {
	name := NodeFilesRouteName(nodes...)
	rhn := RouteHandlerName(GET + name)
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
		Method:      GET,
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

func NodeHandlers(handler NodeHandler, nodes ...EmptyNode) []RouteHandler {
	var handlers = []RouteHandler{
		ListChildrenHandler(handler.ListChildren, nodes...),
		ShowNewChildHandler(handler.ShowNewChild, nodes...),
		PostNewChildHandler(handler.PostNewChild, nodes...),
		ShowFilesHandler(handler.ShowFiles, nodes...),
		ShowNodeDetailsHandler(handler.ShowDetails, nodes...),
		ShowEditHandler(handler.ShowEdit, nodes...),
		PostEditHandler(handler.PostEdit, nodes...),
		DeleteHandler(handler.Delete, nodes...),
	}
	return handlers
}

func ShowDetailsURL(handler NodeHandler) string {
	return handler.Router().Reverse(string(ShowNodeDetailsRHN(handler.NodeSet()...)), handler.Params().ToSlice()...)
}

func ListChildrenURL(handler NodeHandler) string {
	return handler.Router().Reverse(string(ListChildrenRHN(handler.NodeSet()...)), handler.Params().ToSlice()...)
}

func ListSiblingsURL(handler NodeHandler) string {
	return handler.Router().Reverse(string(ListChildrenRHN(handler.NodeSet()[:len(handler.NodeSet())-1]...)), handler.Params().ToSlice()...)
}

func ShowFilesURL(handler NodeHandler) string {
	return handler.Router().Reverse(string(ShowNodeFilesRHN(handler.NodeSet()...)), handler.Params().ToSlice()...)
}

func ShowEditNodeURL(handler NodeHandler) string {
	return handler.Router().Reverse(string(ShowEditRHN(handler.NodeSet()...)), handler.Params().ToSlice()...)
}

func PostEditNodeURL(handler NodeHandler) string {
	return handler.Router().Reverse(string(PostEditRHN(handler.NodeSet()...)), handler.Params().ToSlice()...)
}

func NodeDetailsPage(h NodeHandler, isEdit bool) mt.NodeDetailsPage {
	return mt.NodeDetailsPage{
		Params:          h.Params(),
		Node:            h.Node(),
		GetEditNodeURL:  ShowEditNodeURL(h),
		PostEditNodeURL: PostEditNodeURL(h),
		ListChildrenURL: ListChildrenURL(h),
		UpNavURL:        ListSiblingsURL(h),
		ServerFilesURL:  ShowFilesURL(h),
		BreadCrumbsData: BreadCrumbs(h.Router(), h.Params(), h.AncestorPath()...),
		IsEdit:          isEdit,
	}
}

func NodeListPage(h NodeHandler) mt.NodeListPage {
	return mt.NodeListPage{
		Params:           h.Params(),
		ParentNode:       h.Node(),
		Children:         h.Node().Children(),
		ChildDetailsRHN:  string(ShowNodeDetailsRHN(emptyNodes[:len(h.NodeSet())+1]...)),
		CreateChildRHN:   ShowNewCourse.String(),
		ChildChildrenRHN: ListCourseUnits.String(),
		DeleteChildRHN:   DeleteCourse.String(),
		UpNavURL:         h.Router().Reverse(ListTerms.String(), h.Params().ToSlice()...),
		E:                h.Router(),
		BreadCrumbsData:  BreadCrumbs(h.Router(), h.Params(), h.AncestorPath()...),
	}
}
