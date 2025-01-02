package data

import (
	"gh_static_portfolio/cmd/domain"
	"log"
	"path/filepath"
	"testing"
)

func TestSaveFiles(t *testing.T) {
	srcPath, err := filepath.Abs("hello.txt")
	if err != nil {
		t.Error(err)
	}
	file := domain.NewFile("hello", "test .txt file", srcPath)
	id, err := cr.SaveFile(file)
	file.ID = id
	if err != nil {
		t.Error(err)
	}
}

func TestAbsPath(t *testing.T) {
	path := "./files"
	absPath, err := filepath.Abs(path)
	if err != nil {
		t.Error(err)
	}
	log.Println(absPath)
}
