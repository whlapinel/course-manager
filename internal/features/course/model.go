package course

import "gh_static_portfolio/internal/ports"

type Course struct {
	ports.BaseNode[int, int]
}

type CourseType int

type Courses []Course
