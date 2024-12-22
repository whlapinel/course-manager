package main

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewCourseInstanceTree(templates []domain.CourseInstance, svc CourseService) InstanceTree {

	var tt InstanceTree
	tt.Instances = templates
	tt.service = svc
	var templateMap = make(map[int]domain.CourseInstance)
	var unitMap = make(map[int]domain.Unit)
	var lessonMap = make(map[int]domain.Lesson)
	for _, inst := range templates {
		templateMap[inst.CourseTemplate.ID] = inst
		for _, u := range inst.Units {
			unitMap[u.ID] = u
			for _, l := range u.Lessons {
				lessonMap[l.ID] = l
			}
		}
	}
	tt.InstanceMap = templateMap
	tt.UnitMap = unitMap
	tt.LessonMap = lessonMap
	tt.Tree = widget.NewTree(
		tt.childFunc,
		tt.isBranchFunc,
		tt.createNode,
		tt.updateNode,
	)
	return tt
}

type InstanceTree struct {
	Tree        *widget.Tree
	service     CourseService
	Instances   []domain.CourseInstance
	InstanceMap map[int]domain.CourseInstance
	UnitMap     map[int]domain.Unit
	LessonMap   map[int]domain.Lesson
}

func (tt InstanceTree) instanceNodeID(templateID int) string {
	return fmt.Sprintf("T:%d", templateID)
}

func (tt InstanceTree) unitNodeID(unitID int) string {
	return fmt.Sprintf("U:%d", unitID)
}

func (tt InstanceTree) lessonNodeID(lessonID int) string {
	return fmt.Sprintf("L:%d", lessonID)
}

// Parse back from node ID string to (type, intID)
func (t InstanceTree) parseNodeID(nodeID string) (typ string, id int) {
	parts := strings.Split(nodeID, ":")
	if len(parts) != 2 {
		return "", 0
	}
	typ = parts[0] // "T" or "U" or "L"
	id, _ = strconv.Atoi(parts[1])
	return
}

// ChildUIDs function
func (t InstanceTree) childFunc(nodeID string) []string {
	// If nodeID = "", return all top-level templates
	if nodeID == "" {
		var topLevel []string
		for _, tpl := range t.Instances {
			topLevel = append(topLevel, t.instanceNodeID(tpl.CourseTemplate.ID))
		}
		return topLevel
	}

	// Otherwise, parse the nodeID to see if it is a Template or a Unit
	typ, id := t.parseNodeID(nodeID)

	if typ == "T" {
		// It's a template node; return the unit IDs
		units := t.InstanceMap[id].Units
		var unitIDs []string
		for _, u := range units {
			unitIDs = append(unitIDs, t.unitNodeID(u.ID))
		}
		return unitIDs
	}
	if typ == "U" {
		lessons := t.UnitMap[id].Lessons
		var lessonIDs []string
		for _, l := range lessons {
			lessonIDs = append(lessonIDs, t.lessonNodeID(l.ID))
		}
		return lessonIDs
	}

	// If it's a Lesson or unknown, we have no further children
	return nil
}

// IsBranch function
func (t InstanceTree) isBranchFunc(nodeID string) bool {
	// A template node is a branch (if it has units).
	// A unit node is a leaf (no children).
	typ, id := t.parseNodeID(nodeID)

	if nodeID == "" {
		return true // the invisible root
	}
	if typ == "T" {
		// Check if this template has any units
		units := t.InstanceMap[id].Units
		if len(units) > 0 {
			return true
		}
	}
	if typ == "U" {
		// check if this unit has any lessons
		lessons := t.UnitMap[id].Lessons
		if len(lessons) > 0 {
			return true
		}

	}
	return false
}

// CreateNode (factory for node widgets)
func (t InstanceTree) createNode(_ bool) fyne.CanvasObject {
	// Just return a label. If you need different widgets for branch vs. leaf, check `branch` here.
	nameLabel := widget.NewLabel("name")
	descrLabel := widget.NewLabel("description")
	editBtn := widget.NewButton("Edit", nil)
	hbox := container.NewHBox(editBtn, nameLabel, descrLabel)

	return hbox
}

// UpdateNode (sets the label text based on nodeID)
func (t InstanceTree) updateNode(nodeID string, _ bool, obj fyne.CanvasObject) {
	hbox := obj.(*fyne.Container)
	editBtn := hbox.Objects[0].(*widget.Button)
	nameLabel := hbox.Objects[1].(*widget.Label)
	descrLabel := hbox.Objects[2].(*widget.Label)

	// Root node "" will never be rendered, but let's be safe:
	if nodeID == "" {
		nameLabel.SetText("ROOT")  // or something else
		descrLabel.SetText("ROOT") // or something else

		return
	}

	typ, id := t.parseNodeID(nodeID)
	switch typ {
	case "T": // Template node

		inst := t.InstanceMap[id]
		nameLabel.SetText(inst.CourseTemplate.Name)
		descrLabel.SetText(inst.Description)
		editBtn.OnTapped = func() {
			w := fyne.CurrentApp().NewWindow("Edit Template")
			formLayout := layout.NewFormLayout()
			nameEditLabel := widget.NewLabel("Enter new name:")
			nameEdit := widget.NewEntry()
			nameEdit.SetText(inst.CourseTemplate.Name)
			submitBtn := widget.NewButton("Submit", func() {
				log.Println("value of submitted text:", nameEdit.Text)
				inst.CourseTemplate.Name = nameEdit.Text
				err := t.service.UpdateCourseInstance(inst)
				if err != nil {
					log.Fatalf("error updating template: %s", err)
				}
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
		u := t.UnitMap[id]
		nameLabel.SetText(u.Name)
		descrLabel.SetText(u.Description)
		editBtn.OnTapped = func() {
			w := fyne.CurrentApp().NewWindow("Edit Template")
			formLayout := layout.NewFormLayout()
			nameEditLabel := widget.NewLabel("Enter new name:")
			nameEdit := widget.NewEntry()
			nameEdit.SetText(u.Name)
			submitBtn := widget.NewButton("Submit", func() {
				log.Println("value of submitted text:", nameEdit.Text)
				u.Name = nameEdit.Text
				err := t.service.UpdateUnit(u)
				if err != nil {
					log.Fatalf("error updating template: %s", err)
				}

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
		l := t.LessonMap[id]
		nameLabel.SetText(l.Name)
		descrLabel.SetText(l.Description)
		editBtn.OnTapped = func() {
			w := fyne.CurrentApp().NewWindow("Edit Template")
			formLayout := layout.NewFormLayout()
			nameEditLabel := widget.NewLabel("Enter new name:")
			nameEdit := widget.NewEntry()
			nameEdit.SetText(l.Name)
			submitBtn := widget.NewButton("Submit", func() {
				log.Println("value of submitted text:", nameEdit.Text)
				l.Name = nameEdit.Text
				err := t.service.UpdateLesson(l)
				if err != nil {
					log.Fatalf("error updating template: %s", err)
				}
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

func (svc CourseService) GetInstances(termID int) ([]domain.CourseInstance, error) {
	return svc.repo.GetInstances(termID)
}

func (svc CourseService) UpdateCourseInstance(instance domain.CourseInstance) error {
	return svc.repo.UpdateInstance(instance)
}
