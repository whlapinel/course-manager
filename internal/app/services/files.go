package services

import (
	"fmt"
	fileviews "gh_static_portfolio/internal/app/views/files"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/util"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileService struct {
	fileRepo    ports.FileRepository
	pathService ports.PathingService
}

func NewFileService(
	fileRepo ports.FileRepository,
	pathService ports.PathingService,
) *FileService {
	return &FileService{
		fileRepo:    fileRepo,
		pathService: pathService,
	}
}

func (svc *FileService) WriteMarkdown(relPath string, content []byte, nodes ports.Nodes) error {
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	err := svc.fileRepo.Update(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *FileService) UpdateSlides(nodes ports.Nodes, content []byte) error {
	root := svc.pathService.NodeDirPath(nodes.ToSlice()...)
	slidesPath := svc.pathService.NodeSlidesMarkdownPath(nodes.ToSlice()...)
	return svc.fileRepo.Update(content, root, filepath.Base(slidesPath))
}

// for slides editor
func (svc *FileService) SlidesContent(nodes ports.Nodes) ([]byte, error) {
	if svc == nil {
		return nil, fmt.Errorf("service is nil")
	}
	if svc.pathService == nil {
		return nil, fmt.Errorf("path service is nil")
	}

	root := svc.pathService.NodeDirPath(nodes.ToSlice()...)
	slidesPath := svc.pathService.NodeSlidesMarkdownPath(nodes.ToSlice()...)
	content, err := svc.fileRepo.Read(root, filepath.Base(slidesPath))
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (svc *FileService) FileContent(relPath string, nodes ports.Nodes) ([]byte, error) {
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	content, err := svc.fileRepo.Read(root, relPath)
	if err != nil {
		return nil, err
	}
	return content, nil

}

func (svc *FileService) FileInfo(relPath string, nodes ports.Nodes) (ports.FileInfo, error) {
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	fileInfo, err := svc.fileRepo.FileInfo(relPath, root)
	if err != nil {
		return fileInfo, err
	}
	return fileInfo, nil
}

func (svc *FileService) Update(content []byte, relPath string, nodes ports.Nodes) error {
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	err := svc.fileRepo.Update(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *FileService) Save(header *multipart.FileHeader, relPath string, nodes ports.Nodes) error {
	file, err := header.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	err = svc.fileRepo.Save(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *FileService) NodeFiles(path string, nodes ...ports.Node) ([]fileviews.FilesPageItem, error) {
	root := svc.pathService.NodeFilesDirPath(nodes...)
	entries, err := svc.fileRepo.List(root, path)
	if err != nil {
		return nil, err
	}

	var items []fileviews.FilesPageItem
	for _, entry := range entries {
		var item fileviews.FilesPageItem
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

func (svc *FileService) DeleteFile(path string, nodes ...ports.Node) error {
	root := svc.pathService.NodeFilesDirPath(nodes...)
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
