package managertemplates

import "gh_static_portfolio/internal/templates/components"

type Page interface {
	ComponentData
	PageLayout() components.PageLayout
	BreadCrumbs() BreadCrumbs
}
