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

type courseRepo struct {
	queries *database.Queries
}

// SaveInstance implements domain.CourseRepo.
func (c *courseRepo) SaveInstance(course *domain.CourseInstance) error {
	ctx := context.Background()
	savedCourse, err := c.queries.SaveCourse(ctx, database.SaveCourseParams{
		TemplateID: sql.NullInt64{
			Int64: int64(course.TemplateID),
			Valid: true,
		},
		TermID: sql.NullInt64{
			Int64: int64(course.TermID),
			Valid: false,
		},
		Name: course.Name,
		Description: sql.NullString{
			String: course.Description,
			Valid:  course.Description != "",
		},
	})
	if err != nil {
		return fmt.Errorf("courseRepo.SaveCourse(): %s", err)
	}
	course.ID = int(savedCourse.ID)
	for _, unit := range course.Units {
		unit.CourseID = int(savedCourse.ID)
		log.Println("unit template ID: ", unit.TemplateID)
		log.Println("unit ID: ", unit.ID)
		savedUnit, err := c.SaveUnit(unit)
		log.Println("savedUnit template ID: ", savedUnit.TemplateID)
		log.Println("savedUnit ID: ", savedUnit.ID)
		if err != nil {
			log.Println("savedUnit:", unit.CourseID,
				"\nNumber:", unit.Number,
				"\nName:", unit.Name,
				"\nTemplateID:", unit.TemplateID,
				"\nCourseID:", course.ID,
				"\nCourse TemplateID:", course.TemplateID,
				"\nCourse Name:", course.Name,
				"\nTerm Name:", course.TermName,
			)
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

// GetInstances implements domain.CourseRepository.
func (c *courseRepo) GetInstances() ([]*domain.Course, error) {
	panic("unimplemented")
}

// GetTemplates implements domain.CourseRepository.
func (c *courseRepo) GetTemplates() ([]*domain.Course, error) {
	dbCourses, err := c.queries.GetTemplates(context.Background())
	if err != nil {
		return nil, err
	}
	var courses []*domain.Course
	for _, dbCourse := range dbCourses {
		course := &domain.Course{
			ID:          int(dbCourse.CourseID),
			Name:        dbCourse.CourseName,
			Description: dbCourse.CourseDescr.String,
		}
		dbUnits, err := c.queries.GetUnits(context.Background(), dbCourse.CourseID)
		if err != nil {
			return nil, err
		}
		var units []domain.Unit
		for _, dbUnit := range dbUnits {
			unit := domain.Unit{
				ID:          int(dbUnit.ID),
				CourseID:    int(dbCourse.CourseID),
				Number:      int(dbUnit.Number),
				Name:        dbUnit.Name,
				Description: dbUnit.Description.String,
			}
			dbLessons, err := c.queries.GetLessons(context.Background(), dbUnit.ID)
			if err != nil {
				return nil, err
			}
			var lessons []domain.Lesson
			for _, dbLesson := range dbLessons {
				lesson := domain.Lesson{
					ID:          int(dbLesson.ID),
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
		courses = append(courses, course)
	}
	return courses, nil
}

func NewCourseRepo(db *database.Queries) domain.CourseRepo {
	return &courseRepo{queries: db}
}

func (c *courseRepo) All() ([]*domain.Course, error) {
	panic("not implemented")
}

// Save updates the ID for Course, Units, and Lessons
func (c *courseRepo) SaveTemplate(course *domain.Course) (*domain.Course, error) {
	ctx := context.Background()
	dbCourse, err := c.queries.SaveCourse(ctx, database.SaveCourseParams{
		TemplateID: sql.NullInt64{Valid: false},
		TermID:     sql.NullInt64{Valid: false},
		Name:       course.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("c.queries.SaveCourse(): %s", err)
	}
	course.ID = int(dbCourse.ID)
	for i, unit := range course.Units {
		unit.CourseID = course.ID
		savedUnit, err := c.SaveUnit(unit)
		if err != nil {
			return nil, fmt.Errorf("c.SaveUnit(): %s", err)
		}
		course.Units[i] = *savedUnit
		unit = *savedUnit
		for j, lesson := range unit.Lessons {
			lesson.UnitID = unit.ID
			savedLesson, err := c.SaveLessonTemplate(lesson)
			if err != nil {
				return nil, fmt.Errorf("c.SaveLessonTemplate(): %s", err)
			}
			unit.Lessons[j] = *savedLesson
		}

	}
	return course, nil

}

func (c *courseRepo) SaveUnit(unit domain.Unit) (*domain.Unit, error) {
	log.Println("SaveUnit(): ", "templateID", unit.TemplateID, "ID", unit.ID)
	var hasDescr = unit.Description == ""
	currUnit := database.Unit{
		CourseID: int64(unit.CourseID),
		TemplateID: sql.NullInt64{
			Int64: int64(unit.TemplateID),
			Valid: unit.TemplateID != 0,
		},
		Number:   int64(unit.Number),
		Sequence: int64(unit.SequenceNum),
		Name:     unit.Name,
		Description: sql.NullString{
			String: unit.Description,
			Valid:  hasDescr,
		},
	}
	if currUnit.Sequence == 0 {
		return nil, fmt.Errorf("currUnit sequence is 0")
	}
	currUnit, err := c.queries.SaveUnit(context.Background(), database.SaveUnitParams{
		Number:      currUnit.Number,
		Sequence:    currUnit.Sequence,
		TemplateID:  currUnit.TemplateID,
		Name:        currUnit.Name,
		Description: currUnit.Description,
		CourseID:    currUnit.CourseID,
	})
	if err != nil {
		return nil, fmt.Errorf("courseRepo.SaveUnit(): %s", err)
	}
	unit.ID = int(currUnit.ID)
	log.Println("unit sequence:", unit.SequenceNum)
	if unit.SequenceNum == 0 {
		return nil, fmt.Errorf("unit sequence is 0")
	}
	return &unit, nil

}

// Should include date
func (c *courseRepo) SaveLessonInstance(lesson domain.Lesson) (*domain.Lesson, error) {
	if lesson.TemplateID == 0 {
		return nil, fmt.Errorf("lesson instance template ID is 0")
	}
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

// fetches the date id for a
func (c *courseRepo) SaveLessonDate(lesson domain.Lesson) error {
	dbDate, err := c.queries.GetDate(context.Background(), lesson.Date.Format(time.DateOnly))
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
	return nil
}

func (c *courseRepo) SaveLesson(lesson domain.Lesson) (*domain.Lesson, error) {
	log.Println("lesson name:", lesson.Name, "lesson number: ", lesson.Number)
	dbLesson := database.Lesson{
		UnitID: int64(lesson.UnitID),
		TemplateID: sql.NullInt64{
			Int64: int64(lesson.TemplateID),
			Valid: lesson.TemplateID != 0,
		},
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
		TemplateID:  dbLesson.TemplateID,
		Description: dbLesson.Description,
		UnitID:      dbLesson.UnitID,
	})
	lesson.ID = int(savedLesson.ID)
	if err != nil {
		log.Println(
			"Lesson Unit ID:", lesson.UnitID,
			"\nLesson Template ID:", lesson.TemplateID,
			"\nLesson ID:", lesson.ID,
		)
		return nil, fmt.Errorf("c.queries.SaveLesson(): %v, unit id: %s, lesson #: %s",
			err, strconv.Itoa(int(savedLesson.UnitID)), strconv.Itoa(int(dbLesson.Number)),
		)
	}
	return &lesson, nil

}

// Should not include date
func (c *courseRepo) SaveLessonTemplate(lesson domain.Lesson) (*domain.Lesson, error) {
	savedLesson, err := c.SaveLesson(lesson)
	if err != nil {
		return nil, fmt.Errorf("c.SaveLesson():%s", err)
	}
	return savedLesson, nil
}
