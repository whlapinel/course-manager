package domain

import "time"

type DailyScheduleRepo interface {
	GetSchedule(instance CourseInstance) (CourseSchedule, error)
}

type DailySchedule struct {
	Date    time.Time
	Lessons []Lesson
	Units   []Unit // only the unit name
}

type CourseSchedule struct {
	Course   CourseTemplate
	Term     Term
	Schedule []DailySchedule
}

func (cs CourseSchedule) GetSchedule(date time.Time) *DailySchedule {
	y, m, d := date.Date()
	for _, schedule := range cs.Schedule {
		y2, m2, d2 := schedule.Date.Date()
		if y2 != y {
			continue
		}
		if m2 != m {
			continue
		}
		if d2 != d {
			continue
		}
		return &schedule
	}
	return nil
}
