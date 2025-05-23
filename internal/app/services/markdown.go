package services

import "gh_static_portfolio/internal/ports"

type MarkdownService struct {
	renderer ports.MarkdownRenderer
	paths    ports.PathingService
	files    ports.FileRepository
}

func NewMarkdownService(
	renderer ports.MarkdownRenderer,
	paths ports.PathingService,
	files ports.FileRepository,

) *MarkdownService {
	return &MarkdownService{
		renderer: renderer,
		paths:    paths,
		files:    files,
	}
}

func (s *MarkdownService) ViewMarkdown(relpath string, nodes ...ports.Node) (string, error) {
	root := s.paths.NodeFilesDirPath(nodes...)
	content, err := s.files.Read(root, relpath)
	if err != nil {
		return "", err
	}
	html, err := s.renderer.Render(content)
	if err != nil {
		return "", err
	}
	return html, nil
}
