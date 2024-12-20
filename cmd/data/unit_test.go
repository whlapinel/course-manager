package data

import (
	"log"
	"testing"
)

func TestGetUnits(t *testing.T) {
	units, err := cr.GetUnits(1)
	if err != nil {
		t.Error()
	}
	for _, unit := range units {
		log.Println(unit.Name, unit.Description)
	}
}
