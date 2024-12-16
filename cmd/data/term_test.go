package data

import (
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestSaveTerm(t *testing.T) {
	terms, err := tr.ReadFromCSV()
	if err != nil {
		t.Errorf("error reading from CSV: %s", err)
	}
	for _, term := range terms {
		termID, err := tr.Save(term)
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
