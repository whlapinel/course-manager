package main

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/service"
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

func NewCourseHandler(w fyne.Window, svc service.CourseService) *CourseHandler {
	return &CourseHandler{w, svc}
}

type CourseHandler struct {
	w   fyne.Window
	svc service.CourseService
}

func (ch *CourseHandler) ShowCourseTree() {
	templates, err := ch.svc.GetTemplates()
	if err != nil {
		ch.w.SetContent(ErrorMsg(err))
	}
	tree := ch.NewCourseTree(templates, ch.svc)
	ch.w.SetContent(tree.Tree)
}

func (ch *CourseHandler) NewCourseTree(courses []domain.Course, svc service.CourseService) CourseTree {

	var ct CourseTree
	ct.Courses = courses
	ct.service = svc
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
	ct.Tree = widget.NewTree(
		ct.childFunc,
		ct.isBranchFunc,
		ct.createNode,
		ct.updateNode,
	)
	return ct
}

type CourseTree struct {
	Tree               *widget.Tree
	Instances          domain.Instances
	service            service.CourseService
	ShowLessonDates    func(domain.Lesson)
	ShowCourseCalendar func(instance domain.CourseInstance)
	Courses            []domain.Course
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
	return fmt.Sprintf("T:%d", courseID)
}

func (ct *CourseTree) unitNodeID(unitID int) string {
	return fmt.Sprintf("U:%d", unitID)
}

func (ct *CourseTree) lessonNodeID(lessonID int) string {
	return fmt.Sprintf("L:%d", lessonID)
}

// Parse back from node ID string to (type, intID)
func (ct *CourseTree) parseNodeID(nodeID string) (typ string, id int) {
	parts := strings.Split(nodeID, ":")
	if len(parts) != 2 {
		return "", 0
	}
	typ = parts[0] // "T" or "U" or "L"
	id, _ = strconv.Atoi(parts[1])
	return
}

// ChildUIDs function
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

	if typ == "T" {
		// It's a template node; return the unit IDs
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

// IsBranch function
func (ct *CourseTree) isBranchFunc(nodeID string) bool {
	// A template node is a branch (if it has units).
	// A unit node is a leaf (no children).
	typ, id := ct.parseNodeID(nodeID)

	if nodeID == "" {
		return true // the invisible root
	}
	if typ == "T" {
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

// CreateNode (factory for node widgets)
func (ct *CourseTree) createNode(_ bool) fyne.CanvasObject {
	// Just return a label. If you need different widgets for branch vs. leaf, check `branch` here.
	nameLabel := widget.NewLabel("name")
	descrLabel := widget.NewLabel("description")
	editBtn := widget.NewButton("Edit", nil)
	var hbox *fyne.Container
	if len(ct.Instances) != 0 {
		calendarButton := widget.NewButton("Course Calendar", nil)
		hbox = container.NewHBox(calendarButton, editBtn, nameLabel, descrLabel)
	} else {
		hbox = container.NewHBox(editBtn, nameLabel, descrLabel)
	}
	return hbox
}

// UpdateNode (sets the label text based on nodeID)
func (ct *CourseTree) updateNode(nodeID string, _ bool, obj fyne.CanvasObject) {
	hbox := obj.(*fyne.Container)
	var calButton *widget.Button
	var editBtn *widget.Button
	var nameLabel *widget.Label
	var descrLabel *widget.Label
	if len(ct.Instances) != 0 {
		calButton = hbox.Objects[0].(*widget.Button)
		editBtn = hbox.Objects[1].(*widget.Button)
		nameLabel = hbox.Objects[2].(*widget.Label)
		descrLabel = hbox.Objects[3].(*widget.Label)
	} else {
		editBtn = hbox.Objects[0].(*widget.Button)
		nameLabel = hbox.Objects[1].(*widget.Label)
		descrLabel = hbox.Objects[2].(*widget.Label)
	}

	// Root node "" will never be rendered, but let's be safe:
	if nodeID == "" {
		nameLabel.SetText("ROOT")  // or something else
		descrLabel.SetText("ROOT") // or something else

		return
	}

	typ, id := ct.parseNodeID(nodeID)
	switch typ {
	case "T": // Template node

		holder := ct.CourseMap[id]
		nameLabel.Bind(holder.NameBinding)
		descrLabel.Bind(holder.DescrBinding)
		if calButton != nil {
			calButton.OnTapped = func() {
				log.Println("this will show the course calendar!")
				ct.ShowCourseCalendar(ct.Instances[0]) // should only be one instance if it's an instance tree!
			}
		}
		editBtn.OnTapped = func() {
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
				err = ct.service.UpdateCourseTemplate(holder.course)
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
	case "U": // Unit node
		// you can store units in a map as well for quick lookup
		holder := ct.UnitMap[id]
		nameLabel.Bind(holder.NameBinding)
		descrLabel.Bind(holder.DescrBinding)
		editBtn.OnTapped = func() {
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
		// you can store lessons in a map as well for quick lookup
		holder := ct.LessonMap[id]
		if len(holder.Lesson.Dates) > 0 {
			datesBox := container.NewHBox()
			for _, date := range holder.Lesson.Dates {
				dateLabel := widget.NewLabel(date.Format(time.DateOnly))
				datesBox.Add(dateLabel)
			}
			if len(hbox.Objects) > 3 {
				hbox.Objects[3] = datesBox
			} else {
				hbox.Add(datesBox)
			}
		}
		nameLabel.Bind(holder.NameBinding)
		descrLabel.Bind(holder.DescrBinding)
		editBtn.OnTapped = func() {
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
		nameLabel.SetText("Unknown")
		descrLabel.SetText("Unknown")
	}

}
