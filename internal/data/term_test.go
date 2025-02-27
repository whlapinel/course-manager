package data

import (
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetTermWithDates(t *testing.T) {
	term, err := cr.GetTermWithDates(1)
	if err != nil {
		t.Errorf("error fetching term: %s", err)
	}
	if term.ID != 0 {
		if term.Start.IsZero() {
			t.Error("term Start is zero")
		}
		log.Println(term.Name)
		log.Println(term.ID)
		log.Println("instructional days")
		for _, date := range term.InstructionalDays {
			log.Println(date.Format(time.DateOnly))
		}
		log.Println("non-instructional days")
		for _, date := range term.NonInstructionalDays {
			log.Println(date.Format(time.DateOnly))
		}
	} else {
		log.Println("term id was 0")
	}
}
