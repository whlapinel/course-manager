package hugo

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"path/filepath"
	"strings"
	"time"
)

type CalendarPageData struct {
	dto.Term `json:"term"`
	Path     string
	Months   []*CalendarMonth `json:"months"`
}

type CalendarMonth struct {
	Month     time.Month      `json:"month"`
	PrevMonth string          `json:"prevMonth"`
	NextMonth string          `json:"nextMonth"`
	Path      string          `json:"path"`
	Year      int             `json:"year"`
	Weeks     []*CalendarWeek `json:"weeks"`
}

type CalendarWeek struct {
	Dates []*CalendarDate `json:"dates"`
}

type CalendarDate struct {
	Date      string              `json:"date"`
	ShortDate string              `json:"shortDate"`
	Occasions []*CalendarOccasion `json:"occasions"`
	Lessons   []*CalendarLesson   `json:"lessons"`
}

type CalendarLesson struct {
	Designation string `json:"designation"`
	Name        string `json:"name"`
}

type CalendarOccasion struct {
	Name string `json:"name"`
}

func NewCalendar(term dto.Term) *CalendarPageData {
	var data = CalendarPageData{
		Path: "calendar",

		Term: term,
	}
	var months []*CalendarMonth
	start := term.Start
	firstOfMonth := time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.Local)
	for currMonth := firstOfMonth; !currMonth.After(term.End); currMonth = currMonth.AddDate(0, 1, 0) {
		path := func(month time.Month) string {
			return filepath.Join("calendar", strings.ToLower(month.String()))
		}
		var month = CalendarMonth{
			Path:      path(currMonth.Month()),
			Month:     currMonth.Month(),
			Year:      currMonth.Year(),
			NextMonth: path(currMonth.AddDate(0, 1, 0).Month()),
			PrevMonth: path(currMonth.AddDate(0, -1, 0).Month()),
		}
		firstOfMonth := time.Date(currMonth.Year(), currMonth.Month(), 1, 0, 0, 0, 0, time.Local)
		var sun = firstOfMonth.AddDate(0, 0, -int(firstOfMonth.Weekday()))
		var weeks []*CalendarWeek
		for currWeek := sun; !currWeek.After(firstOfMonth.AddDate(0, 1, 0)); currWeek = currWeek.AddDate(0, 0, 7) {
			var week CalendarWeek
			var dates []*CalendarDate
			for i := range 7 {
				currDate := currWeek.AddDate(0, 0, i)
				var date = CalendarDate{
					Date:      currDate.Format("Jan 2"),
					ShortDate: currDate.Format("02"),
				}
				dates = append(dates, &date)
			}
			week.Dates = dates
			weeks = append(weeks, &week)
		}
		month.Weeks = weeks
		months = append(months, &month)
	}
	data.Months = months
	return &data
}

func (p *CalendarPageData) Children() []Homogenizer {
	var homos []Homogenizer
	for _, month := range p.Months {
		homos = append(homos, month)
	}
	return homos
}

func (p *CalendarPageData) Page() *HomogenizedPageData {
	return nil
}

func (p *CalendarPageData) Section() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = SectionKind
	homoPageData.Type = CalendarType
	homoPageData.Path = p.Path
	homoPageData.Params = struct {
		ParentPath string `json:"parentPath"`
	}{
		ParentPath: "/",
	}
	return &homoPageData
}

func (p *CalendarMonth) Children() []Homogenizer {
	return nil
}

func (p *CalendarMonth) Section() *HomogenizedPageData {
	return nil
}

func (p *CalendarMonth) Page() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = PageKind
	homoPageData.Type = CalendarType
	homoPageData.Path = p.Path
	homoPageData.Weight = int(p.Month)
	homoPageData.Title = fmt.Sprintf("%s %d", p.Month.String(), p.Year)
	homoPageData.Params = struct {
		Weeks         []*CalendarWeek    `json:"weeks"`
		BreadCrumbs   BreadCrumbsPartial `json:"breadCrumbs"`
		YearMonth     string             `json:"yearMonth"`
		PrevMonthPath string             `json:"prevMonthPath"`
		NextMonthPath string             `json:"nextMonthPath"`
	}{
		YearMonth:     time.Date(p.Year, p.Month, 1, 0, 0, 0, 0, time.Local).Format("2006-01"),
		NextMonthPath: p.NextMonth,
		PrevMonthPath: p.PrevMonth,
		Weeks:         p.Weeks,
		BreadCrumbs:   BreadCrumbs(p.Path),
	}
	return &homoPageData

}
