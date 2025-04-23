package managertemplates

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/shared/web"
	cmp "gh_static_portfolio/internal/templates/components/base"
	"log"

	"github.com/a-h/templ"
)

type TermDetailsPage struct {
	NodeDetailsPage
	ShowEditTermDatesURL string
}

func (page TermDetailsPage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: page.PageTitle(),
		UpNav: cmp.UpNav{
			URL:  page.UpNavURL,
			Text: "Back to Term Details",
		},
	}
}

func (page TermDetailsPage) Component() templ.Component {
	return TermDetailsComponent(page)
}

func (page TermDetailsPage) FullComponent(reverse web.Reverse, comp templ.Component, user dto.User) templ.Component {
	return BaseLayout(reverse, comp, user)
}

type TermsListPage struct {
	ShowTermCalendarURL web.AddParams
	NodeListPage
}

func (page TermsListPage) Component() templ.Component {
	var calendarButtons []ComponentData
	for _, term := range page.Children {
		log.Println("term ID: ", term.GetID())
		button := ShowCalendarButton{
			ShowCalendarURL: page.ShowTermCalendarURL(page.NodeListPage.ParentNode.GetID(), term.GetID()),
		}
		calendarButtons = append(calendarButtons, button)

	}
	page.ChildUI = append(page.ChildUI, calendarButtons)
	return TermsListPageComponent(page)

}

type AddNonInstructDayPage struct {
	Term           dto.Term
	GetAddDayURL   string
	PostAddDayURL  string
	TermDetailsURL string
	BreadCrumbs
}

func (page AddNonInstructDayPage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Add Non-Instructional Day",
		UpNav: cmp.UpNav{
			URL:  page.TermDetailsURL,
			Text: "Term Details",
		},
		Crumbs: page.BreadCrumbs.BreadCrumbs(),
	}
}

func (page AddNonInstructDayPage) Component() templ.Component {
	return AddNonInstructDayComponent(page)
}
