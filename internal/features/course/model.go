package course

type Course struct {
	ID          int    `json:"id"`
	ParentID    int    `json:"parentID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CourseType int

type Courses []Course
