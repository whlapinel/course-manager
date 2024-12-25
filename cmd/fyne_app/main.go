package main

import (
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/service"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Course Manager")
	queries, db, err := data.InitDB("test_course_manager.db")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	courseRepo := data.NewCourseRepo(queries)
	courseService := service.NewCourseService(courseRepo)
	templatesHandler := NewCourseHandler(w, courseService)
	instanceHandler := NewInstanceHandler(w, courseService)
	termsHandler := NewTermsHandler(w, courseService, instanceHandler.ShowInstancesTree)
	courseTreeMenuItem := fyne.NewMenuItem("Show Courses", templatesHandler.ShowCourseTree)
	coursesMenu := fyne.NewMenu("Course Templates", courseTreeMenuItem)
	showTermsItem := fyne.NewMenuItem("Show Terms", termsHandler.ShowTermsList)
	termsMenu := fyne.NewMenu("Terms", showTermsItem)
	fileMenu := fyne.NewMenu("File", fyne.NewMenuItem("Import Course From CSV", func() {}))
	mainMenu := fyne.NewMainMenu(fileMenu, coursesMenu, termsMenu)
	w.SetMainMenu(mainMenu)
	w.SetContent(widget.NewLabel("here's something to look at"))
	w.Resize(fyne.Size{Width: 1000, Height: 1000})
	w.ShowAndRun()
}
