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
