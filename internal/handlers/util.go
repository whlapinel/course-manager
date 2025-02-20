package handlers

import (
	"bytes"
	"log"
	"net/url"
	"os"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"

	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// AddQueryParams adds arbitrary query parameters to a given URL
func AddQueryParams(baseURL string, params map[string]string) (string, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// Get the existing query params
	query := parsedURL.Query()

	// Append new params
	for key, value := range params {
		if value != "" {
			query.Set(key, value)
		}
	}
	// Encode and set the updated query string
	parsedURL.RawQuery = query.Encode()
	return parsedURL.String(), nil
}

// this will generate html from markdown file
func RenderMarkdownFile(path string) ([]byte, error) {
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
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", path, err)
	}

	var buf bytes.Buffer
	if err := md.Convert(contents, &buf); err != nil {
		log.Fatalf("Failed to convert Markdown: %v", err)
	}
	return buf.Bytes(), nil
}
