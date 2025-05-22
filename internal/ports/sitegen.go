package ports

type SiteGenerator interface {
	BaseURL(domain string) func(userID string) string
	Build(userID string, termID int) error
}
