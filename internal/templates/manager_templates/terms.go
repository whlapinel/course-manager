package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	cmp "gh_static_portfolio/internal/templates/components"
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

type TermsListPage struct {
	ShowTermCalendarRHN string
	NodeListPage
}

func (page TermsListPage) Component() templ.Component {
	var calendarButtons []ComponentData
	for _, term := range page.Children {
		log.Println("term ID: ", term.GetID())
		button := ShowCalendarButton{
			ShowCalendarURL: page.E.Reverse(page.ShowTermCalendarRHN, page.NodeListPage.ParentNode.GetID().(string), term.GetID()),
		}
		calendarButtons = append(calendarButtons, button)

	}
	page.ChildUI = append(page.ChildUI, calendarButtons)
	return TermsListPageComponent(page)

}

type AddNonInstructDayPage struct {
	Term           domain.Term
	GetAddDayURL   string
	PostAddDayURL  string
	TermDetailsURL string
}

func (page AddNonInstructDayPage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Add Non-Instructional Day",
		UpNav: cmp.UpNav{
			URL:  page.TermDetailsURL,
			Text: "Term Details",
		},
		Crumbs: page.BreadCrumbs().BreadCrumbs(),
	}
}

func (page AddNonInstructDayPage) BreadCrumbs() BreadCrumbs {
	return BreadCrumbs{
		Term:           page.Term,
		TermDetailsURL: page.TermDetailsURL,
	}
}

func (page AddNonInstructDayPage) Component() templ.Component {
	return AddNonInstructDayComponent(page)
}
