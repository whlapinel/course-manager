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
