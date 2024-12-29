package main

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/service"
	"gh_static_portfolio/cmd/util"
	"log"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type CourseCalendar struct {
	Service         service.CourseService
	Course          domain.Course
	CurrentMonthCal *MonthCalendar
	w               fyne.Window
}

func (ch CourseHandler) NewCourseCalendar(svc service.CourseService, course domain.Course, w fyne.Window) (*CourseCalendar, error) {
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
	Lesson       domain.Lesson
	Container    *fyne.Container
	FirstOfMonth time.Time
	Date         time.Time
}

func (cc *CourseCalendar) NewCalLessonHolder(date time.Time, lesson domain.Lesson) *CalLessonHolder {
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
	if err != nil {
		log.Println("error creating resource: ", err)
	}

	lessonContainer := container.NewHBox()
	lessonRemoveButton := widget.NewButton("Remove", func() {
		log.Println("this will remove the current lesson from the current date")
	})
	lessonEdit := widget.NewButton("Edit", func() {
		nameEdit := widget.NewEntry()
		formItem := widget.NewFormItem(lesson.Name, nameEdit)
		form := dialog.NewForm("Edit Lesson", "Ok", "Cancel", []*widget.FormItem{formItem}, func(b bool) {
			if b {
				log.Println("user confirmed!")
			} else {
				log.Println("user canceled!")
			}
			log.Println(nameEdit.Text)
		}, cc.w)
		form.Show()
	})
	shiftRightBtn := widget.NewButton("", onClickShiftLesson(domain.Right))
	shiftLeftBtn := widget.NewButton("", onClickShiftLesson(domain.Left))
	shiftRightBtn.SetIcon(arrowRight)
	shiftLeftBtn.SetIcon(arrowLeft)
	var labelText string
	if len(lesson.Name) > 10 {
		labelText = lesson.Name[:10] + "..."
	} else {
		labelText = lesson.Name
	}
	lessonLabel := widget.NewLabel(labelText)
	lessonContainer.Add(lessonRemoveButton)
	lessonContainer.Add(lessonEdit)
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
			l, newDate, err := cc.Service.Shift(holder.Lesson, cc.Course.Term, cd)
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
				// if it's another month we don't need to update the UI
				if newDate.Month() != holder.Date.Month() {
					return
				}
				holder.Date = newDate
				holder.Lesson = l
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

func (ch *CourseHandler) ShowCourseCalendar(course domain.Course) {
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
