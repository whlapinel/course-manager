package main

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewTermsList(terms []domain.Term, showTermDates func(int)) TermsList {
	tl := TermsList{
		Terms: terms,
	}
	tl.ShowTermDates = showTermDates
	lst := widget.NewList(tl.length, tl.createItem, tl.updateItem)
	tl.List = lst
	return tl

}

func (tl TermsList) length() int {
	return len(tl.Terms)
}

func (tl TermsList) createItem() fyne.CanvasObject {
	termNameLabel := widget.NewLabel("term name")
	termSelectButton := widget.NewButton("select term", nil)
	hBox := container.NewHBox(termSelectButton, termNameLabel)
	return hBox
}

func (tl TermsList) updateItem(id int, o fyne.CanvasObject) {
	o.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
		tl.ShowTermDates(tl.Terms[id].ID)
	}
	o.(*fyne.Container).Objects[1].(*widget.Label).SetText(tl.Terms[id].Name)
}

type TermsList struct {
	ShowTermDates func(int)
	Terms         []domain.Term
	List          *widget.List
}

func NewTermDatesList(term domain.Term) TermDatesList {
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
func NewTermsHandler(w fyne.Window, svc CourseService) *TermsHandler {
	return &TermsHandler{w, svc}
}

type TermsHandler struct {
	w   fyne.Window
	svc CourseService
}

func (h TermsHandler) ShowTermsList() {
	terms, err := h.svc.GetTerms()
	if err != nil {
		h.w.SetContent(widget.NewLabel(fmt.Sprintf("error: %s", err)))
	}
	termsList := NewTermsList(terms, h.ShowTermDates)
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
	datesList := NewTermDatesList(termWithDates)
	borderContainer := container.NewBorder(backButton, nil, nil, nil, datesList.List)
	h.w.SetContent(borderContainer)
}

func (h TermsHandler) ShowInstancesForTerm(templateID int, termID int) {
	instances, err := h.svc.GetInstances(termID)
	if err != nil {
		h.w.SetContent(widget.NewLabel(fmt.Sprintf("yikes! there was a problem: %s", err)))
	}

}
func (svc CourseService) GetTerms() ([]domain.Term, error) {
	terms, err := svc.repo.GetTerms()
	if err != nil {
		return nil, err
	}
	return terms, nil
}
func (svc CourseService) GetTermDates(termID int) (domain.Term, error) {
	var term domain.Term
	term, err := svc.repo.GetTermDates(termID)
	if err != nil {
		return term, err
	}
	if len(term.InstructionalDays) == 0 {
		log.Println("warning: term instructional days was 0. fyne_app courseService GetTermDates")
	}
	return term, nil
}
