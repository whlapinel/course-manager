package main

import (
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// this will be a one-time program to transfer all files and slides into the new lesson_dir
func main() {
	// get all lessons from database
	queries, db, err := data.InitDB("course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	cr := data.NewCourseRepo(queries)
	terms, err := cr.GetTerms()
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
					// get new lesson_dir path (.cmd/data/lessons/lesson_{id})
					newLessonPath := data.LessonDirPath(lesson.ID)

					// for each lesson, get slides path
					oldSlidesPath := data.OldSlidesMarkdownFilePath(*lesson)
					oldFile, err := os.Open(oldSlidesPath)
					if err != nil {
						log.Fatal(err)
					}

					// create new lesson directory
					err = os.Mkdir(newLessonPath, os.ModePerm)
					if err != nil {
						log.Fatal(err)
					}

					// create new slides.md
					newFile, err := os.Create(filepath.Join(newLessonPath, "slides.md"))
					if err != nil {
						log.Fatal(err)
					}

					// copy slides to lesson directory
					_, err = io.Copy(newFile, oldFile)
					if err != nil {
						log.Fatal(err)
					}

					// get old files path
					oldFilesPath := data.OldLessonFilesDirPath(lesson.Files)

					// recursively copy all files from old files to lesson directory

					// Additional Notes (nothing to implement below, must be done manually)
					// files and slides will no longer have their own id and will thus not need to be entered in the database, won't need a domain model either, so no need to update these
					// Also, should just move the slides.html to the same directory
				}
			}
		}
	}
}

func CopyFiles(lesson domain.Lesson, cr data.CourseRepo) error {
	fileDir := lesson.Files
	if fileDir.ID == 0 {
		return nil
	}
	log.Println("Copying fileDir ", fileDir.ID)
	srcRoot := data.OldLessonFilesDirPath(fileDir) // old files path
	log.Println("Source ", srcRoot)
	destRoot := data.NewLessonFilesDirPath(lesson.ID) // new files path
	log.Println("Dest ", destRoot)
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
