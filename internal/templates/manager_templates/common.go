package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates"
	"log"
	"net/url"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type ComponentData interface {
	Component() templ.Component
}

const pageElementID ElementID = "page"

type CourseIDParams struct {
	TermID   NodeIDParam
	CourseID NodeIDParam
	UnitID   NodeIDParam
	LessonID NodeIDParam
}

type NodeIDParam struct {
	Valid bool
	Value int
}

func (params CourseIDParams) ToIntSlice() []interface{} {
	return []interface{}{params.TermID.Value, params.CourseID.Value, params.UnitID.Value, params.LessonID.Value}
}

type ElementID string

// simply prefixes with '#'
func (i ElementID) Selector() string {
	return string("#" + i)
}

func (i ElementID) String() string {
	return string(i)
}

const (
	EditSlidesContainerID ElementID = "slides-editor-container"
	EditSlidesTextAreaID  ElementID = "slides-editor-text-area"
)

type ShiftButton struct {
	Params         CourseIDParams
	Direction      domain.CalendarDirection
	TermID         int
	CourseID       int
	ShiftLessonURL string
	e              *echo.Echo
}

func (data ShiftButton) Component() templ.Component {
	return ShiftButtonComponent(data)
}

func AddQueryParam(path, key, value string) string {
	u, err := url.Parse(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	query := u.Query()
	query.Set(key, value)
	u.RawQuery = query.Encode()
	log.Println(u.String())
	return u.String()
}

type HXButton struct {
	Text      string
	HxConfirm string
	Method    HXMethod
	URL       string
	HxTarget  string
	PushURL   bool
	HxSwap    HxSwap
}

type HxSwap string

const (
	AfterEnd HxSwap = "afterend"
)

func (button HXButton) Component() templ.Component {
	return HxButtonComponent(button)
}
func NewHXButton(method HXMethod, hxSwap HxSwap, url, hxTargetID string, pushURL bool) HXButton {
	return HXButton{
		Method:   method,
		URL:      url,
		HxTarget: hxTargetID,
		PushURL:  pushURL,
		HxSwap:   hxSwap,
	}
}

type HXMethod string

const (
	HxGet    = "hx-get"
	HxPost   = "hx-post"
	HxDelete = "hx-delete"
)

type EditField struct {
	Params           CourseIDParams
	FieldName        string
	Content          string
	GetEditFieldURL  string
	PostEditFieldURL string
	InputComponent   templ.Component
	IsEdit           bool
}

func FieldContainerID(fieldName string) string {
	return templates.KebabCase(fieldName) + "-container"
}

func FieldInputID(fieldName string) string {
	return templates.KebabCase(fieldName)
}

func StaticRootPages() []StaticPage {
	return []StaticPage{
		StaticNewHomePage(),
		StaticNewContactPage(),
		StaticNewCoursesListPage(nil),
	}
}

type StaticPage struct {
	Title     string
	Path      string
	Component templ.Component
}

// for static site
type LessonPage struct {
	Lesson              domain.Lesson
	Unit                domain.Unit
	Course              domain.Course
	HasSlides, HasFiles bool
}

// for static site
type TitleDivData struct {
	title, description, imgPath string
	standards                   []domain.Standard
	Assessments                 []domain.Assessment
	showImg                     bool
	lightBg                     bool
}

func StaticNewHomePage() StaticPage {
	return StaticPage{
		Title:     "Home",
		Path:      filepath.Join(templates.StaticSiteRootDir, "index.html"),
		Component: StaticHomeComponent(),
	}
}

func StaticNewContactPage() StaticPage {
	return StaticPage{
		Title:     "Contact",
		Path:      filepath.Join(templates.StaticSiteRootDir, "contact.html"),
		Component: StaticContactComponent(),
	}
}

func StaticNewCoursesListPage(instances []domain.Course) StaticPage {
	return StaticPage{
		Title:     "Courses",
		Path:      filepath.Join(templates.StaticSiteCoursesDir, "courses.html"),
		Component: StaticCoursesListComponent(instances),
	}

}

func StaticNewCoursePage(course domain.Course) StaticPage {
	return StaticPage{
		Title:     course.Name,
		Path:      templates.CoursePath(course, true),
		Component: CourseComponent(course),
	}
}

func StaticNewCourseCalendarPage(course domain.Course) StaticPage {
	pageTitle := course.Name + " Calendar"
	return StaticPage{
		Title: pageTitle,
		Path:  filepath.Join(templates.CoursePath(course, false), templates.KebabCase(pageTitle+".html")),
		// Component: CourseCalendarComponent(course),
		Component: StaticCourseCalendarComponent(StaticSiteCourseCalendar(course)),
	}
}

func StaticNewUnitPage(unit domain.Unit, course domain.Course) StaticPage {
	return StaticPage{
		Title:     unit.Name,
		Path:      templates.UnitPath(unit, course, true),
		Component: StaticUnitComponent(unit, course),
	}
}

func StaticNewLessonPage(page LessonPage) StaticPage {
	return StaticPage{
		Title:     page.Lesson.Name,
		Path:      templates.LessonPath(page.Lesson, page.Unit, page.Course, true),
		Component: page.Component(),
	}
}
