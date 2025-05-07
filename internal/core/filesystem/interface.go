package filesystem

import "io/fs"

type FileInfo struct {
	IsDir   bool
	AbsPath string
}
type FileRepository interface {
	FileInfo(relPath, root string) (FileInfo, error) // for sending files
	Save(file []byte, root, relPath string) error
	Read(root, relPath string) ([]byte, error)
	Update(contents []byte, rootDir, relPath string) error
	Delete(rootDir, relPath string) error
	List(rootDir, relPath string) ([]fs.DirEntry, error)
}

type MarkdownRenderer interface {
	ToHTML(content []byte) ([]byte, error) // for markdown files, convert to HTML
}
