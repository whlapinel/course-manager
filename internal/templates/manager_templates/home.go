package managertemplates

import "github.com/a-h/templ"

type HomePage struct {
	ListTermsURL string
}

func (page HomePage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: "Home",
	}
}

func (page HomePage) Component() templ.Component {
	return HomePageComponent(page)
}
