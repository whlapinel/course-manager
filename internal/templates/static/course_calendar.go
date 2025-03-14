package templates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	components "gh_static_portfolio/internal/templates/components/base"
	"log"
	"time"

	"github.com/a-h/templ"
)

type CourseCalendarPage struct {
	Nodes         domain.Nodes
	CalendarDates CalendarDates
	Path          string
	AssetsURL     func(relPath string) string
	LessonPageURL func(nodes ...domain.CourseNode) string
}

func (page CourseCalendarPage) Component() templ.Component {
	return components.Layout{
		PageTitle: page.Nodes.Course.Name,
		Page:      CourseCalendarComponent(page),
		Head: Head{
			User:      page.Nodes.User,
			AssetsURL: page.AssetsURL,
		}.Component(),
	}.Component()
}

func (page CourseCalendarPage) Filepath() string {
	return page.Path
}
func (page CourseCalendarPage) Layout() PageData {
	return PageData{
		User:      page.Nodes.User,
		AssetsURL: page.AssetsURL,
	}
}

func MonthDates(term domain.Term) []time.Time {
	dates, _ := term.TermMonths()
	return dates
}

func HasNonZeroWeekDay(week []time.Time) bool {
	for _, day := range week[time.Monday:time.Saturday] {
		if !day.IsZero() {
			return true
		}
	}
	return false
}

func WithinTerm(week []time.Time, term domain.Term) bool {
	for _, day := range week[time.Monday:time.Saturday] {
		if !day.Before(term.Start) && !day.After(term.End) {
			return true
		}
	}
	return false
}

type CalendarDate struct {
	Date      time.Time
	Lessons   []domain.Lesson
	Occasions []domain.Occasion
}

type CalendarDates map[time.Time]CalendarDate

func DateData(date time.Time, page CourseCalendarPage) CalendarDate {
	dataMap := page.GetCalendarDates()
	date = ZeroizeDate(date)
	data := dataMap[date]
	// even if the data was zero, we still want the data.Date
	data.Date = date
	log.Println("DateData():", "date:", date, "data:", data)
	return data
}

// so we can reliably use dates as map keys
func ZeroizeDate(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func (data CourseCalendarPage) GetCalendarDates() CalendarDates {
	return data.CalendarDates
}

type NewCourseCalendarParams struct {
	CourseCalendarPage
}

func NewCourseCalendarPage(params NewCourseCalendarParams) CourseCalendarPage {
	var cc CourseCalendarPage = params.CourseCalendarPage
	cc.CalendarDates = cc.ProcessCalendarDates()
	return cc
}

func (data CourseCalendarPage) ProcessCalendarDates() CalendarDates {
	var datesMap = make(map[time.Time]CalendarDate)
	for _, occasion := range data.Nodes.Term.Occasions {
		data := datesMap[occasion.Date]
		data.Date = occasion.Date
		data.Occasions = append(data.Occasions, occasion)
		datesMap[occasion.Date] = data
	}
	for _, unit := range data.Nodes.Course.Units {
		for _, lesson := range unit.Lessons {
			for _, date := range lesson.Dates {
				data := datesMap[date]
				data.Date = date
				data.Lessons = append(data.Lessons, lesson)
				datesMap[date] = data
			}
		}
	}
	for date, data := range datesMap {
		line := fmt.Sprintf(
			`
datesMap: Key Date: %v
data.Date: %v
data.Occasions: %v
data.Lessons: %v
`,
			date, data.Date, data.Occasions, data.Lessons,
		)
		log.Println(line)
	}
	log.Println("CourseCalendar.ProcessCalendarDates: length of datesMap:", len(datesMap))
	return datesMap
}
