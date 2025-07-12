package calendarviews

import (
	"fmt"
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/basecomponents"
	"gh_static_portfolio/internal/shared/web"
	"time"

	"github.com/a-h/templ"
)

type CourseMonthCalendar struct {
	dto.Term
	CourseDetailsURL         string
	TermDetailsURL           string
	Month                    time.Month
	Year                     int
	Today                    time.Time
	PresentMonthURL          string
	PrevMonthURL             string
	NextMonthURL             string
	CurrentMonthURL          string
	GetCreateOccasionURL     web.AddParams
	RemoveLessonDateURL      web.AddParams
	ShowAddLessonDatePageURL web.AddParams
	LessonDetailURL          web.AddParams
	ShiftLessonURL           web.AddParams
	Weeks                    [][]CalendarDate
	appcomponents.CourseManagerLayout
}

func (data CourseMonthCalendar) LessonContainerID(lessonID int, date time.Time) string {
	return fmt.Sprintf("lesson-id-%d-%s", lessonID, date.Format(time.DateOnly))
}

func (data CourseMonthCalendar) LessonDetailsButton(lessonID int, unitID int, text string) templ.Component {
	button := cmp.Button{
		Text:     text,
		Method:   cmp.HxGet,
		URL:      data.LessonDetailURL(lessonID, unitID),
		HxTarget: "#page",
		PushURL:  true,
	}
	return button.Component()
}

func (data CourseMonthCalendar) AddOccasionButton(date time.Time) templ.Component {
	button := cmp.Button{
		Text:     "+O",
		HxTarget: "#dialog-container",
		Method:   cmp.HxGet,
		URL:      data.GetCreateOccasionURL(date.Format(time.DateOnly)),
	}
	return button.Component()
}

func (data CourseMonthCalendar) AddLessonButton(date time.Time) templ.Component {
	button := cmp.Button{
		Text:     "+L",
		HxTarget: "#page",
		Method:   cmp.HxGet,
		URL:      data.ShowAddLessonDatePageURL(date.Format(time.DateOnly)),
	}
	return button.Component()
}
func (data CourseMonthCalendar) RemoveLessonButton(lesson dto.Lesson, date time.Time) templ.Component {
	button := cmp.Button{
		HxConfirm: "Are you sure you want to remove this lesson date? The lesson itself will not be deleted.",
		Method:    cmp.HxDelete,
		URL:       data.RemoveLessonDateURL(date.Format(time.DateOnly), lesson.ParentID, lesson.ID),
		HxTarget:  fmt.Sprintf("#%s", data.LessonContainerID(lesson.ID, date)),
		Image:     cmp.DeleteImage(),
		Class:     "bg-red-700 p-1 rounded",
	}
	return button.Component()
}

func (data CourseMonthCalendar) ShiftButton(unitID, lessonID int, cd dto.CalendarDirection, date time.Time) templ.Component {
	var image templ.Component
	switch cd {
	case dto.Left:
		image = cmp.ChevronLeft()
	case dto.Right:
		image = cmp.ChevronRight()
	}
	return cmp.Button{
		HxTarget: "#page",
		Method:   cmp.HxPost,
		URL:      data.ShiftLessonURL(unitID, lessonID, cd, date.Format(time.DateOnly)),
		Image:    image,
	}.Component()

}

func (data CourseMonthCalendar) Component() templ.Component {
	return NewNewCalendarComponent(data)
}

func (data CourseMonthCalendar) HTMXResponse() templ.Component {
	return data.Component()
}

func (data CourseMonthCalendar) NonHTMXResponse() templ.Component {
	return data.CourseManagerLayout.WithPage(data.Component())
}
