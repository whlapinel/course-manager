package templates

type MarkdownDocument struct {
	AssetsURL func(relPath string) string
	Title     string
	Content   string
	Static    bool
}
