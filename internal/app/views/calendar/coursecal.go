package calendarviews

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/basecomponents"
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"time"

	"github.com/a-h/templ"
)

type CourseCalendar struct {
	Nodes                 ports.Nodes
	Params                routes.NodePath
	Course                dto.Course
	Term                  dto.Term
	TermDetailsURL        string
	CourseDetailsURL      string
	LessonDetailsFunc     web.AddParams
	ShiftLessonFunc       web.AddParams
	ListTermCoursesURL    string
	ShowAddLessonDateFunc web.AddParams
	RemoveLessonDateFunc  web.AddParams
	CreateOccasionURL     string
	GetEditOccasionURL    web.AddParams
	PostEditOccasionURL   web.AddParams
	DeleteOccasionURL     web.AddParams
	CalendarDates         CalendarDates
	BreadCrumbsData       ac.BreadCrumbs
	ac.CourseManagerLayout
}

func (page CourseCalendar) OccasionEditor(occasion occasion.Occasion) templ.Component {
	return OccasionEditor{
		Occasion:            occasion,
		IsEditing:           false,
		GetEditOccasionURL:  page.GetEditOccasionURL(occasion.ID),
		PostEditOccasionURL: page.PostEditOccasionURL(occasion.ID),
		DeleteOccasionURL:   page.DeleteOccasionURL(occasion.ID),
	}.Component()
}
func (data CourseCalendar) AddOccasionButton(date time.Time) templ.Component {
	return AddOccasionButton{
		Date:              date,
		CreateOccasionURL: data.CreateOccasionURL,
		FormID:            "form-" + date.Format(time.DateOnly),
	}.Component()
}

func (data CourseCalendar) HTMXResponse() templ.Component {
	return data.Component()
}

func (data CourseCalendar) NonHTMXResponse() templ.Component {
	return data.CourseManagerLayout.WithPage(data.Component())
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
	params.UnitID = lesson.ParentID
	params.LessonID = lesson.ID
	container := CalendarLessonContainerNew{
		Date:             date,
		Params:           params,
		Lesson:           lesson,
		LessonDetailsURL: data.LessonDetailsFunc(lesson.ParentID, lesson.ID),
		Course:           data.Course,
		ShiftLessonURL: func(cd, date string) string {
			return data.ShiftLessonFunc(lesson.ParentID, lesson.ID, cd, date)
		},
		RemoveLessonDateURL: data.RemoveLessonDateFunc(date.Format(time.DateOnly), lesson.ParentID, lesson.ID),
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
