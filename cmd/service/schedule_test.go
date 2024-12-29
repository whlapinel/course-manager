package service

import (
	"log"
	"testing"
)

func TestGetSchedule(t *testing.T) {
	terms, err := svc.GetTerms()
	if err != nil {
		t.Error()
	}
	courses, err := svc.GetCourses(terms[0].ID)
	if err != nil {
		t.Error()
	}
	schedule, err := svc.GetSchedule(courses[0])
	if err != nil {
		t.Error()
	}
	for _, dailySchedule := range schedule.Schedule {
		log.Println(dailySchedule.Date, dailySchedule.Lessons[0])
	}

}
