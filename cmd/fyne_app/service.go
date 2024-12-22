package main

import "gh_static_portfolio/cmd/data"


func NewCourseService(repo data.CourseRepo) CourseService {
	return CourseService{repo}
}

type CourseService struct {
	repo data.CourseRepo
}
