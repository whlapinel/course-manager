package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseListPage struct {
	NodeListPage
}

func (page CourseListPage) Component() templ.Component {
	return CourseListComponent(page)
}

type CopyCourseData struct {
	TermID                  int
	CourseID                int
	Terms                   []domain.Term
	E                       *echo.Echo
	PostCopyCourseToTermRHN string
}

func (data CopyCourseData) Component() templ.Component {
	return CopyCourseComponent(data)
}

type CourseDetailsPage struct {
	NodeDetailsPage
	GetCopyCourseURL string
}

func (page CourseDetailsPage) Component() templ.Component {
	return CourseDetailsComponent(page)
}
