package lesson

import (
	"fmt"
	"gh_static_portfolio/internal/ports"
	datespkg "gh_static_portfolio/internal/shared/dates"
	"slices"
	"time"
)

type Lesson struct {
	UnitNumber int         `json:"unitNumber"`
	UnitName   string      `json:"unitName"`
	Dates      []time.Time `json:"dates"`
	ports.BaseNode[int, int]
}

// days is positive if to the right, negative if to the left. returns error if out of range
func (l *Lesson) ShiftDate(days int, oldDate time.Time, instructionalDates []time.Time) (newDate time.Time, err error) {
	currDateIndex := slices.IndexFunc(instructionalDates, func(date time.Time) bool {
		return datespkg.IsSameDate(date, oldDate)
	})
	// if shifting right
	if days > 0 {
		// and resulting index is more than that of last element
		if currDateIndex+days > len(instructionalDates)-1 {
			// return error
			return newDate, fmt.Errorf("out of range: shifting %d days results in attempted index of %d and length of slice is %d", days, currDateIndex+days, len(instructionalDates))
		}
	} else if days < 0 {
		if currDateIndex+days < 0 {
			return newDate, fmt.Errorf("out of range: shifting %d days results in attempted index of %d and length of slice is %d", days, currDateIndex+days, len(instructionalDates))
		}
	}
	newDate = instructionalDates[currDateIndex+days]
	oldLessonDateIndex := slices.IndexFunc(l.Dates, func(date time.Time) bool {
		return datespkg.IsSameDate(date, oldDate)
	})
	l.Dates = slices.Replace(l.Dates, oldLessonDateIndex, oldLessonDateIndex+1, newDate)
	datespkg.Sort(l.Dates)
	l.Dates = slices.CompactFunc(l.Dates, func(date1, date2 time.Time) bool {
		return datespkg.IsSameDate(date1, date2)
	})
	return newDate, nil
}
