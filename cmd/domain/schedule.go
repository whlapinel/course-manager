package domain

import "time"

type DailyScheduleRepo interface {
	GetSchedule(instanceID int) ([]DailySchedule, error)
}

type DailySchedule struct {
	Date    time.Time
	Lessons []Lesson
}

type CourseSchedule struct {
	Course   Course
	Schedule []DailySchedule
}
