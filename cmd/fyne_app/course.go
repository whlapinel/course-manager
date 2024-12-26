package main

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/service"
	"gh_static_portfolio/cmd/util"
	"log"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

func (ih *InstanceHandler) ShowCourseTree(termID int) {
	instances, err := ih.svc.GetCourses(termID)
	if err != nil {
		ih.w.SetContent(ErrorMsg(err))
	}
	tree := ih.NewCourseTree(instances, ih.svc)
	ih.w.SetContent(tree.Tree)
}

func (ih *InstanceHandler) NewCourseTree(instances domain.Courses, svc service.CourseService) CourseTree {
	courses := instances.Courses()
	var ct CourseTree
	ct.Courses = courses
	ct.service = svc
	ct.Courses = instances
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
	ct.ShowCourseCalendar = ih.ShowTermMonthsList
	ct.Tree = widget.NewTree(
		ct.childFunc,
		ct.isBranchFunc,
		ct.createNode,
		ct.updateNode,
	)
	return ct
}

func (ch *InstanceHandler) ShowTermMonthsList(instance domain.Course) {
	schedule, err := ch.svc.GetSchedule(instance)
	if err != nil {
		ch.w.SetContent(ErrorMsg(err))
	}
	cal := ch.NewTermMonthsList(schedule)
	ch.w.SetContent(cal)

}

func (ch *InstanceHandler) NewTermMonthsList(schedule domain.CourseSchedule) *widget.List {
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
	dateGrid := container.NewGridWithColumns(5)
	for _, week := range util.GetMonthDates(month) {
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
			dateGrid.Add(dateContainer)
		}
	}
	cal := container.NewBorder(widget.NewLabel(month.Format(time.DateOnly)), nil, nil, nil, dateGrid)
	ch.w.SetContent(cal)
}

type CourseTree struct {
	Tree               *widget.Tree
	Courses            domain.Courses
	service            service.CourseService
	ShowLessonDates    func(domain.Lesson)
	ShowCourseCalendar func(course domain.Course)
	CourseMap          map[int]CourseHolder
	UnitMap            map[int]UnitHolder
	LessonMap          map[int]LessonHolder
}

type CourseHolder struct {
	course       domain.Course
	NameBinding  binding.String
	DescrBinding binding.String
}

type UnitHolder struct {
	Unit         domain.Unit
	NameBinding  binding.String
	DescrBinding binding.String
}

type LessonHolder struct {
	Lesson       domain.Lesson
	NameBinding  binding.String
	DescrBinding binding.String
}

func (ct CourseTree) NewCourseHolder(tpl domain.Course) CourseHolder {
	holder := CourseHolder{
		course:       tpl,
		NameBinding:  binding.NewString(),
		DescrBinding: binding.NewString(),
	}
	holder.NameBinding.Set(tpl.Name)
	holder.DescrBinding.Set(tpl.Description)
	return holder
}

func (ct *CourseTree) NewUnitHolder(u domain.Unit) UnitHolder {
	holder := UnitHolder{
		Unit:         u,
		NameBinding:  binding.NewString(),
		DescrBinding: binding.NewString(),
	}
	holder.NameBinding.Set(u.Name)
	holder.DescrBinding.Set(u.Description)
	return holder
}
func (ct *CourseTree) NewLessonHolder(l domain.Lesson) LessonHolder {
	holder := LessonHolder{
		Lesson:       l,
		NameBinding:  binding.NewString(),
		DescrBinding: binding.NewString(),
	}
	holder.NameBinding.Set(l.Name)
	holder.DescrBinding.Set(l.Description)
	return holder
}

func (ct *CourseTree) courseNodeID(courseID int) string {
	return fmt.Sprintf("C:%d", courseID)
}

func (ct *CourseTree) unitNodeID(unitID int) string {
	return fmt.Sprintf("U:%d", unitID)
}

func (ct *CourseTree) lessonNodeID(lessonID int) string {
	return fmt.Sprintf("L:%d", lessonID)
}

func (ct *CourseTree) parseNodeID(nodeID string) (typ string, id int) {
	parts := strings.Split(nodeID, ":")
	if len(parts) != 2 {
		return "", 0
	}
	typ = parts[0] // "C" or "U" or "L"
	id, _ = strconv.Atoi(parts[1])
	return
}

