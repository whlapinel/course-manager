package files

import (
	fileviews "gh_static_portfolio/internal/app/views/files"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/util"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Service struct {
	fileRepo    ports.FileRepository
	pathService ports.PathingService
}

func NewFileService(
	fileRepo ports.FileRepository,
	pathService ports.PathingService,

) *Service {
	return &Service{
		fileRepo:    fileRepo,
		pathService: pathService,
	}
}

func (svc *Service) WriteMarkdown(relPath string, content []byte, nodes ports.Nodes) error {
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	err := svc.fileRepo.Update(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) FileContent(relPath string, nodes ports.Nodes) ([]byte, error) {
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	content, err := svc.fileRepo.Read(root, relPath)
	if err != nil {
		return nil, err
	}
	return content, nil

}

func (svc *Service) FileInfo(relPath string, nodes ports.Nodes) (ports.FileInfo, error) {
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	fileInfo, err := svc.fileRepo.FileInfo(relPath, root)
	if err != nil {
		return fileInfo, err
	}
	return fileInfo, nil
}

func (svc *Service) Update(content []byte, relPath string, nodes ports.Nodes) error {
	root := svc.pathService.NodeFilesDirPath(nodes.ToSlice()...)
	err := svc.fileRepo.Update(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Save(header *multipart.FileHeader, relPath string, nodes ports.Nodes) error {
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

func (svc *Service) NodeFiles(path string, nodes ...ports.Node) ([]fileviews.FilesPageItem, error) {
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

func (svc *Service) DeleteFile(path string, nodes ...ports.Node) error {
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
