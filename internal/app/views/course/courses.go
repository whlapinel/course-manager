package courseviews

import (
	ac "gh_static_portfolio/internal/app/components"
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
	ac.NodeListPage
	ac.CourseManagerLayout
}

func (page CoursesListPage) HTMXResponse() templ.Component {
	return page.Component()
}

func (page CoursesListPage) NonHTMXResponse() templ.Component {
	return page.CourseManagerLayout.Component2(page.Component())
}

func (page CoursesListPage) Component() templ.Component {
	var calendarButtons []ac.ComponentData
	var assessmentButtons []ac.ComponentData

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
	ac.NodeDetailsPage
	GetCopyCourseURL         string
	PostSelectStandardSetURL string
	StandardSets             []standard.StandardSet
	ac.CourseManagerLayout
}

func (p CourseDetailsPage) HTMXResponse() templ.Component {
	return p.Component()
}

func (p CourseDetailsPage) NonHTMXResponse() templ.Component {
	return p.CourseManagerLayout.Component2(p.Component())
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
		HxTarget: "#page",
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
		HxTarget: "#page",
		PushURL:  true,
	}
	return button.Component()
}

