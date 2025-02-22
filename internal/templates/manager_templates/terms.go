package managertemplates

import (
	"gh_static_portfolio/internal/domain"
	"log"

	"github.com/a-h/templ"
)

type TermDetailsPage struct {
	NodeDetailsPage
	ShowEditTermDatesURL string
}

func (page TermDetailsPage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: page.PageTitle(),
		UpNav: UpNav{
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
	log.Println("TermsListPage.Component(): RHN, userID, termID: ", page.ShowTermCalendarRHN, page.NodeListPage.ParentNode.GetID().(string), page.Children[0].GetID())
	log.Println("TermsListPage.Component(): ShowCalendarURL: ", page.E.Reverse(page.ShowTermCalendarRHN, page.NodeListPage.ParentNode.GetID().(string), page.Children[0].GetID()))

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

func (page AddNonInstructDayPage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: "Add Non-Instructional Day",
		UpNav: UpNav{
			URL:  page.TermDetailsURL,
			Text: "Term Details",
		},
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
