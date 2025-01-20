package domain

import (
	"time"
)

var testTermStart = time.Date(2025, time.January, 2, 0, 0, 0, 0, time.Local)
var testTermEnd = testTermStart.AddDate(0, 3, 0)
var testLessonDates = []time.Time{time.Date(2025, time.January, 3, 0, 0, 0, 0, time.Local)}
var testNonInstructDays = []time.Time{time.Date(2025, time.January, 6, 0, 0, 0, 0, time.Local)}
