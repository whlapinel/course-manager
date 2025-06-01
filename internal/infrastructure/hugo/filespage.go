package hugo

type FilesPageData struct {
	Path       string     `json:"path"`
	ParentPath string     `json:"parentPath"`
	Files      []File     `json:"files"`
	FilePages  []FilePage `json:"filePages"`
}

type FilePage struct {
	File
	ContentPath string `json:"contentPath"`
}

type File struct {
	Path string `json:"path"`
}

// files page has no child sections
func (d *FilesPageData) Section() *HomogenizedPageData {
	return nil
}

func (d *FilesPageData) Children() []Homogenizer {
	return nil
}

func (d *FilesPageData) Page() *HomogenizedPageData {
	var homoPageData HomogenizedPageData
	homoPageData.Kind = PageKind
	homoPageData.Type = FilesType
	homoPageData.Path = d.Path
	homoPageData.Title = "Files"
	homoPageData.Params = struct {
		Files      []File     `json:"files"`
		FilePages  []FilePage `json:"filePages"`
		ParentPath string     `json:"parentPath"`
	}{
		Files:      d.Files,
		FilePages:  d.FilePages,
		ParentPath: d.ParentPath,
	}
	return &homoPageData
}
