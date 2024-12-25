package main

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/service"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewTermsHandler(w fyne.Window, svc service.CourseService, showInstances func(int)) *TermsHandler {
	return &TermsHandler{w, svc, showInstances}
}

type TermsHandler struct {
	w                 fyne.Window
	svc               service.CourseService
	showInstancesTree func(int)
}

func (h TermsHandler) ShowTermsList() {
	terms, err := h.svc.GetTerms()
	if err != nil {
		h.w.SetContent(widget.NewLabel(fmt.Sprintf("error: %s", err)))
	}
	termsList := h.NewTermsList(terms, h.ShowTermDates, h.showInstancesTree)
	h.w.SetContent(termsList.List)
}

func (h TermsHandler) ShowTermDates(termID int) {
	termWithDates, err := h.svc.GetTermDates(termID)
	log.Println(len(termWithDates.InstructionalDays))
	if err != nil {
		h.w.SetContent(widget.NewLabel(err.Error()))
	}
	if len(termWithDates.InstructionalDays) == 0 {
		h.w.SetContent(widget.NewLabel("instructional days: empty"))
	}
	backButton := widget.NewButton("Back", h.ShowTermsList)
	datesList := h.NewTermDatesList(termWithDates)
	borderContainer := container.NewBorder(backButton, nil, nil, nil, datesList.List)
	h.w.SetContent(borderContainer)
}

func (th TermsHandler) NewTermsList(terms []domain.Term, showTermDates func(int), showInstancesTree func(int)) TermsList {
	tl := TermsList{
		Terms: terms,
	}
	tl.ShowInstancesTree = showInstancesTree
	tl.ShowTermDates = showTermDates
	lst := widget.NewList(tl.length, tl.createItem, tl.updateItem)
	tl.List = lst
	return tl

}

func (tl TermsList) length() int {
	return len(tl.Terms)
}

func (tl TermsList) createItem() fyne.CanvasObject {
	termSelectButton := widget.NewButton("view calendar", nil)
	viewInstanceButton := widget.NewButton("view instances", nil)
	termNameLabel := widget.NewLabel("term name")
	hBox := container.NewHBox(termSelectButton, viewInstanceButton, termNameLabel)
	return hBox
}

func (tl TermsList) updateItem(id int, o fyne.CanvasObject) {
	o.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
		tl.ShowTermDates(tl.Terms[id].ID)
	}
	o.(*fyne.Container).Objects[1].(*widget.Button).OnTapped = func() {
		tl.ShowInstancesTree(tl.Terms[id].ID)

	}
	o.(*fyne.Container).Objects[2].(*widget.Label).SetText(tl.Terms[id].Name)
}

type TermsList struct {
	ShowTermDates     func(int)
	ShowInstancesTree func(int)
	Terms             []domain.Term
	List              *widget.List
}

func (th TermsHandler) NewTermDatesList(term domain.Term) TermDatesList {
	var tdl TermDatesList
	tdl.Term = term
	tdl.List = widget.NewList(tdl.length, tdl.createItem, tdl.updateItem)
	return tdl
}

type TermDatesList struct {
	Term domain.Term
	List *widget.List
}

func (tdl TermDatesList) length() int {
	return len(tdl.Term.InstructionalDays)
}

func (tdl TermDatesList) createItem() fyne.CanvasObject {
	dateLabel := widget.NewLabel("date")
	dateSelectBtn := widget.NewButton("select term", nil)
	hBox := container.NewHBox(dateSelectBtn, dateLabel)
	return hBox
}

func (tdl TermDatesList) updateItem(id int, o fyne.CanvasObject) {
	o.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = nil
	o.(*fyne.Container).Objects[1].(*widget.Label).SetText(tdl.Term.InstructionalDays[id].Format(time.DateOnly))
}
