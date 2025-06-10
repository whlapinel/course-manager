package services

import (
	"gh_static_portfolio/internal/ports"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type MarkdownService struct {
	renderer    ports.MarkdownRenderer
	paths       ports.PathingService
	files       ports.FileRepository
	frontmatter ports.FrontMatterReadWriter
}

func NewMarkdownService(
	frontmatter ports.FrontMatterReadWriter,
	renderer ports.MarkdownRenderer,
	paths ports.PathingService,
	files ports.FileRepository,

) *MarkdownService {
	return &MarkdownService{
		frontmatter: frontmatter,
		renderer:    renderer,
		paths:       paths,
		files:       files,
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

// for create markdown file interface sent via post request
func (svc *MarkdownService) Create(data []byte, relPath string, nodes ports.Nodes) error {
	root := svc.paths.NodeFilesDirPath(nodes.ToSlice()...)
	file, err := svc.setFrontMatter(data, root, relPath)
	if err != nil {
		return err
	}
	data, err = svc.frontmatter.ToBytes(file)
	if err != nil {
		return err
	}
	err = svc.files.Save(data, root, relPath)
	if err != nil {
		return err
	}
	return nil

}

// for upload markdown file
func (svc *MarkdownService) Save(header *multipart.FileHeader, relPath string, nodes ports.Nodes) error {
	file, err := header.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	root := svc.paths.NodeFilesDirPath(nodes.ToSlice()...)
	mdFile, err := svc.setFrontMatter(data, root, relPath)
	if err != nil {
		return err
	}
	data, err = svc.frontmatter.ToBytes(mdFile)
	if err != nil {
		return err
	}
	err = svc.files.Save(data, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *MarkdownService) Update(data []byte, relPath string, nodes ports.Nodes) error {
	root := svc.paths.NodeFilesDirPath(nodes.ToSlice()...)
	file, err := svc.setFrontMatter(data, root, relPath)
	if err != nil {
		return err
	}
	data, err = svc.frontmatter.ToBytes(file)
	if err != nil {
		return err
	}
	err = svc.files.Update(data, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *MarkdownService)helperComments()string{
	return `
	# hi
	`
}

func (svc *MarkdownService) setFrontMatter(data []byte, root, relPath string) (ports.MarkdownFile, error) {
	file, err := svc.frontmatter.ParseFrontMatter(data)
	if err != nil {
		return file, err
	}
	path := filepath.Join(root, relPath)
	file.Path = path
	if file.FrontMatter == nil {
		file.FrontMatter = &ports.FrontMatter{}
	}
	segments := strings.Split(path, "/")
	url := filepath.Join(segments[8:]...)
	url = strings.ReplaceAll(url, "_", "-")
	file.FrontMatter.URL = strings.TrimSuffix(url, ".md")
	rootDirSegments := strings.Split(root, "/")
	parentPath := filepath.Join(rootDirSegments[8:]...)
	file.FrontMatter.Type = "standalone"
	if file.FrontMatter.Params == nil {
		file.FrontMatter.Params = make(map[string]any)
	}
	file.FrontMatter.Params["parentPath"] = parentPath
	return file, nil
}
