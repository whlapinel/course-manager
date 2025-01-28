package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/labstack/echo/v4"
)

type LessonDetailsPage struct {
	NodeDetailsPage
	E                                                           *echo.Echo
	Standards                                                   []domain.Standard
	GetObjectivesURL                                            string
	PostLessonStandardURL, DeleteLessonStandardRHN              string
	GetSlidesURL, EditSlidesURL, GithubFilesURL, ServerFilesURL string
}

func AddParam(params CourseIDParams, param interface{}) []interface{} {
	return append(params.ToIntSlice(), param)
}
func (page LessonDetailsPage) DeleteStandardURL(stdID int) string {
	params := AddParam(page.Params, stdID)
	return page.E.Reverse(page.DeleteLessonStandardRHN, params...)
}

func (page LessonDetailsPage) Lesson() domain.Lesson {
	return page.Node.(domain.Lesson)
}

type ObjectiveSelect struct {
	Objectives []domain.Standard
}
