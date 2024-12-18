package data

import (
	"context"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
	"slices"
	"time"
)

type dailyScheduleRepo struct {
	queries *database.Queries
}

func NewDailyScheduleRepo(queries *database.Queries) domain.DailyScheduleRepo {
	return dailyScheduleRepo{queries: queries}

}

// GetSchedule implements domain.DailyScheduleRepo.
func (d dailyScheduleRepo) GetSchedule(instanceID int) ([]domain.DailySchedule, error) {
	dbSchedules, err := d.queries.GetDailySchedules(context.Background(), int64(instanceID))
	if err != nil {
		return nil, err
	}
	if len(dbSchedules) == 0 {
		return nil, fmt.Errorf("dbSchedules is empty")
	}
	dateMap := make(map[string][]domain.Lesson)
	for _, dbSchedule := range dbSchedules {
		log.Println("dbSchedule: ", dbSchedule.Date, dbSchedule.LessonName)
		lesson := domain.Lesson{
			Name:        dbSchedule.LessonName.String,
			Description: dbSchedule.LessonDescription.String,
		}
		lessons, exists := dateMap[dbSchedule.Date]
		if !exists {
			lessons = []domain.Lesson{
				lesson,
			}
		} else {
			lessons = append(lessons, lesson)
		}
		dateMap[dbSchedule.Date] = lessons
	}
	var keys []string
	for dateString := range dateMap {
		keys = append(keys, dateString)
	}
	dates, err := parseDates(keys)
	if err != nil {
		return nil, err
	}
	sorted := sortDates(dates)
	var schedules []domain.DailySchedule
	for _, date := range sorted {
		dailySchedule := domain.DailySchedule{
			Date:    date,
			Lessons: dateMap[date.Format(time.DateOnly)],
		}
		schedules = append(schedules, dailySchedule)
	}
	return schedules, nil
}

func parseDates(dateStrings []string) ([]time.Time, error) {
	var dates []time.Time
	for _, dateString := range dateStrings {
		date, err := time.Parse(time.DateOnly, dateString)
		if err != nil {
			return nil, err
		}
		dates = append(dates, date)

	}
	return dates, nil
}

func sortDates(dates []time.Time) []time.Time {
	slices.SortFunc(dates, compare)
	return dates
}

func compare(a time.Time, b time.Time) int {
	if a.Before(b) {
		return -1
	} else {
		return 1
	}
}
