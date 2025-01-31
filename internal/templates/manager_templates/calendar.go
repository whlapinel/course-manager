package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CourseCalendar struct {
	Course                        domain.Course
	TermDetailsURL                string
	CourseDetailsURL              string
	LessonDetailsRouteHandlerName string
	ShiftLessonRouteHandlerName   string
	ListTermCoursesRHN            string
	E                             *echo.Echo
}

func (data CourseCalendar) Component() templ.Component {
	return CourseCalendarTemplate(data)
}

func (data CourseCalendar) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: data.Course.Name + " Course Calendar",
		UpNav: UpNav{
			URL:  data.E.Reverse(data.ListTermCoursesRHN),
			Text: "Back to Courses",
		},
	}
}

func (data CourseCalendar) BreadCrumbs() BreadCrumbs {
	return BreadCrumbs{
		Term:             data.Course.Term,
		TermDetailsURL:   data.TermDetailsURL,
		Course:           data.Course,
		CourseDetailsURL: data.CourseDetailsURL,
	}
}
