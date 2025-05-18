package hugo

import "gh_static_portfolio/internal/app/dto"

type CoursePageData struct {
	dto.Course        `json:"course"`
	Path              string          `json:"path"`
	UnitsListPagePath string          `json:"unitListPagePath"`
	UnitPages         []*UnitPageData `json:"unitPages"`
	FilesPage         *FilesPageData  `json:"filesPage"`
}

func (d *CoursePageData) Children() []Homogenizer {
	var homogenized []Homogenizer
	homogenized = append(homogenized, d.FilesPage)

	for _, page := range d.UnitPages {
		homogenized = append(homogenized, page)
	}
	return homogenized
}

func (d *CoursePageData) Page() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = PageKind
	homoPageData.Type = CourseType
	homoPageData.Path = d.Path
	homoPageData.Title = d.Course.Name
	homoPageData.Params = struct {
		ChildSectionPath string             `json:"childSectionPath"`
		FilesPagePath    string             `json:"filesPagePath"`
		BreadCrumbs      BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ChildSectionPath: d.UnitsListPagePath,
		FilesPagePath:    d.FilesPage.Path,
		BreadCrumbs:      BreadCrumbs(d.Path),
	}
	return &homoPageData
}

func (d *CoursePageData) Section() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = SectionKind
	homoPageData.Type = UnitType
	homoPageData.Title = "Units"
	homoPageData.Path = d.UnitsListPagePath
	homoPageData.Params = struct {
		ParentPath  string             `json:"parentPath"`
		BreadCrumbs BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ParentPath:  d.Path,
		BreadCrumbs: BreadCrumbs(d.Path),
	}
	return &homoPageData
}
