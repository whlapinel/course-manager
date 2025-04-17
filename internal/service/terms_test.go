package service

import (
	"gh_static_portfolio/internal/domain"
	"log"
	"testing"
)

func TestGetTerm(t *testing.T) {
	term, err := cr.GetTermWithDates(1)
	if err != nil {
		t.Error(err)
	}

	log.Println(term.Name)
	log.Println(term.Description)
	log.Println(term.Description)
	for _, date := range term.InstructionalDays {
		log.Println(date)
	}

}

func TestDates(t *testing.T) {
	term, err := cr.GetTermWithDates(1)
	if err != nil {
		t.Error(err)
	}
	selected := term.InstructionalDays[10]
	log.Println(selected)
	filtered := svc.Dates(selected, term.InstructionalDays, domain.Left)
	log.Println("len(filtered)", len(filtered))
	for _, date := range filtered {
		log.Println(date)
	}
}
