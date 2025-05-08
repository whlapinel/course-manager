package files

import (
	"fmt"
	fileviews "gh_static_portfolio/internal/app/views/files"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/util"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type Service struct {
	fileRepo FileRepository
}

func NewFileService(
	fileRepo FileRepository,

) *Service {
	return &Service{
		fileRepo: fileRepo,
	}
}

func (svc *Service) WriteMarkdown(relPath string, content []byte, nodes node.Nodes) error {
	root := NodeFilesDirPath(nodes.ToSlice()...)
	err := svc.fileRepo.Update(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) FileContent(relPath string, nodes node.Nodes) ([]byte, error) {
	root := NodeFilesDirPath(nodes.ToSlice()...)
	content, err := svc.fileRepo.Read(root, relPath)
	if err != nil {
		return nil, err
	}
	return content, nil

}

func (svc *Service) FileInfo(relPath string, nodes node.Nodes) (FileInfo, error) {
	root := NodeFilesDirPath(nodes.ToSlice()...)
	fileInfo, err := svc.fileRepo.FileInfo(relPath, root)
	if err != nil {
		return fileInfo, err
	}
	return fileInfo, nil
}

func (svc *Service) Update(content []byte, relPath string, nodes node.Nodes) error {
	root := NodeFilesDirPath(nodes.ToSlice()...)
	err := svc.fileRepo.Update(content, root, relPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Save(header *multipart.FileHeader, relPath string, nodes node.Nodes) error {
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

func (svc *Service) NodeFiles(path string, nodes ...node.Node) ([]fileviews.FilesPageItem, error) {
	root := NodeFilesDirPath(nodes...)
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

func (svc *Service) DeleteFile(path string, nodes ...node.Node) error {
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

// nodes is path from root, where term is the root
func NodeDirPath(nodes ...node.Node) string {
	var path = "./internal/data"
	for _, node := range nodes {
		log.Println("node", node)
		if node == nil {
			break
		}
		path = filepath.Join(path, strings.ToLower(node.TypeName()+"s"))
		var dirName string
		if id, ok := node.GetID().(string); ok {
			dirName = fmt.Sprintf("%s_%s", strings.ToLower(node.TypeName()), id)
		} else if id, ok := node.GetID().(int); ok {
			dirName = fmt.Sprintf("%s_%d", strings.ToLower(node.TypeName()), id)
		}
		path = filepath.Join(
			path,
			dirName,
		)
	}
	return path
}

func NodeFilesDirPath(nodes ...node.Node) string {
	return filepath.Join(NodeDirPath(nodes...), "files")
}

func NodeImagePath(nodes ...node.Node) string {
	return filepath.Join(NodeDirPath(nodes...), "image.png")
}
