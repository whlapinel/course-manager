package data

import (
	"log"
	"testing"
)

func TestGetDailySchedules(t *testing.T) {
	instances, err := cr.GetInstances()
	if err != nil {
		t.Errorf("error getting instances: %s", err)
	}
	for _, instance := range instances {
		log.Println("instance name: ", instance.Name)
		schedules, err := dsr.GetSchedule(instance.ID)
		if err != nil {
			t.Errorf("error getting schedules: %s", err)
		}
		log.Println(len(schedules))
		for _, schedule := range schedules {
			log.Println(schedule.Date, ":", len(schedule.Lessons))
			for _, lesson := range schedule.Lessons {
				log.Println(lesson.Name)
			}
		}
	}

}
