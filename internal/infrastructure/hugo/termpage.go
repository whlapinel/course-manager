package hugo

import "gh_static_portfolio/internal/app/dto"

type TermPageData struct {
	Term               dto.Term          `json:"term"`
	Path               string            `json:"path"`
	CourseListPagePath string            `json:"courseListPagePath"`
	CoursePages        []*CoursePageData `json:"coursePages"`
	FilesPage          *FilesPageData    `json:"filesPage"`
}

func (d *TermPageData) Children() []Homogenizer {
	var homoSlice []Homogenizer
	homoSlice = append(homoSlice, d.FilesPage)
	for _, child := range d.CoursePages {
		homoSlice = append(homoSlice, child)
	}
	return homoSlice
}
func (d *TermPageData) Page() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = PageKind
	homoPageData.Type = TermType
	homoPageData.Path = d.Path
	homoPageData.Title = d.Term.Name
	homoPageData.Params = struct {
		ChildSectionPath string             `json:"childSectionPath"`
		FilesPagePath    string             `json:"filesPagePath"`
		BreadCrumbs      BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ChildSectionPath: d.CourseListPagePath,
		FilesPagePath:    d.FilesPage.Path,
		BreadCrumbs:      BreadCrumbs(d.Path),
	}
	return &homoPageData
}

func (d *TermPageData) Section() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = SectionKind
	homoPageData.Type = CourseType
	homoPageData.Title = "Courses"
	homoPageData.Path = d.CourseListPagePath
	homoPageData.Params = struct {
		ParentPath  string             `json:"parentPath"`
		BreadCrumbs BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ParentPath:  d.Path,
		BreadCrumbs: BreadCrumbs(d.Path),
	}
	return &homoPageData
}
