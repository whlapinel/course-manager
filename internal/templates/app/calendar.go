package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
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
	GetTerm() domain.Term
	IsStatic() bool
	Page
}

func MonthDates(term domain.Term) []time.Time {
	dates, _ := term.TermMonths()
	return dates
}

type TermCalendar struct {
	Params               domain.NodePath
	Term                 domain.Term
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

func (term TermCalendar) GetTerm() domain.Term {
	return term.Term
}

func (page TermCalendar) Occasions(date time.Time) []domain.Occasion {
	var occasions []domain.Occasion
	for _, occ := range page.Term.Occasions {
		if util.IsSameDate(occ.Date, date) {
			occasions = append(occasions, occ)
		}
	}
	return occasions
}

func (page TermCalendar) TermOccasionEditor(occasion domain.Occasion) templ.Component {
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
	Admin                         bool
	Static                        bool
	Nodes                         domain.Nodes
	Params                        domain.NodePath
	Course                        domain.Course
	TermDetailsURL                string
	CourseDetailsURL              string
	LessonDetailsRouteHandlerName string
	ShiftLessonRouteHandlerName   string
	ListTermCoursesRHN            string
	CreateOccasionRHN             string
	ShowAddLessonDateRHN          string
	RemoveLessonDateRHN           string
	E                             *echo.Echo
	CalendarDates                 CalendarDates
	BreadCrumbsData               BreadCrumbs
}

func (data CourseCalendar) GetCalendarDates() CalendarDates {
	return data.CalendarDates
}

func (data CourseCalendar) ProcessCalendarDates() CalendarDates {
	var datesMap = make(map[time.Time]CalendarDate)
	for _, occasion := range data.Course.Term.Occasions {
		data := datesMap[occasion.Date]
		data.Date = occasion.Date
		data.Occasions = append(data.Occasions, occasion)
		datesMap[occasion.Date] = data
	}
	for _, unit := range data.Course.Units {
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

func (page CourseCalendar) IsStatic() bool {
	return page.Static
}
func (page CourseCalendar) GetTerm() domain.Term {
	return page.Course.Term
}

func WithinTerm(week []time.Time, term domain.Term) bool {
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

func Occasions(date time.Time, page CalendarPage) []domain.Occasion {
	var occasions []domain.Occasion
	for _, occ := range page.GetTerm().Occasions {
		if util.IsSameDate(occ.Date, date) {
			occasions = append(occasions, occ)
		}
	}
	return occasions
}

type TermOccasionEditor struct {
	Params              domain.NodePath
	Occasion            domain.Occasion
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
	data.CalendarDates = data.ProcessCalendarDates()
	return DynamicCourseCalendarComponent(data)
}

func (data CourseCalendar) CalendarLessonContainer(lesson domain.Lesson, date time.Time) templ.Component {
	params := data.Params
	params.UnitID = lesson.UnitID
	params.LessonID = lesson.ID
	container := CalendarLessonContainerNew{
		Date:                date,
		Params:              params,
		lesson:              lesson,
		LessonDetailsURL:    data.E.Reverse(data.LessonDetailsRouteHandlerName, params.ToSlice()...),
		Course:              data.Course,
		ShiftLessonRHN:      data.ShiftLessonRouteHandlerName,
		RemoveLessonDateURL: data.E.Reverse(data.RemoveLessonDateRHN, AddParams(params, date.Format(time.DateOnly))...),
		E:                   data.E,
	}
	if data.Static {
		container.Static = true
	}
	return container.Component()
}

func (data CourseCalendar) ShowAddLessonDatePageURL(date time.Time) string {
	return data.E.Reverse(
		data.ShowAddLessonDateRHN,
		data.Params.UserID,
		data.Params.TermID,
		data.Params.CourseID,
		date.Format(time.DateOnly),
	)
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
			URL:  data.E.Reverse(data.ListTermCoursesRHN, data.Params.ToSlice()...),
			Text: "Back to Courses",
		},
		Crumbs: data.BreadCrumbsData.BreadCrumbs(),
	}
}

type CalendarLessonContainerNew struct {
	Params              domain.NodePath
	Static              bool
	Date                time.Time
	Course              domain.Course
	lesson              domain.Lesson
	LessonDetailsURL    string
	RemoveLessonDateURL string
	ShiftLessonRHN      string
	E                   *echo.Echo
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
	//	<div class="flex flex-col gap-1 items-center">
	//	<div class="flex">
	//		if !data.Static {
	//			<button
	//				class="bg-red-700 p-1 rounded"
	//				hx-confirm={ "Are you sure you want to remove this lesson date? The lesson itself will not be deleted." }
	//				hx-delete={ data.RemoveLessonDateURL }
	//				hx-target="#page"
	//			>
	//				Remove Lesson
	//			</button>
	//		}
	//		@data.LinkWithInfoDialog()
	//	</div>
	//
	// </div>
}

func (data CourseCalendar) BreadCrumbs() BreadCrumbs {
	return data.BreadCrumbsData
}

func (data CalendarLessonContainerNew) ShiftButton(cd domain.CalendarDirection) templ.Component {
	button := ShiftButton{
		Direction:      cd,
		Params:         data.Params,
		ShiftLessonURL: data.E.Reverse(data.ShiftLessonRHN, AddParams(data.Params, cd.String())...),
		e:              data.E,
	}
	return button.Component()
}

type AddLessonToDatePage struct {
	Date              time.Time
	Nodes             domain.Nodes
	Params            domain.NodePath
	Course            domain.Course
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

func StaticSiteCourseCalendar(course domain.Course) CourseCalendar {
	cc := CourseCalendar{
		Admin:  false,
		Static: true,
		Course: course,
	}
	cc.CalendarDates = cc.ProcessCalendarDates()
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
	Params         domain.NodePath
	Units          []domain.Unit
	ListLessonsRHN string
	Echo           *echo.Echo
}

func (data UnitPicker) ListLessonsButton(unit domain.Unit) templ.Component {
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
	Params          domain.NodePath
	ListUnitsURL    string
	Lessons         []domain.Lesson
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

func (data LessonPicker) SelectLessonButton(lesson domain.Lesson) templ.Component {
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
	Lessons   []domain.Lesson
	Occasions []domain.Occasion
}

type CalendarDates map[time.Time]CalendarDate
