package templates

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
)

const (
	rootDir      = "./python/docs/"
	aboutDir     = "./python/docs/about/"
	educationDir = "./python/docs/about/education/"
	blogDir      = "./python/docs/blog/"
	coursesDir   = "./python/docs/courses/"
)

func courseFilePath(course domain.Course) string {
	return fmt.Sprintf("%s%s", coursesDir, DirName(course))
}

func unitFilePath(unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("%s%s%s", coursesDir, DirName(course), DirName(unit))
}

func lessonFilePath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("%s%s%s%s", coursesDir, DirName(course), DirName(unit), DirName(lesson))
}

func filesFilePath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("https://github.com/whlapinel/python/tree/main/docs/courses/%s%s%sfiles", DirName(course), DirName(unit), DirName(lesson))
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
		aboutDir,
		educationDir,
		blogDir,
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

type page struct {
	title     string
	directory string
	component templ.Component
}

func (p *page) Templify() templ.Component {
	return p.component
}

func (p *page) GetTitle() string {
	return p.title
}

func (p *page) Directory() string {
	return p.directory
}

func NewHomePage() Templifier {
	return &page{
		title:     "Home",
		directory: rootDir,
		component: HomeComponent(),
	}
}

func NewContactPage() Templifier {
	return &page{
		title:     "Contact",
		directory: rootDir,
		component: ContactComponent(),
	}
}

func NewCoursesListPage(instances []*domain.Course) Templifier {
	return &page{
		title:     "Courses",
		directory: coursesDir,
		component: CoursesListComponent(instances),
	}

}

func NewCoursePage(instance domain.Course) Templifier {
	return &page{
		title:     instance.GetTitle(),
		directory: courseFilePath(instance),
		component: CourseComponent(instance),
	}
}

func NewCourseCalendarPage(course domain.Course) Templifier {
	if course.Name == "" {
		log.Fatal("schedule.Course.Name is empty string")
	}
	return &page{
		title:     course.Name + " Calendar",
		directory: courseFilePath(course),
		component: CourseCalendarComponent(course),
	}
}

func NewUnitPage(unit domain.Unit, instance domain.Course) Templifier {
	return &page{
		title:     unit.GetTitle(),
		directory: unitFilePath(unit, instance),
		component: UnitComponent(unit, instance),
	}
}

func NewLessonPage(lesson domain.Lesson, unit domain.Unit, instance domain.Course) Templifier {
	return &page{
		title:     lesson.GetTitle(),
		directory: lessonFilePath(lesson, unit, instance),
		component: LessonComponent(lesson, unit, instance),
	}
}

// same as FileName() but with "/" at the end instead of the .html extension
func DirName(t Titler) string {
	return strings.ReplaceAll(strings.ToLower(t.GetTitle()), " ", "-") + "/"
}

// replaces spaces with dashes, makes lowercase, and adds .html file extension
func FileName(t Titler) string {
	if t.GetTitle() == "Home" {
		return "index.html"
	}
	return strings.ReplaceAll(strings.ToLower(t.GetTitle()), " ", "-") + ".html"
}

func FilePathToURL(directory string) templ.SafeURL {
	route := RemoveDocsFromPath(directory)
	route = MakeAbsolute(route)
	route = PrefixWithRoot(route)
	return Sanitize(route)
}

func Sanitize(route string) templ.SafeURL {
	return templ.SafeURL(route)
}
func RemoveDocsFromPath(directory string) string {
	return strings.ReplaceAll(directory, "/python/docs", "")
}

func MakeAbsolute(route string) string {
	return strings.ReplaceAll(route, "./", "/")
}

func PrefixWithRoot(route string) string {
	return "/python" + route
}

func rootPages() []Templifier {
	return []Templifier{
		NewHomePage(),
		// NewAboutPage(),
		// NewBlogsListPage(nil),
		NewContactPage(),
		NewCoursesListPage(nil),
	}
}
