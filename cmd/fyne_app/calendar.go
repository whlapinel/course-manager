package main

import (
	"errors"
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/service"
	"gh_static_portfolio/cmd/util"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type CourseCalendar struct {
	Service         service.CourseService
	Course          *domain.Course
	CurrentMonthCal *MonthCalendar
	w               fyne.Window
}

func (ch CourseHandler) NewCourseCalendar(svc service.CourseService, course *domain.Course, w fyne.Window) (*CourseCalendar, error) {
	var cc CourseCalendar
	cc.Service = svc
	cc.Course = course
	cc.w = w
	return &cc, nil
}

type MonthCalendar struct {
	Calendar *fyne.Container
	Month    time.Time           // first of month
	DateMap  map[int]*LessonList // map of date (day) to LessonHolder
}

type CalLessonHolder struct {
	Lesson       *domain.Lesson
	Container    *fyne.Container
	FirstOfMonth time.Time
	Date         time.Time
}

func (cc *CourseCalendar) NewCalLessonHolder(date time.Time, lesson *domain.Lesson) *CalLessonHolder {
	var lh CalLessonHolder
	lh.Date = date
	var onClickShiftLesson = cc.OnClickShiftLesson(&lh)
	var file, err = os.ReadFile("./cmd/fyne_app/images/arrow_right.png")
	if err != nil {
		log.Println("error reading file: ", err)
	}
	arrowRight := fyne.NewStaticResource("arrow_right", file)
	file, err = os.ReadFile("./cmd/fyne_app/images/arrow_left.png")
	if err != nil {
		log.Println("error creating resource: ", err)
	}

	arrowLeft := fyne.NewStaticResource("arrow_left", file)
	file, err = os.ReadFile("./cmd/fyne_app/images/delete.png")
	if err != nil {
		log.Println("error creating resource: ", err)
	}

	deleteIcon := fyne.NewStaticResource("delete_icon", file)
	file, err = os.ReadFile("./cmd/fyne_app/images/edit.png")
	if err != nil {
		log.Println("error creating resource: ", err)
	}

	editIcon := fyne.NewStaticResource("edit_icon", file)
	if err != nil {
		log.Println("error creating resource: ", err)
	}

	lessonContainer := container.NewHBox()
	lessonRemoveBtn := widget.NewButton("", func() {
		log.Println("this will remove the current lesson from the current date")
	})
	lessonRemoveBtn.SetIcon(deleteIcon)
	viewFilesBtn2 := widget.NewButton("F2", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if reader == nil {
				return
			}
			path := reader.URI().Path()
			if err != nil {
				dialog.ShowError(err, cc.w)
			}
			err = exec.Command("code", path).Start()
			if err != nil {
				dialog.ShowError(err, cc.w)
			}

		}, cc.w)
		if lesson.Files.ID == 0 {
			dialog.ShowInformation("Info", "No files have been added to this lesson", cc.w)
		}
		path := data.OldLessonFilesDirPath(lesson.Files)
		uri, err := storage.ParseURI("file://" + path)
		if err != nil {
			dialog.ShowError(err, cc.w)
		}
		if uri == nil {
			dialog.ShowError(fmt.Errorf("URI is nil"), cc.w)
		} else {
			listableURI, err := storage.ListerForURI(uri)
			if err != nil {
				dialog.ShowError(err, cc.w)
			}
			fileDialog.SetLocation(listableURI)
			fileDialog.Show()
		}
	})
	if lesson.Files.ID == 0 {
		viewFilesBtn2.Disable()
	}
	viewFilesBtn := widget.NewButton("Files", func() {
		var path = data.OldLessonFilesDirPath(lesson.Files)
		hasFiles := lesson.Files.ID != 0
		_, err := os.Stat(path)
		if !hasFiles || errors.Is(err, fs.ErrNotExist) {
			dialog.ShowInformation("Info", "Directory does not exist, creating and registering directory.", cc.w)
			fileDir, err := cc.Service.CreateNewLessonFileDir(lesson)
			if err != nil {
				dialog.ShowError(err, cc.w)
			}
			path = data.OldLessonFilesDirPath(fileDir)
		}
		log.Println(path)
		// Should open file in new window in VS Code
		infoFile, err := os.Create(filepath.Join(path, "secret_info.txt"))
		if err != nil {
			dialog.ShowError(err, cc.w)
		} else {
			info := fmt.Sprintf(
				`Lesson ID: %d
				Lesson Name: %s
				Lesson Description: %s
				Date this info last updated: %s`,
				lesson.ID, lesson.Name, lesson.Description, time.Now().Format(time.DateOnly))
			infoFile.WriteString(info)
		}
		if lesson.Files.ID != 0 {
			err = exec.Command("code", path).Start()
			if err != nil {
				dialog.ShowError(err, cc.w)
			}
		}
	})
	viewSlidesBtn := widget.NewButton("Slides", func() {
		var path = data.SlidesMarkdownFilePath(lesson.Slides)
		_, err := os.Stat(path)
		if errors.Is(err, fs.ErrNotExist) {
			dialog.ShowInformation("Info", "Slides file does not exist, creating and registering file.", cc.w)
			slides, err := cc.Service.CreateNewLessonSlides(lesson)
			if err != nil {
				dialog.ShowError(err, cc.w)
			}
			path = data.SlidesMarkdownFilePath(slides)
		}
		err = exec.Command("code", path).Start()
		if err != nil {
			dialog.ShowError(err, cc.w)
		}
	})
	lessonEditBtn := widget.NewButton("", func() {
		nameEdit := widget.NewEntry()
		nameEdit.SetText(lesson.Name)
		descrEdit := widget.NewMultiLineEntry()
		descrEdit.SetText(lesson.Description)
		nameItem := widget.NewFormItem("Name", nameEdit)
		descrItem := widget.NewFormItem("Description", descrEdit)
		form := dialog.NewForm("Edit Lesson", "Ok", "Cancel", []*widget.FormItem{nameItem, descrItem}, func(b bool) {
			if b {
				log.Println("user confirmed!")
				log.Println("ID: ", lesson.ID)
				lesson.Name = nameEdit.Text
				log.Println("Name:", lesson.Name)
				lesson.Description = descrEdit.Text
				log.Println("Description:", lesson.Description)
				err := cc.Service.UpdateLesson(*lesson)
				if err != nil {
					dialog.ShowError(err, cc.w)
				} else {
					dialog.ShowInformation("Confirmation", "Lesson updated!", cc.w)
				}
			} else {
				log.Println("user canceled!")
			}
		}, cc.w)
		form.Resize(form.MinSize().Add(fyne.NewSize(200, 200)))
		form.Show()
	})
	lessonEditBtn.SetIcon(editIcon)
	shiftRightBtn := widget.NewButton("", onClickShiftLesson(domain.Right))
	shiftLeftBtn := widget.NewButton("", onClickShiftLesson(domain.Left))
	shiftRightBtn.SetIcon(arrowRight)
	shiftLeftBtn.SetIcon(arrowLeft)
	var labelText string
	if len(lesson.Name) > 15 {
		labelText = lesson.Name[:15] + "..."
	} else {
		labelText = lesson.Name
	}
	lessonLabel := widget.NewLabel(labelText)
	lessonContainer.Add(lessonRemoveBtn)
	lessonContainer.Add(lessonEditBtn)
	lessonContainer.Add(viewSlidesBtn)
	lessonContainer.Add(viewFilesBtn)
	lessonContainer.Add(viewFilesBtn2)
	lessonContainer.Add(shiftLeftBtn)
	lessonContainer.Add(lessonLabel)
	lessonContainer.Add(shiftRightBtn)
	lh.Lesson = lesson
	lh.Container = lessonContainer
	return &lh
}

