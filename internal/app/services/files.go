package services

import (
	"fmt"
	mt "gh_static_portfolio/internal/newtemplates/app"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/util"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileService struct {
	MarkdownToHTML func(srcPath string) error
}

func NewFileService(MarkdownToHTML func(srcPath string) error) *FileService {
	return &FileService{
		MarkdownToHTML: MarkdownToHTML,
	}
}

func (svc FileService) WriteToMarkdown(relPath, content string, nodes node.Nodes) error {
	nodeDirPath := NodeFilesDirPath(nodes.ToSlice()...)
	dstPath := filepath.Join(nodeDirPath, relPath)
	log.Println("content", content)
	// Create a destination file
	log.Println("dstPath:", dstPath)
	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %s", err)
	}
	defer dst.Close()
	bytes, err := dst.Write([]byte(content))
	if err != nil {
		return err
	}
	log.Printf("%d bytes written to file at %s", bytes, dstPath)
	err = svc.MarkdownToHTML(dstPath)
	if err != nil {
		return err
	}
	return nil
}

func (svc FileService) WriteFile(file *multipart.FileHeader, relPath string, nodes node.Nodes) error {
	dstPath, err := svc.NodeFilePath(relPath, nodes.ToSlice()...)
	if err != nil {
		return err
	}
	// Open the file
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	defer src.Close()
	// Create a destination file
	log.Println("dstPath:", dstPath)
	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %s", err)
	}
	defer dst.Close()
	// Copy the content of the uploaded file to the destination
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to save file: %s", err)
	}
	if util.IsMarkdown(dstPath) {
		log.Println("this is a markdown file")
		err = svc.MarkdownToHTML(dstPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsSecurePath(root, path string) bool {
	rootClean := filepath.Clean(root)
	cleaned := filepath.Clean(filepath.Join(rootClean, path))

	// Check if cleaned is inside root
	rel, err := filepath.Rel(rootClean, cleaned)
	if err != nil {
		log.Printf("SECURITY WARNING: error resolving path: %v", err)
		return false
	}
	if strings.HasPrefix(rel, "..") || strings.HasPrefix(filepath.ToSlash(rel), "../") {
		log.Printf("SECURITY WARNING: path traversal attempt detected: %q resolves to %q (rel: %q)", path, cleaned, rel)
		return false
	}
	return true
}

func (svc FileService) NodeFilePath(path string, nodes ...node.Node) (string, error) {
	root := NodeFilesDirPath(nodes...)
	cleaned := filepath.Join(root, filepath.Clean(path))

	// Normalize root too
	rootClean := filepath.Clean(root)

	// Ensure the resolved path is still under root
	if !strings.HasPrefix(cleaned, rootClean+string(os.PathSeparator)) {
		log.Printf("SECURITY WARNING: path traversal attempt detected: %q", path)
		return "", fmt.Errorf("SECURITY WARNING: path traversal attempt detected")
	}
	log.Println("root:   ", rootClean)
	log.Println("joined: ", cleaned)

	return cleaned, nil
}

func (svc FileService) IsDir(path string, nodes ...node.Node) (bool, error) {
	root := NodeFilesDirPath(nodes...)
	path = filepath.Join(root, path)
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if info.IsDir() {
		return true, nil
	}
	return false, nil
}

func (svc FileService) NodeFiles(path string, nodes ...node.Node) ([]mt.FilesPageItem, error) {
	root := NodeFilesDirPath(nodes...)
	path = filepath.Join(root, path)
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Println("error in CourseService.LessonFiles(): ReadDir", path)
		return nil, err
	}
	var markdownItems = make(map[string]mt.FilesPageItem)
	var htmlItems = make(map[string]mt.FilesPageItem)
	var items []mt.FilesPageItem
	for _, entry := range entries {
		var item mt.FilesPageItem
		item.Path = entry.Name()
		item.Name = entry.Name()
		if entry.IsDir() {
			item.IsDir = true
		}
		// don't want to show html versions of markdown files
		// which have same name as markdown
		if util.IsMarkdown(item.Path) {
			name := strings.TrimSuffix(item.Path, ".md")
			markdownItems[name] = item
			items = append(items, item)
			// make sure html file doesn't have same name as any markdown file
		} else if strings.HasSuffix(item.Path, ".html") {
			name := strings.TrimSuffix(item.Path, ".html")
			htmlItems[name] = item
		} else {
			items = append(items, item)
		}
	}
	// append html file names that don't match markdown file names
	for name, item := range htmlItems {
		if _, ok := markdownItems[name]; !ok {
			items = append(items, item)
		}
	}
	return items, nil
}

func (svc FileService) CreateNodeFilesDir(nodes ...node.Node) error {
	root := NodeFilesDirPath(nodes...)
	err := os.MkdirAll(root, os.ModePerm)
	if err != nil {
		return err
	}
	return err
}

// func (svc CourseService) PostLessonFile(lessonID int, path string)

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
