package appcomponents

import (
	cmp "gh_static_portfolio/internal/basecomponents"
	"gh_static_portfolio/internal/ports"

	"github.com/a-h/templ"
)

type BreadCrumbs struct {
	ports.Nodes
	UserDetailsURL   string
	TermDetailsURL   string
	CourseDetailsURL string
	UnitDetailsURL   string
	LessonDetailsURL string
}

func (data BreadCrumbs) Component() templ.Component {
	return NewBreadCrumbsComponent(data)
}

func (data BreadCrumbs) BreadCrumbs() cmp.BreadCrumbs {
	var items []cmp.BreadCrumbsItem
	if data.User == nil {
		return cmp.BreadCrumbs{}
	}
	if data.User.GetID() != "" {
		item := cmp.BreadCrumbsItem{
			NavItem: cmp.NavItem{
				Text:   data.User.GetName(),
				URL:    data.UserDetailsURL,
				Method: cmp.HxGet,
			},
		}
		items = append(items, item)

		if data.Term != nil && data.Term.GetID() != 0 {
			item := cmp.BreadCrumbsItem{
				NavItem: cmp.NavItem{
					Text:   data.Term.GetName(),
					URL:    data.TermDetailsURL,
					Method: cmp.HxGet,
				},
			}
			items = append(items, item)

			if data.Course != nil && data.Course.GetID() != 0 {
				item := cmp.BreadCrumbsItem{
					NavItem: cmp.NavItem{
						Text:   data.Course.GetName(),
						URL:    data.CourseDetailsURL,
						Method: cmp.HxGet,
					},
				}
				items = append(items, item)

				if data.Unit != nil && data.Unit.GetID() != 0 {
					item := cmp.BreadCrumbsItem{
						NavItem: cmp.NavItem{
							Text:   data.Unit.Designation(),
							URL:    data.UnitDetailsURL,
							Method: cmp.HxGet,
						},
					}
					items = append(items, item)

					if data.Lesson != nil && data.Lesson.GetID() != 0 {
						item := cmp.BreadCrumbsItem{
							NavItem: cmp.NavItem{
								Text:   data.Lesson.Designation(),
								URL:    data.LessonDetailsURL,
								Method: cmp.HxGet,
							},
						}
						items = append(items, item)

					}
				}
			}
		}
	}
	bc := cmp.BreadCrumbs{
		Items: items,
	}
	return bc
}

func NewBreadCrumbsComponent(data BreadCrumbs) templ.Component {
	bc := data.BreadCrumbs()
	return bc.Component()
}
