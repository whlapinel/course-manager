package service

import "gh_static_portfolio/internal/domain"

type SaveCourseParams struct {
	TermID      int
	Name        string
	Description string
}

func (svc CourseService) SaveCourse(params SaveCourseParams) (domain.Course, error) {

	course := domain.NewCourse(domain.NewCourseParams{
		TermID:      params.TermID,
		Name:        params.Name,
		Description: params.Description,
	})
	id, err := svc.repo.SaveCourse(course)
	if err != nil {
		return domain.Course{}, err
	}
	course.ID = id
	return course, nil

}
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

func (svc CourseService) DeleteCourse(courseID int) error {
	return svc.repo.DeleteCourse(courseID)
}

func (svc CourseService) CopyCourseToTerm(courseID int, termID int) (domain.Course, error) {
	// get course
	course, err := svc.GetCourse(courseID)
	if err != nil {
		return domain.Course{}, err
	}
	term, err := svc.GetTerm(termID)
	if err != nil {
		return domain.Course{}, err
	}
	newCourse := course.FitToTerm(term)
	svc.SaveCourse(SaveCourseParams{
		TermID:      termID,
		Name:        newCourse.Name,
		Description: newCourse.Description,
	})
	for _, unit := range newCourse.Units {
		svc.SaveUnit(SaveUnitParams{
			Unit: *unit,
		})
		for _, lesson := range unit.Lessons {
			svc.SaveLesson(SaveLessonParams{
				Lesson: *lesson,
			})
		}
	}
	return newCourse, nil
}