func (ct *CourseTree) childFunc(nodeID string) []string {
	// If nodeID = "", return all top-level templates
	if nodeID == "" {
		var topLevel []string
		for _, tpl := range ct.Courses {
			topLevel = append(topLevel, ct.courseNodeID(tpl.ID))
		}
		return topLevel
	}

	// Otherwise, parse the nodeID to see if it is a Template or a Unit
	typ, id := ct.parseNodeID(nodeID)

	if typ == "C" {
		units := ct.CourseMap[id].course.Units
		var unitIDs []string
		for _, u := range units {
			unitIDs = append(unitIDs, ct.unitNodeID(u.ID))
		}
		return unitIDs
	}
	if typ == "U" {
		lessons := ct.UnitMap[id].Unit.Lessons
		var lessonIDs []string
		for _, l := range lessons {
			lessonIDs = append(lessonIDs, ct.lessonNodeID(l.ID))
		}
		return lessonIDs
	}

	// If it's a Lesson or unknown, we have no further children
	return nil
}

func (ct *CourseTree) isBranchFunc(nodeID string) bool {
	// A template node is a branch (if it has units).
	// A unit node is a leaf (no children).
	typ, id := ct.parseNodeID(nodeID)

	if nodeID == "" {
		return true // the invisible root
	}
	if typ == "C" {
		// Check if this template has any units
		units := ct.CourseMap[id].course.Units
		if len(units) > 0 {
			return true
		}
	}
	if typ == "U" {
		// check if this unit has any lessons
		lessons := ct.UnitMap[id].Unit.Lessons
		if len(lessons) > 0 {
			return true
		}

	}
	return false
}

type courseTreeRow struct {
	box         *fyne.Container
	calendarBtn *widget.Button
	editBtn     *widget.Button
	nameLabel   *widget.Label
	descrLabel  *widget.Label
	datesBox    *fyne.Container
}

func (ct *CourseTree) createNode(_ bool) fyne.CanvasObject {
	// Just return a label. If you need different widgets for branch vs. leaf, check `branch` here.
	row := &courseTreeRow{}
	row.calendarBtn = widget.NewButton("Calendar", nil)
	row.editBtn = widget.NewButton("Edit", nil)
	row.nameLabel = widget.NewLabel("")
	row.descrLabel = widget.NewLabel("")
	row.datesBox = container.NewHBox()
	row.box = container.NewHBox(
		row.calendarBtn,
		row.editBtn,
		row.nameLabel,
		row.descrLabel,
		row.datesBox,
	)
	// hide or show as needed
	return row.box
}

func (ct *CourseTree) getCourseRowFromContainer(container *fyne.Container) courseTreeRow {
	var row courseTreeRow
	row.box = container
	row.calendarBtn = container.Objects[0].(*widget.Button)
	row.editBtn = container.Objects[1].(*widget.Button)
	row.nameLabel = container.Objects[2].(*widget.Label)
	row.descrLabel = container.Objects[3].(*widget.Label)
	row.datesBox = container.Objects[4].(*fyne.Container)
	return row

}

