package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseCalendar struct {
	Params                        CourseIDParams
	Course                        domain.Course
	TermDetailsURL                string
	CourseDetailsURL              string
	LessonDetailsRouteHandlerName string
	ShiftLessonRouteHandlerName   string
	ListTermCoursesRHN            string
	ShowAddLessonDateRHN          string
	RemoveLessonDateRHN           string
	E                             *echo.Echo
}

func (data CourseCalendar) Component() templ.Component {
	return CourseCalendarTemplate(data)
}

func (data CourseCalendar) CalendarLessonContainer(lesson domain.Lesson, date time.Time) templ.Component {
	params := data.Params
	params.UnitID.Value = lesson.UnitID
	params.LessonID.Value = lesson.ID
	container := CalendarLessonContainerNew{
		Date:                date,
		Params:              params,
		lesson:              lesson,
		LessonDetailsURL:    data.E.Reverse(data.LessonDetailsRouteHandlerName, params.ToIntSlice()...),
		Course:              data.Course,
		ShiftLessonRHN:      data.ShiftLessonRouteHandlerName,
		RemoveLessonDateURL: data.E.Reverse(data.RemoveLessonDateRHN, AddParam(params, date.Format(time.DateOnly))...),
		E:                   data.E,
	}
	return container.Component()
}

func (data CourseCalendar) ShowAddLessonDatePageURL(date time.Time) string {
	return data.E.Reverse(
		data.ShowAddLessonDateRHN,
		data.Params.TermID.Value,
		data.Params.CourseID.Value,
		date.Format(time.DateOnly),
	)
}

func (data CourseCalendar) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: data.Course.Name + " Course Calendar",
		UpNav: UpNav{
			URL:  data.E.Reverse(data.ListTermCoursesRHN, data.Params.ToIntSlice()...),
			Text: "Back to Courses",
		},
	}
}

type CalendarLessonContainerNew struct {
	Params              CourseIDParams
	Date                time.Time
	Course              domain.Course
	lesson              domain.Lesson
	LessonDetailsURL    string
	RemoveLessonDateURL string
	ShiftLessonRHN      string
	E                   *echo.Echo
}

func (data CalendarLessonContainerNew) Component() templ.Component {
	return CalendarLessonContainer(data)
}

func (data CourseCalendar) BreadCrumbs() BreadCrumbs {
	return BreadCrumbs{
		Term:             data.Course.Term,
		TermDetailsURL:   data.TermDetailsURL,
		Course:           data.Course,
		CourseDetailsURL: data.CourseDetailsURL,
	}
}

func (data CalendarLessonContainerNew) ShiftButton(cd domain.CalendarDirection) templ.Component {
	button := ShiftButton{
		Direction:      cd,
		Params:         data.Params,
		ShiftLessonURL: data.E.Reverse(data.ShiftLessonRHN, AddParam(data.Params, cd.String())...),
		e:              data.E,
	}
	return button.Component()
}

type AddLessonToDatePage struct {
	Date             time.Time
	Params           CourseIDParams
	Course           domain.Course
	AddLessonDateRHN string
	E                *echo.Echo
}

func (page AddLessonToDatePage) Component() templ.Component {
	return AddLessonToDatePageComponent(page)

}

func (data AddLessonToDatePage) AddLessonDateURL(unitID, lessonID int) string {
	data.Params.UnitID.Value = unitID
	data.Params.LessonID.Value = lessonID
	return data.E.Reverse(data.AddLessonDateRHN, data.Params.ToIntSlice()...)
}
