package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/internal/data/database"
	"gh_static_portfolio/internal/domain"
	"os"
	"path/filepath"
)

func CourseDirPath(courseID int) string {
	return fmt.Sprintf("./internal/data/courses/course_%d", courseID)
}

func CourseFilesDirPath(courseID int) string {
	return filepath.Join(CourseDirPath(courseID), "files")
}

func CourseImagePath(courseID int) string {
	return filepath.Join(CourseDirPath(courseID), "image.png")
}

func (c CourseRepo) SaveCourse(course domain.Course) (id int, err error) {
	ctx := context.Background()
	savedCourse, err := c.queries.SaveCourse(ctx, database.SaveCourseParams{
		TermID: int64(course.Term.ID),
		Name:   course.Name,
		Description: sql.NullString{
			String: course.Description,
			Valid:  course.Description != "",
		},
		StdSetID: sql.NullInt64{
			Valid: course.StandardSet.ID != 0,
			Int64: int64(course.StandardSet.ID),
		},
	})
	if err != nil {
		return 0, fmt.Errorf("courseRepo.SaveCourse(): %s", err)
	}
	return int(savedCourse.ID), nil
}
func (cr CourseRepo) GetCourse(courseID int) (domain.Course, error) {
	dbCourse, err := cr.queries.GetCourseByCourseID(context.Background(), int64(courseID))
	if err != nil {
		return domain.Course{}, err
	}
	course := domain.Course{
		ID:          int(dbCourse.ID),
		Name:        dbCourse.Name,
		Description: dbCourse.Description.String,
		Term: domain.Term{
			ID: int(dbCourse.TermID),
		},
		StandardSet: domain.StandardSet{
			ID: int(dbCourse.StdSetID.Int64),
		},
	}
	term, err := cr.GetTermByID(course.Term.ID)
	if err != nil {
		return domain.Course{}, err
	}

	termWithDates, err := cr.GetTermWithDates(term.ID)
	if err != nil {
		return domain.Course{}, err
	}
	term.InstructionalDays = termWithDates.InstructionalDays
	occasions, err := cr.GetTermOccasions(term.ID)
	if err != nil {
		return domain.Course{}, err
	}
	term.Occasions = occasions
	course.Term = term
	if course.StandardSet.ID != 0 {
		standardSet, err := cr.GetStandardSetByID(course.StandardSet.ID)
		if err != nil {
			return domain.Course{}, err
		}
		course.StandardSet = standardSet
	}
	return course, nil

}
func (cr CourseRepo) GetCourses(termID int) ([]domain.Course, error) {
	dbCourses, err := cr.queries.GetCoursesByTermID(context.Background(), int64(termID))
	if err != nil {
		return nil, err
	}
	term, err := cr.GetTermByID(termID)
	if err != nil {
		return nil, err
	}
	termWithDates, err := cr.GetTermWithDates(termID)
	if err != nil {
		return nil, err
	}
	term.InstructionalDays = termWithDates.InstructionalDays
	var courses []domain.Course
	for _, dbCourse := range dbCourses {
		course := domain.Course{
			ID:          int(dbCourse.ID),
			Name:        dbCourse.Name,
			Description: dbCourse.Description.String,
			Term:        term,
		}
		units, err := cr.GetUnits(course.ID)
		if err != nil {
			return nil, err
		}
		for i, unit := range units {
			units[i].Lessons, err = cr.GetLessons(unit.ID)
			if err != nil {
				return nil, err
			}
		}
		course.Units = units
		courses = append(courses, course)

	}
	return courses, nil
}

func (c CourseRepo) UpdateCourse(course domain.Course) error {
	err := c.queries.UpdateCourse(context.Background(), database.UpdateCourseParams{
		ID:   int64(course.ID),
		Name: course.Name,
		Description: sql.NullString{
			Valid:  course.Description != "",
			String: course.Description,
		},
		StdSetID: sql.NullInt64{
			Valid: course.StandardSet.ID != 0,
			Int64: int64(course.StandardSet.ID),
		},
	})
	if err != nil {
		return err
	}
	return nil

}

func (cr CourseRepo) DeleteCourse(course domain.Course) error {
	for _, unit := range course.Units {
		for _, lesson := range unit.Lessons {
			cr.DeleteLesson(lesson)
		}
		cr.DeleteUnit(unit)
	}
	err := cr.deleteCourseDir(course.ID)
	if err != nil {
		return err
	}
	_, err = cr.queries.DeleteCourse(context.Background(), int64(course.ID))
	if err != nil {
		return err
	}
	return nil
}

func (cr CourseRepo) deleteCourseDir(courseID int) error {
	path := CourseDirPath(courseID)
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