func (ct *CourseTree) updateNode(nodeID string, _ bool, obj fyne.CanvasObject) {
	row := ct.getCourseRowFromContainer(obj.(*fyne.Container))

	if nodeID == "" {
		row.nameLabel.SetText("ROOT")  // or something else
		row.descrLabel.SetText("ROOT") // or something else

		return
	}

	typ, id := ct.parseNodeID(nodeID)
	switch typ {
	case "C": // Template node
		holder := ct.CourseMap[id]
		row.datesBox.Hide()
		row.nameLabel.Bind(holder.NameBinding)
		row.descrLabel.Bind(holder.DescrBinding)
		row.calendarBtn.OnTapped = func() {
			ct.ShowCourseCalendar(ct.Courses[0]) // should only be one instance if it's an instance tree!
		}
		if row.editBtn.OnTapped == nil {
			row.editBtn.OnTapped = func() {
				w := fyne.CurrentApp().NewWindow("Edit Template")
				formLayout := layout.NewFormLayout()
				nameEditLabel := widget.NewLabel("Enter new name:")
				name, err := holder.NameBinding.Get()
				if err != nil {
					log.Fatalf("error retrieving name from binding: %s", err)
				}
				nameEdit := widget.NewEntry()
				nameEdit.SetText(name)
				submitBtn := widget.NewButton("Submit", func() {
					log.Println("value of submitted text:", nameEdit.Text)
					// Need to do some sort of input validation here
					holder.course.Name = nameEdit.Text
					err = ct.service.UpdateCourse(holder.course)
					if err != nil {
						log.Fatalf("error updating template: %s", err)
					}
					holder.NameBinding.Set(holder.course.Name)
					ct.Tree.Refresh()
					w.Close()
				})
				formContainer := container.New(formLayout, nameEditLabel, nameEdit)
				vBox := container.NewVBox(formContainer, submitBtn)
				w.Resize(fyne.Size{Width: 500, Height: 500})
				w.SetContent(vBox)
				w.Show()
			}
			return
		}
	case "U": // Unit node
		// you can store units in a map as well for quick lookup
		holder := ct.UnitMap[id]
		row.calendarBtn.Hide()
		row.nameLabel.Bind(holder.NameBinding)
		row.descrLabel.Bind(holder.DescrBinding)
		row.editBtn.OnTapped = func() {
			w := fyne.CurrentApp().NewWindow("Edit Unit")
			formLayout := layout.NewFormLayout()
			nameEditLabel := widget.NewLabel("Enter new name:")
			name, err := holder.NameBinding.Get()
			if err != nil {
				log.Fatalf("error retrieving name from binding: %s", err)
			}
			nameEdit := widget.NewEntry()
			nameEdit.SetText(name)
			submitBtn := widget.NewButton("Submit", func() {
				log.Println("value of submitted text:", nameEdit.Text)
				// Need to do some sort of input validation here
				holder.Unit.Name = nameEdit.Text
				err = ct.service.UpdateUnit(holder.Unit)
				if err != nil {
					log.Fatalf("error updating unit: %s", err)
				}
				holder.NameBinding.Set(holder.Unit.Name)
				ct.Tree.Refresh()
				w.Close()
			})
			formContainer := container.New(formLayout, nameEditLabel, nameEdit)
			vBox := container.NewVBox(formContainer, submitBtn)
			w.Resize(fyne.Size{Width: 500, Height: 500})
			w.SetContent(vBox)
			w.Show()
		}

		return
	case "L": // Lesson node
		row.calendarBtn.Hide()
		holder := ct.LessonMap[id]
		if len(holder.Lesson.Dates) > 0 {
			newDatesBox := container.NewHBox()
			for _, date := range holder.Lesson.Dates {
				log.Println(date.Format(time.DateOnly))
				dateLabel := widget.NewLabel(date.Format(time.DateOnly))
				newDatesBox.Add(dateLabel)
			}
			for i, child := range row.box.Objects {
				if child == row.datesBox {
					row.box.Objects[i] = newDatesBox
					break
				}
			}
		}
		row.nameLabel.Bind(holder.NameBinding)
		row.descrLabel.Bind(holder.DescrBinding)
		row.editBtn.OnTapped = func() {
			w := fyne.CurrentApp().NewWindow("Edit Unit")
			formLayout := layout.NewFormLayout()
			nameEditLabel := widget.NewLabel("Enter new name:")
			name, err := holder.NameBinding.Get()
			if err != nil {
				log.Fatalf("error retrieving name from binding: %s", err)
			}
			nameEdit := widget.NewEntry()
			nameEdit.SetText(name)
			submitBtn := widget.NewButton("Submit", func() {
				log.Println("value of submitted text:", nameEdit.Text)
				// Need to do some sort of input validation here
				holder.Lesson.Name = nameEdit.Text
				err = ct.service.UpdateLesson(holder.Lesson)
				if err != nil {
					log.Fatalf("error updating unit: %s", err)
				}
				holder.NameBinding.Set(holder.Lesson.Name)
				ct.Tree.Refresh()
				w.Close()
			})
			formContainer := container.New(formLayout, nameEditLabel, nameEdit)
			vBox := container.NewVBox(formContainer, submitBtn)
			w.Resize(fyne.Size{Width: 500, Height: 500})
			w.SetContent(vBox)
			w.Show()
		}
		return
	default:
		row.nameLabel.SetText("Unknown")
		row.descrLabel.SetText("Unknown")
	}

}
