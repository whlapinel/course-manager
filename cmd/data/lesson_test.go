package data

import (
	"log"
	"testing"
)

func TestGetLessonDates(t *testing.T) {
	dates, err := cr.GetLessonDates(1)
	if err != nil {
		t.Error(err)
	}
	for _, date := range dates {
		log.Println(date)
	}

}
