package handlers

import (
	managertemplates "gh_static_portfolio/internal/newtemplates/app"
	templates "gh_static_portfolio/internal/newtemplates/shared"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
)

func BreadCrumbs(nodes templates.Nodes, path routes.NodePath, reverse web.Reverse) managertemplates.BreadCrumbs {
	return managertemplates.BreadCrumbs{
		Nodes:            nodes,
		UserDetailsURL:   userDetailsURL(path, reverse),
		TermDetailsURL:   termDetailsURL(path, reverse),
		CourseDetailsURL: courseDetailsURL(path, reverse),
		UnitDetailsURL:   unitDetailsURL(path, reverse),
		LessonDetailsURL: lessonDetailsURL(path, reverse),
	}

}

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
