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

func (svc *CourseCalendarService) CalendarDates(courseID int) (calendarviews.CalendarDates, error) {
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
	var datesMap = make(calendarviews.CalendarDates)
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
