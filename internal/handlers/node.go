package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/util"
	"path/filepath"
)

// e.g. "/users/:user-id/terms/:term-id/courses/:course-id/units"
func ChildNodesRouteName(nodes ...domain.CourseNode) RouteName {
	var path = string(NodeRouteName(nodes...))
	path = filepath.Join(path, util.KebabCase(fmt.Sprintf("%ss", nodes[len(nodes)-1].ChildTypeName())))
	return RouteName(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id"
func NodeRouteName(nodes ...domain.CourseNode) RouteName {
	var path string = "/"
	for _, node := range nodes {
		path = filepath.Join(path, util.KebabCase(node.TypeName()+"s"))
		path = filepath.Join(path, fmt.Sprintf("/:%s-id", util.KebabCase(node.TypeName())))
	}
	return RouteName(path)

}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/new
func NewNodeRouteName(nodes ...domain.CourseNode) RouteName {
	var path = string(ChildNodesRouteName(nodes...))
	path = filepath.Join(path, "/new")
	return RouteName(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id/edit
func EditNodeRouteName(nodes ...domain.CourseNode) RouteName {
	var path = string(NodeRouteName(nodes...))
	path = filepath.Join(path, "/edit")
	return RouteName(path)
}

// e.g. /users/:user-id/terms/:term-id/courses/:course-id/units/:unit-id/files/*
func NodeFilesRouteName(nodes ...domain.CourseNode) RouteName {
	var path = string(ChildNodesRouteName(nodes...))
	path = filepath.Join(path, "/files/*")
	return RouteName(path)
}

func ListChildrenRHN(nodes ...domain.CourseNode) RouteHandlerName {
	return RouteHandlerName(GET + ChildNodesRouteName(nodes...))
}

func ShowNodeDetailsRHN(nodes ...domain.CourseNode) RouteHandlerName {
	return RouteHandlerName(GET + NodeRouteName(nodes...))
}

func ShowNodeFiles(nodes ...domain.CourseNode) RouteHandlerName {
	return RouteHandlerName(GET + NodeFilesRouteName(nodes...))
}