func (cc *CourseCalendar) OnClickShiftLesson(holder *CalLessonHolder) func(domain.CalendarDirection) func() {
	return func(cd domain.CalendarDirection) func() {
		return func() {
			log.Println(cd)
			l, newDate, err := cc.Service.Shift(*holder.Lesson, cc.Course.Term, cd)
			if err != nil {
				log.Println("error in OnClickShiftLesson", err)
				return
			}
			log.Println("cc.CurrentMonthCal is", cc.CurrentMonthCal)
			if cc.CurrentMonthCal != nil {
				log.Println("cc.CurrentMonthCal.DateMap is", cc.CurrentMonthCal.DateMap)
			}
			currList := cc.CurrentMonthCal.DateMap[holder.Date.Day()]
			if currList == nil {
				log.Fatal("currList is nil")
			}
			err = currList.RemoveLesson(holder.Lesson.ID)
			if err != nil {
				log.Println("error in OnClickShiftLesson", err)
			}
			if holder == nil {
				log.Println(("error: holder is nil after RemoveLesson(); unable to add to new LessonList"))
			} else {
				// need to update the cc.Courses with updated lesson
				if l.ID == 0 {
					dialog.ShowError(fmt.Errorf("shifted lesson ID is 0"), cc.w)
				}
				log.Println("shifted lesson id: ", l.ID)
			outerLoop:
				for i, unit := range cc.Course.Units {
					for k, lesson := range unit.Lessons {
						log.Println("lesson id: ", lesson.ID, "l.ID:", l.ID)
						if lesson.ID == l.ID {
							log.Println("updating cc.Course")
							*cc.Course.Units[i].Lessons[k] = l
							log.Println("updated cc.Course Lesson:", cc.Course.Units[i].Lessons[k].Dates)
							break outerLoop
						} else {
							log.Println("not equal")
						}
					}
				}

				// if it's another month we don't need to update the UI
				if newDate.Month() != holder.Date.Month() {
					return
				}
				holder.Date = newDate
				*holder.Lesson = l

				newList := cc.CurrentMonthCal.DateMap[newDate.Day()]
				newList.AddLesson(holder)
				cc.CurrentMonthCal.Calendar.Refresh()
			}
		}
	}

}

