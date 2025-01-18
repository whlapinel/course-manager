package util

import (
	"log"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	weeks := GetMonthDates(time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local))
	for i, week := range weeks {
		log.Println("Week: ", i)
		for j, date := range week {
			log.Println(time.Weekday(j).String())
			log.Println(date)
		}

	}

}
