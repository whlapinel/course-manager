package lesson

import (
	"log"
	"testing"
	"time"
)

func TestShiftLesson(t *testing.T) {
	lesson := Lesson{
		Dates: []time.Time{
			time.Now(),
			time.Now().Add(time.Hour * 24),
			time.Now().Add(time.Hour * 24 * 2),
		},
	}
	termDates := []time.Time{
		time.Now().Add(time.Hour * 24 * -2),
		time.Now().Add(time.Hour * 24 * -1),
		time.Now(),
		time.Now().Add(time.Hour * 24),
		time.Now().Add(time.Hour * 24 * 2),
		time.Now().Add(time.Hour * 24 * 3),
		time.Now().Add(time.Hour * 24 * 4),
		time.Now().Add(time.Hour * 24 * 5),
		time.Now().Add(time.Hour * 24 * 6),
		time.Now().Add(time.Hour * 24 * 7),
		time.Now().Add(time.Hour * 24 * 8),
		time.Now().Add(time.Hour * 24 * 10),
	}
	log.Println("Before shift")
	for _, date := range lesson.Dates {
		log.Println(date.Format(time.DateOnly))
	}
	_, err := lesson.ShiftDate(-2, time.Now(), termDates)
	if err != nil {
		t.Error(err)
	}
	log.Println("After shift:")
	for _, date := range lesson.Dates {
		log.Println(date.Format(time.DateOnly))
	}

}
