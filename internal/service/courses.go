package service

import (
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"time"
)

type SaveCourseParams struct {
	TermID          int
	Name            string
	Description     string
	InstructionDays []time.Time
}

func (svc CourseService) SaveCourse(params SaveCourseParams) (domain.Course, error) {
	course := domain.NewCourse(domain.NewCourseParams{
		TermID:      params.TermID,
		Name:        params.Name,
		Description: params.Description,
	})
	if params.InstructionDays != nil {
		course.InstructionalDays = params.InstructionDays
	}
	id, err := svc.repo.SaveCourse(course)
	if err != nil {
		return domain.Course{}, err
	}
	course.ID = id
	return course, nil
}
func (svc CourseService) GetCourseForCalendar(courseID int) (domain.Course, error) {
	course, err := svc.repo.GetCourse(courseID)
	if err != nil {
		return domain.Course{}, err
	}
	units, err := svc.repo.GetUnits(courseID)
	if err != nil {
		return domain.Course{}, err
	}
	for i, unit := range units {
		lessons, err := svc.repo.GetLessons(unit.ID)
		if err != nil {
			return domain.Course{}, err
		}
		units[i].Lessons = lessons
	}
	course.Units = units
	return domain.Course{}, nil
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
	course, err := svc.GetCourse(courseID)
	if err != nil {
		return err
	}
	units, err := svc.GetUnits(courseID)
	if err != nil {
		return err
	}
	for _, unit := range units {
		lessons, err := svc.GetLessons(unit.ID)
		if err != nil {
			return err
		}
		unit.Lessons = append(unit.Lessons, lessons...)
	}
	course.Units = append(course.Units, units...)
	return svc.repo.DeleteCourse(*course)
}

func (svc CourseService) CopyCourseToTerm(courseID int, termID int) (*domain.Course, error) {
	// get oldCourse
	oldCourse, err := svc.GetCourse(courseID)
	if err != nil {
		return nil, err
	}
	// get units
	units, err := svc.GetUnits(courseID)
	if err != nil {
		return nil, err
	}
	// for each unit get lessons
	for _, unit := range units {
		lessons, err := svc.GetLessons(unit.ID)
		if err != nil {
			return nil, err
		}
		// for each lesson append to unit.Lessons
		unit.Lessons = append(unit.Lessons, lessons...)
		oldCourse.Units = append(oldCourse.Units, unit)
	}
	oldTerm, err := svc.GetTerm(oldCourse.Term.ID)

	// get newTerm
	newTerm, err := svc.GetTerm(termID)
	if err != nil {
		return nil, err
	}
	// fit course to term
	newCourse := oldCourse.FitToTerm(newTerm)

	// save course
	newCourse, err = svc.SaveCourse(SaveCourseParams{
		TermID:      newTerm.ID,
		Name:        newCourse.Name,
		Description: newCourse.Description,
	})
	if err != nil {
		return nil, err
	}
	srcRoot := data.CourseDirPath(oldCourse.ID)
	destRoot := data.CourseDirPath(newCourse.ID)
	err = data.CopyNodeDir(srcRoot, destRoot)
	if err != nil {
		return nil, err
	}
	// save unit with modified CourseID
	for _, oldUnit := range oldCourse.Units {
		oldUnit.CourseID = newCourse.ID
		newUnit, err := svc.SaveUnit(SaveUnitParams{
			Unit: oldUnit,
		})
		if err != nil {
			return nil, err
		}
		srcRoot := data.UnitDirPath(oldUnit.ID)
		destRoot := data.UnitDirPath(newUnit.ID)
		err = data.CopyNodeDir(srcRoot, destRoot)
		if err != nil {
			return nil, err
		}
		// save lesson with modified UnitID
		for _, oldLesson := range oldUnit.Lessons {
			oldLesson.UnitID = newUnit.ID
			newLesson, err := svc.SaveLesson(SaveLessonParams{
				Lesson: oldLesson,
			})
			if err != nil {
				return nil, err
			}
			srcRoot := data.LessonDirPath(oldTerm, oldCourse, oldUnit, oldLesson)
			destRoot := data.LessonDirPath(newTerm, newCourse, newUnit, newLesson)
			err = data.CopyNodeDir(srcRoot, destRoot)
			if err != nil {
				return nil, err
			}

		}
	}
	return &newCourse, nil
}
