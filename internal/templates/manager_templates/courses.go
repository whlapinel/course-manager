package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseListPage struct {
	ShowCalendarRHN string
	NodeListPage
}

func (page CourseListPage) Component() templ.Component {
	var calendarButtons []ComponentData
	for _, course := range page.Children {
		button := ShowCalendarButton{
			ShowCalendarURL: page.E.Reverse(page.ShowCalendarRHN, course.GetParentID(), course.GetID()),
		}
		calendarButtons = append(calendarButtons, button)

	}
	page.ChildUI = calendarButtons
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
	button := HXButton{
		Text:     "Calendar",
		Method:   HxGet,
		URL:      data.ShowCalendarURL,
		HxTarget: pageElementID.Selector(),
		PushURL:  true,
	}
	return button.Component()
}
