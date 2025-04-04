package service

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	mt "gh_static_portfolio/internal/templates/app"
	"gh_static_portfolio/internal/util"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func (svc CourseService) WriteToMarkdown(relPath, content string, nodes domain.Nodes) error {
	nodeDirPath := data.NodeFilesDirPath(nodes.ToSlice()...)
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

func (svc CourseService) WriteFile(file *multipart.FileHeader, relPath string, nodes domain.Nodes) error {
	nodeDirPath := data.NodeFilesDirPath(nodes.ToSlice()...)
	dstPathDir := filepath.Join(nodeDirPath, relPath)
	dstPath := filepath.Join(dstPathDir, file.Filename)
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

func (svc CourseService) NodeFilePath(path string, nodes ...domain.CourseNode) string {
	root := data.NodeFilesDirPath(nodes...)
	path = filepath.Join(root, path)
	return path
}

func (svc CourseService) IsDir(path string, nodes ...domain.CourseNode) (bool, error) {
	root := data.NodeFilesDirPath(nodes...)
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

func (svc CourseService) NodeFiles(path string, nodes ...domain.CourseNode) ([]mt.FilesPageItem, error) {
	root := data.NodeFilesDirPath(nodes...)
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

func (svc CourseService) CreateNodeFilesDir(nodes ...domain.CourseNode) error {
	root := data.NodeFilesDirPath(nodes...)
	err := os.MkdirAll(root, os.ModePerm)
	if err != nil {
		return err
	}
	return err
}

// func (svc CourseService) PostLessonFile(lessonID int, path string)

func (svc CourseService) DeleteFile(path string, nodes ...domain.CourseNode) error {
	root := data.NodeFilesDirPath(nodes...)
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
