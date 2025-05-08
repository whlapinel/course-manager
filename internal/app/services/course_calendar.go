package services

import (
	"gh_static_portfolio/internal/app/dto"
	calendarviews "gh_static_portfolio/internal/app/views/calendar"
)

type CourseCalendarService struct {
	getCourse  func(courseID int) (dto.Course, error)
	getUnits   func(courseID int) ([]dto.Unit, error)
	getLessons func(unitID int) ([]dto.Lesson, error)
	getTerm    func(termID int) (dto.Term, error)
}

func NewCourseCalendarService(
	getCourse func(courseID int) (dto.Course, error),
	getUnits func(courseID int) ([]dto.Unit, error),
	getLessons func(unitID int) ([]dto.Lesson, error),
	getTerm func(termID int) (dto.Term, error),

) *CourseCalendarService {
	return &CourseCalendarService{
		getCourse:  getCourse,
		getUnits:   getUnits,
		getLessons: getLessons,
		getTerm:    getTerm,
	}
}

func (svc *CourseCalendarService) Course(courseID int) (dto.Course, error) {
	courseDTO, err := svc.getCourse(courseID)
	if err != nil {
		return dto.Course{}, err
	}
	return courseDTO, nil
}

func (svc *CourseCalendarService) CalendarDates(courseID int) (calendarviews.CalendarDates, error) {
	courseDTO, err := svc.getCourse(courseID)
	if err != nil {
		return nil, err
	}
	termDTO, err := svc.getTerm(courseDTO.ParentID)
	if err != nil {
		return nil, err
	}
	units, err := svc.getUnits(courseID)
	if err != nil {
		return nil, err
	}
	var datesMap = make(calendarviews.CalendarDates)
	for _, occ := range termDTO.Occasions {
		item := datesMap[occ.Date]
		item.Date = occ.Date
		item.Occasions = append(item.Occasions, occ)
		datesMap[occ.Date] = item
	}
	for _, unit := range units {
		lessons, err := svc.getLessons(unit.ID)
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
