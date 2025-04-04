package service

import (
	"gh_static_portfolio/internal/data"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
)

func NewCourseService(repo data.CourseRepo) CourseService {
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
	return CourseService{
		repo: repo,
		md:   md,
	}
}

type CourseService struct {
	md   goldmark.Markdown
	repo data.CourseRepo
}
