package ports

type MarkdownRenderer interface {
	Render(content []byte) (string, error)
}
