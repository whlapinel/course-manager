package data

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

const (
	courseNameCol = iota
	dayNumCol
	unitNumCol
	unitSequenceCol
	unitDescrCol
	lessonNumCol
	lessonDescrCol
	stdNumCol
	stdDescrCol
	scheduleDateCol
	scheduleTermIdCol
	scheduleTermNameCol
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

func (c *courseRepo) WriteToCSV(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

// TODO: Convert this to import directly from CSV rather than using instance import
func (c *courseRepo) ReadFromCSV() ([]*domain.Course, error) {
	courses, err := importCoursesFromCSV()
	if err != nil {
		return nil, err
	}
	return courses, nil

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

func importCoursesFromCSV() ([]*domain.Course, error) {
	file, err := os.Open(scheduleCsvDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		log.Fatalf("file is empty")
	}
	type UnitMap map[int]domain.Unit
	type CourseHolder struct {
		Course domain.Course
		Units  UnitMap
	}
	courseMap := make(map[string]CourseHolder)
	for _, record := range records[1:] {
		courseName := record[courseNameCol]
		courseHolder, exists := courseMap[courseName]
		termName := record[scheduleTermNameCol]
		if !exists {
			course := domain.Course{
				Name:     courseName,
				TermName: termName,
			}
			holder := CourseHolder{
				Course: course,
				Units:  UnitMap{},
			}
			courseMap[courseName] = holder
			courseHolder = holder
		}
		unitNum := 0
		if record[unitNumCol] != "" {
			unitNum, err = strconv.Atoi(record[unitNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading unit number from csv")
			}
		} else {
			return nil, fmt.Errorf("unit number field is blank")
		}
		unitSequence, err := strconv.Atoi(record[unitSequenceCol])
		if err != nil {
			return nil, fmt.Errorf("error reading unit sequence number from csv")
		}
		unitName := fmt.Sprintf("Unit %d", unitNum)
		if unitNum < 0 {
			unitName = record[unitDescrCol]
		}
		unitDescr := record[unitDescrCol]
		unit, exists := courseHolder.Units[unitSequence]
		if !exists {
			courseHolder.Units[unitSequence] = domain.Unit{
				Number:      unitNum,
				SequenceNum: unitSequence,
				Name:        unitName,
				Description: unitDescr,
			}
			unit = courseHolder.Units[unitSequence]
		}
		lessonNum := 0
		if record[lessonNumCol] != "" {
			lessonNum, err = strconv.Atoi(record[lessonNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading lesson number from csv")
			}
		}
		lessonName := fmt.Sprintf("Lesson %d.%d", unitNum, lessonNum)
		if unitNum < 0 {
			lessonName = fmt.Sprintf("%s Day %d", unitName, lessonNum)

		}
		lessonDescr := record[lessonDescrCol]
		lessonDate, err := time.Parse(time.DateOnly, record[scheduleDateCol])
		if err != nil {
			return nil, fmt.Errorf("error parsing date: %s", err)
		}
		lesson := domain.NewLesson(lessonNum, unit.ID, lessonName, lessonDescr, lessonDate)
		unit.Lessons = append(unit.Lessons, lesson)
		courseHolder.Units[unitSequence] = unit
		courseMap[courseName] = courseHolder
	}
	var courses []*domain.Course
	for _, holder := range courseMap {
		course := holder.Course
		nums := sortedBySequence(holder.Units)
		for _, num := range nums {
			course.Units = append(course.Units, holder.Units[num])
		}
		courses = append(courses, &course)

	}
	return courses, nil
}

func sortedBySequence(unitMap map[int]domain.Unit) []int {
	keys := make([]int, 0, len(unitMap))
	for sequence := range unitMap {
		keys = append(keys, sequence)
	}

	slices.Sort(keys)
	return keys

}

// This imports a course template and a term and generates a course instance
func GenerateCourseInstancesFromCSV2(date time.Time) ([]*domain.Course, error) {
	courses, err := importCoursesFromCSV()
	if err != nil {
		return nil, err
	}
	terms, err := TermsLoader()
	if err != nil {
		return nil, err
	}
	var currentTerm *domain.Term
	for _, term := range terms {
		if date.After(term.Start) && date.Before(term.End) {
			currentTerm = term
		}
	}
	for i, course := range courses {
		dateNum := 0
		currDate := currentTerm.InstructionalDays[dateNum]
		courses[i].TermName = currentTerm.Name
		for j, unit := range course.Units {
			for k, lesson := range unit.Lessons {
				log.Printf("Assigning %v to lesson %v", currDate, lesson.Name) // Log assignment
				courses[i].Units[j].Lessons[k].Date = currDate
				if dateNum != len(currentTerm.InstructionalDays)-1 {
					dateNum++
					currDate = currentTerm.InstructionalDays[dateNum]
				} else {
					currDate = time.Time{}
				}
			}
		}
	}
	return courses, nil
}

func WriteCourseInstancesToCSV(instances []*domain.Course) error {
	file, err := os.Create(newScheduleDir)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	var rows [][]string
	rows = append(rows, []string{
		"course_name",
		"day_num",
		"unit_num",
		"unit_descr",
		"lesson_num",
		"lesson_descr",
		"std_num",
		"std_descr",
		"date",
		"term_id",
		"term_name",
	})
	for _, instance := range instances {
		dayNum := 0
		for _, unit := range instance.Units {
			for _, lesson := range unit.Lessons {
				dayNum++
				courseName := instance.Name
				unitNum := unit.Number
				unitDescr := unit.Description
				lessonNum := lesson.Number
				stdNum := ""
				stdDescr := ""
				date := lesson.Date.Format(time.DateOnly)
				if lesson.Date.IsZero() {
					date = ""
				}
				termID := ""
				termName := instance.TermName
				row := []string{
					courseName,
					strconv.Itoa(dayNum),
					strconv.Itoa(unitNum),
					unitDescr,
					strconv.Itoa(lessonNum),
					lesson.Description,
					stdNum,
					stdDescr,
					date,
					termID,
					termName,
				}
				rows = append(rows, row)

			}

		}
	}
	writer.WriteAll(rows)
	return nil

}
