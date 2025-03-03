package components

import "github.com/a-h/templ"

type Component interface {
	Component() templ.Component
}

type Form struct {
	Element
	Title      string
	Subtitle   string
	Components []Component
	PostURL    string
	CancelURL  string
	HxTarget   string
	Editing    bool
}

type NewFormParams struct {
	Title     string
	Subtitle  string
	PostURL   string
	CancelURL string
	HxTarget  string
	Editing   bool
}

func NewForm(params NewFormParams) Form {
	return Form{
		Element: Element{
			ID: Kebab(params.Title),
		},
		Title:     params.Title,
		Subtitle:  params.Subtitle,
		CancelURL: params.CancelURL,
		PostURL:   params.PostURL,
		HxTarget:  params.HxTarget,
		Editing:   params.Editing,
	}
}

func (el Form) AddElement(newComps ...Component) Form {
	components := append(el.Components, newComps...)
	el.Components = components
	return el
}

func (el Form) Component() templ.Component {
	return FormComponent(el)
}
