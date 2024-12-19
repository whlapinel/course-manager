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

type DailyScheduleRepo struct {
	queries *database.Queries
}

func NewDailyScheduleRepo(queries *database.Queries) DailyScheduleRepo {
	return DailyScheduleRepo{queries: queries}

}

// GetSchedule implements domain.DailyScheduleRepo.
func (d DailyScheduleRepo) GetSchedule(instance domain.CourseInstance) (domain.CourseSchedule, error) {
	if instance.Term.Start.IsZero() {
		log.Fatal("GetSchedule(): term not initialized")
	}
	var schedule = domain.CourseSchedule{
		Course: instance.CourseTemplate,
		Term:   instance.Term,
	}
	dbSchedules, err := d.queries.GetDailySchedules(context.Background(), int64(instance.CourseTemplate.ID))
	if err != nil {
		return schedule, err
	}
	type ScheduleHolder struct {
		Lesson domain.Lesson
		Unit   domain.Unit
	}
	dateMap := make(map[string][]ScheduleHolder)
	for _, dbSchedule := range dbSchedules {
		log.Println("dbSchedule: ", dbSchedule.Date, dbSchedule.LessonName)
		lesson := domain.Lesson{
			Name:        dbSchedule.LessonName.String,
			Description: dbSchedule.LessonDescription.String,
		}
		unit := domain.Unit{
			Name: dbSchedule.UnitName,
		}
		holder := ScheduleHolder{
			Lesson: lesson,
			Unit:   unit,
		}
		holders, exists := dateMap[dbSchedule.Date]
		if !exists {
			holders = []ScheduleHolder{
				holder,
			}
		} else {
			holders = append(holders, holder)
		}
		dateMap[dbSchedule.Date] = holders
	}
	var keys []string
	for dateString := range dateMap {
		keys = append(keys, dateString)
	}
	dates, err := parseDates(keys)
	if err != nil {
		return schedule, err
	}
	sorted := sortDates(dates)
	var schedules []domain.DailySchedule
	for _, date := range sorted {
		holders := dateMap[date.Format(time.DateOnly)]
		var lessons []domain.Lesson
		var units []domain.Unit
		for _, holder := range holders {
			lessons = append(lessons, holder.Lesson)
			units = append(units, holder.Unit)
		}
		dailySchedule := domain.DailySchedule{
			Date:    date,
			Lessons: lessons,
			Units:   units,
		}
		schedules = append(schedules, dailySchedule)
	}
	schedule.Schedule = schedules
	if schedule.Term.Start.IsZero() {
		return domain.CourseSchedule{}, fmt.Errorf("GetSchedule(): term not initialized")
	}
	return schedule, nil
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
