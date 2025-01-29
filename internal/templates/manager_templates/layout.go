package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseManagerLayout struct {
	PageTitle                                  string
	Page                                       templ.Component
	ListTermsRHN, GenerateSiteRHN, SyncSiteRHN string
	E                                          *echo.Echo
}

type BreadCrumbs struct {
	Term             domain.Term
	Course           domain.Course
	Unit             domain.Unit
	Lesson           domain.Lesson
	TermDetailsURL   string
	CourseDetailsURL string
	UnitDetailsURL   string
	LessonDetailsURL string
}

func (data BreadCrumbs) Component() templ.Component {
	return BreadCrumbsComponent(data)
}
