package templates

import (
	"gh_static_portfolio/internal/templates/components"

	"github.com/a-h/templ"
)

type StaticHomePageParams struct {
	HomePage
}

func NewHomePage(params StaticHomePageParams) HomePage {
	return HomePage{
		Path: params.Path,
		StaticLayout: StaticLayout{
			User:      params.User,
			AssetsURL: params.AssetsURL,
		},
	}
}

type HomePage struct {
	StaticLayout
	Path string
}

// Component implements StaticPage.
func (data HomePage) Component() templ.Component {
	return components.Layout{
		Head: Head{
			User:      data.User,
			AssetsURL: data.AssetsURL,
		}.Component(),
		Page: HomeComponent(data),
	}.Component()
}

// Filepath implements StaticPage.
func (h HomePage) Filepath() string {
	return h.Path
}

// Layout implements StaticPage.
func (h HomePage) Layout() StaticLayout {
	return h.StaticLayout
}
