package hugo

type PageData struct {
	*TermPageData `json:"termPageData"`
}

func (d *PageData) Children() []Homogenizer {
	return []Homogenizer{
		d.TermPageData,
	}
}

func (d *PageData) Page() *HomogenizedPageData {
	return nil
}

func (d *PageData) Section() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = SectionKind
	homoPageData.Type = TermType
	homoPageData.Path = "terms"
	homoPageData.Title = "Terms"
	homoPageData.Params = struct {
		ParentPath  string             `json:"parentPath"`
		BreadCrumbs BreadCrumbsPartial `json:"breadCrumbs"`
	}{
		ParentPath:  "",
		BreadCrumbs: BreadCrumbs(d.Path),
	}
	return &homoPageData
}
