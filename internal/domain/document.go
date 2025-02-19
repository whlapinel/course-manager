package domain

// for files that have metadata registered in DB
type Document struct {
	Title string
	Path  string // must be relative to node `files` directory
}
