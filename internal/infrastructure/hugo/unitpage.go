package hugo

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/ports"
	"strings"
)

type UnitPageData struct {
	dto.Unit            `json:"unit"`
	Designation         string            `json:"designation"`
	Path                string            `json:"path"`
	LessonsListPagePath string            `json:"lessonsListPagePath"`
	LessonPages         []*LessonPageData `json:"lessonPages"`
	FilesPage           *FilesPageData    `json:"filesPage"`
	ports.BreadCrumbsMaker
}

func (d *UnitPageData) Children() []Homogenizer {
	var homogenized []Homogenizer
	homogenized = append(homogenized, d.FilesPage)
	for _, page := range d.LessonPages {
		homogenized = append(homogenized, page)
	}
	return homogenized
}

func (d *UnitPageData) Page() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = PageKind
	homoPageData.Type = UnitType
	homoPageData.Path = d.Path
	homoPageData.URL = strings.ReplaceAll(d.Path, "_", "-")
	homoPageData.Weight = d.Sequence
	homoPageData.Title = fmt.Sprintf("%s: %s", d.Designation, d.Name)
	homoPageData.Params = struct {
		ChildSectionPath string                   `json:"childSectionPath"`
		FilesPagePath    string                   `json:"filesPagePath"`
		BreadCrumbs      ports.BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ChildSectionPath: d.LessonsListPagePath,
		FilesPagePath:    d.FilesPage.Path,
		BreadCrumbs:      d.BreadCrumbsMaker.BreadCrumbs(d.Path),
	}
	return &homoPageData
}

func (d *UnitPageData) Section() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = SectionKind
	homoPageData.Type = LessonType
	homoPageData.Path = d.LessonsListPagePath
	homoPageData.URL = strings.ReplaceAll(d.LessonsListPagePath, "_", "-")
	homoPageData.Title = "Lessons"
	homoPageData.Params = struct {
		ParentPath  string                   `json:"parentPath"`
		BreadCrumbs ports.BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ParentPath:  d.Path,
		BreadCrumbs: d.BreadCrumbsMaker.BreadCrumbs(d.LessonsListPagePath),
	}
	return &homoPageData
}
