package data

import (
	"context"
	"database/sql"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"time"
)

type LessonRepo struct {
	queries *database.Queries
}

func NewLessonRepo(queries *database.Queries) LessonRepo {
	return LessonRepo{queries: queries}
}

func (lr LessonRepo) All(unitID int) ([]domain.Lesson, error) {
	dbLessons, err := lr.queries.GetLessons(context.Background(), int64(unitID))
	if err != nil {
		return nil, err
	}
	var lessons []domain.Lesson
	for _, dbLesson := range dbLessons {
		lesson := domain.Lesson{
			ID:          int(dbLesson.ID),
			UnitID:      unitID,
			TemplateID:  int(dbLesson.TemplateID.Int64),
			Number:      int(dbLesson.Number),
			Name:        dbLesson.Name.String,
			Description: dbLesson.Description.String,
		}
		dbLessonDates, err := lr.queries.GetLessonDates(context.Background(), int64(lesson.ID))
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
	return lessons, nil
}

func (lr LessonRepo) Cancel(lesson domain.Lesson, date time.Time) error {
	err := lr.queries.DeleteLessonDates(context.Background(), int64(lesson.ID))
	return err
}

func (lr LessonRepo) Schedule(lesson domain.Lesson, date time.Time) error {
	dbDate, err := lr.queries.GetDate(context.Background(), date.Format(time.DateOnly))
	if err != nil {
		return err
	}
	lr.queries.SaveLessonDate(context.Background(), database.SaveLessonDateParams{
		LessonID: int64(lesson.ID),
		DateID:   dbDate.ID,
	})
	return nil
}

func (lr LessonRepo) Update(lesson domain.Lesson) error {
	err := lr.queries.UpdateLesson(context.Background(), database.UpdateLessonParams{
		Name: sql.NullString{
			Valid:  lesson.Name != "",
			String: lesson.Name,
		},
		Number: int64(lesson.Number),
		Description: sql.NullString{
			Valid:  lesson.Description != "",
			String: lesson.Description,
		},
	})
	return err
}

func (lr LessonRepo) Delete(lesson domain.Lesson) error {
	err := lr.queries.DeleteLesson(context.Background(), int64(lesson.ID))
	return err
}
