package templates

import (
	"gh_static_portfolio/internal/domain"
	"log"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
)

const (
	GithubRoot           = "https://github.com/whlapinel/python/tree/main/docs/courses"
	StaticSiteRootDir    = "./python/docs/"
	StaticSiteCoursesDir = "./python/docs/courses/"
)

// Student-facing site
func KebabCase(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, " ", "-"))
}

// Student-facing site
func CoursePath(course domain.Course, page bool) string {
	dirPath := filepath.Join(StaticSiteCoursesDir, KebabCase(course.Name))
	if page {
		return filepath.Join(dirPath, KebabCase(course.Name+".html"))
	}
	return dirPath
}

// Student-facing site
func CourseImagePath(course domain.Course) string {
	dir := CoursePath(course, false)
	return filepath.Join(dir, "image.png")
}

// Student-facing site
func UnitPath(unit domain.Unit, course domain.Course, page bool) string {
	dirPath := filepath.Join(CoursePath(course, false), KebabCase(unit.Name))
	if page {
		return filepath.Join(dirPath, KebabCase(unit.Name+".html"))
	}
	return dirPath
}

// Student-facing site
func UnitImagePath(unit domain.Unit, course domain.Course) string {
	dir := UnitPath(unit, course, false)
	return filepath.Join(dir, "image.png")
}

// Student-facing site
func LessonPath(lesson domain.Lesson, unit domain.Unit, course domain.Course, page bool) string {
	dirPath := filepath.Join(UnitPath(unit, course, false), KebabCase(lesson.Name))
	if page {
		return filepath.Join(dirPath, KebabCase(lesson.Name+".html"))
	}
	return dirPath
}

// Student-facing site
func LessonImagePath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	dir := LessonPath(lesson, unit, course, false)
	return filepath.Join(dir, "image.png")
}

// Student-facing site
func LessonFilesURL(lesson domain.Lesson, unit domain.Unit, course domain.Course) templ.SafeURL {
	filePath, err := url.JoinPath(GithubRoot, LessonFilesPath(lesson, unit, course))
	if err != nil {
		log.Println("error generating LessonFilesURL:", err)
	}
	return templ.SafeURL(strings.ReplaceAll(filePath, "/python/docs/courses", ""))
}

// Student-facing site
func LessonFilesPath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(LessonPath(lesson, unit, course, false), "files")
}

// Student-facing site
func SlidesPath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(LessonPath(lesson, unit, course, false), "slides.html")
}

// This should not be used except for a one-time transfer
func SlidesMarkdownPath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return filepath.Join(LessonPath(lesson, unit, course, false), "slides.md")
}

// list of directories to be cleared (used for clearing html files only)
func DirectoriesClearList() []string {
	return []string{
		StaticSiteRootDir,
		StaticSiteCoursesDir,
	}
}

// courses directory to be deleted completely
func DeleteDirList() []string {
	return []string{StaticSiteCoursesDir}
}
