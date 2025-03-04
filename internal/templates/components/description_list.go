package components

import (
	"github.com/a-h/templ"
)

type NewDescriptionListParams struct {
	Title       string
	Subtitle    string
	GetEditURL  string
	PostEditURL string
}

func NewDescriptionList(params NewDescriptionListParams) DescriptionList {
	return DescriptionList{
		Title:       params.Title,
		Subtitle:    params.Subtitle,
		GetEditURL:  params.GetEditURL,
		PostEditURL: params.PostEditURL,
	}
}

func (el DescriptionList) AddItems(items ...DescriptionListItem) DescriptionList {
	el.Items = append(el.Items, items...)
	return el
}

func NewDescriptionListItem(name, value string, editable, isEditing bool) DescriptionListItem {
	return DescriptionListItem{
		Element: Element{
			ID: Kebab(name),
		},
		Name:      name,
		Value:     value,
		Editable:  editable,
		IsEditing: isEditing,
	}
}

type DescriptionList struct {
	Title, Subtitle string
	GetEditURL      string
	PostEditURL     string
	Items           []DescriptionListItem
}

type DescriptionListItem struct {
	Element
	Name      string
	Editable  bool
	IsEditing bool
	Value     string
}

func (el DescriptionList) Component() templ.Component {
	return DescriptionListComponent(el)
}
