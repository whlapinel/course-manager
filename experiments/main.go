package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	AddFile("hello.txt")
}

func AddFile(path string) {
	err := os.MkdirAll("./lesson_1", 0777)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	newFile, err := os.Create(filepath.Join("./lesson_1", filepath.Base(path)))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(newFile, file)

}
