package managertemplates

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	cmp "gh_static_portfolio/internal/templates/components/base"
	components "gh_static_portfolio/internal/templates/components/base"
	templates "gh_static_portfolio/internal/templates/shared"
	"gh_static_portfolio/internal/util"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CalendarPage interface {
	GetCalendarDates() CalendarDates
	GetTerm() dto.Term
	IsStatic() bool
	Page
}

func MonthDates(term dto.Term) []time.Time {
	dates, _ := term.TermMonths()
	return dates
}

type TermCalendar struct {
	Params               routes.NodePath
	Term                 dto.Term
	ListTermsURL         string
	TermDetailsURL       string
	CreateOccasionURL    string
	GetEditOccasionRHN   string
	PostEditOccasionRHN  string
	DeleteOccasionRHN    string
	CurrentOccasionIndex int
	E                    *echo.Echo
	CalendarDates        CalendarDates
	BreadCrumbsData      BreadCrumbs
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

func (data TermCalendar) GetCalendarDates() CalendarDates {
	return data.CalendarDates
}

func (data TermCalendar) ProcessCalendarDates() CalendarDates {
	var calendarDates CalendarDates = make(map[time.Time]CalendarDate)
	for _, occasion := range data.Term.Occasions {
		data := calendarDates[occasion.Date]
		data.Date = occasion.Date
		data.Occasions = append(data.Occasions, occasion)
		calendarDates[occasion.Date] = data
	}
	return calendarDates
}

func (term TermCalendar) GetTerm() dto.Term {
	return term.Term
}

func (page TermCalendar) Occasions(date time.Time) []occasion.Occasion {
	var occasions []occasion.Occasion
	for _, occ := range page.Term.Occasions {
		if util.IsSameDate(occ.Date, date) {
			occasions = append(occasions, occ)
		}
	}
	return occasions
}

func (page TermCalendar) TermOccasionEditor(occasion occasion.Occasion) templ.Component {
	return TermOccasionEditor{
		Params:              page.Params,
		Occasion:            occasion,
		IsEditing:           false,
		GetEditOccasionURL:  page.E.Reverse(page.GetEditOccasionRHN, AddParams(page.Params, occasion.ID)...),
		PostEditOccasionURL: page.E.Reverse(page.PostEditOccasionRHN, AddParams(page.Params, occasion.ID)...),
		DeleteOccasionURL:   page.E.Reverse(page.DeleteOccasionRHN, AddParams(page.Params, occasion.ID)...),
	}.Component()

}

type AddOccasionButton struct {
	Date              time.Time
	CreateOccasionURL string
	FormID            string
}

func (data TermCalendar) AddOccasionButton(date time.Time) templ.Component {
	return AddOccasionButton{
		Date:              date,
		CreateOccasionURL: data.CreateOccasionURL,
		FormID:            "form-" + date.Format(time.DateOnly),
	}.Component()
}

func (button AddOccasionButton) Component() templ.Component {
	return AddOccasionButtonComponent(button)
}

func (data TermCalendar) Component() templ.Component {
	data.CalendarDates = data.ProcessCalendarDates()
	return TermCalendarComponent(data)
}

func (data TermCalendar) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: data.Term.Name + " Calendar",
		UpNav: cmp.UpNav{
			URL:  data.ListTermsURL,
			Text: "Back to Terms",
		},
		Crumbs: data.BreadCrumbs().BreadCrumbs(),
	}

}

func (page TermCalendar) IsStatic() bool {
	return false
}

func (data TermCalendar) BreadCrumbs() BreadCrumbs {
	return data.BreadCrumbsData
}

type CourseCalendar struct {
	Admin                 bool
	Static                bool
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
	BreadCrumbsData       BreadCrumbs
}

func (data CourseCalendar) GetCalendarDates() CalendarDates {
	return data.CalendarDates
}

func (page CourseCalendar) IsStatic() bool {
	return page.Static
}
func (page CourseCalendar) GetTerm() dto.Term {
	return page.Term
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

type TermOccasionEditor struct {
	Params              routes.NodePath
	Occasion            occasion.Occasion
	IsEditing           bool
	GetEditOccasionURL  string
	PostEditOccasionURL string
	DeleteOccasionURL   string
}

func (data TermOccasionEditor) ComponentID() string {
	var id string
	id = "occasion " + strconv.Itoa(data.Occasion.ID) + " editor"
	id = strings.ToLower(id)
	id = strings.ReplaceAll(id, " ", "-")
	return id
}

func (data TermOccasionEditor) Component() templ.Component {
	return TermOccasionEditorComponent(data)
}

func (data CourseCalendar) Component() templ.Component {
	return DynamicCourseCalendarComponent(data)
}

func (data CourseCalendar) CalendarLessonContainer(lesson dto.Lesson, date time.Time) templ.Component {
	params := data.Params
	params.UnitID = lesson.UnitID
	params.LessonID = lesson.ID
	container := CalendarLessonContainerNew{
		Date:             date,
		Params:           params,
		lesson:           lesson,
		LessonDetailsURL: data.LessonDetailsFunc(lesson.UnitID, lesson.ID),
		Course:           data.Course,
		ShiftLessonURL: func(cd string) string {
			return data.ShiftLessonFunc(lesson.UnitID, lesson.ID, cd)
		},
		RemoveLessonDateURL: data.RemoveLessonDateFunc(lesson.UnitID, lesson.ID, date.Format(time.DateOnly)),
	}
	if data.Static {
		container.Static = true
	}
	return container.Component()
}

func (data CourseCalendar) ShowAddLessonDatePageURL(date time.Time) string {
	return data.ShowAddLessonDateFunc(date.Format(time.DateOnly))
}

func (data CourseCalendar) PageLayout() cmp.PageLayout {
	if data.Static {
		return cmp.PageLayout{
			PageTitle: data.Course.Name + " Course Calendar",
			UpNav: cmp.UpNav{
				URL:  RemoveDocsFromPath(templates.StaticSiteCoursesDir + "courses.html"),
				Text: "Back to Courses",
			},
		}
	}
	return cmp.PageLayout{
		PageTitle: data.Course.Name + " Course Calendar",
		UpNav: cmp.UpNav{
			URL:  data.ListTermCoursesURL,
			Text: "Back to Courses",
		},
		Crumbs: data.BreadCrumbsData.BreadCrumbs(),
	}
}

type CalendarLessonContainerNew struct {
	Params              routes.NodePath
	Static              bool
	Date                time.Time
	Course              dto.Course
	lesson              dto.Lesson
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
		Text:    data.lesson.Designation(),
		Details: data.lesson.Name,
	}.Component()
}

func (data CalendarLessonContainerNew) Component() templ.Component {
	return CalendarLessonContainer(data)
}

func (data CalendarLessonContainerNew) RemoveLessonButton() templ.Component {
	button := components.Button{
		HxConfirm: "Are you sure you want to remove this lesson date? The lesson itself will not be deleted.",
		Method:    cmp.HxDelete,
		URL:       data.RemoveLessonDateURL,
		HxTarget:  "#page",
		Image:     cmp.DeleteImage(),
		Class:     "bg-red-700 p-1 rounded",
	}
	return button.Component()
}

func (data CourseCalendar) BreadCrumbs() BreadCrumbs {
	return data.BreadCrumbsData
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
	BreadCrumbsData   BreadCrumbs
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

func StaticSiteCourseCalendar(course dto.Course) CourseCalendar {
	cc := CourseCalendar{
		Admin:  false,
		Static: true,
		Course: course,
	}
	return cc
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

func (data AddLessonToDatePage) BreadCrumbs() BreadCrumbs {
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
