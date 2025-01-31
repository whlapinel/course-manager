package service

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"log"
	"os"
)

type SaveLessonParams struct {
	domain.Lesson
}

func (svc CourseService) SaveLesson(params SaveLessonParams) (*domain.Lesson, error) {
	newLesson := domain.NewLesson(domain.NewLessonParams{
		Lesson: params.Lesson,
	})
	lesson, err := svc.repo.SaveLesson(newLesson)
	if err != nil {
		return &domain.Lesson{}, err
	}
	return lesson, nil
}

func (svc CourseService) CreateNewLessonSlides(nodes ...domain.CourseNode) error {
	path := data.SlidesMarkdownFilePath(nodes...)
	// make sure file does not exist
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("CourseService.CreateNewLessonSlides: file exists already at %s", path)
	}
	_, err = os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func (svc CourseService) CreateNewLessonFileDir(nodes ...domain.CourseNode) error {
	path := data.NodeFilesDirPath(nodes...)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (svc CourseService) GetLesson(lessonID int) (domain.Lesson, error) {
	lesson, err := svc.repo.GetLesson(lessonID)
	if err != nil {
		return domain.Lesson{}, fmt.Errorf("CourseService.GetLesson: %d", lessonID)
	}
	lessonDates, err := svc.repo.GetLessonDates(lessonID)
	if err != nil {
		return domain.Lesson{}, err
	}
	lesson.Dates = lessonDates
	objectives, err := svc.repo.GetLessonObjectives(lesson)
	if err != nil {
		return domain.Lesson{}, err
	}
	lesson.Standards = objectives
	return lesson, nil

}

func (svc CourseService) GetLessons(unitID int) ([]domain.Lesson, error) {
	lessons, err := svc.repo.GetLessons(unitID)
	if err != nil {
		return nil, err
	}
	for i, lesson := range lessons {
		lessonDates, err := svc.repo.GetLessonDates(lesson.ID)
		if err != nil {
			return nil, err
		}
		lesson.Dates = lessonDates
		objectives, err := svc.repo.GetLessonObjectives(lesson)
		if err != nil {
			return nil, err
		}
		lesson.Standards = objectives
		lessons[i] = lesson
	}
	return lessons, nil
}

func (svc CourseService) UpdateLesson(l domain.Lesson) error {
	log.Println("CourseService.UpdateLesson: ")
	log.Println("Descr:", l.Description)
	err := svc.repo.UpdateLesson(l)
	if err != nil {
		return err
	}
	return nil
}

// This version is for the web app
func (svc CourseService) WebShift(termID, courseID, lessonID int, cd domain.CalendarDirection) error {
	lesson, err := svc.GetLesson(lessonID)
	if err != nil {
		return err
	}
	term, err := svc.repo.GetTermByID(termID)
	if err != nil {
		return err
	}
	termWithDates, err := svc.repo.GetTermWithDates(termID)
	if err != nil {
		return err
	}
	term.InstructionalDays = termWithDates.InstructionalDays
	newLesson, _, err := lesson.Shift(cd, term)
	if err != nil {
		return err
	}
	err = svc.UpdateLesson(newLesson)
	if err != nil {
		return err
	}
	return err
}

func (svc CourseService) Extend(lesson domain.Lesson, term domain.Term, direction domain.CalendarDirection) (domain.Lesson, error) {
	shifted, err := lesson.Extend(direction, term)
	if err != nil {
		return domain.Lesson{}, err
	}
	err = svc.UpdateLesson(lesson)
	if err != nil {
		return domain.Lesson{}, err
	}
	return shifted, nil
}

func (svc CourseService) SetLessonObjective(lessonID, stdID int) error {
	lesson, err := svc.repo.GetLesson(lessonID)
	if err != nil {
		return err
	}
	objective, err := svc.repo.GetStandardByID(stdID)
	if err != nil {
		return err
	}
	return svc.repo.SaveLessonStandard(lesson, objective)
}

func (svc CourseService) DeleteLessonObjective(lessonID, stdID int) error {
	lesson, err := svc.GetLesson(lessonID)
	if err != nil {
		return err
	}
	objective, err := svc.repo.GetStandardByID(stdID)
	if err != nil {
		return err
	}
	return svc.repo.DeleteLessonStandard(lesson, objective)
}
