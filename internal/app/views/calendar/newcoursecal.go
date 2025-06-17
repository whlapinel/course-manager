package calendarviews

import (
	appcomponents "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"time"

	"github.com/a-h/templ"
)

type MonthCalendar struct {
	dto.Term
	CourseDetailsURL string
	Month            time.Month
	Year             int
	PrevMonthURL     string
	NextMonthURL     string
	CurrentMonthURL  string
	Weeks            [][]CalendarDate
	appcomponents.CourseManagerLayout
}

func (data MonthCalendar) Component() templ.Component {
	return NewNewCalendarComponent(data)
}

func (data MonthCalendar) HTMXResponse() templ.Component {
	return data.Component()
}

func (data MonthCalendar) NonHTMXResponse() templ.Component {
	return data.CourseManagerLayout.WithPage(data.Component())
}
