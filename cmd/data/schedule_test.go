package data

import (
	"log"
	"testing"
	"time"
)

func TestGetDailySchedules(t *testing.T) {
	currentTerm, err := cr.GetTerm(time.Now().AddDate(0, 3, 0))
	if err != nil {
		t.Errorf("error getting term: %s", err)
	}
	log.Println(currentTerm.Name)
	courses, err := cr.GetCourses(currentTerm.ID)
	if err != nil {
		t.Errorf("error getting instances: %s", err)
	}
	for _, course := range courses {
		log.Println("Term name: ", course.Term.Name)
		log.Println("Course name: ", course.Name)
		schedule, err := cr.GetSchedule(course)
		if err != nil {
			t.Errorf("error getting schedules: %s", err)
		}
		log.Println(len(schedule.Schedule))
		log.Println("Term Name: ", schedule.Term.Name)
		for _, dailySchedule := range schedule.Schedule {
			log.Println(dailySchedule.Date, ":", len(dailySchedule.Lessons))
			for _, lesson := range dailySchedule.Lessons {
				log.Println(lesson.Name)
				for _, date := range lesson.Dates {
					log.Println("Date:", date.Format(time.DateOnly))
				}
			}
		}
	}

}
