package services

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	calendarviews "gh_static_portfolio/internal/app/views/calendar"
	"gh_static_portfolio/internal/features/courseoccasion"
	"gh_static_portfolio/internal/features/lesson"
	"log"
	"time"
)

type CourseCalendarService struct {
	terms         *TermService
	courses       *CourseService
	units         *UnitService
	lessons       *LessonService
	occasions     *courseoccasion.Service
	lessonService *lesson.Service
}

func NewCourseCalendarService(
	terms *TermService,
	courses *CourseService,
	units *UnitService,
	lessons *LessonService,
	occasions *courseoccasion.Service,
	lessonService *lesson.Service,

) *CourseCalendarService {
	return &CourseCalendarService{
		terms:         terms,
		courses:       courses,
		units:         units,
		lessons:       lessons,
		occasions:     occasions,
		lessonService: lessonService,
	}
}

func (svc *CourseCalendarService) AddLessonToDate(date time.Time, lessonID, termID int) error {
	return svc.lessonService.AddLessonToDate(date, lessonID, termID)
}

func (svc *CourseCalendarService) DeleteLessonDate(date time.Time, lessonID, termID int) error {
	return svc.lessonService.DeleteLessonDate(date, lessonID, termID)
}

func (svc *CourseCalendarService) ShiftLesson(currDate time.Time, lessonID int, termID int, direction dto.CalendarDirection) error {
	// instructional days
	termInstance, err := svc.terms.ByID(int(termID))
	if err != nil {
		log.Println("error retrieving term")
		return fmt.Errorf("error retrieving term instance of ID %d: %w", termID, err)
	}
	shiftMagnitude := 1
	if direction == dto.Left {
		shiftMagnitude = -1
	}
	err = svc.lessonService.Shift(lessonID, termID, shiftMagnitude, currDate, termInstance.InstructionalDays)
	if err != nil {
		return err
	}
	return nil
}

func (svc *CourseCalendarService) Course(courseID int) (dto.Course, error) {
	courseDTO, err := svc.courses.ByID(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	return courseDTO, nil
}

func (svc *CourseCalendarService) CalendarDates(courseID int) (calendarviews.DatesMap, error) {
	courseDTO, err := svc.courses.ByID(courseID)
	if err != nil {
		return nil, err
	}
	courseOccasions, err := svc.occasions.ByCourseID(courseID)
	if err != nil {
		return nil, err
	}
	termDTO, err := svc.terms.ByID(courseDTO.ParentID)
	if err != nil {
		return nil, err
	}
	units, err := svc.units.ByParentID(courseID)
	if err != nil {
		return nil, err
	}
	var datesMap = make(calendarviews.DatesMap)
	for _, occ := range termDTO.Occasions {
		item := datesMap[occ.Date]
		item.Date = occ.Date
		item.TermOccasions = append(item.TermOccasions, occ)
		datesMap[occ.Date] = item
	}
	for _, occ := range courseOccasions {
		item := datesMap[occ.Date]
		item.Date = occ.Date
		item.CourseOccasions = append(item.CourseOccasions, occ)
		datesMap[occ.Date] = item
	}
	for _, unit := range units {
		lessons, err := svc.lessons.ByParentID(unit.ID)
		if err != nil {
			return nil, err
		}
		for _, lesson := range lessons {
			for _, date := range lesson.Dates {
				item := datesMap[date]
				item.Date = date
				item.Lessons = append(item.Lessons, lesson)
				datesMap[date] = item
			}
		}
	}
	return datesMap, nil
}

func (svc *CourseCalendarService) MonthWeeks(month, year, courseID int) ([][]calendarviews.CalendarDate, error) {
	datesMap, err := svc.CalendarDates(courseID)
	if err != nil {
		return nil, err
	}
	monthWeeks := svc.calendarWeeks(month, year, datesMap)
	for i, week := range monthWeeks {
		log.Println("Week: ", i, week)
		for j, date := range week {
			log.Println("Date: ", j, date)
			log.Println("Date: ", j, date.Date)
		}
	}
	return monthWeeks, nil
}

func (svc *CourseCalendarService) calendarWeeks(month, year int, datesMap calendarviews.DatesMap) [][]calendarviews.CalendarDate {
	weeks := svc.weeks(month, year)
	calWeeks := make([][]calendarviews.CalendarDate, 0)
	for _, week := range weeks {
		calWeek := make([]calendarviews.CalendarDate, 0)
		for _, date := range week {
			var calDate calendarviews.CalendarDate
			calDate.Date = date
			if data, ok := datesMap[date]; ok {
				calDate = data
			}
			calWeek = append(calWeek, calDate)
		}
		calWeeks = append(calWeeks, calWeek)
	}
	return calWeeks
}

func (svc *CourseCalendarService) weeks(month, year int) [][]time.Time {
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	var sun = firstOfMonth.AddDate(0, 0, -int(firstOfMonth.Weekday()))
	var weeks [][]time.Time
	for currWeek := sun; !currWeek.After(firstOfMonth.AddDate(0, 1, 0)); currWeek = currWeek.AddDate(0, 0, 7) {
		var week []time.Time
		for i := range 7 {
			currDate := currWeek.AddDate(0, 0, i)
			week = append(week, currDate)
		}
		weeks = append(weeks, week)
	}
	return weeks
}
