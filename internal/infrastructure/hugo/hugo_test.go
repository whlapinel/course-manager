package hugo

import (
	"log"
	"testing"
)

func TestFilePages(t *testing.T) {
	paths := []string{
		"terms/term_2/files/assignment_1_8_4.md",
		"terms/term_2/files/testing.txt",
	}
	pages, paths := FilePages(paths)
	log.Println("pages:", pages)
	log.Println("paths:", paths)

}
