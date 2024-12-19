package data

import (
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

type CSVHeader int

const (
	courseNameCol CSVHeader = iota
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

func (c CSVHeader) String() string {
	return []string{
		"course_name",
		"day_num",
		"unit_num",
		"unit_sequence",
		"unit_descr",
		"lesson_num",
		"lesson_descr",
		"std_num",
		"std_descr",
		"date",
		"term_id",
		"term_name",
	}[c]
}

func CSVHeaders() []string {
	return []string{
		courseNameCol.String(),
		dayNumCol.String(),
		unitNumCol.String(),
		unitDescrCol.String(),
		lessonNumCol.String(),
		lessonDescrCol.String(),
		stdNumCol.String(),
		stdDescrCol.String(),
		scheduleDateCol.String(),
		scheduleTermIdCol.String(),
		scheduleTermNameCol.String(),
	}

}

func (c CourseRepo) WriteToCSV(course domain.CourseTemplate) error {
	return fmt.Errorf("not implemented")
}

type LessonMap map[int]domain.Lesson
type UnitHolder struct {
	Unit    domain.Unit
	Lessons LessonMap
}
type UnitMap map[int]UnitHolder
type CourseInstanceHolder struct {
	Instance domain.CourseInstance
	Units    UnitMap
}
type CourseTemplateHolder struct {
	Template domain.CourseTemplate
	Units    UnitMap
}
type InstanceMap map[string]CourseInstanceHolder
type TemplateMap map[string]CourseTemplateHolder

func (c CourseRepo) ImportTemplatesFromCSV() ([]domain.CourseTemplate, error) {
	return importTemplatesFromCSV()
}

func importTemplatesFromCSV() ([]domain.CourseTemplate, error) {
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
	courseMap := TemplateMap{}
	for _, record := range records[1:] {
		courseName := record[courseNameCol]
		courseHolder, exists := courseMap[courseName]
		if !exists {
			template := domain.CourseTemplate{
				Name: courseName,
			}
			holder := CourseTemplateHolder{
				Template: template,

				Units: UnitMap{},
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
			unit = UnitHolder{
				Unit: domain.Unit{
					Number:      unitNum,
					SequenceNum: unitSequence,
					Name:        unitName,
					Description: unitDescr,
				},
				Lessons: LessonMap{},
			}
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
		lesson, exists := unit.Lessons[lessonNum]
		if !exists {
			lesson = domain.NewLesson(lessonNum, unit.Unit.ID, lessonName, lessonDescr, []time.Time{lessonDate})
		} else {
			lesson.Dates = append(lesson.Dates, lessonDate)
		}
		unit.Lessons[lessonNum] = lesson
		courseHolder.Units[unitSequence] = unit
		courseMap[courseName] = courseHolder
	}
	var courses []domain.CourseTemplate
	for _, courseHolder := range courseMap {
		course := courseHolder.Template
		unitNums := sortUnitMapKeys(courseHolder.Units)
		for _, unitNum := range unitNums {
			unit := courseHolder.Units[unitNum].Unit
			unitHolder := courseHolder.Units[unitNum]
			lessonNums := sortLessonMapKeys(unitHolder.Lessons)
			for _, lessonNum := range lessonNums {
				lesson := unitHolder.Lessons[lessonNum]
				unit.Lessons = append(unit.Lessons, lesson)
			}
			course.Units = append(course.Units, unit)
		}
		courses = append(courses, course)

	}
	return courses, nil
}
func importInstancesFromCSV() ([]domain.CourseInstance, error) {
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
	courseMap := InstanceMap{}
	for _, record := range records[1:] {
		courseName := record[courseNameCol]
		termName := record[scheduleTermNameCol]
		courseHolder, exists := courseMap[courseName]
		if !exists {
			instance := domain.CourseInstance{
				CourseTemplate: domain.CourseTemplate{
					Name: courseName,
				},
				Term: domain.Term{
					Name: termName,
				},
			}
			holder := CourseInstanceHolder{
				Instance: instance,
				Units:    UnitMap{},
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
			unit = UnitHolder{
				Unit: domain.Unit{
					Number:      unitNum,
					SequenceNum: unitSequence,
					Name:        unitName,
					Description: unitDescr,
				},
				Lessons: LessonMap{},
			}
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
		lesson, exists := unit.Lessons[lessonNum]
		if !exists {
			lesson = domain.NewLesson(lessonNum, unit.Unit.ID, lessonName, lessonDescr, []time.Time{lessonDate})
		} else {
			lesson.Dates = append(lesson.Dates, lessonDate)
		}
		unit.Lessons[lessonNum] = lesson
		courseHolder.Units[unitSequence] = unit
		courseMap[courseName] = courseHolder
	}
	var courses []domain.CourseInstance
	for _, courseHolder := range courseMap {
		course := courseHolder.Instance
		unitNums := sortUnitMapKeys(courseHolder.Units)
		for _, unitNum := range unitNums {
			unit := courseHolder.Units[unitNum].Unit
			unitHolder := courseHolder.Units[unitNum]
			lessonNums := sortLessonMapKeys(unitHolder.Lessons)
			for _, lessonNum := range lessonNums {
				lesson := unitHolder.Lessons[lessonNum]
				unit.Lessons = append(unit.Lessons, lesson)
			}
			course.Units = append(course.Units, unit)
		}
		courses = append(courses, course)

	}
	return courses, nil
}

func sortUnitMapKeys(unitMap UnitMap) []int {
	keys := make([]int, 0, len(unitMap))
	for sequence := range unitMap {
		keys = append(keys, sequence)
	}

	slices.Sort(keys)
	return keys

}
func sortLessonMapKeys(lessonMap LessonMap) []int {
	keys := make([]int, 0, len(lessonMap))
	for sequence := range lessonMap {
		keys = append(keys, sequence)
	}

	slices.Sort(keys)
	return keys

}

// This imports a course template and a term and generates a course instance
func GenerateCourseInstancesFromCSV2(date time.Time) ([]domain.CourseInstance, error) {
	instances, err := importInstancesFromCSV()
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
	for i, course := range instances {
		dateNum := 0
		currDate := currentTerm.InstructionalDays[dateNum]
		instances[i].Term.Name = currentTerm.Name
		for j, unit := range course.Units {
			for k, lesson := range unit.Lessons {
				log.Printf("Assigning %v to lesson %v", currDate, lesson.Name) // Log assignment
				instances[i].Units[j].Lessons[k].Dates[0] = currDate
				if dateNum != len(currentTerm.InstructionalDays)-1 {
					dateNum++
					currDate = currentTerm.InstructionalDays[dateNum]
				} else {
					currDate = time.Time{}
				}
			}
		}
	}
	return instances, nil
}

func WriteCourseInstancesToCSV(instances []domain.CourseInstance) error {
	file, err := os.Create(newScheduleDir)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	var rows [][]string
	rows = append(rows, CSVHeaders())
	for _, instance := range instances {
		dayNum := 0
		for _, unit := range instance.Units {
			for _, lesson := range unit.Lessons {
				dayNum++
				courseName := instance.CourseTemplate.Name
				unitNum := unit.Number
				unitDescr := unit.Description
				lessonNum := lesson.Number
				stdNum := ""
				stdDescr := ""
				for _, date := range lesson.Dates {
					dateString := date.Format(time.DateOnly)
					if date.IsZero() {
						dateString = ""
					}
					termID := ""
					termName := instance.Term.Name
					row := []string{
						courseName,
						strconv.Itoa(dayNum),
						strconv.Itoa(unitNum),
						unitDescr,
						strconv.Itoa(lessonNum),
						lesson.Description,
						stdNum,
						stdDescr,
						dateString,
						termID,
						termName,
					}
					rows = append(rows, row)
				}

			}

		}
	}
	writer.WriteAll(rows)
	return nil

}
