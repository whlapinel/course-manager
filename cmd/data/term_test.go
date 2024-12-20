package data

import (
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestSaveTerm(t *testing.T) {
	terms, err := cr.ReadFromCSV()
	if err != nil {
		t.Errorf("error reading from CSV: %s", err)
	}
	for _, term := range terms {
		termID, err := cr.SaveTerm(term)
		if err != nil {
			t.Errorf("error saving term: %s", err)
		}
		if termID == 0 {
			t.Errorf("term ID is 0")
		}
		log.Println(term.Name)
		for _, date := range term.InstructionalDays {
			log.Println("date: ", date.Format(time.DateOnly))

		}
	}

}

func TestGetTerm(t *testing.T) {
	term, err := cr.GetTerm(time.Now())
	if err != nil {
		t.Errorf("error fetching term: %s", err)
	}
	if term.Start.IsZero() {
		t.Error("term Start is zero")
	}
	log.Println(term.Name)
	log.Println(term.ID)
	monthDates := term.TermMonths()
	for _, date := range monthDates {
		log.Println(date.Format(time.DateOnly))
	}
}
