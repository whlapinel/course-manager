package ports

type MarkdownFile struct {
	Path string
	*FrontMatter
	Content string
}

type FrontMatterReadWriter interface {
	ParseFrontMatter(data []byte) (MarkdownFile, error)
	ToBytes(file MarkdownFile) ([]byte, error)
}

type FrontMatter struct {
	Title       string   `toml:"title,omitempty"`
	Type        string   `toml:"type,omitempty"`
	Date        string   `toml:"date,omitempty"`
	Draft       bool     `toml:"draft,omitempty"`
	Tags        []string `toml:"tags,omitempty"`
	Description string   `toml:"description,omitempty"`
	Slug        string   `toml:"slug,omitempty"`
	URL         string   `toml:"url,omitempty"`
	// Add other fields as needed
	Params map[string]any `toml:"params,inline"` // Catch any other fields
}
