package lesson

import (
	"fmt"
	"log"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (svc *Service) Save(lesson Lesson) error {
	_, err := svc.repo.Save(lesson)
	if err != nil {
		return err
	}
	return nil
}

// for date shifts use the Shift method
func (svc *Service) Update(lesson Lesson) error {
	return svc.repo.Update(lesson)
}

func (svc *Service) Shift(lessonID, termID int, days int, oldDate time.Time, instructionalDates []time.Time) error {
	// fetch lesson
	lesson, err := svc.ByID(lessonID)
	if err != nil {
		log.Println("error retrieving lesson")
		return fmt.Errorf("error retrieving lesson with ID %d: %w", lessonID, err)
	}
	// shift date
	newDate, err := lesson.ShiftDate(days, oldDate, instructionalDates)
	if err != nil {
		return fmt.Errorf("error in lesson.ShiftDate: %w", err)
	}
	// persist changes
	err = svc.repo.AddLessonDate(lessonID, termID, newDate)
	if err != nil {
		return fmt.Errorf("error in repo.AddLessonDate: %w", err)
	}
	err = svc.repo.RemoveLessonDate(lessonID, termID, oldDate)
	if err != nil {
		return fmt.Errorf("error in repo.RemoveLessonDate: %w", err)
	}
	return nil
}

func (svc *Service) ByID(lessonID int) (Lesson, error) {
	return svc.repo.ByID(lessonID)
}

func (svc *Service) ByUnitID(unitID int) ([]Lesson, error) {
	return svc.repo.ByUnitID(unitID)
}

func (svc *Service) Delete(lessonID int) error {
	return svc.repo.Delete(lessonID)
}
