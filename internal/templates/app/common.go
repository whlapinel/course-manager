package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	templates "gh_static_portfolio/internal/templates/shared"
	"log"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type ComponentData interface {
	Component() templ.Component
}

const pageElementID ElementID = "page"

type NodePath struct {
	UserID   NodeIDParam
	TermID   NodeIDParam
	CourseID NodeIDParam
	UnitID   NodeIDParam
	LessonID NodeIDParam
}

type NodeIDParam struct {
	Valid bool
	Value interface{}
}

func AddNodeChildIDToParams(params NodePath, childID any) NodePath {
	var newParams NodePath
	if params.UserID.Value.(string) == "" {
		newParams.UserID = NodeIDParam{Value: childID.(string)}
		return newParams
	} else if params.TermID.Value == nil {
		newParams = params
		newParams.TermID = NodeIDParam{Value: childID}
		return newParams
	} else if params.CourseID.Value == nil {
		newParams = params
		newParams.CourseID = NodeIDParam{Value: childID}
		return newParams
	} else if params.UnitID.Value == nil {
		newParams = params
		newParams.UnitID = NodeIDParam{Value: childID}
		return newParams
	} else if params.LessonID.Value == nil {
		newParams = params
		newParams.LessonID = NodeIDParam{Value: childID}
		return newParams
	}
	return params
}

// converts params into a slice of interfaces
func (params NodePath) ToSlice(additionalParams ...interface{}) []interface{} {
	var base []interface{}
	paramSlice := []interface{}{
		params.UserID.Value,
		params.TermID.Value,
		params.CourseID.Value,
		params.UnitID.Value,
		params.LessonID.Value,
	}
	for _, param := range paramSlice {
		if param != nil {
			base = append(base, param)
		}
	}
	return append(base, additionalParams...)
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
	Params         domain.NodePath
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

type EditField struct {
	Params           domain.NodePath
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
	StaticFilesURL              templ.SafeURL
	GitHubFilesURL              templ.SafeURL
	showImg                     bool
	lightBg                     bool
}

func (page LessonPage) Component() templ.Component {
	return StaticLessonComponent(page)
}

func (page LessonPage) TitleDivNew() templ.Component {
	return TitleDivData{
		title:          fmt.Sprintf("%s: %s", page.Lesson.Designation(), page.Lesson.Name),
		description:    page.Lesson.Description,
		StaticFilesURL: templ.SafeURL(templates.LessonFilesPath(page.Lesson, page.Unit, page.Course)),
		GitHubFilesURL: templates.LessonFilesURL(page.Lesson, page.Unit, page.Course),
		imgPath:        filepath.Join(templates.LessonPath(page.Lesson, page.Unit, page.Course, false), page.Lesson.Image.BasePath),
		standards:      page.Lesson.Standards,
		Assessments:    page.Lesson.Assessments,
	}.Component()

}

// Github files URL
func (data TitleDivData) FileURL(filepath string) templ.SafeURL {
	url, _ := url.JoinPath(string(data.GitHubFilesURL), filepath)
	return templ.SafeURL(url)
}

func (data TitleDivData) ViewMarkdownURL(filepath string) templ.SafeURL {
	filepath = strings.ReplaceAll(filepath, ".md", ".html")
	url, _ := url.JoinPath(string(data.StaticFilesURL), filepath)
	url = RemoveDocsFromPath(url)
	url = MakeAbsolute(url)
	return templ.SafeURL(url)
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
