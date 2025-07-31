package ports

type SiteGenerator interface {
	StaticSiteURL(lastName string, courseID int) string
	Build(user, term, course Node) error
	BreadCrumbsMaker
}
type BreadCrumbsMaker interface {
	BreadCrumbs(url string) BreadCrumbsPartial
}
type BreadCrumbsPartial struct {
	Items []BreadCrumbsItem `json:"items" toml:"items"`
}

type BreadCrumbsItem struct {
	URL string `json:"url" toml:"url"`
}
