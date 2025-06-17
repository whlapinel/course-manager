package calendarviews

import (
	ac "gh_static_portfolio/internal/app/components"
	"gh_static_portfolio/internal/app/dto"
	cmp "gh_static_portfolio/internal/basecomponents"
	"gh_static_portfolio/internal/core/occasion"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/util"
	"gh_static_portfolio/internal/shared/web"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
)

type TermCalendar struct {
	Params              routes.NodePath
	Term                dto.Term
	ListTermsURL        string
	TermDetailsURL      string
	CreateOccasionURL   string
	GetEditOccasionURL  web.AddParams
	PostEditOccasionURL web.AddParams
	DeleteOccasionURL   web.AddParams
	CalendarDates       DatesMap
	BreadCrumbsData     ac.BreadCrumbs
	ac.CourseManagerLayout
}

func (p TermCalendar) HTMXResponse() templ.Component {
	return p.Component()
}

func (p TermCalendar) NonHTMXResponse() templ.Component {
	return p.CourseManagerLayout.WithPage(p.Component())
}

func (data TermCalendar) GetCalendarDates() DatesMap {
	return data.CalendarDates
}

func (term TermCalendar) GetTerm() dto.Term {
	return term.Term
}

func (page TermCalendar) Occasions(date time.Time) []occasion.Occasion {
	var occasions []occasion.Occasion
	for _, occ := range page.Term.Occasions {
		if util.IsSameDate(occ.Date, date) {
			occasions = append(occasions, occ)
		}
	}
	return occasions
}

func (page TermCalendar) OccasionEditor(occasion occasion.Occasion) templ.Component {
	return OccasionEditor{
		Occasion:            occasion,
		IsEditing:           false,
		GetEditOccasionURL:  page.GetEditOccasionURL(occasion.ID),
		PostEditOccasionURL: page.PostEditOccasionURL(occasion.ID),
		DeleteOccasionURL:   page.DeleteOccasionURL(occasion.ID),
	}.Component()

}

func (data TermCalendar) AddOccasionButton(date time.Time) templ.Component {
	return AddOccasionButton{
		Date:              date,
		CreateOccasionURL: data.CreateOccasionURL,
		FormID:            "form-" + date.Format(time.DateOnly),
	}.Component()
}

func (data TermCalendar) Component() templ.Component {
	return TermCalendarComponent(data)
}

func (data TermCalendar) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: data.Term.Name + " Calendar",
		UpNav: cmp.UpNav{
			URL:  data.ListTermsURL,
			Text: "Back to Terms",
		},
		Crumbs: data.BreadCrumbs().BreadCrumbs(),
	}

}

func (data TermCalendar) BreadCrumbs() ac.BreadCrumbs {
	return data.BreadCrumbsData
}

type OccasionEditor struct {
	Occasion            occasion.Occasion
	IsEditing           bool
	GetEditOccasionURL  string
	PostEditOccasionURL string
	DeleteOccasionURL   string
}

func (data OccasionEditor) ComponentID() string {
	var id string
	id = "occasion " + strconv.Itoa(data.Occasion.ID) + " editor"
	id = strings.ToLower(id)
	id = strings.ReplaceAll(id, " ", "-")
	return id
}

func (data OccasionEditor) Component() templ.Component {
	return OccasionEditorComponent(data)
}
