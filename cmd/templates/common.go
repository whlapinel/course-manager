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
	rootDir      = "./whlapinel.github.io/docs/"
	aboutDir     = "./whlapinel.github.io/docs/about/"
	educationDir = "./whlapinel.github.io/docs/about/education/"
	blogDir      = "./whlapinel.github.io/docs/blog/"
	coursesDir   = "./whlapinel.github.io/docs/courses/"
)

func courseRoutePath(course domain.Course) string {
	return fmt.Sprintf("%s%s", coursesDir, DirName(course))
}

func unitRoutePath(unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("%s%s%s", coursesDir, DirName(course), DirName(unit))
}

func lessonRoutePath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("%s%s%s%s", coursesDir, DirName(course), DirName(unit), DirName(lesson))
}

func filesRoutePath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("https://github.com/whlapinel/whlapinel.github.io/tree/main/docs/courses/%s%s%sfiles", DirName(course), DirName(unit), DirName(lesson))
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

func NewAboutPage() Templifier {
	return &page{
		title:     "About",
		directory: rootDir,
		component: AboutComponent(),
	}
}

func NewEducationPage(items []*domain.EducationItem) Templifier {
	return &page{
		title:     "Education",
		directory: aboutDir,
		component: EducationListComponent(items),
	}
}

func NewClassesPage(item *domain.EducationItem) Templifier {
	return &page{
		title:     item.School,
		directory: educationDir,
		component: ClassListComponent(item),
	}
}

func NewWorkHistoryPage(items []*domain.WorkHistoryItem) Templifier {
	return &page{
		title:     "Work History",
		directory: aboutDir,
		component: WorkHistoryComponent(items),
	}
}

func NewProjectsPage(items []*domain.ProjectItem) Templifier {
	return &page{
		title:     "Projects",
		directory: aboutDir,
		component: ProjectListComponent(items),
	}
}

func NewSkillsPage(items []*domain.SkillItem) Templifier {
	return &page{
		title:     "Skills",
		directory: aboutDir,
		component: SkillsListComponent(items),
	}
}

func NewBlogsListPage(items []*domain.Blog) Templifier {
	return &page{
		title:     "Blog",
		directory: blogDir,
		component: BlogsListComponent(items),
	}
}

func NewBlogPage(item *domain.Blog) Templifier {
	if item == nil {
		return &page{
			directory: blogDir,
			component: BlogComponent(nil),
		}
	}
	return &page{
		title:     item.GetTitle(),
		directory: blogDir,
		component: BlogComponent(item),
	}
}

func NewPersonalPage() Templifier {
	return &page{
		title:     "Personal",
		directory: aboutDir,
		component: PersonalComponent(),
	}
}

func NewCoursesListPage(courses []*domain.Course) Templifier {
	return &page{
		title:     "Courses",
		directory: coursesDir,
		component: CoursesListComponent(courses),
	}

}

func NewCoursePage(course *domain.Course) Templifier {
	return &page{
		title:     course.GetTitle(),
		directory: courseRoutePath(*course),
		component: CourseComponent(*course),
	}
}

func NewCourseCalendarPage(course domain.Course) Templifier {
	return &page{
		title:     course.GetTitle() + " Calendar",
		directory: courseRoutePath(course),
		component: CourseCalendarComponent(course),
	}
}

func NewUnitPage(unit domain.Unit, course domain.Course) Templifier {
	return &page{
		title:     unit.GetTitle(),
		directory: unitRoutePath(unit, course),
		component: UnitComponent(unit, course),
	}
}

func NewLessonPage(lesson domain.Lesson, unit domain.Unit, course domain.Course) Templifier {
	return &page{
		title:     lesson.GetTitle(),
		directory: lessonRoutePath(lesson, unit, course),
		component: LessonComponent(lesson, unit, course),
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

func RemoveDocsFromPath(directory string) string {
	return strings.ReplaceAll(directory, "./whlapinel.github.io/docs", "")
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
