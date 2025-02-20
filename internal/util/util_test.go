package util

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	weeks := GetMonthDates(time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local))
	for i, week := range weeks {
		log.Println("Week: ", i)
		for j, date := range week {
			log.Println(time.Weekday(j).String())
			log.Println(date)
		}

	}

}

func TestIsMarkdown(t *testing.T) {
	path := "/test.md"
	got := IsMarkdown(path)
	expected := true
	if got != expected {
		report := fmt.Sprintf("expected %s, got %s", strconv.FormatBool(expected), strconv.FormatBool(got))
		t.Error(report)
	}
}
