package service

import (
	"gh_static_portfolio/internal/data"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"os"
	"path/filepath"
)

func (svc CourseService) LessonFilePath(lessonID int, path string) string {
	root := data.LessonFilesDirPath(lessonID)
	path = filepath.Join(root, path)
	return path
}

func (svc CourseService) IsDir(lessonID int, path string) (bool, error) {
	root := data.LessonFilesDirPath(lessonID)
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

func (svc CourseService) LessonFiles(lessonID int, path string) ([]mt.FilesPageItem, error) {
	root := data.LessonFilesDirPath(lessonID)
	path = filepath.Join(root, path)
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var items []mt.FilesPageItem
	for _, entry := range entries {
		var item mt.FilesPageItem
		item.Path = entry.Name()
		if entry.IsDir() {
			item.IsDir = true
		}
		items = append(items, item)
	}
	return items, nil
}

// func (svc CourseService) PostLessonFile(lessonID int, path string)
