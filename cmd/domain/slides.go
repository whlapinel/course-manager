package domain

type Slides struct {
	ID          int
	Name        string
	Description string
	SourcePath  string
}

func NewSlides(name, descr, srcPath string) Slides {
	return Slides{Name: name, Description: descr, SourcePath: srcPath}
}
