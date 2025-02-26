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
	Slides                                                       string
	E                                                            *echo.Echo
	Standards                                                    []domain.Standard
	GetObjectivesURL                                             string
	PostLessonStandardURL, DeleteLessonStandardRHN               string
	GetEditAssessmentRHN, PostAssessmentURL, DeleteAssessmentRHN string
	GetSlidesURL, EditSlidesURL                                  string
}

func (page LessonDetailsPage) DeleteStandardURL(stdID int) string {
	return page.E.Reverse(page.DeleteLessonStandardRHN, AddParams(page.Params, stdID)...)
}

func (page LessonDetailsPage) DeleteAssessmentURL(assessmentID int) string {
	return page.E.Reverse(page.DeleteAssessmentRHN, AddParams(page.Params, assessmentID)...)
}

func (page LessonDetailsPage) Lesson() domain.Lesson {
	return page.Node.(domain.Lesson)
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

type EditAssessmentForm struct {
	Params                NodePath
	Assessment            domain.Assessment
	PostEditAssessmentURL string
	LessonDetailsURL      string
}

func (data EditAssessmentForm) Component() templ.Component {
	return EditAssessmentFormComponent(data)
}
