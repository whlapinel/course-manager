package service

import "gh_static_portfolio/internal/domain"

func (svc CourseService) GetStandardSets() ([]domain.StandardSet, error) {
	return svc.repo.GetStandardSets()
}

func (svc CourseService) SetStandardSet(courseID, setID int) error {
	course, err := svc.GetCourse(courseID)
	if err != nil {
		return err
	}
	_, err = svc.repo.GetStandardSetByID(setID)
	if err != nil {
		return err
	}
	course.StandardSet.ID = setID
	return svc.repo.UpdateCourse(course)
}

func (svc CourseService) GetCourseStandardsWithObjectives(course domain.Course) ([]domain.Standard, error) {
	return svc.repo.GetCourseStandardsWithObjectives(course.StandardSet)
}

func (svc CourseService) GetCourseStandards(course domain.Course) ([]domain.Standard, error) {
	return svc.repo.GetCourseStandards(course.StandardSet)
}

func (svc CourseService) GetLessonStandards(lesson domain.Lesson) ([]domain.Standard, error) {
	return svc.repo.GetLessonObjectives(lesson)
}
