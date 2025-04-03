package components

import "github.com/a-h/templ"

type Component interface {
	Component() templ.Component
}

type Form struct {
	Element
	Title      string
	Subtitle   string
	Components []Component // should only be input elements
	PostURL    string
	CancelURL  string
	HxTarget   string
	Editing    bool
}

type EditableInfo struct {
	Element
	Title      string
	Subtitle   string
	GetEditURL string
	Components []EditableInfoItem // should only be text elements, not input
}

func (el EditableInfo) Component() templ.Component {
	return EditableInfoComponent(el)
}

type EditableInfoItem struct {
	Element
	Field string
	Value string
	
}

func (el EditableInfoItem) Component() templ.Component {
	return EditableInfoItemComponent(el)
}

type NewFormParams struct {
	Title     string
	Subtitle  string
	PostURL   string
	CancelURL string
	HxTarget  string
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
