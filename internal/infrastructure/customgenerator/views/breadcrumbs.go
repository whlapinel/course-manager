package templates

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/shared/node"
	components "gh_static_portfolio/internal/newtemplates/components/base"
)

func BreadCrumbs(nodes node.Nodes, urlFunc func(...node.Node) string) components.BreadCrumbs {
	var items []components.BreadCrumbsItem
	nodeSlice := nodes.ToSlice()
	for i, node := range nodeSlice {
		if user, ok := node.(dto.User); ok {
			item := components.BreadCrumbsItem{
				NavItem: components.NavItem{
					Text: user.Username(),
					URL:  "/",
				},
			}
			items = append(items, item)
		} else {
			if node.GetID().(int) != 0 {
				if i > len(nodeSlice) {
					break
				}
				item := components.BreadCrumbsItem{
					NavItem: components.NavItem{
						Text: node.GetName(),
						URL:  urlFunc(nodeSlice[:i+1]...),
					},
				}
				items = append(items, item)
			} else {
				break
			}
		}
	}
	return components.BreadCrumbs{
		Items: items,
	}

}
