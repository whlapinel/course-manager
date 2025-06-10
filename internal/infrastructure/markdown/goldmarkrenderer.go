package markdown

import (
	"bytes"
	"gh_static_portfolio/internal/ports"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
)

type goldmarkRenderer struct {
	goldmark.Markdown
}

func New() ports.MarkdownRenderer {
	// create goldmark.Markdown
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			// this prevents front matter from being rendered along with content in app
			&frontmatter.Extender{
				Mode: frontmatter.SetMetadata,
			},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	return goldmarkRenderer{
		Markdown: md,
	}
}

// Render implements ports.MarkdownRenderer.
func (g goldmarkRenderer) Render(content []byte) (string, error) {
	var buf bytes.Buffer
	err := g.Markdown.Convert(content, &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
