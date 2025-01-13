package main

import (
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/data/csv"
	"gh_static_portfolio/cmd/domain"
	sitegenerator "gh_static_portfolio/cmd/gen_site"
	"gh_static_portfolio/cmd/service"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("something isn't right!")
	myApp := app.New()
	w := myApp.NewWindow("Course Manager")
	queries, db, err := data.InitDB("course_manager.db")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	courseRepo := data.NewCourseRepo(queries)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := NewCourseHandler(w, courseService)
	termsHandler := NewTermsHandler(w, courseService, courseHandler.ShowCourseTree)
	showTermsItem := fyne.NewMenuItem("Show Terms", termsHandler.ShowTermsList)
	termsMenu := fyne.NewMenu("Terms", showTermsItem)
	importCSVItem := fyne.NewMenuItem("Import Course From CSV", func() {
		selectFileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if reader == nil {
				// user canceled
				return
			}
			defer reader.Close()

			courses, err := csv.ImportCoursesFromCSVReaderV2(reader)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			// Now do something with these courses, e.g. store them in your app, re-draw the UI, etc.
			log.Printf("Successfully imported %d courses\n", len(courses))
		}, w)
		selectFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		selectFileDialog.Show()
	})
	exportCSVItem := fyne.NewMenuItem("Export Course To CSV", func() {
		selectFileDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if writer == nil {
				// user canceled
				return
			}
			defer w.Close()

			// Now we can pass `w` to the CSV writer function
			var courses domain.Courses
			terms, err := courseRepo.GetTerms()
			if err != nil {
				dialog.ShowError(err, w)
			}
			for _, term := range terms {
				termCourses, err := courseRepo.GetCourses(term.ID)
				if err != nil {
					dialog.ShowError(err, w)
				}
				courses = append(courses, termCourses...)
			}
			if err := csv.WriteCoursesToCSVV2(courses, writer); err != nil {
				dialog.ShowError(err, w)
				return
			}
			dialog.ShowInformation("Success", "Courses exported to CSV", w)
		}, w)
		selectFileDialog.SetFileName(fmt.Sprintf("courses_export_%s", strings.ReplaceAll(" ", time.Now().Format(time.DateTime), "_")))
		selectFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		selectFileDialog.Show()
	})
	testButtonItem := fyne.NewMenuItem("Test File Open", func() {
		file, err := os.Create("hello_world.md")
		if err != nil {
			dialog.ShowError(err, w)
		}
		defer file.Close()
		content := []byte("My name is joe and I work in a button factory")
		file.Write(content)
		err = exec.Command("code", file.Name()).Start()
		if err != nil {
			dialog.ShowError(err, w)
		}

	})
	fileMenu := fyne.NewMenu("File", importCSVItem, exportCSVItem, testButtonItem)
	generateMenu := fyne.NewMenu("Generate", fyne.NewMenuItem("Generate Site", func() {
		log.Println("generating static site")
		err := sitegenerator.Generate(courseRepo)
		if err != nil {
			dialog.ShowError(err, w)
		}
	}))
	mainMenu := fyne.NewMainMenu(fileMenu, termsMenu, generateMenu)
	w.SetMainMenu(mainMenu)
	w.SetContent(widget.NewLabel("here's something to look at"))
	w.Resize(fyne.Size{Width: 1000, Height: 1000})
	w.ShowAndRun()
}
