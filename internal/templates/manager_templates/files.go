package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type FilesPage struct {
	Params             CourseIDParams
	CurrentPath        string
	OpenFileRHN        string
	UploadFileRHN      string
	Node               domain.CourseNode
	Files              []FilesPageItem
	E                  *echo.Echo
	PopRouteSegmentRHN string
	BreadCrumbsData    BreadCrumbs
}

type FilesPageItem struct {
	Path  string
	IsDir bool
}

func (data FilesPage) BreadCrumbs() BreadCrumbs {
	return data.BreadCrumbsData
}

func (data FilesPage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: "Files for " + data.Node.GetName(),
		UpNav: UpNav{
			URL:  data.E.Reverse(data.PopRouteSegmentRHN, data.Params.ToIntSlice()...),
			Text: "Up one level",
		},
	}
}
func (data FilesPage) Component() templ.Component {
	return FilesComponent(data)
}
