package service

import "gh_static_portfolio/internal/data"

func NewCourseService(repo data.CourseRepo) CourseService {
	return CourseService{repo}
}

type CourseService struct {
	repo data.CourseRepo
}
