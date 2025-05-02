package filesystem

type File struct {
	Name         string // last element of path including extension
	AbsolutePath string
	RootDir      string
	RelativePath string // relative to node file's root (should not allow escaping)
}
