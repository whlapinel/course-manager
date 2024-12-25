package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
)

func (c CourseRepo) GetTemplate(id int) (domain.Course, error) {
	dbCourse, err := c.queries.GetTemplate(context.Background(), int64(id))
	if err != nil {
		return domain.Course{}, err
	}
	course := domain.Course{
		ID:          int(dbCourse.CourseID),
		Name:        dbCourse.CourseName,
		Description: dbCourse.CourseDescr.String,
	}
	dbUnits, err := c.queries.GetUnits(context.Background(), dbCourse.CourseID)
	if err != nil {
		return domain.Course{}, err
	}
	var units []domain.Unit
	for _, dbUnit := range dbUnits {
		unit := domain.Unit{
			ID:          int(dbUnit.ID),
			CourseID:    int(dbUnit.CourseID),
			Number:      int(dbUnit.Number),
			Name:        dbUnit.Name,
			Description: dbUnit.Description.String,
		}
		dbLessons, err := c.queries.GetLessons(context.Background(), dbUnit.ID)
		if err != nil {
			return domain.Course{}, err
		}
		var lessons []domain.Lesson
		for _, dbLesson := range dbLessons {
			lesson := domain.Lesson{
				ID:          int(dbLesson.ID),
				TemplateID:  int(dbLesson.TemplateID.Int64),
				Number:      int(dbLesson.Number),
				Name:        dbLesson.Name.String,
				Description: dbLesson.Description.String,
			}
			lessons = append(lessons, lesson)
		}
		unit.Lessons = lessons
		units = append(units, unit)
	}
	course.Units = units

	return course, nil

}

func (c CourseRepo) GetTemplates() ([]domain.Course, error) {
	dbCourses, err := c.queries.GetTemplates(context.Background())
	if err != nil {
		return nil, err
	}
	var courses []domain.Course
	for _, dbCourse := range dbCourses {
		course := domain.Course{
			ID:          int(dbCourse.CourseID),
			Name:        dbCourse.CourseName,
			Description: dbCourse.CourseDescr.String,
		}
		units, err := c.GetUnits(int(dbCourse.CourseID))
		if err != nil {
			return nil, err
		}
		for i, unit := range units {
			lessons, err := c.GetLessons(unit.ID)
			if err != nil {
				return nil, err
			}
			units[i].Lessons = lessons
		}
		course.Units = units
		courses = append(courses, course)
	}
	return courses, nil
}

func (c CourseRepo) SaveTemplate(course domain.Course) (domain.Course, error) {
	ctx := context.Background()
	dbCourse, err := c.queries.SaveTemplate(ctx, database.SaveTemplateParams{
		Name: course.Name,
		Description: sql.NullString{
			Valid:  course.Description != "",
			String: course.Description,
		},
	})
	if err != nil {
		return domain.Course{}, fmt.Errorf("c.queries.SaveCourse(): %s", err)
	}
	course.ID = int(dbCourse.ID)
	for i, unit := range course.Units {
		unit.CourseID = course.ID
		savedUnit, err := c.SaveUnit(unit)
		if err != nil {
			return domain.Course{}, fmt.Errorf("c.SaveUnit(): %s", err)
		}
		course.Units[i] = *savedUnit
		unit = *savedUnit
		for j, lesson := range unit.Lessons {
			lesson.UnitID = unit.ID
			savedLesson, err := c.SaveLessonTemplate(lesson)
			if err != nil {
				return domain.Course{}, fmt.Errorf("c.SaveLessonTemplate(): %s", err)
			}
			unit.Lessons[j] = *savedLesson
		}

	}
	return course, nil

}

// This is only for updating the course template, not units or lessons
func (c CourseRepo) UpdateTemplate(tpl domain.Course) error {
	err := c.queries.UpdateTemplate(context.Background(), database.UpdateTemplateParams{
		ID:   int64(tpl.ID),
		Name: tpl.Name,
		Description: sql.NullString{
			Valid:  tpl.Description != "",
			String: tpl.Description,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
