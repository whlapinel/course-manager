package ports

type SiteGenerator interface {
	Build(userID string, termID int) error
}
