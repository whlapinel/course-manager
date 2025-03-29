package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type FilesPage struct {
	Root                bool
	ParentDirectory     FilesPageItem
	CurrentDirectory    FilesPageItem
	Params              domain.NodePath
	CurrentPath         string
	OpenFileRHN         string
	UploadFileRHN       string
	Node                domain.CourseNode
	Files               []FilesPageItem
	E                   *echo.Echo
	PopRouteSegmentRHN  string
	BreadCrumbsData     BreadCrumbs
	ViewMarkdownRHN     string
	EditMarkdownFileURL func(relPath string) string
}

type FilesPageItem struct {
	Name  string
	URL   string
	Path  string
	IsDir bool
}

type MarkdownEditor struct {
	Params          domain.NodePath
	Path            string
	Contents        string
	PostEditFileURL func(relPath string) string
	E               *echo.Echo
}

func (data MarkdownEditor) Component() templ.Component {
	return MarkdownEditorComponent(data)
}

func (data FilesPage) BreadCrumbs() BreadCrumbs {
	return data.BreadCrumbsData
}

func AddParams(params domain.NodePath, additionalParams ...any) []any {
	pathSlice := params.ToSlice()
	pathSlice = append(pathSlice, additionalParams...)
	return pathSlice
}
func (data FilesPage) FileURL(file FilesPageItem) string {
	return data.E.Reverse(data.ViewMarkdownRHN, AddParams(data.Params, file.Path)...)

}

func (data FilesPage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Files for " + data.Node.GetName(),
		UpNav: cmp.UpNav{
			URL:  data.E.Reverse(data.PopRouteSegmentRHN, data.Params.ToSlice()...),
			Text: "Up one level",
		},
		Crumbs: data.BreadCrumbs().BreadCrumbs(),
	}
}

func (data FilesPage) EditMarkdownButton(file FilesPageItem) templ.Component {
	return cmp.Button{
		Text:     "Edit",
		Method:   cmp.HxGet,
		HxTarget: "#markdown",
		PushURL:  true,
		URL:      data.EditMarkdownFileURL(file.Path),
		Image:    cmp.EditImage(),
	}.Component()
}

// view as HTML
func (data FilesPage) ViewMarkdownButton(file FilesPageItem) templ.Component {
	return cmp.Button{
		Text:     "View As HTML",
		Method:   cmp.HxGet,
		HxTarget: "#markdown",
		URL:      data.FileURL(file),
		PushURL:  true,
	}.Component()

}
func (data FilesPage) Component() templ.Component {
	return NewFilesComponent(data)
}
