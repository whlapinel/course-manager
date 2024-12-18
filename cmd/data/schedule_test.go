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
		schedule, err := dsr.GetSchedule(*instance)
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
