package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/manager_templates/components"
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
	EditField         string
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
	return NewNodeDetailsComponent(page)
}

func (page NodeDetailsPage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: page.PageTitle(),
		UpNav: UpNav{
			URL:  page.UpNavURL,
			Text: page.upNavText(),
		},
	}
}

func (page NodeDetailsPage) CalendarButton() templ.Component {
	return HXButton{
		Text:     "Calendar",
		Method:   HxGet,
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
		parentPageText = "Home"
		return fmt.Sprintf("Up to %s", parentPageText)
	}
	return fmt.Sprintf("Up to %ss", parentPageText)
}

func (page NodeDetailsPage) ListChildrenButton() templ.Component {
	return HXButton{
		Text:     fmt.Sprintf("%ss", page.Node.ChildTypeName()),
		Method:   HxGet,
		URL:      page.ListChildrenURL,
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

func (page NodeDetailsPage) ViewFilesButton() templ.Component {
	return HXButton{
		Method:   HxGet,
		URL:      page.ServerFilesURL,
		Text:     "Files",
		HxTarget: "#page",
		PushURL:  true,
	}.Component()
}

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
