package service

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"log"
	"os"
	"time"
)

type SaveLessonParams struct {
	domain.Lesson
}

func (svc CourseService) AddLessonDate(lessonID, termID int, dateToAdd time.Time) error {
	lesson, err := svc.GetLesson(lessonID)
	if err != nil {
		return err
	}
	term, err := svc.GetTerm(termID)
	if err != nil {
		return err
	}
	// make sure date is an instructional day
	isValidDate := false
	for _, date := range term.InstructionalDays {
		if domain.IsSameDate(date, dateToAdd) {
			isValidDate = true
		}
	}
	if isValidDate {
		lesson.Dates = append(lesson.Dates, dateToAdd)
		svc.UpdateLesson(lesson)
	} else {
		return fmt.Errorf("attempt to add date not in term instructional days: %s, instructional days: %s", dateToAdd, term.InstructionalDays)
	}
	return nil
}

func (svc CourseService) RemoveLessonDate(lessonID int, dateToRemove time.Time) error {
	lesson, err := svc.GetLesson(lessonID)
	if err != nil {
		return err
	}
	for i, date := range lesson.Dates {
		if domain.IsSameDate(date, dateToRemove) {
			lesson.Dates = append(lesson.Dates[:i], lesson.Dates[i+1:]...)
		}
	}
	err = svc.UpdateLesson(lesson)
	if err != nil {
		return err
	}
	return nil
}

func (svc CourseService) SaveLesson(params SaveLessonParams) (*domain.Lesson, error) {
	newLesson := domain.NewLesson(domain.NewLessonParams{
		Lesson: params.Lesson,
	})
	lesson, err := svc.repo.SaveLesson(newLesson)
	if err != nil {
		return &domain.Lesson{}, err
	}
	log.Println("CourseService.SaveLesson: lesson ID: ", lesson.ID)
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
	domain.SortByNumber(lesson.Standards)
	assessments, err := svc.repo.GetLessonAssessments(lessonID)
	if err != nil {
		return domain.Lesson{}, err
	}
	lesson.Assessments = assessments
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
		assessments, err := svc.GetLessonAssessments(lesson.ID)
		if err != nil {
			return nil, err
		}
		lesson.Assessments = assessments
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

func (svc CourseService) DeleteLesson(lessonID int) error {
	lesson, err := svc.GetLesson(lessonID)
	if err != nil {
		return err
	}
	return svc.repo.DeleteLesson(lesson)
}

// returns dates left or right of date
func (svc CourseService) Dates(selectedDate time.Time, dates []time.Time, cd domain.CalendarDirection) []time.Time {
	var filtered []time.Time
	switch cd {
	case domain.Right:
		for _, date := range dates {
			if date.After(selectedDate) {
				filtered = append(filtered, date)
			}
		}
	case domain.Left:
		for _, date := range dates {
			if date.Before(selectedDate) {
				filtered = append(filtered, date)
			}
		}
	}
	return filtered

}

// // all lessons on or after the date provided will be shifted one day left or right
// func (svc CourseService) ShiftAll(termID, courseID int, date time.Time, cd domain.CalendarDirection) error {
// 	term, err := svc.repo.GetTermWithDates(termID)
// 	if err != nil {
// 		return err
// 	}
// 	// get slice of all dates left or right
// 	dates := svc.Dates(date, term.InstructionalDays, cd)
// 	for _, date := range dates {
// 		// get lessons on date
// 		lessons, err := svc.repo.GetLessonsOnDateForCourse(date, courseID)
// 		if err != nil {
// 			return err
// 		}
// 		for _, lesson := range lessons {

// 		}

// 		// shift lessons in direction cd
// 	}

// }

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

func (svc CourseService) Extend(termID, courseID, lessonID int, cd domain.CalendarDirection) error {
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

	_, err = lesson.Extend(cd, term)
	if err != nil {
		return err
	}
	err = svc.UpdateLesson(lesson)
	if err != nil {
		return err
	}
	return nil
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
