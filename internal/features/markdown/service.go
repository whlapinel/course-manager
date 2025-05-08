package markdown

import (
	"bytes"
	"log"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type Service struct {
	goldmark.Markdown
}

func NewService() *Service {
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

	return &Service{
		Markdown: md,
	}
}

// ToHTML implements file.MarkdownRenderer.
func (g *Service) ToHTML(markdown []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := g.Markdown.Convert(markdown, &buf); err != nil {
		log.Fatalf("Failed to convert Markdown: %v", err)
	}
	return buf.Bytes(), nil
}
