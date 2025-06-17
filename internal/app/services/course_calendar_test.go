package services

import (
	"log"
	"testing"
)

func TestWeeks(t *testing.T) {
	svc := NewCourseCalendarService(nil, nil, nil, nil, nil, nil)
	weeks := svc.weeks(6, 2025)
	for i, week := range weeks {
		log.Println("Week: ", i+1)
		for _, date := range week {
			log.Println(date)
		}
	}
}
