package services

import (
	"gh_static_portfolio/internal/core/filesystem"
	mt "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/util"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileService struct {
	fileRepo         filesystem.FileRepository
	markdownRenderer filesystem.MarkdownRenderer
}

func NewFileService(
	fileRepo filesystem.FileRepository,
	markdownRenderer filesystem.MarkdownRenderer,
) *FileService {
	return &FileService{
		fileRepo:         fileRepo,
		markdownRenderer: markdownRenderer,
	}
}

func (svc FileService) ViewMarkdown(relPath string, nodes node.Nodes) ([]byte, error) {
	root := NodeFilesDirPath(nodes.ToSlice()...)
	content, err := svc.fileRepo.Read(root, relPath)
	if err != nil {
		return nil, err
	}
	html, err := svc.markdownRenderer.ToHTML(content)
	if err != nil {
		return nil, err
	}
	return html, nil
}

func (svc FileService) Update(content []byte, relPath string, nodes node.Nodes) error {
	root := NodeFilesDirPath(nodes.ToSlice()...)
	err := svc.fileRepo.Update(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc FileService) Save(header *multipart.FileHeader, relPath string, nodes node.Nodes) error {
	file, err := header.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	root := NodeFilesDirPath(nodes.ToSlice()...)
	err = svc.fileRepo.Save(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc FileService) NodeFiles(path string, nodes ...node.Node) ([]mt.FilesPageItem, error) {
	root := NodeFilesDirPath(nodes...)
	entries, err := svc.fileRepo.List(root, path)
	if err != nil {
		return nil, err
	}
	var items []mt.FilesPageItem
	for _, entry := range entries {
		var item mt.FilesPageItem
		item.Path = entry.Name()
		item.Name = entry.Name()
		if entry.IsDir() {
			item.IsDir = true
		}
		if util.IsMarkdown(item.Path) {
			item.IsMarkdown = true
		}
		items = append(items, item)
	}
	return items, nil
}

func (svc FileService) DeleteFile(path string, nodes ...node.Node) error {
	root := NodeFilesDirPath(nodes...)
	path = filepath.Join(root, path)
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
