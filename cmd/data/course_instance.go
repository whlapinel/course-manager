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

func (c CourseRepo) SaveInstance(instance domain.CourseInstance) error {
	ctx := context.Background()
	savedCourse, err := c.queries.SaveInstance(ctx, database.SaveInstanceParams{
		TemplateID: sql.NullInt64{
			Int64: int64(instance.TemplateID),
			Valid: true,
		},
		TermID: sql.NullInt64{
			Int64: int64(instance.Term.ID),
			Valid: instance.Term.ID != 0,
		},
		Name: instance.CourseTemplate.Name,
		Description: sql.NullString{
			String: instance.Description,
			Valid:  instance.Description != "",
		},
	})
	if err != nil {
		return fmt.Errorf("courseRepo.SaveCourse(): %s", err)
	}
	instance.CourseTemplate.ID = int(savedCourse.ID)
	for _, unit := range instance.Units {
		unit.CourseID = int(savedCourse.ID)
		savedUnit, err := c.SaveUnit(unit)
		if err != nil {
			return fmt.Errorf("error in c.SaveUnit(): %s", err)
		}
		unit = *savedUnit
		log.Println("lesson count: ", len(unit.Lessons), "for ", unit.Name)
		for _, lesson := range unit.Lessons {
			lesson.UnitID = unit.ID
			_, err := c.SaveLessonInstance(lesson)
			if err != nil {
				return fmt.Errorf("error in SaveLessonInstance():%s", err)
			}
		}

	}
	return nil

}

func (c CourseRepo) GetInstances(termID int) ([]domain.CourseInstance, error) {
	dbInstances, err := c.queries.GetInstances(context.Background(), sql.NullInt64{Valid: true, Int64: int64(termID)})
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
	var instances []domain.CourseInstance
	for _, dbInstance := range dbInstances {
		instance := domain.CourseInstance{
			CourseTemplate: domain.CourseTemplate{
				ID:          int(dbInstance.CourseID),
				Name:        dbInstance.CourseName,
				Description: dbInstance.CourseDescr.String,
			},
			Term: term,
		}
		dbUnits, err := c.queries.GetUnits(context.Background(), int64(instance.CourseTemplate.ID))
		if err != nil {
			return nil, err
		}
		var units []domain.Unit
		for _, dbUnit := range dbUnits {
			unit := domain.Unit{
				ID:          int(dbUnit.ID),
				CourseID:    int(dbUnit.CourseID),
				TemplateID:  int(dbUnit.TemplateID.Int64),
				Number:      int(dbUnit.Number),
				SequenceNum: int(dbUnit.Sequence),
				Name:        dbUnit.Name,
				Description: dbUnit.Description.String,
			}
			dbLessons, err := c.queries.GetLessons(context.Background(), int64(unit.ID))
			if err != nil {
				return nil, err
			}
			var lessons []domain.Lesson
			for _, dbLesson := range dbLessons {
				lesson := domain.Lesson{
					ID:          int(dbLesson.ID),
					UnitID:      unit.ID,
					TemplateID:  int(dbLesson.TemplateID.Int64),
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
				lessons = append(lessons, lesson)
			}
			unit.Lessons = lessons
			units = append(units, unit)
		}
		instance.Units = units
		if instance.Term.Start.IsZero() {
			log.Fatal("term not initialized")
		}
		instances = append(instances, instance)

	}
	return instances, nil
}

func (c CourseRepo) GetInstance(templateID int, termID int) (domain.CourseInstance, error) {
	var instance domain.CourseInstance
	dbInstance, err := c.queries.GetInstance(context.Background(), database.GetInstanceParams{
		TermID: sql.NullInt64{
			Valid: termID != 0,
			Int64: int64(termID),
		},
		TemplateID: sql.NullInt64{
			Valid: templateID != 0,
			Int64: int64(templateID),
		},
	})
	if err != nil {
		return instance, nil
	}
	instance = domain.CourseInstance{
		CourseTemplate: domain.CourseTemplate{
			ID:          int(dbInstance.CourseID),
			Name:        dbInstance.CourseName,
			Description: dbInstance.CourseDescr.String,
		},
		TemplateID: templateID,
	}
	dbUnits, err := c.queries.GetUnits(context.Background(), int64(instance.CourseTemplate.ID))
	if err != nil {
		return instance, err
	}
	var units []domain.Unit
	for _, dbUnit := range dbUnits {
		unit := domain.Unit{
			ID:          int(dbUnit.ID),
			CourseID:    int(dbUnit.CourseID),
			TemplateID:  int(dbUnit.TemplateID.Int64),
			Number:      int(dbUnit.Number),
			SequenceNum: int(dbUnit.Sequence),
			Name:        dbUnit.Name,
			Description: dbUnit.Description.String,
		}
		dbLessons, err := c.queries.GetLessons(context.Background(), int64(unit.ID))
		if err != nil {
			return instance, err
		}
		var lessons []domain.Lesson
		for _, dbLesson := range dbLessons {
			lesson := domain.Lesson{
				ID:          int(dbLesson.ID),
				UnitID:      unit.ID,
				TemplateID:  int(dbLesson.TemplateID.Int64),
				Number:      int(dbLesson.Number),
				Name:        dbLesson.Name.String,
				Description: dbLesson.Description.String,
			}
			dbLessonDates, err := c.queries.GetLessonDates(context.Background(), int64(lesson.ID))
			if err != nil {
				return instance, err
			}
			var lessonDates []time.Time
			for _, dbLessonDate := range dbLessonDates {
				lessonDate, err := time.Parse(time.DateOnly, dbLessonDate)
				if err != nil {
					return instance, err
				}
				lessonDates = append(lessonDates, lessonDate)
			}
			lesson.Dates = lessonDates
			lessons = append(lessons, lesson)
		}
		unit.Lessons = lessons
		units = append(units, unit)
	}
	instance.Units = units
	return instance, nil
}

func (c CourseRepo) UpdateInstance(instance domain.CourseInstance) error {
	err := c.queries.UpdateInstance(context.Background(), database.UpdateInstanceParams{
		ID:   int64(instance.CourseTemplate.ID),
		Name: instance.CourseTemplate.Name,
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
