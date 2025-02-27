package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates"
	"gh_static_portfolio/internal/util"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CalendarPage interface {
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
	CurrentOccasionIndex int
	E                    *echo.Echo
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
	return TermCalendarComponent(data)
}

func (data TermCalendar) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: data.Term.Name + " Calendar",
		UpNav: UpNav{
			URL:  data.ListTermsURL,
			Text: "Back to Terms",
		},
	}

}

func (page TermCalendar) IsStatic() bool {
	return false
}

func (data TermCalendar) BreadCrumbs() BreadCrumbs {
	return BreadCrumbs{
		Term:           data.Term,
		TermDetailsURL: data.TermDetailsURL,
	}
}

type CourseCalendar struct {
	Admin                         bool
	Static                        bool
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

func (data CourseCalendar) PageLayout() PageLayout {
	if data.Static {
		return PageLayout{
			PageTitle: data.Course.Name + " Course Calendar",
			UpNav: UpNav{
				URL:  RemoveDocsFromPath(templates.StaticSiteCoursesDir + "courses.html"),
				Text: "Back to Courses",
			},
		}
	}
	return PageLayout{
		PageTitle: data.Course.Name + " Course Calendar",
		UpNav: UpNav{
			URL:  data.E.Reverse(data.ListTermCoursesRHN, data.Params.ToSlice()...),
			Text: "Back to Courses",
		},
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
		ShiftLessonURL: data.E.Reverse(data.ShiftLessonRHN, AddParams(data.Params, cd.String())...),
		e:              data.E,
	}
	return button.Component()
}

type AddLessonToDatePage struct {
	Date             time.Time
	Params           domain.NodePath
	Course           domain.Course
	AddLessonDateRHN string
	E                *echo.Echo
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
	return CourseCalendar{
		Admin:  false,
		Static: true,
		Course: course,
	}
}
