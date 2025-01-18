package domain

import (
	"log"
	"testing"
	"time"
)

func TestTermMonths(t *testing.T) {
	term, err := NewTerm(time.Now(), time.Now().AddDate(0, 4, 24), []time.Time{}, Semester, 1, "")
	if err != nil {
		t.Error()
	}
	months := term.TermMonths()
	if len(months) == 0 {
		t.Error()
	}
	for i, month := range months {
		log.Println("Month", i, month.String())
	}

}
