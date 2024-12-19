package data

import (
	"log"
	"testing"
	"time"
)

func TestGetDailySchedules(t *testing.T) {
	currentTerm, err := tr.GetTerm(time.Now())
	if err != nil {
		t.Errorf("error getting term: %s", err)
	}
	instances, err := cr.GetInstances(*currentTerm)
	if err != nil {
		t.Errorf("error getting instances: %s", err)
	}
	for _, instance := range instances {
		log.Println("instance name: ", instance.CourseTemplate.Name)
		schedule, err := dsr.GetSchedule(instance)
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
