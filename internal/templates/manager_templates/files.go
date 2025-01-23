package managertemplates

import (
	"gh_static_portfolio/internal/domain"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type FilesPage struct {
	Params        CourseIDParams
	CurrentPath   string
	OpenFileRHN   string
	UploadFileRHN string
	Node          domain.CourseNode
	Files         []FilesPageItem
	E             *echo.Echo
}

type FilesPageItem struct {
	Path  string
	IsDir bool
}

func (data FilesPage) Component() templ.Component {
	return FilesComponent(data)

}
