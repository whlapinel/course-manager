package appcomponents

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/base"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/util"

	"github.com/a-h/templ"
	"github.com/labstack/gommon/log"
)

type NodeCreatePage struct {
	ParentNode        ports.Node
	NodeType          dto.NodeTypeName
	Params            routes.NodePath
	PostCreateNodeURL string
	CancelURL         string
	BreadCrumbsData   BreadCrumbs
	CourseManagerLayout
}

func (p NodeCreatePage) HTMXResponse() templ.Component {
	return p.Component()
}

func (p NodeCreatePage) NonHTMXResponse() templ.Component {
	return p.Component2(p.Component())
}

func (page NodeCreatePage) Component() templ.Component {
	return NodeCreatePageComponent(page)
}

func (page NodeCreatePage) FormComponent() templ.Component {
	title := fmt.Sprintf("New %s", page.NodeType)
	subtitle := fmt.Sprintf("Enter %s details here", page.NodeType)
	form := cmp.NewForm(cmp.NewFormParams{
		Title:     title,
		Subtitle:  subtitle,
		PostURL:   page.PostCreateNodeURL,
		CancelURL: page.CancelURL,
		HxTarget:  "#page",
	})
	nameInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
		Name: "Name",
		Type: cmp.Text,
	})
	descTextArea := cmp.NewTextAreaWithLabel(cmp.TextAreaWithLabelParams{
		Name: "Description",
	})
	switch page.ParentNode.(type) {
	case dto.Course, dto.Unit: // for create unit and create lesson page
		numInput := cmp.NewInputWithLabel(cmp.InputWithLabelParams{
			Name: "Number",
			Type: cmp.Number,
		})
		form = form.AddElement(numInput)
		form = form.AddElement(nameInput, descTextArea)
	case dto.User: // for create term page
		form = form.AddElement(nameInput, descTextArea)
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

func (page NodeCreatePage) PageLayout() cmp.PageLayout {
	if page.ParentNode == nil {
		log.Errorf("page.ParentNode is nil! params: %v", page.Params.ToSlice()...)
	}
	return cmp.PageLayout{
		PageTitle: fmt.Sprintf("New %s for %s", page.NodeType.String(), page.ParentNode.GetName()),
		UpNav: cmp.UpNav{
			URL:  page.CancelURL,
			Text: "Cancel",
		},
		Crumbs: page.BreadCrumbs().BreadCrumbs(),
	}
}

func (page NodeCreatePage) BreadCrumbs() BreadCrumbs {
	return page.BreadCrumbsData

}

func NodeCreateFormID(nodeType dto.NodeTypeName) string {
	return util.KebabCase(nodeType.String()) + "-form"
}
