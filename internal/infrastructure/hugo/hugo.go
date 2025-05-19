package hugo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/ports"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Params struct {
	Domain                   string
	HugoURL                  string
	SitesRootDir             string
	DataFilesRoot            string
	DataPathingService       ports.PathingService
	StaticSitePathingService ports.PathingService
	GetUser                  func(userID string) (dto.User, error)
	GetTerm                  func(termID int) (dto.Term, error)
	GetCourses               func(termID int) ([]dto.Course, error)
	GetUnits                 func(courseID int) ([]dto.Unit, error)
	GetLessons               func(unitID int) ([]dto.Lesson, error)
}

type hugoGenerator struct {
	Params
	DataFilesSymlinkRoot string
}

func (h *hugoGenerator) configure(user dto.User) error {
	userDataPath, err := filepath.Abs(h.DataPathingService.NodeDirPath(user))
	if err != nil {
		return err
	}
	config := NewConfig(user, NewHugoConfigParams{
		Domain:       h.Domain,
		Title:        user.Username(),
		UserDataPath: userDataPath,
		ConfigPath:   filepath.Join(h.SitesRootDir, user.ID, "hugo.toml"),
	})
	return config.Write()
}

func New(params Params) (ports.SiteGenerator, error) {
	return &hugoGenerator{
		Params: params,
	}, nil
}

