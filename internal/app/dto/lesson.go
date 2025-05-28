package dto

import (
	"fmt"
	"gh_static_portfolio/internal/core/assessment"
	"gh_static_portfolio/internal/core/standard"
	"gh_static_portfolio/internal/features/lesson"
	"gh_static_portfolio/internal/ports"
	"slices"
	"time"
)

type Lesson struct {
	lesson.Lesson `json:"lesson"`
	Standards     []standard.Standard     `json:"standards"`
	Assessments   []assessment.Assessment `json:"assessments"`
}

func (l Lesson) GetChildren() []ports.Node {
	return []ports.Node{}
}

func (l Lesson) GetTypeName() string {
	return LessonTypeName.String()
}

func (l Lesson) GetParentTypeName() string {
	return UnitTypeName.String()
}

func LessonDesignator(lesson Lesson, unit Unit) string {
	if unit.Number < 0 {
		return fmt.Sprintf("%s %d", unit.Name, lesson.Number)
	}
	return fmt.Sprintf("Lesson %d.%d", unit.Number, lesson.Number)
}

type CalendarDirection int

const (
	Right CalendarDirection = iota
	Left
)

var dirStringList = []string{"right", "left"}

func (d CalendarDirection) String() string {
	return dirStringList[d]
}

func ParseDirection(cd string) (CalendarDirection, error) {
	for i, word := range dirStringList {
		if cd == word {
			return CalendarDirection(i), nil
		}
	}
	return 0, fmt.Errorf("invalid direction value")
}

func (l Lesson) DedupeDates() Lesson {
	var dateMap = make(map[string]time.Time)
	for _, date := range l.Dates {
		dateMap[date.Format(time.DateOnly)] = date
	}
	var deduped []time.Time
	for _, date := range dateMap {
		deduped = append(deduped, date)
	}
	l.Dates = deduped
	return l
}

func (l Lesson) SortDates() Lesson {
	slices.SortFunc(l.Dates, func(d1 time.Time, d2 time.Time) int {
		if d1.Before(d2) {
			return 1
		}
		return -1
	})
	return l
}

// returns designation e.g. Lesson 1.2
func (l Lesson) Designation() string {
	if l.UnitNumber >= 0 {
		return fmt.Sprintf("Lesson %d.%d", l.UnitNumber, l.Number)
	}
	return fmt.Sprintf("%s: Day %d", l.UnitName, l.Number)
}
