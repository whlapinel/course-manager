package templates

import (
	"log"
	"testing"
	"time"
)

func TestGetMonthDates(t *testing.T) {
	today := time.Now().AddDate(0, -2, 0)
	monthSlice := GetMonthDates(today)
	for _, week := range monthSlice {
		log.Println("new week: week of ", week[0].Format("Mon 1/2/06"))
		for _, date := range week {
			log.Println(date.Format("Mon 1/2/06"))
		}
	}

}
