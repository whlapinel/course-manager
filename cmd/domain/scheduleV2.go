package domain

import "time"

type DailyScheduleV2 struct {
	Date    time.Time
	Lessons []Lesson
	Units   []Unit // only the unit name
}

type CourseScheduleV2 struct {
	Course   Course
	Term     Term
	Schedule map[int]MonthSchedule
}

type MonthSchedule struct {
	
}
