package managertemplates

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	tpl "gh_static_portfolio/internal/templates"

	cmp "gh_static_portfolio/internal/templates/components"

	"github.com/a-h/templ"
)

type NodeCreatePage struct {
	ParentNode        domain.CourseNode
	NodeType          domain.NodeTypeName
	Params            domain.NodePath
	PostCreateNodeURL string
	CancelURL         string
	BreadCrumbsData   BreadCrumbs
}

func (page NodeCreatePage) Component() templ.Component {
	return NodeCreatePageComponent(page)
}

func (page NodeCreatePage) FormComponent() templ.Component {
	return CreateNodeFormComponent(page)
}

func (page NodeCreatePage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: fmt.Sprintf("New %s for %s", page.NodeType.String(), page.ParentNode.GetName()),
		UpNav: UpNav{
			URL:  page.CancelURL,
			Text: "Cancel",
		},
	}
}

func (page NodeCreatePage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData

}

func NodeCreateFormID(nodeType domain.NodeTypeName) string {
	return tpl.KebabCase(nodeType.String()) + "-form"
}

func CreateNodeFormComponent(page NodeCreatePage) templ.Component {
	title := fmt.Sprintf("New %s", page.NodeType)
	subtitle := fmt.Sprintf("Enter %s details here", page.NodeType)
	form := cmp.NewForm(cmp.NewFormParams{
		Title:     title,
		Subtitle:  subtitle,
		PostURL:   page.PostCreateNodeURL,
		CancelURL: page.CancelURL,
		HxTarget:  "#page",
	})
	_, userOk := page.ParentNode.(domain.User)
	_, courseOk := page.ParentNode.(domain.Course)
	_, unitOk := page.ParentNode.(domain.Unit)
	if courseOk || unitOk {
		numInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
			Name: "Number",
			Type: cmp.Number,
		})
		form = form.AddElement(numInput)
	}
	nameInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "Name",
		Type: cmp.Text,
	})
	descTextArea := cmp.NewTextAreaWithLabel(cmp.TextAreaWithLabelParams{
		Name: "Description",
	})
	form = form.AddElement(nameInput, descTextArea)
	if userOk {
		startInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
			Name: "Start Date",
			Type: cmp.Date,
		})
		endInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
			Name: "End Date",
			Type: cmp.Date,
		})
		form = form.AddElement(startInput, endInput)
	}
	return form.Component()
}
