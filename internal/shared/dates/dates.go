package dates

import (
	"slices"
	"time"
)

func Sort(dates []time.Time) {
	slices.SortFunc(dates, func(a, b time.Time) int {
		if a.Before(b) {
			return 1
		}
		if b.Before(a) {
			return -1
		} else {
			return 0
		}
	})
}
