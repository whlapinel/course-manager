package domain

type BlogRepository interface {
	GetAll() ([]*Blog, error)
}

type Blog struct {
	title   string
	content string
}

func NewBlog(title, content string) *Blog {
	return &Blog{
		title:   title,
		content: content,
	}
}

func (b *Blog) GetTitle() string {
	return b.title
}

func (b *Blog) Content() string {
	return b.content
}

type EducationRepository interface {
	GetAll() ([]*EducationItem, error)
}

type EducationItem struct {
	Image   string
	Year    string
	School  string
	Degree  string
	Minor   string
	Classes []*ClassItem
}

func (e *EducationItem) GetTitle() string {
	return e.School
}

type ProjectRepository interface {
	GetAll() ([]*ProjectItem, error)
}

type ProjectItem struct {
	Image       string
	Title       string
	Subtitle    string
	Description string
	Link        string
}

type ClassRepository interface {
	GetAll() ([]*ClassItem, error)
}

type ClassItem struct {
	Image       string
	Number      string
	Title       string
	Subtitle    string
	Description string
	Link        string
}

type WorkHistoryRepository interface {
	GetAll() ([]*WorkHistoryItem, error)
}

type WorkHistoryItem struct {
	Image       string
	Year        string
	Company     string
	Position    string
	Description string
}

type SkillRepository interface {
	GetAll() ([]*SkillItem, error)
}

type SkillItem struct {
	Image       string
	Title       string
	Subtitle    string
	Description string
}
