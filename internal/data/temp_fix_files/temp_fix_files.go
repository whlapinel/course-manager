package main

import (
	"gh_static_portfolio/internal/data"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	queries, db, err := data.InitDB("internal/data/database/course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	cr := data.NewCourseRepo(queries)
	terms, err := cr.GetTerms()
	if err != nil {
		log.Fatal(err)
	}
	for _, term := range terms {
		courses, err := cr.GetCourses(term.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, course := range courses {
			units, err := cr.GetUnits(course.ID)
			if err != nil {
				log.Fatal(err)
			}
			for _, unit := range units {
				lessons, err := cr.GetLessons(unit.ID)
				if err != nil {
					log.Fatal(err)
				}
				for _, lesson := range lessons {
					err := CopyLessonFiles(data.LessonDirPath(lesson.ID), data.NodeDirPath(term, course, unit, lesson))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}

}

func CopyLessonFiles(srcRoot, destRoot string) error {
	return filepath.WalkDir(srcRoot, func(srcPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(srcRoot, srcPath)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destRoot, relPath)
		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}
		return copyFile(srcPath, destPath)
	})
}

func copyFile(srcPath, destPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}
