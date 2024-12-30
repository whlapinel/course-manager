package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
	"time"
)

func (c CourseRepo) SaveCourse(course domain.Course) (id int, err error) {
	ctx := context.Background()
	savedCourse, err := c.queries.SaveCourse(ctx, database.SaveCourseParams{
		TermID: sql.NullInt64{
			Int64: int64(course.Term.ID),
			Valid: course.Term.ID != 0,
		},
		Name: course.Name,
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
		*unit = *savedUnit
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

func (c CourseRepo) GetCourses(termID int) ([]*domain.Course, error) {
	dbCourses, err := c.queries.GetCourses(context.Background(), sql.NullInt64{Valid: true, Int64: int64(termID)})
	if err != nil {
		return nil, err
	}
	term, err := c.GetTermByID(termID)
	if err != nil {
		return nil, err
	}
	termWithDates, err := c.GetTermDates(termID)
	if err != nil {
		return nil, err
	}
	// THIS IS KINDA JACKED UP
	term.InstructionalDays = termWithDates.InstructionalDays
	var courses []*domain.Course
	for _, dbCourse := range dbCourses {
		course := domain.Course{
			ID:          int(dbCourse.CourseID),
			Name:        dbCourse.CourseName,
			Description: dbCourse.CourseDescr.String,
			Term:        term,
		}
		dbUnits, err := c.queries.GetUnits(context.Background(), int64(course.ID))
		if err != nil {
			return nil, err
		}
		var units []*domain.Unit
		for _, dbUnit := range dbUnits {
			unit := domain.Unit{
				ID:          int(dbUnit.ID),
				CourseID:    int(dbUnit.CourseID),
				Number:      int(dbUnit.Number),
				SequenceNum: int(dbUnit.Sequence),
				Name:        dbUnit.Name,
				Description: dbUnit.Description.String,
			}
			dbLessons, err := c.queries.GetLessons(context.Background(), int64(unit.ID))
			if err != nil {
				return nil, err
			}
			var lessons []*domain.Lesson
			for _, dbLesson := range dbLessons {
				lesson := domain.Lesson{
					ID:          int(dbLesson.ID),
					UnitID:      unit.ID,
					Number:      int(dbLesson.Number),
					Name:        dbLesson.Name.String,
					Description: dbLesson.Description.String,
				}
				dbLessonDates, err := c.queries.GetLessonDates(context.Background(), int64(lesson.ID))
				if err != nil {
					return nil, err
				}
				var lessonDates []time.Time
				for _, dbLessonDate := range dbLessonDates {
					lessonDate, err := time.Parse(time.DateOnly, dbLessonDate)
					if err != nil {
						return nil, err
					}
					lessonDates = append(lessonDates, lessonDate)
				}
				lesson.Dates = lessonDates
				lessons = append(lessons, &lesson)
			}
			unit.Lessons = lessons
			units = append(units, &unit)
		}
		course.Units = units
		if course.Term.Start.IsZero() {
			log.Fatal("term not initialized")
		}
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
