package hugo

type Kind string

const (
	HomeKind    Kind = "home"
	PageKind    Kind = "page"
	SectionKind Kind = "section"
)

type Type string

const (
	CalendarType Type = "calendar"
	CourseType   Type = "course"
	UnitType     Type = "unit"
	LessonType   Type = "lesson"
	FilesType    Type = "files"
)

type HomogenizedPageData struct {
	Kind   `json:"kind"`
	Type   `json:"type"`
	Path   string `json:"path"`
	URL    string `json:"url"`
	Params any    `json:"params"`
	Title  string `json:"title"`
	Weight int    `json:"weight"`
}
