package templates

import "github.com/a-h/templ"

type FilesPageSection struct {
	Root             bool
	Empty            bool
	Path             string
	ParentDirectory  File
	CurrentDirectory File
	Files            []File
}

// Filepath implements StaticPage.
func (data FilesPageSection) Filepath() string {
	return data.Path
}

func NewFilesPageSection() StaticPage {
	return FilesPageSection{}
}

type File struct {
	Name  string
	Path  string
	URL   string
	IsDir bool
}

func (data FilesPageSection) Component() templ.Component {
	return FilesPageSectionComponent(data)
}
