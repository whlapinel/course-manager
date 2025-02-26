package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type FilesPage struct {
	Params             domain.NodePath
	CurrentPath        string
	OpenFileRHN        string
	UploadFileRHN      string
	Node               domain.CourseNode
	Files              []FilesPageItem
	E                  *echo.Echo
	PopRouteSegmentRHN string
	BreadCrumbsData    BreadCrumbs
	ViewMarkdownRHN    string
}

type FilesPageItem struct {
	Path  string
	IsDir bool
}

func (data FilesPage) BreadCrumbs() BreadCrumbs {
	return data.BreadCrumbsData
}

func AddParams(params domain.NodePath, additionalParams ...any) []any {
	pathSlice := params.ToSlice()
	for _, param := range additionalParams {
		pathSlice = append(pathSlice, param)
	}
	return pathSlice
}
func (data FilesPage) FileURL(file FilesPageItem) string {
	return data.E.Reverse(data.ViewMarkdownRHN, AddParams(data.Params, file.Path)...)

}

func (data FilesPage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: "Files for " + data.Node.GetName(),
		UpNav: UpNav{
			URL:  data.E.Reverse(data.PopRouteSegmentRHN, data.Params.ToSlice()...),
			Text: "Up one level",
		},
	}
}

// view as HTML
func (data FilesPage) ViewMarkdownButton(file FilesPageItem) templ.Component {
	return HXButton{
		Text:     "View As HTML",
		Method:   HxGet,
		HxTarget: "#markdown",
		URL:      data.FileURL(file),
		PushURL:  true,
	}.Component()

}
func (data FilesPage) Component() templ.Component {
	return FilesComponent(data)
}
