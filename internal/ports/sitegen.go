package ports

type SiteGenerator interface {
	StaticSiteURL(lastName string, courseID int) string
	Build(user, term, course Node) error
}
