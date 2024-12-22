package main

import (
	"gh_static_portfolio/cmd/data"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Course Template Tree")
	queries, db, err := data.InitDB("test_course_manager.db")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	courseRepo := data.NewCourseRepo(queries)
	courseService := NewCourseService(courseRepo)
	templates, err := courseService.GetTemplates()
	t := NewCourseTemplateTree(templates, courseService)
	showCoursesButton := widget.NewButton("Courses", func() {
		if err != nil {
			log.Fatal()
		}
		w.SetContent(t.Tree)
	})
	termsHandler := NewTermsHandler(w, courseService)
	showTermsButton := widget.NewButton("Terms", termsHandler.ShowTermsList)

	otherButton := widget.NewButton("some other button", func() {
		log.Println("this doesn't really do anything sorry")
	})
	content := container.NewVBox(showCoursesButton, showTermsButton, otherButton)
	w.SetContent(content)
	w.Resize(fyne.Size{Width: 1000, Height: 1000})
	w.ShowAndRun()
}
