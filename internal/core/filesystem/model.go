package filesystem

type File struct {
	Name         string `json:"name"` // last element of path including extension
	AbsolutePath string `json:"absolutePath"`
	RootDir      string `json:"rootDir"`
	RelativePath string `json:"relativePath"` // relative to node file's root (should not allow escaping)
}
