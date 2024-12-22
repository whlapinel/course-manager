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

func NewCourseTemplateTree(templates []domain.CourseTemplate, svc CourseService) TemplateTree {

	var tt TemplateTree
	tt.Templates = templates
	tt.service = svc
	var templateMap = make(map[int]domain.CourseTemplate)
	var unitMap = make(map[int]domain.Unit)
	var lessonMap = make(map[int]domain.Lesson)
	for _, tpl := range templates {
		templateMap[tpl.ID] = tpl
		for _, u := range tpl.Units {
			unitMap[u.ID] = u
			for _, l := range u.Lessons {
				lessonMap[l.ID] = l
			}
		}
	}
	tt.TemplateMap = templateMap
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

type TemplateTree struct {
	Tree        *widget.Tree
	service     CourseService
	Templates   []domain.CourseTemplate
	TemplateMap map[int]domain.CourseTemplate
	UnitMap     map[int]domain.Unit
	LessonMap   map[int]domain.Lesson
}

func (tt TemplateTree) templateNodeID(templateID int) string {
	return fmt.Sprintf("T:%d", templateID)
}

func (tt TemplateTree) unitNodeID(unitID int) string {
	return fmt.Sprintf("U:%d", unitID)
}

func (tt TemplateTree) lessonNodeID(lessonID int) string {
	return fmt.Sprintf("L:%d", lessonID)
}

// Parse back from node ID string to (type, intID)
func (t TemplateTree) parseNodeID(nodeID string) (typ string, id int) {
	parts := strings.Split(nodeID, ":")
	if len(parts) != 2 {
		return "", 0
	}
	typ = parts[0] // "T" or "U" or "L"
	id, _ = strconv.Atoi(parts[1])
	return
}

// ChildUIDs function
func (t TemplateTree) childFunc(nodeID string) []string {
	// If nodeID = "", return all top-level templates
	if nodeID == "" {
		var topLevel []string
		for _, tpl := range t.Templates {
			topLevel = append(topLevel, t.templateNodeID(tpl.ID))
		}
		return topLevel
	}

	// Otherwise, parse the nodeID to see if it is a Template or a Unit
	typ, id := t.parseNodeID(nodeID)

	if typ == "T" {
		// It's a template node; return the unit IDs
		units := t.TemplateMap[id].Units
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
func (t TemplateTree) isBranchFunc(nodeID string) bool {
	// A template node is a branch (if it has units).
	// A unit node is a leaf (no children).
	typ, id := t.parseNodeID(nodeID)

	if nodeID == "" {
		return true // the invisible root
	}
	if typ == "T" {
		// Check if this template has any units
		units := t.TemplateMap[id].Units
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
func (t TemplateTree) createNode(_ bool) fyne.CanvasObject {
	// Just return a label. If you need different widgets for branch vs. leaf, check `branch` here.
	nameLabel := widget.NewLabel("name")
	descrLabel := widget.NewLabel("description")
	editBtn := widget.NewButton("Edit", nil)
	hbox := container.NewHBox(editBtn, nameLabel, descrLabel)

	return hbox
}

// UpdateNode (sets the label text based on nodeID)
func (t TemplateTree) updateNode(nodeID string, _ bool, obj fyne.CanvasObject) {
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

		tpl := t.TemplateMap[id]
		nameLabel.SetText(tpl.Name)
		descrLabel.SetText(tpl.Description)
		editBtn.OnTapped = func() {
			w := fyne.CurrentApp().NewWindow("Edit Template")
			formLayout := layout.NewFormLayout()
			nameEditLabel := widget.NewLabel("Enter new name:")
			nameEdit := widget.NewEntry()
			nameEdit.SetText(tpl.Name)
			submitBtn := widget.NewButton("Submit", func() {
				log.Println("value of submitted text:", nameEdit.Text)
				tpl.Name = nameEdit.Text
				err := t.service.UpdateCourseTemplate(tpl)
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

func (svc CourseService) GetTemplates() ([]domain.CourseTemplate, error) {
	templates, err := svc.repo.GetTemplates()
	if err != nil {
		return nil, err
	}
	log.Println(len(templates))
	return templates, nil
}
func (svc CourseService) UpdateCourseTemplate(tpl domain.CourseTemplate) error {
	err := svc.repo.UpdateTemplate(tpl)
	if err != nil {
		return err
	}
	return nil
}
