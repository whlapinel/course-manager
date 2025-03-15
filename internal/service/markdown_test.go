package service

import (
	"os"
	"testing"
)

func TestMarkdownToHTML(t *testing.T) {
	markdown := "# Hello world!\n## This is me!"
	path := "./markdown_test/test_file.md"
	err := os.MkdirAll("./markdown_test", os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	_, err = file.Write([]byte(markdown))
	if err != nil {
		t.Error(err)
	}
	svc.MarkdownToHTML(path)

}
