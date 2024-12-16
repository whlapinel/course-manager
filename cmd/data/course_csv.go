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
	rows = append(rows, CSVHeaders())
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
