package hugo

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"strings"
)

type LessonPageData struct {
	dto.Lesson     `json:"lesson"`
	Designation    string        `json:"designation"`
	Path           string        `json:"path"`
	Content        string        `json:"content"`
	FilesPage      FilesPageData `json:"filesPage"`
	SlidesPath     string        `json:"slidesPath"`
	SlidesDataPath string        `json:"slidesDataPath"`
}

func (d *LessonPageData) Children() []Homogenizer {
	return []Homogenizer{&d.FilesPage}
}

func (d *LessonPageData) Page() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = PageKind
	homoPageData.Type = LessonType
	homoPageData.Path = d.Path
	homoPageData.URL = strings.ReplaceAll(d.Path, "_", "-")
	homoPageData.Weight = d.Number
	homoPageData.Title = fmt.Sprintf("%s: %s", d.Designation, d.Name)
	homoPageData.Params = struct {
		FilesPagePath  string             `json:"filesPagePath"`
		SlidesPath     string             `json:"slidesPath"`
		SlidesDataPath string             `json:"slidesDataPath"`
		BreadCrumbs    BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		SlidesPath:     d.SlidesPath,
		SlidesDataPath: d.SlidesDataPath,
		FilesPagePath:  d.FilesPage.Path,
		BreadCrumbs:    BreadCrumbs(d.Path),
	}
	return &homoPageData
}

func (d *LessonPageData) Section() *HomogenizedPageData {
	return nil
}
