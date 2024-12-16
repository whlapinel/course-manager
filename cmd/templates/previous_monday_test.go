package templates

import (
	"log"
	"testing"
	"time"
)

func TestPreviousMonday(t *testing.T) {
	date := time.Now().Add(24 * time.Hour * 3)
	prevMon := PreviousMonday(date)
	if prevMon.Weekday() != time.Monday {
		t.Errorf("weekday is %s not Monday", prevMon.Weekday().String())
	}
	log.Println("Previous Monday is: ", prevMon.String()[:10])
}
