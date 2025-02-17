package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseManagerLayout struct {
	PageTitle string
	User      domain.User
	Page      templ.Component
	E         *echo.Echo
}

func (cml CourseManagerLayout) Component() templ.Component {
	return CourseManagerLayoutComponent(cml)
}

type BreadCrumbs struct {
	User             domain.User
	Term             domain.Term
	Course           domain.Course
	Unit             domain.Unit
	Lesson           domain.Lesson
	UserDetailsURL   string
	TermDetailsURL   string
	CourseDetailsURL string
	UnitDetailsURL   string
	LessonDetailsURL string
}

func (data BreadCrumbs) Component() templ.Component {
	return BreadCrumbsComponent(data)
}
