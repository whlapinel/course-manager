package main

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/service"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewTermsHandler(w fyne.Window, svc service.CourseService, showCourses func(termID int)) *TermsHandler {
	return &TermsHandler{w, svc, showCourses}
}

type TermsHandler struct {
	w           fyne.Window
	svc         service.CourseService
	showCourses func(termID int)
}

func (h TermsHandler) ShowTermsList() {
	terms, err := h.svc.GetTerms()
	if err != nil {
		h.w.SetContent(widget.NewLabel(fmt.Sprintf("error: %s", err)))
	}
	termsList := h.NewTermsList(terms, h.showCourses)
	h.w.SetContent(termsList.List)
}

func (th TermsHandler) NewTermsList(terms []domain.Term, showCourses func(termID int)) TermsList {
	tl := TermsList{
		Terms: terms,
	}
	tl.ShowCourses = showCourses
	lst := widget.NewList(tl.length, tl.createItem, tl.updateItem)
	tl.List = lst
	return tl

}

func (tl TermsList) length() int {
	return len(tl.Terms)
}

func (tl TermsList) createItem() fyne.CanvasObject {
	selectTermBtn := widget.NewButton("View Courses", nil)
	termNameLabel := widget.NewLabel("term name")
	hBox := container.NewHBox(selectTermBtn, termNameLabel)
	return hBox
}

func (tl TermsList) updateItem(id int, o fyne.CanvasObject) {
	o.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
		log.Println("Showing courses for Term:", tl.Terms[id].Name, "id:", tl.Terms[id].ID)
		tl.ShowCourses(tl.Terms[id].ID)

	}
	o.(*fyne.Container).Objects[1].(*widget.Label).SetText(tl.Terms[id].Name)
}

type TermsList struct {
	ShowCourses func(termID int)
	Terms       []domain.Term
	List        *widget.List
}
