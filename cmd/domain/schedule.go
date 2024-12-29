package domain

import "time"

type DailySchedule struct {
	Date    time.Time
	Lessons []Lesson
	Units   []Unit // only the unit name
}

type CourseSchedule struct {
	Course   Course
	Term     Term
	Schedule []DailySchedule
}

func (cs CourseSchedule) GetSchedule(date time.Time) DailySchedule {
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
		return schedule
	}
	return DailySchedule{Date: date}
}

func (cs *CourseSchedule) AddLesson(date time.Time, lesson Lesson) DailySchedule {
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
		schedule.Lessons = append(schedule.Lessons, lesson)
		return schedule
	}
	return DailySchedule{Date: date}
}

func (cs *CourseSchedule) RemoveLesson(date time.Time, removedLesson Lesson) DailySchedule {
	y, m, d := date.Date()
	for i, schedule := range cs.Schedule {
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
		for j, lesson := range schedule.Lessons {
			if lesson.ID == removedLesson.ID {
				schedule.Lessons = append(schedule.Lessons[:j], schedule.Lessons[j+1:]...)
			}
		}
		cs.Schedule[i] = schedule
		return schedule
	}
	return DailySchedule{Date: date}

}
