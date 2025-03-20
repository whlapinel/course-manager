package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseListPage struct {
	ShowCalendarRHN    string
	ShowAssessmentsRHN string
	NodeListPage
}

func (page CourseListPage) Component() templ.Component {
	var calendarButtons []ComponentData
	var assessmentButtons []ComponentData

	for _, course := range page.Children {
		button := ShowCalendarButton{
			ShowCalendarURL: page.E.Reverse(page.ShowCalendarRHN, AddParams(page.Params, course.GetID())...),
		}
		calendarButtons = append(calendarButtons, button)
		button2 := ShowAssessmentsButton{
			ShowAssessmentsURL: page.E.Reverse(page.ShowAssessmentsRHN, AddParams(page.Params, course.GetID())...),
		}
		assessmentButtons = append(assessmentButtons, button2)

	}
	page.ChildUI = append(page.ChildUI, calendarButtons)
	page.ChildUI = append(page.ChildUI, assessmentButtons)
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
	GetCopyCourseURL         string
	PostSelectStandardSetURL string
	StandardSets             []domain.StandardSet
}

func (page CourseDetailsPage) Course() domain.Course {
	return page.Node.(domain.Course)
}

func (page CourseDetailsPage) Component() templ.Component {
	return CourseDetailsComponent(page)
}

type ShowCalendarButton struct {
	ShowCalendarURL string
}

func (data ShowCalendarButton) Component() templ.Component {
	button := cmp.Button{
		Text:     "Calendar",
		Method:   cmp.HxGet,
		URL:      data.ShowCalendarURL,
		HxTarget: pageElementID.Selector(),
		PushURL:  true,
	}
	return button.Component()
}

type ShowAssessmentsButton struct {
	ShowAssessmentsURL string
}

func (data ShowAssessmentsButton) Component() templ.Component {
	button := cmp.Button{
		Text:     "Assessments",
		Method:   cmp.HxGet,
		URL:      data.ShowAssessmentsURL,
		HxTarget: pageElementID.Selector(),
		PushURL:  true,
	}
	return button.Component()
}
