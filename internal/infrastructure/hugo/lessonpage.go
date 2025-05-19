package hugo

import "gh_static_portfolio/internal/app/dto"

type LessonPageData struct {
	dto.Lesson  `json:"lesson"`
	Designation string        `json:"designation"`
	Path        string        `json:"path"`
	Content     string        `json:"content"`
	FilesPage   FilesPageData `json:"filesPage"`
}

func (d *LessonPageData) Children() []Homogenizer {
	return []Homogenizer{&d.FilesPage}
}

func (d *LessonPageData) Page() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = PageKind
	homoPageData.Type = LessonType
	homoPageData.Path = d.Path
	homoPageData.Title = d.Name
	homoPageData.Params = struct {
		FilesPagePath string             `json:"filesPagePath"`
		BreadCrumbs   BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		FilesPagePath: d.FilesPage.Path,
		BreadCrumbs:   BreadCrumbs(d.Path),
	}
	return &homoPageData
}

func (d *LessonPageData) Section() *HomogenizedPageData {
	return nil
}
