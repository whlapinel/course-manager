package calendarviews

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/core/occasion"
	cmp "gh_static_portfolio/internal/newtemplates/components/base"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/util"
	"log"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CalendarPage interface {
	GetCalendarDates() CalendarDates
	GetTerm() dto.Term
	ac.Page
}

func MonthDates(term dto.Term) []time.Time {
	dates, _ := term.TermMonths()
	return dates
}

func DateData(date time.Time, page CalendarPage) CalendarDate {
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

type AddOccasionButton struct {
	Date              time.Time
	CreateOccasionURL string
	FormID            string
}

func (button AddOccasionButton) Component() templ.Component {
	return AddOccasionButtonComponent(button)
}

func WithinTerm(week []time.Time, term dto.Term) bool {
	for _, day := range week[time.Monday:time.Saturday] {
		if !day.Before(term.Start) && !day.After(term.End) {
			return true
		}
	}
	return false
}

func HasNonZeroWeekDay(week []time.Time) bool {
	for _, day := range week[time.Monday:time.Saturday] {
		if !day.IsZero() {
			return true
		}
	}
	return false
}

func Occasions(date time.Time, page CalendarPage) []occasion.Occasion {
	var occasions []occasion.Occasion
	for _, occ := range page.GetTerm().Occasions {
		if util.IsSameDate(occ.Date, date) {
			occasions = append(occasions, occ)
		}
	}
	return occasions
}

type CalendarLessonContainerNew struct {
	Params              routes.NodePath
	Date                time.Time
	Course              dto.Course
	Lesson              dto.Lesson
	LessonDetailsURL    string
	RemoveLessonDateURL string
	ShiftLessonURL      func(cd string) string
}

type LinkWithInfoDialog struct {
	Static  bool
	URL     string
	Target  string
	PushURL bool
	// what will display without hover
	Text string
	// what will display on hover
	Details string
}

func (data LinkWithInfoDialog) Component() templ.Component {
	return LinkWithInfoDialogComponent(data)
}

func (data CalendarLessonContainerNew) LinkWithInfoDialog() templ.Component {
	return LinkWithInfoDialog{
		URL:     data.LessonDetailsURL,
		Target:  "#page",
		PushURL: true,
		Text:    data.Lesson.Designation(),
		Details: data.Lesson.Name,
	}.Component()
}

func (data CalendarLessonContainerNew) Component() templ.Component {
	return CalendarLessonContainer(data)
}

func (data CalendarLessonContainerNew) RemoveLessonButton() templ.Component {
	button := cmp.Button{
		HxConfirm: "Are you sure you want to remove this lesson date? The lesson itself will not be deleted.",
		Method:    cmp.HxDelete,
		URL:       data.RemoveLessonDateURL,
		HxTarget:  "#page",
		Image:     cmp.DeleteImage(),
		Class:     "bg-red-700 p-1 rounded",
	}
	return button.Component()
}

type ShiftButton struct {
	Params         routes.NodePath
	Direction      dto.CalendarDirection
	TermID         int
	CourseID       int
	ShiftLessonURL string
	e              *echo.Echo
}

func (data ShiftButton) Component() templ.Component {
	return ShiftButtonComponent(data)
}

func (data CalendarLessonContainerNew) ShiftButton(cd dto.CalendarDirection) templ.Component {
	button := ShiftButton{
		Direction:      cd,
		Params:         data.Params,
		ShiftLessonURL: data.ShiftLessonURL(cd.String()),
	}
	return button.Component()
}

type AddLessonToDatePage struct {
	Date              time.Time
	Nodes             node.Nodes
	Params            routes.NodePath
	Course            dto.Course
	ListLessonsRHN    string
	AddLessonDateRHN  string
	E                 *echo.Echo
	TermDetailsURL    string
	BreadCrumbsData   ac.BreadCrumbs
	CourseDetailsURL  string
	CourseCalendarURL string
}

func (page AddLessonToDatePage) Component() templ.Component {
	return AddLessonToDatePageComponent(page)

}

func (data AddLessonToDatePage) AddLessonDateURL(unitID, lessonID int) string {
	data.Params.UnitID = unitID
	data.Params.LessonID = lessonID
	return data.E.Reverse(data.AddLessonDateRHN, AddParams(data.Params, data.Date.Format(time.DateOnly))...)
}

func (data AddLessonToDatePage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Add Lesson to " + data.Date.Format(time.DateOnly),
		UpNav: cmp.UpNav{
			URL:  data.CourseCalendarURL,
			Text: "Back to Calendar",
		},
	}

}

func (data AddLessonToDatePage) BreadCrumbs() ac.BreadCrumbs {
	return data.BreadCrumbsData
}

func (data AddLessonToDatePage) UnitPicker() templ.Component {
	return UnitPicker{
		Date:           data.Date,
		Params:         data.Params,
		Units:          data.Course.Units,
		ListLessonsRHN: data.ListLessonsRHN,
		Echo:           data.E,
	}.Component()
}

type UnitPicker struct {
	Date           time.Time
	Params         routes.NodePath
	Units          []dto.Unit
	ListLessonsRHN string
	Echo           *echo.Echo
}

func (data UnitPicker) ListLessonsButton(unit dto.Unit) templ.Component {
	return cmp.Button{
		Text:     unit.Designation(),
		Method:   cmp.HxGet,
		URL:      data.SelectUnitURL(unit.ID),
		HxTarget: "#picker",
	}.Component()
}

func (data UnitPicker) SelectUnitURL(unitID int) string {
	return data.Echo.Reverse(data.ListLessonsRHN, AddParams(data.Params, data.Date.Format(time.DateOnly), unitID)...)
}

func (data UnitPicker) Component() templ.Component {
	return UnitPickerComponent(data)
}

type LessonPicker struct {
	Date            time.Time
	Params          routes.NodePath
	ListUnitsURL    string
	Lessons         []dto.Lesson
	SelectLessonRHN string
	Echo            *echo.Echo
}

func (data LessonPicker) ListUnitsButton() templ.Component {
	return cmp.Button{
		Text:     "Back to units",
		Method:   cmp.HxGet,
		URL:      data.ListUnitsURL,
		HxTarget: "#page",
	}.Component()
}

func (data LessonPicker) Component() templ.Component {
	return LessonPickerComponent(data)
}

func (data LessonPicker) SelectLessonButton(lesson dto.Lesson) templ.Component {
	return cmp.Button{
		Text:     lesson.Designation(),
		Method:   cmp.HxPost,
		URL:      data.SelectLessonURL(lesson.ID),
		HxTarget: "#picker",
	}.Component()
}

func (data LessonPicker) SelectLessonURL(lessonID int) string {
	return data.Echo.Reverse(data.SelectLessonRHN, AddParams(data.Params, lessonID, data.Date.Format(time.DateOnly))...)
}

type CalendarDate struct {
	Date      time.Time
	Lessons   []dto.Lesson
	Occasions []occasion.Occasion
}

type CalendarDates map[time.Time]CalendarDate

func AddParams(params routes.NodePath, additionalParams ...any) []any {
	pathSlice := params.ToSlice()
	pathSlice = append(pathSlice, additionalParams...)
	return pathSlice
}
