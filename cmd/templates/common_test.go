package templates

import (
	"log"
	"testing"
)

func TestFilePathToURL(t *testing.T) {
	route := FilePathToURL(coursesDir)
	log.Println(route)
	route = FilePathToURL(rootDir)
	log.Println(route)
}
