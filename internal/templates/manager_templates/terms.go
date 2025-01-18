package managertemplates

import (
	"gh_static_portfolio/internal/domain"

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

func (page AddNonInstructDayPage) Component() templ.Component {
	return AddNonInstructDayComponent(page)
}
