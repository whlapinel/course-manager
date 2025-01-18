package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gh_static_portfolio/internal/data/database"
	"gh_static_portfolio/internal/domain"
	"log"
)

func (c CourseRepo) SaveCourse(course domain.Course) (id int, err error) {
	ctx := context.Background()
	savedCourse, err := c.queries.SaveCourse(ctx, database.SaveCourseParams{
		TermID: int64(course.Term.ID),
		Name:   course.Name,
		Description: sql.NullString{
			String: course.Description,
			Valid:  course.Description != "",
		},
	})
	if err != nil {
		return 0, fmt.Errorf("courseRepo.SaveCourse(): %s", err)
	}
	course.ID = int(savedCourse.ID)
	for _, unit := range course.Units {
		unit.CourseID = int(savedCourse.ID)
		savedUnit, err := c.SaveUnit(*unit)
		if err != nil {
			return 0, fmt.Errorf("error in c.SaveUnit(): %s", err)
		}
		*unit = savedUnit
		log.Println("lesson count: ", len(unit.Lessons), "for ", unit.Name)
		for _, lesson := range unit.Lessons {
			lesson.UnitID = unit.ID
			_, err := c.SaveLessonInstance(*lesson)
			if err != nil {
				return 0, fmt.Errorf("error in SaveLessonInstance():%s", err)
			}
		}

	}
	return course.ID, nil

}
func (cr CourseRepo) GetCourse(courseID int) (*domain.Course, error) {
	dbCourse, err := cr.queries.GetCourseByCourseID(context.Background(), int64(courseID))
	if err != nil {
		return nil, err
	}
	course := domain.Course{
		ID:   int(dbCourse.ID),
		Name: dbCourse.Name,
		Term: domain.Term{
			ID: int(dbCourse.TermID),
		},
	}
	term, err := cr.GetTermByID(course.Term.ID)
	if err != nil {
		return nil, err
	}
	termWithDates, err := cr.GetTermDates(term.ID)
	if err != nil {
		return nil, err
	}
	term.InstructionalDays = termWithDates.InstructionalDays
	course.Term = term
	return &course, nil

}
func (cr CourseRepo) GetCourses(termID int) ([]*domain.Course, error) {
	dbCourses, err := cr.queries.GetCourses(context.Background(), int64(termID))
	if err != nil {
		return nil, err
	}
	term, err := cr.GetTermByID(termID)
	if err != nil {
		return nil, err
	}
	termWithDates, err := cr.GetTermDates(termID)
	if err != nil {
		return nil, err
	}
	term.InstructionalDays = termWithDates.InstructionalDays
	var courses []*domain.Course
	for _, dbCourse := range dbCourses {
		course := domain.Course{
			ID:          int(dbCourse.CourseID),
			Name:        dbCourse.CourseName,
			Description: dbCourse.CourseDescr.String,
			Term:        term,
		}
		dbImage, err := cr.queries.GetCourseImage(context.Background(), dbCourse.CourseID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
		}
		course.Image = domain.Image{
			ID:          int(dbImage.ID),
			Name:        dbImage.Name,
			Description: dbImage.Description.String,
			BasePath:    dbImage.BasePath,
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
		courses = append(courses, &course)

	}
	return courses, nil
}

func (c CourseRepo) UpdateCourse(instance domain.Course) error {
	err := c.queries.UpdateCourse(context.Background(), database.UpdateCourseParams{
		ID:   int64(instance.ID),
		Name: instance.Name,
		Description: sql.NullString{
			Valid:  instance.Description != "",
			String: instance.Description,
		},
	})
	if err != nil {
		return err
	}
	return nil

}
