package ports

type SiteGenerator interface {
	StaticSiteURL(userID string) string
	Build(userID string, termID int) error
}
