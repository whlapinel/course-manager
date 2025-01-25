package managertemplates

type Page interface {
	ComponentData
	PageLayout() PageLayout
	BreadCrumbs() BreadCrumbs
}

type UpNav struct {
	URL  string
	Text string
}

type PageLayout struct {
	PageTitle string
	UpNav     UpNav
}
