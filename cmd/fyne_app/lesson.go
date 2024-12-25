package main

import (
	"gh_static_portfolio/cmd/domain"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (svc CourseService) UpdateLesson(l domain.Lesson) error {
	err := svc.repo.UpdateLesson(l)
	if err != nil {
		return err
	}
	return nil
}

func (ch InstanceHandler) NewLessonDatesList(lesson domain.Lesson) LessonDatesList {
	var ldl LessonDatesList
	ldl.Lesson = lesson
	ldl.List = widget.NewList(ldl.length, ldl.createItem, ldl.updateItem)
	return ldl
}

type LessonDatesList struct {
	Lesson domain.Lesson
	List   *widget.List
}

func (ldl *LessonDatesList) length() int {
	return len(ldl.Lesson.Dates)
}

func (ldl *LessonDatesList) createItem() fyne.CanvasObject {
	dateLabel := widget.NewLabel("date")
	hBox := container.NewHBox(dateLabel)
	return hBox
}

func (ldl *LessonDatesList) updateItem(id int, o fyne.CanvasObject) {
	o.(*fyne.Container).Objects[0].(*widget.Label).SetText(ldl.Lesson.Dates[id].Format(time.DateOnly))
}
