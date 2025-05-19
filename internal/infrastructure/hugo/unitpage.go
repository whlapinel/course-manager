package hugo

import "gh_static_portfolio/internal/app/dto"

type UnitPageData struct {
	dto.Unit            `json:"unit"`
	Designation         string            `json:"designation"`
	Path                string            `json:"path"`
	LessonsListPagePath string            `json:"lessonsListPagePath"`
	LessonPages         []*LessonPageData `json:"lessonPages"`
	FilesPage           *FilesPageData    `json:"filesPage"`
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
	homoPageData.Title = d.Name
	homoPageData.Params = struct {
		ChildSectionPath string             `json:"childSectionPath"`
		FilesPagePath    string             `json:"filesPagePath"`
		BreadCrumbs      BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ChildSectionPath: d.LessonsListPagePath,
		FilesPagePath:    d.FilesPage.Path,
		BreadCrumbs:      BreadCrumbs(d.Path),
	}
	return &homoPageData
}

func (d *UnitPageData) Section() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = SectionKind
	homoPageData.Type = LessonType
	homoPageData.Path = d.LessonsListPagePath
	homoPageData.Title = "Lessons"
	homoPageData.Params = struct {
		ParentPath  string             `json:"parentPath"`
		BreadCrumbs BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ParentPath:  d.Path,
		BreadCrumbs: BreadCrumbs(d.Path),
	}
	return &homoPageData
}