func (h *hugoGenerator) SinglePagePath(svc ports.PathingService, nodes ...ports.Node) (string, error) {
	path := svc.NodeDirPath(nodes...)
	path, err := h.contentPath(path)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (h *hugoGenerator) ListPagePath(svc ports.PathingService, nodes ...ports.Node) (string, error) {
	path := svc.NodeChildrenDirPath(nodes...)
	path, err := h.contentPath(path)
	if err != nil {
		return "", err
	}
	return path, nil
}

// removes first two route segments static root and user directory,
// e.g. "hugosites/123556/users/user_123/terms" becomes "terms"
func (h *hugoGenerator) contentPath(path string) (string, error) {
	segments := strings.Split(path, "/")
	if len(segments) < 2 {
		return "", fmt.Errorf("less than 2 segments in argument path: %s has %d segments", path, len(segments))
	}
	new := filepath.Join(segments[2:]...)
	return new, nil
}

// this is for json data directory
func (h *hugoGenerator) dataDir(userID string) string {
	return filepath.Join(h.SitesRootDir, userID, "data")
}

func (h *hugoGenerator) Build(userID string, termID int) error {
	user, err := h.GetUser(userID)
	if err != nil {
		return err
	}
	err = h.configure(user)
	if err != nil {
		return err
	}
	pageData, err := h.PageData(user, termID)
	if err != nil {
		return err
	}
	homogenized := h.HomogenizedData(pageData)

	log.Println("len(homogenized): ", len(homogenized))
	// encode to json
	content, err := json.Marshal(homogenized)
	if err != nil {
		return err
	}
	dataDir := h.dataDir(userID)
	// write data to site data folder
	err = os.MkdirAll(dataDir, os.ModePerm)
	if err != nil {
		return err
	}
	path := filepath.Join(dataDir, "data.json")
	log.Println("writing to path:", path)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	size, err := file.Write(content)
	if err != nil {
		return err
	}
	log.Println(size, "bytes written")
	err = file.Sync()
	if err != nil {
		return err
	}

	// check to see if site folder exists
	siteRoot := filepath.Join(h.SitesRootDir, userID)
	_, err = os.Stat(siteRoot)
	if err != nil {
		return err
	}

	cmd := exec.Command("hugo")

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	cmd.Dir = siteRoot
	err = cmd.Run()
	log.Println("STDOUT:", outBuf.String())
	if err != nil {
		log.Println("Hugo error:", err)
		log.Println("STDERR:", errBuf.String())
		return err
	}
	return nil
}

func (h *hugoGenerator) HomogenizedData(pageData Homogenizer) []*HomogenizedPageData {
	var homogenized []*HomogenizedPageData
	if pageData.Page() != nil {
		homogenized = append(homogenized, pageData.Page())
	}
	if pageData.Section() != nil {
		homogenized = append(homogenized, pageData.Section())
	}
	for _, page := range pageData.Children() {
		//recursive case
		pages := h.HomogenizedData(page)
		if pages != nil {
			homogenized = append(homogenized, pages...)
		}
	}
	// base case
	return homogenized
}

func (h *hugoGenerator) PageData(user dto.User, termID int) (*PageData, error) {
	sitePathingService := h.StaticSitePathingService.WithSegment(user.ID)
	term, err := h.GetTerm(termID)
	if err != nil {
		return nil, err
	}
	termPagePath, err := h.SinglePagePath(sitePathingService, term)
	if err != nil {
		return nil, err
	}
	coursesListPagePath, err := h.ListPagePath(sitePathingService, term)
	if err != nil {
		return nil, err
	}
	var nodes = []ports.Node{term}
	filesDirPath := filepath.Join(termPagePath, "files")
	files, err := h.FilePaths(append([]ports.Node{user}, nodes...)...)
	pages, raw := FilePages(files)

	if err != nil {
		return nil, err
	}
	var pageData = &PageData{
		TermPageData: &TermPageData{
			Term:               term,
			Path:               termPagePath,
			CourseListPagePath: coursesListPagePath,
			FilesPage: &FilesPageData{
				Path:       filesDirPath,
				ParentPath: termPagePath,
				Files:      raw,
				FilePages:  pages,
			},
		},
	}
	courses, err := h.GetCourses(termID)
	if err != nil {
		return nil, err
	}
	var coursePages []*CoursePageData
	for _, course := range courses {
		var nodes = []ports.Node{term, course}
		coursePagePath, err := h.SinglePagePath(sitePathingService, nodes...)
		if err != nil {
			return nil, err
		}

		unitsListPagePath, err := h.ListPagePath(sitePathingService, nodes...)
		if err != nil {
			return nil, err
		}
		filesDirPath := filepath.Join(coursePagePath, "files")
		files, err := h.FilePaths(append([]ports.Node{user}, nodes...)...)
		pages, raw := FilePages(files)

		if err != nil {
			return nil, err
		}
		coursePage := &CoursePageData{
			Course:            course,
			Path:              coursePagePath,
			UnitsListPagePath: unitsListPagePath,
			FilesPage: &FilesPageData{
				Path:       filesDirPath,
				ParentPath: coursePagePath,
				Files:      raw,
				FilePages:  pages,
			},
		}
		var unitPages []*UnitPageData
		units, err := h.GetUnits(coursePage.Course.Course.ID)
		if err != nil {
			return nil, err
		}
		for _, unit := range units {
			var nodes = []ports.Node{term, course, unit}
			unitPagePath, err := h.SinglePagePath(sitePathingService, nodes...)
			if err != nil {
				return nil, err
			}
			lessonsListPagePath, err := h.ListPagePath(sitePathingService, nodes...)
			if err != nil {
				return nil, err
			}
			files, err := h.FilePaths(append([]ports.Node{user}, nodes...)...)
			pages, raw := FilePages(files)
			filesDirPath := filepath.Join(unitPagePath, "files")
			if err != nil {
				return nil, err
			}
			unitPage := &UnitPageData{
				Unit:                unit,
				Path:                unitPagePath,
				LessonsListPagePath: lessonsListPagePath,
				FilesPage: &FilesPageData{
					Path:       filesDirPath,
					ParentPath: unitPagePath,
					Files:      raw,
					FilePages:  pages,
				},
			}
			lessons, err := h.GetLessons(unit.ID)
			if err != nil {
				return nil, err
			}
			var lessonPages []*LessonPageData
			for _, lesson := range lessons {
				var nodes = []ports.Node{term, course, unit, lesson}
				lessonPagePath, err := h.SinglePagePath(sitePathingService, nodes...)
				if err != nil {
					return nil, err
				}
				filesDirPath := filepath.Join(lessonPagePath, "files")
				files, err := h.FilePaths(append([]ports.Node{user}, nodes...)...)
				pages, raw := FilePages(files)

				if err != nil {
					return nil, err
				}

				lessonPage := &LessonPageData{
					Lesson: lesson,
					Path:   lessonPagePath,
					FilesPage: FilesPageData{
						Path:       filesDirPath,
						ParentPath: lessonPagePath,
						Files:      raw,
						FilePages:  pages,
					},
					Content: strings.ReplaceAll(
						`
						# This is a test of the markdown system!					
						`,
						"\t", "",
					),
				}
				lessonPages = append(lessonPages, lessonPage)
			}
			unitPage.LessonPages = lessonPages
			unitPages = append(unitPages, unitPage)
		}
		coursePage.UnitPages = unitPages
		coursePages = append(coursePages, coursePage)
	}
	pageData.CoursePages = coursePages
	return pageData, nil
}

type Homogenizer interface {
	// every page type will create a page (of kind page)
	Page() *HomogenizedPageData
	// first page of each type will create one section page to contain its children,
	// if children is nil this will also be nil. for example course page will create a section page for units
	// but lesson page will return nil
	Section() *HomogenizedPageData
	// e.g. term page has course children
	Children() []Homogenizer
}

func (h *hugoGenerator) FilePaths(nodes ...ports.Node) ([]string, error) {
	dirPath := h.DataPathingService.NodeFilesDirPath(nodes...)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	var paths []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(dirPath, entry.Name())
		path = contentPath(path)
		paths = append(paths, path)
	}
	return paths, nil
}

// splits the paths into pages and raw files, where pages have the .md extension
// since these should be rendered by hugo as pages with a path stripped of the extension
func FilePages(paths []string) (pages []string, raw []string) {
	for _, path := range paths {
		if filepath.Ext(path) == ".md" {
			pages = append(pages, path[:len(path)-3])
		} else {
			raw = append(raw, path)
		}
	}
	return pages, raw
}

// strips the first 4 segments off the root
// internal/data/users/user_123/terms/term_2 becomes terms/term_2
func contentPath(path string) string {
	segments := strings.Split(path, "/")
	return filepath.Join(segments[4:]...)
}
