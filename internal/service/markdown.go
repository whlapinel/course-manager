package service

import (
	"bytes"
	"context"
	templates "gh_static_portfolio/internal/templates/static"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"golang.org/x/sync/errgroup"
)

func (svc CourseService) MarkdownToHTML(srcPath string) error {
	content, err := svc.RenderMarkdownFile(srcPath)
	if err != nil {
		return err
	}
	fileName := filepath.Base(srcPath)
	outputPath := filepath.Join(filepath.Dir(srcPath), strings.TrimSuffix(fileName, ".md")+".html")
	output, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create output file %s: %v", outputPath, err)
	}
	defer output.Close()

	log.Println("Writing:", outputPath)
	data := templates.MarkdownDocument{
		AssetsURL: StaticAssetsURL,
		Title:     "No title",
		Content:   string(content),
		Static:    true,
	}
	err = templates.DocLayout(data).Render(context.Background(), output)
	if err != nil {
		return err
	}
	return nil

}

// this will generate html from markdown file
func (svc CourseService) RenderMarkdownFile(path string) ([]byte, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", path, err)
	}

	var buf bytes.Buffer
	if err := svc.md.Convert(contents, &buf); err != nil {
		log.Fatalf("Failed to convert Markdown: %v", err)
	}
	return buf.Bytes(), nil
}

// this will generate html from all markdown files within the files directory
func (svc CourseService) renderMarkdownFiles(title, filesPath string) error {
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

	var markdownGroup errgroup.Group
	for _, entry := range mdFiles {
		log.Println("rendering markdown for:", entry.Name())

		inputPath := filepath.Join(filesPath, entry.Name())
		markdownGroup.Go(func() error {
			err := renderMarkdown(svc.md, entry, title, inputPath, filesPath)
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

func renderMarkdown(md goldmark.Markdown, entry os.DirEntry, title, inputPath, filesPath string) error {
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
	data := templates.MarkdownDocument{
		AssetsURL: StaticAssetsURL,
		Title:     title,
		Content:   buf.String(),
		Static:    true,
	}
	err = templates.DocLayout(data).Render(context.Background(), output)
	if err != nil {
		return err
	}
	return nil
}
