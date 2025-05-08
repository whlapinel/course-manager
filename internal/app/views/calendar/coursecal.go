package calendarviews

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/newtemplates/components/base"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseCalendar struct {
	Nodes                 node.Nodes
	Params                routes.NodePath
	Course                dto.Course
	Term                  dto.Term
	TermDetailsURL        string
	CourseDetailsURL      string
	LessonDetailsFunc     web.AddParams
	ShiftLessonFunc       web.AddParams
	ListTermCoursesURL    string
	CreateOccasionFunc    web.AddParams
	ShowAddLessonDateFunc web.AddParams
	RemoveLessonDateFunc  web.AddParams
	E                     *echo.Echo
	CalendarDates         CalendarDates
	BreadCrumbsData       ac.BreadCrumbs
	ac.CourseManagerLayout
}

func (data CourseCalendar) HTMXResponse() templ.Component {
	return data.Component()
}

func (data CourseCalendar) NonHTMXResponse() templ.Component {
	return data.CourseManagerLayout.Component2(data.Component())
}

func (data CourseCalendar) GetCalendarDates() CalendarDates {
	return data.CalendarDates
}

func (page CourseCalendar) GetTerm() dto.Term {
	return page.Term
}

func (data CourseCalendar) Component() templ.Component {
	return NewCalendarComponent(data)
}

func (data CourseCalendar) BreadCrumbs() ac.BreadCrumbs {
	return data.BreadCrumbsData
}

func (data CourseCalendar) CalendarLessonContainer(lesson dto.Lesson, date time.Time) templ.Component {
	params := data.Params
	params.UnitID = lesson.UnitID
	params.LessonID = lesson.ID
	container := CalendarLessonContainerNew{
		Date:             date,
		Params:           params,
		Lesson:           lesson,
		LessonDetailsURL: data.LessonDetailsFunc(lesson.UnitID, lesson.ID),
		Course:           data.Course,
		ShiftLessonURL: func(cd string) string {
			return data.ShiftLessonFunc(lesson.UnitID, lesson.ID, cd)
		},
		RemoveLessonDateURL: data.RemoveLessonDateFunc(lesson.UnitID, lesson.ID, date.Format(time.DateOnly)),
	}
	return container.Component()
}

func (data CourseCalendar) ShowAddLessonDatePageURL(date time.Time) string {
	return data.ShowAddLessonDateFunc(date.Format(time.DateOnly))
}

func (data CourseCalendar) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: data.Course.Name + " Course Calendar",
		UpNav: cmp.UpNav{
			URL:  data.ListTermCoursesURL,
			Text: "Back to Courses",
		},
		Crumbs: data.BreadCrumbsData.BreadCrumbs(),
	}
}
