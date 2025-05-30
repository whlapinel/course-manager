package basecomponents

import "github.com/a-h/templ"

type TabSet struct {
	Tabs          []TabLink
	AssetsURLFunc func(string) string
}

func (data TabSet) Component() templ.Component {
	return TabSetComponent(data)
}

type TabLink struct {
	Name string
	URL  string
}

func (data TabLink) Component() templ.Component {
	return TabLinkComponent(data)
}
