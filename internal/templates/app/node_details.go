package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components/base"
	"strconv"

	"github.com/a-h/templ"
)

type NodeDetailsPage struct {
	Params            domain.NodePath
	ParentNode        domain.CourseNode
	Node              domain.CourseNode
	GetEditNodeURL    string
	PostEditNodeURL   string
	ListChildrenURL   string
	UpNavURL          string
	CancelEditURL     string
	IsEdit            bool
	NodeImageURL      func() string
	BreadCrumbsData   BreadCrumbs
	CourseCalendarURL string
	ServerFilesURL    string
}

func (page NodeDetailsPage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData
}

func (page NodeDetailsPage) Component() templ.Component {
	return NodeDetailsComponent(page)
}

func (page NodeDetailsPage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: page.PageTitle(),
		UpNav: cmp.UpNav{
			URL:  page.UpNavURL,
			Text: page.upNavText(),
		},
		Crumbs: page.BreadCrumbs().BreadCrumbs(),
	}
}

func (page NodeDetailsPage) CalendarButton() templ.Component {
	return cmp.Button{
		Text:     "📅",
		Method:   cmp.HxGet,
		URL:      page.CourseCalendarURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeDetailsPage) PageTitle() string {
	desig := page.Node.Designation()
	if desig != "" {
		return fmt.Sprintf("%s (%s)", desig, page.Node.GetName())
	}
	return page.Node.GetName()

}

func (page NodeDetailsPage) upNavText() string {
	parentPageText := page.Node.TypeName()
	if page.Node.ParentTypeName() == string(domain.RootTypeName) {
		parentPageText = "🏠"
		return fmt.Sprintf("Up to %s", parentPageText)
	}
	return fmt.Sprintf("Up to %ss", parentPageText)
}

func (page NodeDetailsPage) ListChildrenButton() templ.Component {
	return cmp.Button{
		Text:     fmt.Sprintf("%ss", page.Node.ChildTypeName()),
		Method:   cmp.HxGet,
		URL:      page.ListChildrenURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeDetailsPage) ViewFilesButton() templ.Component {
	return cmp.Button{
		Method:   cmp.HxGet,
		URL:      page.ServerFilesURL,
		Text:     "📂",
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeDetailsPage) DetailsFormComponent(editing bool) templ.Component {
	if editing {
		var comps []cmp.Component
		numItem := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
			Name:  "Number",
			Type:  cmp.Number,
			Value: strconv.Itoa(page.Node.GetNumber()),
		})
		comps = append(comps, numItem)
		nameItem := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
			Name:  "Name",
			Type:  cmp.Text,
			Value: page.Node.GetName(),
		})
		comps = append(comps, nameItem)
		descriptionItem := cmp.NewTextAreaWithLabel(cmp.TextAreaWithLabelParams{
			Name:  "Description",
			Value: page.Node.GetDescription(),
		})
		comps = append(comps, descriptionItem)
		form := cmp.NewForm(cmp.NewFormParams{
			PostURL:   page.PostEditNodeURL,
			CancelURL: page.CancelEditURL,
			HxTarget:  "#page",
		})
		form = form.AddElement(comps...)
		return form.Component()
	} else {
		var numItem, nameItem, descriptionItem cmp.EditableInfoItem
		var components []cmp.EditableInfoItem
		if page.Node.GetNumber() >= 0 {
			numItem = cmp.EditableInfoItem{
				Element: cmp.Element{ID: cmp.Kebab("Number")},
				Field:   "Number",
				Value:   strconv.Itoa(page.Node.GetNumber()),
			}
			components = append(components, numItem)
		}
		nameItem = cmp.EditableInfoItem{
			Element: cmp.Element{ID: cmp.Kebab("Name")},
			Field:   "Name",
			Value:   page.Node.GetName(),
		}
		components = append(components, nameItem)
		descriptionItem = cmp.EditableInfoItem{
			Element: cmp.Element{ID: cmp.Kebab("Description")},
			Field:   "Description",
			Value:   page.Node.GetDescription(),
		}
		components = append(components, descriptionItem)
		info := cmp.EditableInfo{
			Element: cmp.Element{
				ID: cmp.Kebab("lesson-form"),
			},
			GetEditURL: page.GetEditNodeURL,
			Components: components,
		}

		return info.Component()
	}

}

// newer version, not working, needs to be replaced
func (page NodeDetailsPage) FieldsComponent(editField string) templ.Component {
	list := cmp.NewDescriptionList(cmp.NewDescriptionListParams{
		Title:       page.PageTitle(),
		GetEditURL:  page.GetEditNodeURL,
		PostEditURL: page.PostEditNodeURL,
	})
	var numItem, nameItem, descriptionItem cmp.DescriptionListItem
	numItem = cmp.NewDescriptionListItem("Number", strconv.Itoa(page.Node.GetNumber()), true, false)
	nameItem = cmp.NewDescriptionListItem("Name", page.Node.GetName(), true, false)
	descriptionItem = cmp.NewDescriptionListItem("Description", page.Node.GetDescription(), true, false)
	switch editField {
	case "number":
		numItem.IsEditing = true
	case "name":
		nameItem.IsEditing = true
	case "description":
		descriptionItem.IsEditing = true
	}
	list = list.AddItems(
		numItem,
		nameItem,
		descriptionItem,
	)
	return list.Component()
}
