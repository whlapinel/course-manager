package domain

import "path/filepath"

type File struct {
	ID          int
	Name        string
	Description string
	SourcePath  string // original source path of file, not persisted
	BasePath    string // the last segment of the original path including extension, prefixed with file ID when saved
}

func NewFile(name, descr, srcPath string) File {
	return File{Name: name, Description: descr, SourcePath: srcPath, BasePath: filepath.Base(srcPath)}
}
