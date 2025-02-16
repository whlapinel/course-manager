package statictemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/templates"
	managertemplates "gh_static_portfolio/internal/templates/manager_templates"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
)

// Student-facing site
func SlidesPath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(templates.LessonPath(lesson, unit, course, false), "slides.html")
}

// This should not be used except for a one-time transfer
func SlidesMarkdownPath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(templates.LessonPath(lesson, unit, course, false), "slides.md")
}

// list of directories to be cleared (used for clearing html files only)
func DirectoriesClearList() []string {
	return []string{
		templates.StaticSiteRootDir,
		templates.StaticSiteCoursesDir,
	}
}

// courses directory to be deleted completely
func DeleteDirList() []string {
	return []string{templates.StaticSiteCoursesDir}
}

type Page struct {
	Title     string
	Path      string
	Component templ.Component
}

func NewHomePage() Page {
	return Page{
		Title:     "Home",
		Path:      filepath.Join(templates.StaticSiteRootDir, "index.html"),
		Component: HomeComponent(),
	}
}

func NewContactPage() Page {
	return Page{
		Title:     "Contact",
		Path:      filepath.Join(templates.StaticSiteRootDir, "contact.html"),
		Component: ContactComponent(),
	}
}

func NewCoursesListPage(instances []domain.Course) Page {
	return Page{
		Title:     "Courses",
		Path:      filepath.Join(templates.StaticSiteCoursesDir, "courses.html"),
		Component: CoursesListComponent(instances),
	}

}

func NewCoursePage(course domain.Course) Page {
	return Page{
		Title:     course.Name,
		Path:      templates.CoursePath(course, true),
		Component: CourseComponent(course),
	}
}

func NodePage(nodes ...domain.CourseNode) Page {
	leafNode := nodes[len(nodes)-1]
	return Page{
		Title: leafNode.GetName(),
		Path:  templates.NodePage(nodes...),
	}
}

func NewCourseCalendarPage(course domain.Course) Page {
	pageTitle := course.Name + " Calendar"
	return Page{
		Title: pageTitle,
		Path:  filepath.Join(templates.CoursePath(course, false), templates.KebabCase(pageTitle+".html")),
		// Component: CourseCalendarComponent(course),
		Component: managertemplates.StaticSiteCourseCalendar(course).Component(),
	}
}

func NewUnitPage(unit domain.Unit, course domain.Course) Page {
	return Page{
		Title:     unit.Name,
		Path:      templates.UnitPath(unit, course, true),
		Component: UnitComponent(unit, course),
	}
}

type LessonPage struct {
	Lesson              domain.Lesson
	Unit                domain.Unit
	Course              domain.Course
	HasSlides, HasFiles bool
}

func (page LessonPage) Component() templ.Component {
	return LessonComponent(page)
}

// 	@TitleDiv(fmt.Sprintf("%s: %s", page.Lesson.Designation(), page.Lesson.Name), page.Lesson.Description, filepath.Join(templates.LessonPath(page.Lesson, page.Unit, page.Course, false), page.Lesson.Image.BasePath), page.Lesson.Image.ID != 0, false, page.Lesson.Standards...)

func (page LessonPage) TitleDivNew() templ.Component {
	return TitleDivData{
		title:       fmt.Sprintf("%s: %s", page.Lesson.Designation(), page.Lesson.Name),
		description: page.Lesson.Description,
		imgPath:     filepath.Join(templates.LessonPath(page.Lesson, page.Unit, page.Course, false), page.Lesson.Image.BasePath),
		standards:   page.Lesson.Standards,
		Assessments: page.Lesson.Assessments,
	}.Component()

}

func NewLessonPage(page LessonPage) Page {
	return Page{
		Title:     page.Lesson.Name,
		Path:      templates.LessonPath(page.Lesson, page.Unit, page.Course, true),
		Component: page.Component(),
	}
}

func FilePathToURL(directory string) templ.SafeURL {
	route := RemoveDocsFromPath(directory)
	route = MakeAbsolute(route)
	return Sanitize(route)
}

func Sanitize(route string) templ.SafeURL {
	return templ.SafeURL(route)
}
func RemoveDocsFromPath(directory string) string {
	return strings.ReplaceAll(directory, "/docs", "")
}

func MakeAbsolute(route string) string {
	return "/" + route
}

func rootPages() []Page {
	return []Page{
		NewHomePage(),
		NewContactPage(),
		NewCoursesListPage(nil),
	}
}

type TitleDivData struct {
	title, description, imgPath string
	standards                   []domain.Standard
	Assessments                 []domain.Assessment
	showImg                     bool
	lightBg                     bool
}

func (data TitleDivData) Component() templ.Component {
	return TitleDivNewComponent(data)
}
