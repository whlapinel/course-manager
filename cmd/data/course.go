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

type CourseRepo struct {
	queries *database.Queries
}

func NewCourseRepo(db *database.Queries) CourseRepo {
	return CourseRepo{queries: db}
}

// SaveInstance implements domain.CourseRepo.
func (c CourseRepo) SaveInstance(instance *domain.CourseInstance) error {
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
				"\nCourseID:", instance.CourseTemplate.ID,
				"\nCourse TemplateID:", instance.TemplateID,
				"\nCourse Name:", instance.CourseTemplate.Name,
				"\nTerm Name:", instance.Term.Name,
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
func (c CourseRepo) GetInstances(term domain.Term) ([]domain.CourseInstance, error) {
	dbInstances, err := c.queries.GetInstances(context.Background(), sql.NullInt64{Valid: true, Int64: int64(term.ID)})
	if err != nil {
		return nil, err
	}
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

func (c CourseRepo) GetTemplate(id int) (domain.CourseTemplate, error) {
	dbCourse, err := c.queries.GetTemplate(context.Background(), int64(id))
	if err != nil {
		return domain.CourseTemplate{}, err
	}
	course := domain.CourseTemplate{
		ID:          int(dbCourse.CourseID),
		Name:        dbCourse.CourseName,
		Description: dbCourse.CourseDescr.String,
	}
	dbUnits, err := c.queries.GetUnits(context.Background(), dbCourse.CourseID)
	if err != nil {
		return domain.CourseTemplate{}, err
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
			return domain.CourseTemplate{}, err
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

// GetTemplates implements domain.CourseRepository.
func (c CourseRepo) GetTemplates() ([]domain.CourseTemplate, error) {
	dbCourses, err := c.queries.GetTemplates(context.Background())
	if err != nil {
		return nil, err
	}
	var courses []domain.CourseTemplate
	for _, dbCourse := range dbCourses {
		course := domain.CourseTemplate{
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
				CourseID:    int(dbUnit.CourseID),
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
		courses = append(courses, course)
	}
	return courses, nil
}

// Save updates the ID for Course, Units, and Lessons
func (c CourseRepo) SaveTemplate(course domain.CourseTemplate) (domain.CourseTemplate, error) {
	ctx := context.Background()
	dbCourse, err := c.queries.SaveTemplate(ctx, database.SaveTemplateParams{
		Name: course.Name,
		Description: sql.NullString{
			Valid:  course.Description != "",
			String: course.Description,
		},
	})
	if err != nil {
		return domain.CourseTemplate{}, fmt.Errorf("c.queries.SaveCourse(): %s", err)
	}
	course.ID = int(dbCourse.ID)
	for i, unit := range course.Units {
		unit.CourseID = course.ID
		savedUnit, err := c.SaveUnit(unit)
		if err != nil {
			return domain.CourseTemplate{}, fmt.Errorf("c.SaveUnit(): %s", err)
		}
		course.Units[i] = *savedUnit
		unit = *savedUnit
		for j, lesson := range unit.Lessons {
			lesson.UnitID = unit.ID
			savedLesson, err := c.SaveLessonTemplate(lesson)
			if err != nil {
				return domain.CourseTemplate{}, fmt.Errorf("c.SaveLessonTemplate(): %s", err)
			}
			unit.Lessons[j] = *savedLesson
		}

	}
	return course, nil

}

func (c CourseRepo) SaveUnit(unit domain.Unit) (*domain.Unit, error) {
	log.Println("SaveUnit(): ", "templateID", unit.TemplateID, "ID", unit.ID)
	var hasDescr = unit.Description != ""
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
func (c CourseRepo) SaveLessonInstance(lesson domain.Lesson) (*domain.Lesson, error) {
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
func (c CourseRepo) SaveLessonTemplate(lesson domain.Lesson) (*domain.Lesson, error) {
	savedLesson, err := c.SaveLesson(lesson)
	if err != nil {
		return nil, fmt.Errorf("c.SaveLesson():%s", err)
	}
	return savedLesson, nil
}
