package course

type Course struct {
	ID          int
	ParentID    int
	Name        string
	Description string
}

type CourseType int

type Courses []Course
