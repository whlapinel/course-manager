package main

import (
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/service"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewInstanceHandler(w fyne.Window, svc service.CourseService) *InstanceHandler {
	return &InstanceHandler{w, svc}
}

type InstanceHandler struct {
	w   fyne.Window
	svc service.CourseService
}

func (ih *InstanceHandler) ShowInstancesTree(termID int) {
	instances, err := ih.svc.GetInstances(termID)
	if err != nil {
		ih.w.SetContent(ErrorMsg(err))
	}
	tree := ih.NewCourseInstanceTree(instances, ih.svc)
	ih.w.SetContent(tree.Tree)
}

func (ih *InstanceHandler) NewCourseInstanceTree(instances domain.Instances, svc service.CourseService) CourseTree {
	courses := instances.Courses()
	var ct CourseTree
	ct.Courses = courses
	ct.service = svc
	ct.ShowLessonDates = ih.ShowLessonDates
	ct.Instances = instances
	var courseMap = make(map[int]CourseHolder)
	var unitMap = make(map[int]UnitHolder)
	var lessonMap = make(map[int]LessonHolder)
	for _, course := range courses {
		courseMap[course.ID] = ct.NewCourseHolder(course)
		for _, u := range course.Units {
			unitMap[u.ID] = ct.NewUnitHolder(u)
			for _, l := range u.Lessons {
				lessonMap[l.ID] = ct.NewLessonHolder(l)
			}
		}
	}
	ct.CourseMap = courseMap
	ct.UnitMap = unitMap
	ct.LessonMap = lessonMap
	ct.ShowCourseCalendar = ih.ShowCalendar
	ct.Tree = widget.NewTree(
		ct.childFunc,
		ct.isBranchFunc,
		ct.createNode,
		ct.updateNode,
	)
	return ct
}

func (ch *InstanceHandler) ShowLessonDates(lesson domain.Lesson) {
	datesList := ch.NewLessonDatesList(lesson)
	ch.w.SetContent(datesList.List)
}

func (ch *InstanceHandler) ShowCalendar(instance domain.CourseInstance) {
	schedule, err := ch.svc.GetSchedule(instance)
	if err != nil {
		ch.w.SetContent(ErrorMsg(err))
	}
	cal := ch.NewInstanceCalendar(schedule)
	ch.w.SetContent(cal)

}

func (ch *InstanceHandler) NewInstanceCalendar(schedule domain.CourseSchedule) *widget.List {
	list := widget.NewList(
		func() int {
			return len(schedule.Term.TermMonths())
		},
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel("month header"), container.NewVBox())
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			monthBox := co.(*fyne.Container)
			day := schedule.Term.TermMonths()[lii]
			monthLabel := monthBox.Objects[0].(*widget.Label)
			monthLabel.SetText(day.Month().String() + strconv.Itoa(day.Year()))
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		monthDate := schedule.Term.TermMonths()[id]
		ch.ShowMonthCalendar(schedule, monthDate)
	}
	return list
}

func (ch *InstanceHandler) ShowMonthCalendar(schedule domain.CourseSchedule, month time.Time) {
	weeksContainer := container.NewVBox()
	for _, week := range GetMonthDates(month) {
		weekContainer := container.NewHBox()
		for _, date := range week[1:6] {
			dailySchedule := schedule.GetSchedule(date)
			dateContainer := container.NewVBox()
			if !date.IsZero() {
				dateHeader := widget.NewLabel(date.Format("Mon 1/02/06"))
				addLessonButton := widget.NewButton("Add Lesson", func() {
					log.Println("this will open a dialog to add an existing lesson to the current day")
				})
				dateContainer.Add(dateHeader)
				dateContainer.Add(addLessonButton)
			}
			if dailySchedule != nil {
				lessonList := container.NewVBox()
				for _, lesson := range dailySchedule.Lessons {
					lessonContainer := container.NewHBox()
					lessonRemoveButton := widget.NewButton("Remove", func() {
						log.Println("this will remove the current lesson from the current date")
					})
					lessonEdit := widget.NewButton("Edit", func() {
						nameLabel := widget.NewLabel(lesson.Name)
						nameEdit := widget.NewEntry()
						lessonEditForm := container.New(layout.NewFormLayout(), nameLabel, nameEdit)
						newWindow := fyne.CurrentApp().NewWindow("Edit Lesson")
						newWindow.SetContent(lessonEditForm)
						newWindow.Show()
					})
					lessonLabel := widget.NewLabel(lesson.GetTitle())
					lessonContainer.Add(lessonRemoveButton)
					lessonContainer.Add(lessonEdit)
					lessonContainer.Add(lessonLabel)
					lessonList.Add(lessonContainer)
				}
				dateContainer.Add(lessonList)
			}
			weekContainer.Add(dateContainer)
		}
		weeksContainer.Add(weekContainer)
	}
	cal := container.NewBorder(widget.NewLabel(month.Format(time.DateOnly)), nil, nil, nil, weeksContainer)
	ch.w.SetContent(cal)
}