type LessonList struct {
	Date          time.Time
	ListContainer *fyne.Container
	LessonHolders []*CalLessonHolder
}

func (cc *MonthCalendar) NewLessonList(date time.Time) *LessonList {
	return &LessonList{
		Date:          date,
		ListContainer: container.NewVBox(),
		LessonHolders: []*CalLessonHolder{},
	}
}

func (ll *LessonList) RemoveLesson(id int) error {
	log.Println("length of ll.LessonHolders: ", len(ll.LessonHolders))
	if len(ll.LessonHolders) == 0 {
		return fmt.Errorf("LessonHolders empty")
	}
	for i, lessonHolder := range ll.LessonHolders {
		if lessonHolder.Lesson.ID == id {
			ll.LessonHolders = append(ll.LessonHolders[:i], ll.LessonHolders[i+1:]...)
			ll.ListContainer.Remove(lessonHolder.Container)
			ll.ListContainer.Refresh()
			return nil
		}
	}
	return fmt.Errorf("lesson holder not found in this lesson list")

}

func (ll *LessonList) AddLesson(l *CalLessonHolder) {
	if ll == nil {
		log.Fatal("LessonList is nil")
	}
	if ll.LessonHolders == nil {
		log.Fatal("LessonHolders is nil")
	}
	// make sure this isn't a duplicate
	for _, lessonHolder := range ll.LessonHolders {
		if lessonHolder.Lesson.ID == l.Lesson.ID {
			return
		}
	}
	if l.Container == nil {
		log.Fatal("ll.AddLesson: l.Container is nil")
	}
	ll.LessonHolders = append(ll.LessonHolders, l)
	ll.ListContainer.Add(l.Container)
}

