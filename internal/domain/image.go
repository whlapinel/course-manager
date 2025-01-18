package domain

import "path/filepath"

type Image struct {
	ID          int
	Name        string
	Description string
	SourcePath  string
	BasePath    string
}

func NewImage(name, descr, srcPath string) Image {
	return Image{Name: name, Description: descr, SourcePath: srcPath, BasePath: filepath.Base(srcPath)}
}

