package templates

import (
	"log"
	"time"
)

// This gives all weekdates beginning on the 1st the current month and ending on the last day.
// It's a 2D slice with each sub-slice being for a week, index number is the weekday and the value is the date
// week starts with Sunday
func GetMonthDates(date time.Time) [][]time.Time {
	year, month, _ := date.Date()
	first := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	last := first.AddDate(0, 1, 0).AddDate(0, 0, -1)
	log.Println("last:", last.Format(time.DateOnly))
	monthSlice := [][]time.Time{}
	currDate := first
	currWeekIndex := 0
	currWeekSlice := make([]time.Time, 7)
	for !currDate.After(last) {
		currWeekSlice[currDate.Weekday()] = currDate
		// if it's Sunday, append the current week and start a new week
		if currDate.Weekday() == 6 {
			monthSlice = append(monthSlice, currWeekSlice)
			currWeekIndex = 0
			currWeekSlice = make([]time.Time, 7)
		}
		currWeekIndex++
		currDate = currDate.AddDate(0, 0, 1)
	}
	if last.Weekday() != time.Sunday {
		monthSlice = append(monthSlice, currWeekSlice)
	}
	return monthSlice
}
