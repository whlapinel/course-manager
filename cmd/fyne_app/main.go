package main

import (
	"gh_static_portfolio/cmd/data"
	sitegenerator "gh_static_portfolio/cmd/gen_site"
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
	instanceHandler := NewInstanceHandler(w, courseService)
	termsHandler := NewTermsHandler(w, courseService, instanceHandler.ShowCourseTree)
	showTermsItem := fyne.NewMenuItem("Show Terms", termsHandler.ShowTermsList)
	termsMenu := fyne.NewMenu("Terms", showTermsItem)
	fileMenu := fyne.NewMenu("File", fyne.NewMenuItem("Import Course From CSV", func() {}))
	generateMenu := fyne.NewMenu("Generate", fyne.NewMenuItem("Generate Site", func() {
		sitegenerator.Generate()
	}))
	mainMenu := fyne.NewMainMenu(fileMenu, termsMenu, generateMenu)
	w.SetMainMenu(mainMenu)
	w.SetContent(widget.NewLabel("here's something to look at"))
	w.Resize(fyne.Size{Width: 1000, Height: 1000})
	w.ShowAndRun()
}
