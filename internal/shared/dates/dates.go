package dates

import (
	"slices"
	"time"
)

func Sort(dates []time.Time) {
	slices.SortFunc(dates, func(a, b time.Time) int {
		if a.Before(b) {
			return -1
		}
		if b.Before(a) {
			return 1
		} else {
			return 0
		}
	})
}

func IsSameDate(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	if y1 == y2 && m1 == m2 && d1 == d2 {
		return true
	}
	return false
}