func (ch *CourseHandler) ShowCourseCalendar(course *domain.Course) {
	calendar, err := ch.NewCourseCalendar(ch.svc, course, ch.w)
	if err != nil {
		ch.w.SetContent(ErrorMsg(err))
	}
	ch.w.SetContent(calendar.NewTermMonthsList())
	ch.w.SetFixedSize(true)
}

func (cc *CourseCalendar) NewTermMonthsList() *widget.List {
	list := widget.NewList(
		func() int {
			return len(cc.Course.Term.TermMonths())
		},
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel("month header"), container.NewVBox())
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			monthBox := co.(*fyne.Container)
			day := cc.Course.Term.TermMonths()[lii]
			monthLabel := monthBox.Objects[0].(*widget.Label)
			monthLabel.SetText(day.Month().String() + strconv.Itoa(day.Year()))
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		cc.NewMonthCalendar(cc.Course.TermMonths()[id])
		cc.ShowMonthCalendar()
	}
	return list
}

func (cc *CourseCalendar) ShowMonthCalendar() {
	cc.w.SetContent(cc.CurrentMonthCal.Calendar)
}

func (cc *CourseCalendar) NewMonthCalendar(month time.Time) {
	mc := MonthCalendar{}
	mc.Month = month
	dateGrid := container.NewGridWithColumns(5)
	dateMap := make(map[int]*LessonList)
	// GetMonthDates returns a slice of weeks for the months
	// Weeks are slices of dates where the index corresponds to the weekday and the value is the date
	// Begins with Sunday
	for _, week := range util.GetMonthDates(month) {
		for _, calDate := range week[1:6] {
			dateContainer := container.NewVBox()
			// GetMonthDates puts a zero date when date is not within month
			if !calDate.IsZero() {
				dateHeader := widget.NewLabel(calDate.Format("Mon 1/02/06"))
				dateContainer.Add(dateHeader)
				lessonList := mc.NewLessonList(calDate)
				isInstructDay := cc.Course.Term.IsInstructionDay(calDate)
				if isInstructDay {
					addLessonButton := widget.NewButton("Add Lesson", func() {
						log.Println("this will open a dialog to add an existing lesson to the current day")
					})
					dateContainer.Add(addLessonButton)

					for _, unit := range cc.Course.Units {
						for _, lesson := range unit.Lessons {
							for _, lessonDate := range lesson.Dates {
								if domain.IsSameDate(lessonDate, calDate) {
									lessonHolder := cc.NewCalLessonHolder(calDate, lesson)
									lessonList.AddLesson(lessonHolder)
								}
							}
						}
					}
					dateContainer.Add(lessonList.ListContainer)
					dateMap[calDate.Day()] = lessonList
					log.Println(calDate.Format(time.DateOnly))
				} else {
					holidayLabel := widget.NewLabel("No Class")
					dateContainer.Add(holidayLabel)
				}
			}
			dateGrid.Add(dateContainer)

		}

	}
	calHeader := container.NewVBox()
	monthLabel := widget.NewLabel(fmt.Sprintf("%s %s", month.Month().String(), strconv.Itoa(month.Year())))
	calHeader.Add(monthLabel)
	nextMonthBtn := widget.NewButton("Next Month", func() {
		cc.NewMonthCalendar(month.AddDate(0, 1, 0))
		cc.ShowMonthCalendar()
	})
	prevMonthBtn := widget.NewButton("Previous Month", func() {
		cc.NewMonthCalendar(month.AddDate(0, -1, 0))
		cc.ShowMonthCalendar()
	})
	calHeader.Add(prevMonthBtn)
	calHeader.Add(nextMonthBtn)
	cal := container.NewBorder(calHeader, nil, nil, nil, dateGrid)
	mc.Calendar = cal
	mc.DateMap = dateMap
	cc.CurrentMonthCal = &mc
}
