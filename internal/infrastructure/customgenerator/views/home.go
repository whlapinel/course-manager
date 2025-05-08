package templates

import (
	"gh_static_portfolio/internal/app/dto"
	components "gh_static_portfolio/internal/newtemplates/components/base"

	"github.com/a-h/templ"
)

type StaticHomePageParams struct {
	HomePage
}

func NewHomePage(params StaticHomePageParams) HomePage {
	return HomePage{
		Path: params.Path,
		PageData: PageData{
			User:      params.User,
			AssetsURL: params.AssetsURL,
		},
		Term:    params.Term,
		TermURL: params.TermURL,
	}
}

type HomePage struct {
	PageData
	Path    string
	Term    dto.Term
	TermURL string
}

// Component implements StaticPage.
func (data HomePage) Component() templ.Component {
	return Layout{
		Head: Head{
			User:      data.User,
			AssetsURL: data.AssetsURL,
		}.Component(),
		Page: HomeComponent(data),
	}.Component()
}

func (h HomePage) ViewTermLink() templ.Component {
	return components.Link{
		Text: "View " + h.Term.Name,
		URL:  h.TermURL,
	}.Component()
}

// Filepath implements StaticPage.
func (h HomePage) Filepath() string {
	return h.Path
}

// Layout implements StaticPage.
func (h HomePage) Layout() PageData {
	return h.PageData
}
