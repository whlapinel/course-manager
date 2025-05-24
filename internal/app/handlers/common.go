package handlers

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	managertemplates "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"

	"github.com/labstack/echo/v4"
)

func BreadCrumbs(nodes ports.Nodes, path routes.NodePath, reverse web.Reverse) appcomponents.BreadCrumbs {
	return managertemplates.BreadCrumbs{
		Nodes:            nodes,
		UserDetailsURL:   userDetailsURL(path, reverse),
		TermDetailsURL:   termDetailsURL(path, reverse),
		CourseDetailsURL: courseDetailsURL(path, reverse),
		UnitDetailsURL:   unitDetailsURL(path, reverse),
		LessonDetailsURL: lessonDetailsURL(path, reverse),
	}

}
// func BreadCrumbs2(nodes ports.Nodes, path routes.NodePath, reverse web.Reverse) appcomponents.BreadCrumbs {
// 	return appcomponents.BreadCrumbs{
// 		Nodes:            nodes,
// 		UserDetailsURL:   userDetailsURL(path, reverse),
// 		TermDetailsURL:   termDetailsURL(path, reverse),
// 		CourseDetailsURL: courseDetailsURL(path, reverse),
// 		UnitDetailsURL:   unitDetailsURL(path, reverse),
// 		LessonDetailsURL: lessonDetailsURL(path, reverse),
// 	}

// }

func userDetailsURL(path routes.NodePath, reverse web.Reverse) string {
	return reverse(routes.GetUser.String(), path.ToSlice()...)
}

func termDetailsURL(path routes.NodePath, reverse web.Reverse) string {
	return reverse(routes.GetTerm.String(), path.ToSlice()...)
}

func courseDetailsURL(path routes.NodePath, reverse web.Reverse) string {
	return reverse(routes.GetCourse.String(), path.ToSlice()...)
}

func unitDetailsURL(path routes.NodePath, reverse web.Reverse) string {
	return reverse(routes.GetUnit.String(), path.ToSlice()...)
}

func lessonDetailsURL(path routes.NodePath, reverse web.Reverse) string {
	return reverse(routes.GetLesson.String(), path.ToSlice()...)
}

type nodeInfo struct {
	ports.Nodes
	routes.NodePath
}

func parseNodeInfo(c echo.Context, service *services.NodeService) (nodeInfo, error) {
	var info nodeInfo
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return info, err
	}
	nodes, err := service.Nodes(path)
	if err != nil {
		return info, err
	}
	return nodeInfo{
		Nodes:    nodes,
		NodePath: path,
	}, nil

}
