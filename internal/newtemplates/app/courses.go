package managertemplates

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/core/standard"
	"gh_static_portfolio/internal/core/term"
	"gh_static_portfolio/internal/shared/web"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CoursesListPage struct {
	ShowCourseCalendarURL web.AddParams
	ShowAssessmentsURL    web.AddParams
	NodeListPage
}

func (page CoursesListPage) Component() templ.Component {
	var calendarButtons []ComponentData
	var assessmentButtons []ComponentData

	for _, course := range page.Children {
		button := ShowCalendarButton{
			ShowCalendarURL: page.ShowCourseCalendarURL(course.GetID()),
		}
		calendarButtons = append(calendarButtons, button)
		button2 := ShowAssessmentsButton{
			ShowAssessmentsURL: page.ShowAssessmentsURL(course.GetID()),
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
	Terms                   []term.Term
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
	StandardSets             []standard.StandardSet
}

func (page CourseDetailsPage) Course() dto.Course {
	return page.Node.(dto.Course)
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
