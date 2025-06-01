package hugo

type PageData struct {
	Units []*UnitPageData `json:"unitPages"`
	Files *FilesPageData  `json:"filesPage"`
}

func (d *PageData) Children() []Homogenizer {
	var homos []Homogenizer
	homos = append(homos, d.Files)
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
		ParentPath  string             `json:"parentPath"`
		BreadCrumbs BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ParentPath:  "/",
		BreadCrumbs: BreadCrumbs("/units"),
	}
	return &homoPageData
}
