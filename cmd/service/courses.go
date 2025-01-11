package service

import "gh_static_portfolio/cmd/domain"

func (svc CourseService) GetCourseForCalendar(courseID int) (*domain.Course, error) {
	course, err := svc.repo.GetCourse(courseID)
	if err != nil {
		return nil, err
	}
	units, err := svc.repo.GetUnits(courseID)
	if err != nil {
		return nil, err
	}
	for i, unit := range units {
		lessons, err := svc.repo.GetLessons(unit.ID)
		if err != nil {
			return nil, err
		}
		units[i].Lessons = lessons
	}
	course.Units = units
	return course, nil
}

func (svc CourseService) GetCourse(courseID int) (*domain.Course, error) {
	return svc.repo.GetCourse(courseID)
}
func (svc CourseService) GetCourses(termID int) (domain.Courses, error) {
	return svc.repo.GetCourses(termID)
}

func (svc CourseService) UpdateCourse(instance domain.Course) error {
	return svc.repo.UpdateCourse(instance)
}
