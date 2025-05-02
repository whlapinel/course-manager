package markdownrenderer

import (
	"bytes"
	"gh_static_portfolio/internal/core/filesystem"
	"log"

	"github.com/yuin/goldmark"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"

	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type goldmarkRenderer struct {
	goldmark.Markdown
}

func New() filesystem.MarkdownRenderer {
	// create goldmark.Markdown
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
			),
		)),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	return &goldmarkRenderer{
		Markdown: md,
	}
}

// ToHTML implements file.MarkdownRenderer.
func (g *goldmarkRenderer) ToHTML(markdown []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := g.Markdown.Convert(markdown, &buf); err != nil {
		log.Fatalf("Failed to convert Markdown: %v", err)
	}
	return buf.Bytes(), nil
}
