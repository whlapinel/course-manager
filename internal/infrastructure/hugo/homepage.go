package hugo

import "gh_static_portfolio/internal/ports"

type PageData struct {
	Units    []*UnitPageData `json:"unitPages"`
	Calendar *CalendarPageData
	Files    *FilesPageData `json:"filesPage"`
	ports.BreadCrumbsMaker
}

func (d *PageData) Children() []Homogenizer {
	var homos []Homogenizer
	homos = append(homos, d.Files, d.Calendar)

	for _, unitPage := range d.Units {
		homos = append(homos, unitPage)
	}
	return homos
}

func (d *PageData) Page() *HomogenizedPageData {
	return nil
}

func (d *PageData) Section() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = SectionKind
	homoPageData.Type = UnitType
	homoPageData.Path = "units"
	homoPageData.Title = "Units"
	homoPageData.Params = struct {
		ParentPath  string                   `json:"parentPath"`
		BreadCrumbs ports.BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ParentPath:  "/",
		BreadCrumbs: d.BreadCrumbsMaker.BreadCrumbs("units"),
	}
	return &homoPageData
}
