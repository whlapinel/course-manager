package hugo

import (
	"log"
	"testing"
)

func TestBreadCrumbs(t *testing.T) {
	testURL := "units"
	bc := BreadCrumbs(testURL)
	log.Println(bc)
}
