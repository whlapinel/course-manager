package templates

import (
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
)

const (
	githubRoot = "https://github.com/whlapinel/python/tree/main/docs/courses"
	rootDir    = "./python/docs/"
	coursesDir = "./python/docs/courses/"
)

type NameHaver interface {
	GetName() string
}

func kebabCase(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, " ", "-"))
}

func coursePath(course domain.Course, page bool) string {
	dirPath := filepath.Join(coursesDir, kebabCase(course.Name))
	if page {
		return filepath.Join(dirPath, kebabCase(course.Name+".html"))
	}
	return dirPath
}

func unitPath(unit domain.Unit, course domain.Course, page bool) string {
	dirPath := filepath.Join(coursePath(course, false), kebabCase(unit.Name))
	if page {
		return filepath.Join(dirPath, kebabCase(unit.Name+".html"))
	}
	return dirPath
}

func lessonPath(lesson domain.Lesson, unit domain.Unit, course domain.Course, page bool) string {
	dirPath := filepath.Join(unitPath(unit, course, false), kebabCase(lesson.Name))
	if page {
		return filepath.Join(dirPath, kebabCase(lesson.Name+".html"))
	}
	return dirPath
}

func LessonFilesURL(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(githubRoot, LessonFilesPath(lesson, unit, course))
}

func LessonFilesPath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(lessonPath(lesson, unit, course, false), "files")
}

func SlidesPath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(lessonPath(lesson, unit, course, false), "slides.html")
}

// This should not be used except for a one-time transfer
func SlidesMarkdownPath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(lessonPath(lesson, unit, course, false), "slides.md")
}

func hasImage(path string) bool {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("error reading directory:%s", err)
	}
	for _, file := range files {
		if file.Name() == "image.png" {
			return true
		}
	}
	return false

}

// if files are found in "files" subdirectory of lesson directory, this will return true, unless the name is prefixed with "secret"
func hasFilesDir(path string) bool {
	files, err := os.ReadDir(path + "/files")
	if err != nil {
		return false
	}
	for _, file := range files {
		if file.Name()[:6] != "secret" {
			return true
		}
	}
	return false
}

// this returns true if and only if a file is found with name "slides.html"
func hasSlides(path string) bool {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("error reading directory: %s", err)
	}
	for _, file := range files {
		if file.Name() == "slides.html" {
			return true
		}
	}
	return false
}

// list of directories to be cleared (used for clearing html files only)
func DirectoriesClearList() []string {
	return []string{
		rootDir,
		coursesDir,
	}
}

type Templifier interface {
	Templify() templ.Component
	GetTitle() string
	Directory() string
}

type Titler interface {
	GetTitle() string
}

type Page struct {
	Title     string
	Path      string
	Component templ.Component
}

func (p *Page) Templify() templ.Component {
	return p.Component
}

func (p *Page) GetTitle() string {
	return p.Title
}

func NewHomePage() Page {
	return Page{
		Title:     "Home",
		Path:      filepath.Join(rootDir, "index.html"),
		Component: HomeComponent(),
	}
}

func NewContactPage() Page {
	return Page{
		Title:     "Contact",
		Path:      filepath.Join(rootDir, "contact.html"),
		Component: ContactComponent(),
	}
}

func NewCoursesListPage(instances []*domain.Course) Page {
	return Page{
		Title:     "Courses",
		Path:      filepath.Join(coursesDir, "courses.html"),
		Component: CoursesListComponent(instances),
	}

}

func NewCoursePage(course domain.Course) Page {
	return Page{
		Title:     course.Name,
		Path:      coursePath(course, true),
		Component: CourseComponent(course),
	}
}

func NewCourseCalendarPage(course domain.Course) Page {
	pageTitle := course.Name + " Calendar"
	return Page{
		Title:     pageTitle,
		Path:      filepath.Join(coursePath(course, false), kebabCase(pageTitle)),
		Component: CourseCalendarComponent(course),
	}
}

func NewUnitPage(unit domain.Unit, course domain.Course) Page {
	return Page{
		Title:     unit.Name,
		Path:      unitPath(unit, course, true),
		Component: UnitComponent(unit, course),
	}
}

func NewLessonPage(lesson domain.Lesson, unit domain.Unit, course domain.Course) Page {
	return Page{
		Title:     lesson.Name,
		Path:      lessonPath(lesson, unit, course, true),
		Component: LessonComponent(lesson, unit, course),
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
