package localfilesystem

import (
	"gh_static_portfolio/internal/core/filesystem"
	"log"
	"testing"
)

var data []byte
var newData []byte
var fsystem filesystem.FileRepository

func TestMain(m *testing.M) {
	fsystem = New()
	content := "My name is Joe and I work in a button factory"
	newContent := "My name is Sally and I work in a power plant"
	data = []byte(content)
	newData = []byte(newContent)
	m.Run()
}

func TestSave(t *testing.T) {
	err := fsystem.Save(data, "root", "test.md")
	if err != nil {
		t.Error(err)
	}
	err = fsystem.Save(data, "root", "../subdir/test.md")
	if err != nil {
		t.Error(err)
	}
	err = fsystem.Save(data, "root", "subdir/test.md")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	err := fsystem.Update(newData, "root", "subdir/test.md")
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	err := fsystem.Delete("root", "subdir")
	if err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	readContent, err := fsystem.Read("root", "test.md")
	if err != nil {
		t.Error(err)
	}
	log.Println(string(readContent))
}

func TestList(t *testing.T) {
	files, err := fsystem.List("root", "subdir")
	if err != nil {
		t.Error(err)
	}
	for _, file := range files {
		log.Println(file)
	}
}
