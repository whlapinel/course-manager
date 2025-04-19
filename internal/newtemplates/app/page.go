package managertemplates

import components "gh_static_portfolio/internal/templates/components/base"

type Page interface {
	ComponentData
	PageLayout() components.PageLayout
	BreadCrumbs() BreadCrumbs
}
