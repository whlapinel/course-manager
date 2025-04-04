package sitegenerator

import (
	"bytes"
	"context"
	mt "gh_static_portfolio/internal/templates/app"
	"log"
	"os"
	"path/filepath"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"golang.org/x/sync/errgroup"
)

// this will generate html from all markdown files within the files directory
func RenderMarkdownFiles(title, filesPath string) error {
	dirEntries, err := os.ReadDir(filesPath)
	if os.IsNotExist(err) {
		return err
	}
	if err != nil {
		log.Fatal(err)
	}

	// Filter out non-Markdown files
	var mdFiles []os.DirEntry
	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ".md" {
			mdFiles = append(mdFiles, entry)
		}
	}

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
	var markdownGroup errgroup.Group
	for _, entry := range mdFiles {

		inputPath := filepath.Join(filesPath, entry.Name())
		markdownGroup.Go(func() error {
			// Read entire file at once
			contents, err := os.ReadFile(inputPath)
			if err != nil {
				log.Fatalf("Failed to read file %s: %v", inputPath, err)
			}

			var buf bytes.Buffer
			if err := md.Convert(contents, &buf); err != nil {
				log.Fatalf("Failed to convert Markdown: %v", err)
			}

			outputPath := filepath.Join(filesPath, strings.TrimSuffix(entry.Name(), ".md")+".html")
			output, err := os.Create(outputPath)
			if err != nil {
				log.Fatalf("Failed to create output file %s: %v", outputPath, err)
			}
			defer output.Close()

			log.Println("Writing:", outputPath)
			data := mt.MarkdownDocument{
				Title:   title,
				Content: buf.String(),
				Static:  true,
			}
			err = mt.DocLayout(data).Render(context.Background(), output)
			if err != nil {
				return err
			}
			return nil
		})

	}
	if err := markdownGroup.Wait(); err != nil {
		return err
	}
	return nil
}

func RenderMarkdown(md goldmark.Markdown, entry os.DirEntry, title, inputPath, filesPath string) error {
	// Read entire file at once
	contents, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", inputPath, err)
	}

	var buf bytes.Buffer
	if err := md.Convert(contents, &buf); err != nil {
		log.Fatalf("Failed to convert Markdown: %v", err)
	}

	outputPath := filepath.Join(filesPath, strings.TrimSuffix(entry.Name(), ".md")+".html")
	output, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create output file %s: %v", outputPath, err)
	}
	defer output.Close()

	log.Println("Writing:", outputPath)
	data := mt.MarkdownDocument{
		Title:   title,
		Content: buf.String(),
		Static:  true,
	}
	err = mt.DocLayout(data).Render(context.Background(), output)
	if err != nil {
		return err
	}
	return nil
}
