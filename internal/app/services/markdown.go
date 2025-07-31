package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/ports"
	"io"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"strings"
)

type MarkdownService struct {
	renderer      ports.MarkdownRenderer
	paths         ports.PathingService
	files         ports.FileRepository
	frontmatter   ports.FrontMatterReadWriter
	siteGenerator ports.SiteGenerator
	baseURL       SiteBaseURL
}

func NewMarkdownService(
	siteGenerator ports.SiteGenerator,
	frontmatter ports.FrontMatterReadWriter,
	renderer ports.MarkdownRenderer,
	paths ports.PathingService,
	files ports.FileRepository,
	baseURL SiteBaseURL,
) *MarkdownService {
	return &MarkdownService{
		siteGenerator: siteGenerator,
		frontmatter:   frontmatter,
		renderer:      renderer,
		paths:         paths,
		files:         files,
		baseURL:       baseURL,
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
	switch nodes.CurrentNode().(type) {
	case dto.User, dto.Term:
		err := svc.files.Save(data, root, relPath)
		if err != nil {
			return err
		}
	default:
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
		err = svc.siteGenerator.Build(nodes.User.(dto.User), nodes.Term.(dto.Term), nodes.Course.(dto.Course))
		if err != nil {
			return err
		}
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
	switch nodes.CurrentNode().(type) {
	case dto.User, dto.Term:
		err = svc.files.Save(data, root, relPath)
		if err != nil {
			return err
		}
	default:
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
		err = svc.siteGenerator.Build(nodes.User.(dto.User), nodes.Term.(dto.Term), nodes.Course.(dto.Course))
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *MarkdownService) Update(data []byte, newName, oldName string, nodes ports.Nodes) error {
	root := svc.paths.NodeFilesDirPath(nodes.ToSlice()...)
	switch nodes.CurrentNode().(type) {
	case dto.User, dto.Term:
		err := svc.files.Update(data, root, oldName, newName)
		if err != nil {
			return err
		}
	default:
		mdFile, err := svc.setFrontMatter(data, root, newName)
		if err != nil {
			return err
		}
		data, err = svc.frontmatter.ToBytes(mdFile)
		if err != nil {
			return err
		}
		err = svc.files.Update(data, root, oldName, newName)
		if err != nil {
			return err
		}
		err = svc.siteGenerator.Build(nodes.User.(dto.User), nodes.Term.(dto.Term), nodes.Course.(dto.Course))
		if err != nil {
			return err
		}
	}

	return nil
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
	file.URL = svc.fileURL(root, relPath)
	rootDirSegments := strings.Split(root, "/")
	parentPath := filepath.Join(rootDirSegments[8:]...)
	file.Type = "standalone"
	if file.FrontMatter.Params == nil {
		file.FrontMatter.Params = make(map[string]any)
	}
	file.FrontMatter.Params["parentPath"] = parentPath
	file.Params["breadCrumbs"] = svc.siteGenerator.BreadCrumbs(file.URL)
	return file, nil
}

type SiteBaseURL interface {
	StaticSiteURL(lastName string, courseID int) string
}

func (svc *MarkdownService) PreviewFileURL(lastName, relPath string, nodes ports.Nodes) string {
	baseURL := svc.baseURL.StaticSiteURL(lastName, nodes.Course.GetID().(int))
	root := svc.paths.NodeFilesDirPath(nodes.ToSlice()...)
	path := svc.fileURL(root, relPath)
	final, _ := url.JoinPath(baseURL, path)
	return final
}

func (svc *MarkdownService) fileURL(root, relPath string) string {
	relPath = strings.ToLower(relPath)
	relPath = strings.TrimSuffix(relPath, ".md")
	relPath = strings.ReplaceAll(relPath, " ", "-")
	path := filepath.Join(root, relPath)
	segments := strings.Split(path, "/")
	path = filepath.Join(segments[8:]...)
	return strings.ReplaceAll(path, "_", "-")
}
