package domain

import "time"

type File struct {
	ID          int
	Name        string
	Description string
	SourcePath  string
	FileName    string
	Modified    time.Time
}

func NewFile(name, descr, path string) File {
	return File{Name: name, Description: descr, SourcePath: path}
}
