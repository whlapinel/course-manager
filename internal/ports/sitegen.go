package ports

type SiteGenerator interface {
	Build() error
}
