package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type LessonDetailsPage struct {
	NodeDetailsPage
	E                                                           *echo.Echo
	Standards                                                   []domain.Standard
	GetObjectivesURL                                            string
	PostLessonStandardURL, DeleteLessonStandardRHN              string
	PostAssessmentURL, DeleteAssessmentRHN                      string
	GetSlidesURL, EditSlidesURL, GithubFilesURL, ServerFilesURL string
}

func AddParam(params CourseIDParams, param interface{}) []interface{} {
	return append(params.ToIntSlice(), param)
}
func (page LessonDetailsPage) DeleteStandardURL(stdID int) string {
	params := AddParam(page.Params, stdID)
	return page.E.Reverse(page.DeleteLessonStandardRHN, params...)
}

func (page LessonDetailsPage) DeleteAssessmentURL(assessmentID int) string {
	params := AddParam(page.Params, assessmentID)
	return page.E.Reverse(page.DeleteAssessmentRHN, params...)
}

func (page LessonDetailsPage) Lesson() domain.Lesson {
	return page.Node.(domain.Lesson)
}

func (page LessonDetailsPage) ViewSlidesButton() templ.Component {
	return HXButton{
		Text:     "View Slides",
		Method:   HxGet,
		URL:      page.GetSlidesURL,
		HxTarget: "#slides",
		PushURL:  true,
	}.Component()

}

func (page LessonDetailsPage) EditSlidesButton() templ.Component {
	return HXButton{
		Text:     "Edit Slides",
		Method:   HxGet,
		URL:      page.EditSlidesURL,
		HxTarget: EditSlidesContainerID.Selector(),
		PushURL:  true,
	}.Component()

}

func (page LessonDetailsPage) ViewFilesButton() templ.Component {
	return HXButton{
		Method:   HxGet,
		URL:      page.ServerFilesURL,
		Text:     "Files",
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page LessonDetailsPage) Component() templ.Component {
	return LessonDetailsComponent(page)
}

type ObjectiveSelect struct {
	Objectives []domain.Standard
}

// capitalizes first letter
func DisplayCategory(cat domain.AssessmentCategory) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(cat.String()[:1]), cat.String()[1:])
}
