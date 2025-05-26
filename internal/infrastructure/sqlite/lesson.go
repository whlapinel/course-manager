package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/internal/features/lesson"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
	"gh_static_portfolio/internal/shared/dates"
	"time"
)

type lessonRepo struct {
	queries *database.Queries
}

func NewLessonRepo(queries *database.Queries) lesson.Repository {
	return &lessonRepo{
		queries: queries,
	}

}

// ByID implements lesson.Repository.
func (l *lessonRepo) ByID(lessonID int) (lesson.Lesson, error) {
	dbLesson, err := l.queries.GetLesson(context.Background(), int64(lessonID))
	if err != nil {
		return lesson.Lesson{}, err
	}
	dbDates, err := l.queries.GetLessonDates(context.Background(), int64(lessonID))
	if err != nil {
		return lesson.Lesson{}, err
	}
	var lessonDates []time.Time
	for _, dbDate := range dbDates {
		date, err := time.Parse(time.DateOnly, dbDate)
		if err != nil {
			return lesson.Lesson{}, err
		}
		lessonDates = append(lessonDates, date)
	}
	dates.Sort(lessonDates)
	return lesson.Lesson{
		ID:          int(dbLesson.ID),
		UnitID:      int(dbLesson.UnitID),
		UnitNumber:  int(dbLesson.UnitNumber),
		UnitName:    dbLesson.UnitName,
		Number:      int(dbLesson.Number),
		Name:        dbLesson.Name.String,
		Description: dbLesson.Description.String,
		Dates:       lessonDates,
	}, nil
}

// ByUnitID implements lesson.Repository.
func (l *lessonRepo) ByUnitID(unitID int) ([]lesson.Lesson, error) {
	dbLessons, err := l.queries.GetLessons(context.Background(), int64(unitID))
	if err != nil {
		return nil, err
	}
	var lessons []lesson.Lesson
	for _, dbLesson := range dbLessons {
		lesson := lesson.Lesson{
			ID:          int(dbLesson.ID),
			UnitID:      unitID,
			Number:      int(dbLesson.Number),
			Name:        dbLesson.Name.String,
			Description: dbLesson.Description.String,
			UnitNumber:  int(dbLesson.UnitNumber),
			UnitName:    dbLesson.UnitName,
		}
		dbLessonDates, err := l.queries.GetLessonDates(context.Background(), int64(lesson.ID))
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
		dates.Sort(lessonDates)
		lesson.Dates = lessonDates
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

// Delete implements lesson.Repository.
func (l *lessonRepo) Delete(lessonID int) error {
	err := l.queries.DeleteLesson(context.Background(), int64(lessonID))
	if err != nil {
		return err
	}
	return nil
}

// Save implements lesson.Repository.
func (l *lessonRepo) Save(newLesson lesson.Lesson) (int, error) {
	lessonParams := database.SaveLessonParams{
		UnitID: int64(newLesson.UnitID),
		Name: sql.NullString{
			Valid:  newLesson.Name != "",
			String: newLesson.Name,
		},

		Description: sql.NullString{
			Valid:  newLesson.Description != "",
			String: newLesson.Description,
		},
	}
	dbLesson, err := l.queries.SaveLesson(context.Background(), lessonParams)
	if err != nil {
		return 0, fmt.Errorf("lessonRepo.SaveLesson: %s", err)
	}
	return int(dbLesson.ID), nil
}

// Update implements lesson.Repository.
func (l *lessonRepo) Update(updated lesson.Lesson) error {
	panic("unimplemented")
}
