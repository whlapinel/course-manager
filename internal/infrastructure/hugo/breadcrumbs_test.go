package hugo

import (
	"log"
	"testing"
)

func TestBreadCrumbs(t *testing.T) {
	testURL := "/terms/term_2/courses/course_4/units/unit_37/lessons/lesson_268"
	bc := BreadCrumbs(testURL)
	log.Println(bc)
}
