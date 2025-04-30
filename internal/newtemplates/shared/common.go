package templates

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/shared/node"
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

func NodePath(nodes ...node.Node) string {
	path := StaticSiteRootDir
	for _, node := range nodes {
		path = strings.ToLower(filepath.Join(path, node.TypeName()+"s"))
		path = strings.ToLower(filepath.Join(
			path,
			fmt.Sprintf("%s_%d", node.TypeName(), node.GetID()),
		))

	}
	return path
}

func NodePage(nodes ...node.Node) string {
	leafNode := nodes[len(nodes)-1]
	path := NodePath(nodes...)
	path = filepath.Join(path, fmt.Sprintf("%s_%d.html", strings.ToLower(leafNode.TypeName()), leafNode.GetID()))
	return path
}

// Student-facing site
func CoursePath(course dto.Course, page bool) string {
	if page {
		return NodePage(course)
	}
	return NodePath(course)
}

// Student-facing site
func CourseImagePath(course dto.Course) string {
	dir := CoursePath(course, false)
	return filepath.Join(dir, "image.png")
}

// Student-facing site
func UnitPath(unit dto.Unit, course dto.Course, page bool) string {
	if page {
		return NodePage(course, unit)
	}
	return NodePath(course, unit)
}

// Student-facing site
func UnitImagePath(unit dto.Unit, course dto.Course) string {
	dir := UnitPath(unit, course, false)
	return filepath.Join(dir, "image.png")
}

// Student-facing site
func LessonPath(lesson dto.Lesson, unit dto.Unit, course dto.Course, page bool) string {
	if page {
		return NodePage(course, unit, lesson)
	}
	return NodePath(course, unit, lesson)
}

// Student-facing site
func LessonImagePath(lesson dto.Lesson, unit dto.Unit, course dto.Course) string {
	dir := LessonPath(lesson, unit, course, false)
	return filepath.Join(dir, "image.png")
}

// Github files link for student-facing site
func LessonFilesURL(lesson dto.Lesson, unit dto.Unit, course dto.Course) templ.SafeURL {
	filePath, err := url.JoinPath(GithubRoot, LessonFilesPath(lesson, unit, course))
	if err != nil {
		log.Println("error generating LessonFilesURL:", err)
	}
	return templ.SafeURL(strings.ReplaceAll(filePath, "/python/docs/courses", ""))
}
func UnitFilesURL(unit dto.Unit, course dto.Course) templ.SafeURL {
	filePath, err := url.JoinPath(GithubRoot, UnitFilesPath(unit, course))
	if err != nil {
		log.Println("error generating UnitFilesURL:", err)
	}
	return templ.SafeURL(strings.ReplaceAll(filePath, "/python/docs/courses", ""))
}
func CourseFilesURL(course dto.Course) templ.SafeURL {
	filePath, err := url.JoinPath(GithubRoot, CourseFilesPath(course))
	if err != nil {
		log.Println("error generating UnitFilesURL:", err)
	}
	return templ.SafeURL(strings.ReplaceAll(filePath, "/python/docs/courses", ""))
}

// Student-facing site
func LessonFilesPath(lesson dto.Lesson, unit dto.Unit, course dto.Course) string {
	return filepath.Join(LessonPath(lesson, unit, course, false), "files")
}

// Student-facing site
func UnitFilesPath(unit dto.Unit, course dto.Course) string {
	return filepath.Join(UnitPath(unit, course, false), "files")
}

// Student-facing site
func CourseFilesPath(course dto.Course) string {
	return filepath.Join(CoursePath(course, false), "files")
}

// Student-facing site
func SlidesPath(lesson dto.Lesson, unit dto.Unit, course dto.Course) string {
	return filepath.Join(LessonPath(lesson, unit, course, false), "slides.html")
}

// This should not be used except for a one-time transfer
func SlidesMarkdownPath(lesson dto.Lesson, unit dto.Unit, course dto.Course) string {
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
