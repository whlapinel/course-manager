package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
	"strconv"
	"time"
)

func (lr CourseRepo) GetLessons(unitID int) ([]domain.Lesson, error) {
	dbLessons, err := lr.queries.GetLessons(context.Background(), int64(unitID))
	if err != nil {
		return nil, err
	}
	var lessons []domain.Lesson
	for _, dbLesson := range dbLessons {
		lesson := domain.Lesson{
			ID:          int(dbLesson.ID),
			UnitID:      unitID,
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

func (lr CourseRepo) GetLessonDates(lessonID int) ([]time.Time, error) {
	dbDates, err := lr.queries.GetLessonDates(context.Background(), int64(lessonID))
	if err != nil {
		return nil, err
	}
	var dates []time.Time
	for _, dbDate := range dbDates {
		date, err := time.Parse(time.DateOnly, dbDate)
		if err != nil {
			return nil, err
		}
		dates = append(dates, date)

	}
	return dates, nil
}
func (lr CourseRepo) AddLessonDate(lesson domain.Lesson, date time.Time) error {
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

func (lr CourseRepo) DeleteLessonDate(lesson domain.Lesson, date time.Time) error {
	id, err := lr.queries.GetDateID(context.Background(), date.Format(time.DateOnly))
	if err != nil {
		return err
	}
	err = lr.queries.DeleteLessonDate(context.Background(), database.DeleteLessonDateParams{
		LessonID: int64(lesson.ID),
		DateID:   id,
	})
	return err
}

func (c CourseRepo) SaveLessonInstance(lesson domain.Lesson) (*domain.Lesson, error) {
	savedLesson, err := c.SaveLesson(lesson)
	if err != nil {
		return nil, fmt.Errorf("c.SaveLesson():%s", err)
	}
	lesson = *savedLesson
	err = c.SaveLessonDate(lesson)
	if err != nil {
		return nil, fmt.Errorf("c.SaveLessonDate(): %s", err)
	}
	return &lesson, nil
}

func (c CourseRepo) SaveLessonDate(lesson domain.Lesson) error {
	for _, date := range lesson.Dates {
		dbDate, err := c.queries.GetDate(context.Background(), date.Format(time.DateOnly))
		if err != nil {
			return fmt.Errorf("courseRepo.SaveInstance(), c.queries.GetDate(): %s", err)
		}
		log.Println("Saved date: ID:", dbDate.ID, "\nDay Number:", dbDate.DayNumber, "\nTerm ID:", dbDate.TermID)
		lessonDate, err := c.queries.SaveLessonDate(context.Background(), database.SaveLessonDateParams{
			LessonID: int64(lesson.ID),
			DateID:   dbDate.ID,
		})
		if err != nil {
			return fmt.Errorf("courseRepo.SaveInstance(), c.queries.SaveLessonDate: %s", err)
		}
		log.Println("saved lessonDate: \nDate ID:", lessonDate.DateID, "\nLesson ID:", lessonDate.LessonID)
	}
	return nil
}

func (c CourseRepo) SaveLesson(lesson domain.Lesson) (*domain.Lesson, error) {
	log.Println("lesson name:", lesson.Name, "lesson number: ", lesson.Number)
	dbLesson := database.Lesson{
		UnitID: int64(lesson.UnitID),
		Number: int64(lesson.Number),
		Name: sql.NullString{
			String: lesson.Name,
			Valid:  lesson.Name != "",
		},
		Description: sql.NullString{
			String: lesson.Description,
			Valid:  lesson.Description != "",
		},
	}
	savedLesson, err := c.queries.SaveLesson(context.Background(), database.SaveLessonParams{
		Number:      dbLesson.Number,
		Name:        dbLesson.Name,
		Description: dbLesson.Description,
		UnitID:      dbLesson.UnitID,
	})
	lesson.ID = int(savedLesson.ID)
	if err != nil {
		return nil, fmt.Errorf("c.queries.SaveLesson(): %v, unit id: %s, lesson #: %s",
			err, strconv.Itoa(int(savedLesson.UnitID)), strconv.Itoa(int(dbLesson.Number)),
		)
	}
	return &lesson, nil

}

func (c CourseRepo) SaveLessonTemplate(lesson domain.Lesson) (*domain.Lesson, error) {
	savedLesson, err := c.SaveLesson(lesson)
	if err != nil {
		return nil, fmt.Errorf("c.SaveLesson():%s", err)
	}
	return savedLesson, nil
}

func (lr CourseRepo) UpdateLesson(lesson domain.Lesson) error {
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
	if err != nil {
		return err
	}
	err = lr.queries.DeleteAllLessonDates(context.Background(), int64(lesson.ID))
	if err != nil {
		return err
	}
	for _, date := range lesson.Dates {
		id, err := lr.queries.GetDateID(context.Background(), date.Format(time.DateOnly))
		if err != nil {
			return err
		}
		_, err = lr.queries.SaveLessonDate(context.Background(), database.SaveLessonDateParams{
			LessonID: int64(lesson.ID),
			DateID:   id,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (lr CourseRepo) DeleteLesson(lesson domain.Lesson) error {
	err := lr.queries.DeleteLesson(context.Background(), int64(lesson.ID))
	return err
}
