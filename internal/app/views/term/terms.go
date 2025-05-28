package termviews

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/base"
	"gh_static_portfolio/internal/shared/web"
	"log"

	"github.com/a-h/templ"
)

type TermDetailsPage struct {
	ac.NodeDetailsPage
	ShowEditTermDatesURL string
	ac.CourseManagerLayout
}

func (page TermDetailsPage) HTMXResponse() templ.Component {
	return page.Component()
}

func (page TermDetailsPage) NonHTMXResponse() templ.Component {
	return page.CourseManagerLayout.Component2(page.Component())
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
	ShowTermCalendarURL web.AddParams
	ac.NodeListPage
	ac.CourseManagerLayout
}

func (page TermsListPage) Component() templ.Component {
	var calendarButtons []ac.ComponentData
	for _, term := range page.Children {
		log.Println("term ID: ", term.GetID())
		button := ShowCalendarButton{
			ShowCalendarURL: page.ShowTermCalendarURL(term.GetID()),
		}
		calendarButtons = append(calendarButtons, button)

	}
	page.ChildUI = append(page.ChildUI, calendarButtons)
	return TermsListPageComponent(page)

}

func (page TermsListPage) HTMXResponse() templ.Component {
	return page.Component()
}

func (page TermsListPage) NonHTMXResponse() templ.Component {
	return page.CourseManagerLayout.Component2(page.Component())
}

type AddNonInstructDayPage struct {
	Term           dto.Term
	GetAddDayURL   string
	PostAddDayURL  string
	TermDetailsURL string
	ac.BreadCrumbs
	ac.CourseManagerLayout
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

func (page AddNonInstructDayPage) HTMXResponse() templ.Component {
	return page.Component()
}

func (page AddNonInstructDayPage) NonHTMXResponse() templ.Component {
	return page.CourseManagerLayout.Component()
}

type ShowCalendarButton struct {
	ShowCalendarURL string
}

func (data ShowCalendarButton) Component() templ.Component {
	button := cmp.Button{
		Text:     "Calendar",
		Method:   cmp.HxGet,
		URL:      data.ShowCalendarURL,
		HxTarget: "#page",
		PushURL:  true,
	}
	return button.Component()
}
