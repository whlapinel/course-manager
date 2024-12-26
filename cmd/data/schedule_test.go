package data

import (
	"log"
	"testing"
	"time"
)

func TestGetDailySchedules(t *testing.T) {
	currentTerm, err := cr.GetTerm(time.Now())
	if err != nil {
		t.Errorf("error getting term: %s", err)
	}
	instances, err := cr.GetCourses(currentTerm.ID)
	if err != nil {
		t.Errorf("error getting instances: %s", err)
	}
	for _, instance := range instances {
		log.Println("instance name: ", instance.Name)
		schedule, err := cr.GetSchedule(instance)
		if err != nil {
			t.Errorf("error getting schedules: %s", err)
		}
		log.Println(len(schedule.Schedule))
		for _, schedule := range schedule.Schedule {
			log.Println(schedule.Date, ":", len(schedule.Lessons))
			for _, lesson := range schedule.Lessons {
				log.Println(lesson.Name)
			}
		}
	}

}
