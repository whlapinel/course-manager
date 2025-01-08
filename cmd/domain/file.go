package domain

type FilesDir struct {
	ID          int
	Name        string // typically should be lesson name
	Description string // describe contents
	SourcePath  string // original source path of file, not persisted
}

func NewFileDir(name, descr, srcPath string) FilesDir {
	return FilesDir{Name: name, Description: descr, SourcePath: srcPath}
}
