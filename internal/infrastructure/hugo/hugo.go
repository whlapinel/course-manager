package hugo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	calendarviews "gh_static_portfolio/internal/app/views/calendar"
	"gh_static_portfolio/internal/ports"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type CalendarService interface {
	CalendarDates(courseID int, term dto.Term) (calendarviews.DatesMap, error)
}

type Params struct {
	Domain       string
	HugoURL      string
	SitesRootDir string
	CalendarService
	DataFilesRoot            string
	DataPathingService       ports.PathingService
	StaticSitePathingService ports.PathingService
	GetUnits                 func(courseID int) ([]dto.Unit, error)
	GetLessons               func(unitID int) ([]dto.Lesson, error)
}

type hugoGenerator struct {
	Params
	DataFilesSymlinkRoot string
}

func (h *hugoGenerator) StaticSiteURL(lastName string, courseID int) string {
	return staticURLMaker(h.Domain)(lastName, courseID)
}

func (h *hugoGenerator) courseSitePath(lastName string, courseID int) string {
	return filepath.Join(h.SitesRootDir, h.courseSiteDirName(lastName, courseID))
}

func (h *hugoGenerator) courseSiteDirName(lastName string, courseID int) string {
	lastName = strings.ReplaceAll(lastName, " ", "-")
	lastName = strings.ToLower(lastName)
	return fmt.Sprintf("%s-%d", lastName, courseID)
}

func (h *hugoGenerator) configure(user dto.User, term dto.Term, course dto.Course) (*HugoConfig, error) {
	var config HugoConfig
	dataPath, err := filepath.Abs(h.DataPathingService.NodeDirPath(user, term, course))
	if err != nil {
		return nil, err
	}
	siteRoot := h.courseSitePath(user.LastName, course.Course.ID)
	config = NewConfig(user, NewHugoConfigParams{
		Domain:              h.Domain,
		Title:               course.Course.Name,
		Subtitle:            term.Name,
		Username:            fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		SiteRoot:            siteRoot,
		DestinationDataPath: filepath.Join(siteRoot, "data"),
		SourceDataPath:      dataPath,
		ConfigPath:          filepath.Join(siteRoot, "hugo.toml"),
		CourseID:            course.Course.ID,
	})
	return &config, nil
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

func (h *hugoGenerator) Build(user, term, course ports.Node) error {
	courseSiteDir := h.courseSitePath(user.(dto.User).LastName, course.GetID().(int))
	publicDir := filepath.Join(courseSiteDir, "public")
	err := os.RemoveAll(publicDir)
	if err != nil {
		return err
	}
	userDTO, ok := user.(dto.User)
	if !ok {
		return fmt.Errorf("node is not user")
	}
	termDTO, ok := term.(dto.Term)
	if !ok {
		return fmt.Errorf("node is not term")
	}
	courseDTO, ok := course.(dto.Course)
	if !ok {
		return fmt.Errorf("node is not course")
	}
	config, err := h.configure(userDTO, termDTO, courseDTO)
	if err != nil {
		return err
	}
	err = config.Write()
	if err != nil {
		return err
	}
	pageData, err := h.PageData(*config, userDTO, termDTO, courseDTO)
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
	// write data to site data folder
	err = os.MkdirAll(config.DestinationDataPath, os.ModePerm)
	if err != nil {
		return err
	}
	path := filepath.Join(config.DestinationDataPath, "data.json")
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
	_, err = os.Stat(config.SiteRoot)
	if err != nil {
		return err
	}
	cmd := exec.Command("hugo", "--logLevel", "debug")

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	cmd.Dir = config.SiteRoot
	err = cmd.Run()
	log.Println("STDOUT:", outBuf.String())
	log.Println("STDERR:", errBuf.String())
	if err != nil {
		log.Println("Hugo error:", err)
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

func (h *hugoGenerator) PageData(config HugoConfig, user dto.User, term dto.Term, course dto.Course) (*PageData, error) {
	var pageData PageData
	calDates, err := h.CalendarService.CalendarDates(course.Course.ID, term)
	if err != nil {
		return nil, err
	}
	sitePathingService := h.StaticSitePathingService.WithSegment(h.courseSiteDirName(user.LastName, course.Course.ID))
	singlePagePath := func(unit, lesson ports.Node) (string, error) {
		path, err := h.SinglePagePath(sitePathingService, unit, lesson)
		if err != nil {
			return "", err
		}
		return path, nil
	}
	calendar, err := NewCalendar(term, calDates, singlePagePath)
	if err != nil {
		return nil, err
	}
	pageData.Calendar = calendar
	filesDirPath := "files"
	pageData.Files = &FilesPageData{
		Path:       filesDirPath,
		ParentPath: "/",
	}
	var unitPages []*UnitPageData
	units, err := h.GetUnits(course.Course.ID)
	if err != nil {
		return nil, err
	}
	for _, unit := range units {
		var nodes = []ports.Node{unit}
		unitPagePath, err := h.SinglePagePath(sitePathingService, nodes...)
		if err != nil {
			return nil, err
		}
		lessonsListPagePath, err := h.ListPagePath(sitePathingService, nodes...)
		if err != nil {
			return nil, err
		}
		filesDirPath := filepath.Join(unitPagePath, "files")
		unitPage := &UnitPageData{
			Unit:                unit,
			Designation:         unit.Designation(),
			Path:                unitPagePath,
			LessonsListPagePath: lessonsListPagePath,
			FilesPage: &FilesPageData{
				Path:       filesDirPath,
				ParentPath: unitPagePath,
			},
		}
		lessons, err := h.GetLessons(unit.ID)
		if err != nil {
			return nil, err
		}
		var lessonPages []*LessonPageData
		for _, lesson := range lessons {
			var nodes = []ports.Node{unit, lesson}
			lessonPagePath, err := h.SinglePagePath(sitePathingService, nodes...)
			if err != nil {
				return nil, err
			}
			filesDirPath := filepath.Join(lessonPagePath, "files")
			slidesDataPath := h.DataPathingService.NodeSlidesHTMLPath(user, term, course, unit, lesson)
			slidesDataPath = contentPath(slidesDataPath)
			log.Println("slidesDataPath:", slidesDataPath)
			slidesPath := strings.ReplaceAll(slidesDataPath, ".slides.html", "slides.html")
			slidesPath = strings.ReplaceAll(slidesPath, "_", "-")
			lessonPage := &LessonPageData{
				Lesson:      lesson,
				Designation: lesson.Designation(),
				Path:        lessonPagePath,
				FilesPage: FilesPageData{
					Path:       filesDirPath,
					ParentPath: lessonPagePath,
				},
				Content: strings.ReplaceAll(
					`
						# This is a test of the markdown system!					
						`,
					"\t", "",
				),
				SlidesPath:     slidesPath,
				SlidesDataPath: slidesDataPath,
			}
			lessonPages = append(lessonPages, lessonPage)
		}
		unitPage.LessonPages = lessonPages
		unitPages = append(unitPages, unitPage)
	}
	pageData.Units = unitPages
	return &pageData, nil
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

// strips the first 8 segments off the root
// internal/data/users/user_123/terms/term_2/courses/course_5/units/unit_1 units/unit_1
func contentPath(path string) string {
	segments := strings.Split(path, "/")
	return filepath.Join(segments[8:]...)
}
