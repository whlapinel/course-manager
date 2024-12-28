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
	CalendarMap map[int]MonthCalendar // month to calendar
	Service     service.CourseService
	Schedule    domain.CourseSchedule
	w           fyne.Window
}

func (ch CourseHandler) NewCourseCalendar(svc service.CourseService, course domain.Course, w fyne.Window) (CourseCalendar, error) {
	var cc CourseCalendar
	schedule, err := svc.GetSchedule(course)
	if err != nil {
		return CourseCalendar{}, err
	}
	calMap := make(map[int]MonthCalendar)
	for _, month := range schedule.Term.TermMonths() {
		monthCal := cc.NewMonthCalendar(schedule, month)
		calMap[int(month.Month())] = monthCal
	}
	cc.CalendarMap = calMap
	cc.Service = svc
	cc.Schedule = schedule
	cc.w = w
	return cc, nil
}

type MonthCalendar struct {
	Calendar *fyne.Container
	Month    time.Time          // first of month
	DateMap  map[int]LessonList // map of date (day) to LessonHolder
}

type CalLessonHolder struct {
	Lesson       domain.Lesson
	Container    *fyne.Container
	FirstOfMonth time.Time
	Date         time.Time
}

func (cc CourseCalendar) NewCalLessonHolder(date time.Time, lesson domain.Lesson) CalLessonHolder {
	var onClickShiftLesson = cc.OnClickShiftLesson(date, lesson)
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
	lessonLabel := widget.NewLabel(lesson.GetTitle())
	lessonContainer.Add(lessonRemoveButton)
	lessonContainer.Add(lessonEdit)
	lessonContainer.Add(shiftLeftBtn)
	lessonContainer.Add(lessonLabel)
	lessonContainer.Add(shiftRightBtn)

	return CalLessonHolder{
		Lesson:    lesson,
		Container: lessonContainer,
	}

}

func (cc CourseCalendar) OnClickShiftLesson(date time.Time, lesson domain.Lesson) func(domain.CalendarDirection) func() {
	return func(cd domain.CalendarDirection) func() {
		return func() {
			log.Println(cd)
			l, newDate, err := cc.Service.Shift(lesson, cc.Schedule.Term, cd)
			if err != nil {
				log.Println("error in OnClickShiftLesson", err)
			}
			currList := cc.CalendarMap[int(date.Month())].DateMap[date.Day()]
			holder, err := currList.RemoveLesson(lesson.ID)
			if err != nil {
				log.Println("error in OnClickShiftLesson", err)
			}
			holder.Lesson = l
			newList := cc.CalendarMap[int(newDate.Month())].DateMap[newDate.Day()]
			newList.AddLesson(holder)
		}
	}

}

type LessonList struct {
	Date          time.Time
	ListContainer *fyne.Container
	LessonHolders []CalLessonHolder
}

func (cc MonthCalendar) NewLessonList(date time.Time, listContainer *fyne.Container) LessonList {
	return LessonList{
		Date:          date,
		ListContainer: listContainer,
	}
}

func (ll LessonList) RemoveLesson(id int) (CalLessonHolder, error) {
	for _, lessonHolder := range ll.LessonHolders {
		if lessonHolder.Lesson.ID == id {
			ll.ListContainer.Remove(lessonHolder.Container)
			ll.ListContainer.Refresh()
			return lessonHolder, nil
		}
	}
	return CalLessonHolder{}, fmt.Errorf("lesson holder not found in this lesson list")

}

func (ll LessonList) AddLesson(l CalLessonHolder) {
	for _, lessonHolder := range ll.LessonHolders {
		if lessonHolder.Lesson.ID == l.Lesson.ID {
			return
		}
	}
	ll.ListContainer.Add(l.Container)
	ll.ListContainer.Refresh()
}

func (ch *CourseHandler) ShowCourseCalendar(course domain.Course) {
	calendar, err := ch.NewCourseCalendar(ch.svc, course, ch.w)
	if err != nil {
		ch.w.SetContent(ErrorMsg(err))
	}
	ch.w.SetContent(calendar.NewTermMonthsList(calendar.Schedule))
}

func (cc *CourseCalendar) NewTermMonthsList(schedule domain.CourseSchedule) *widget.List {
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
		cal := cc.NewMonthCalendar(schedule, schedule.Course.TermMonths()[id])
		cc.ShowMonthCalendar(cal)
	}
	return list
}

func (cc *CourseCalendar) ShowMonthCalendar(mc MonthCalendar) {
	cc.w.SetContent(mc.Calendar)
}

func (cc *CourseCalendar) NewMonthCalendar(schedule domain.CourseSchedule, month time.Time) MonthCalendar {
	var mc MonthCalendar
	dateGrid := container.NewGridWithColumns(5)
	dateMap := make(map[int]LessonList)
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
			if !dailySchedule.Date.IsZero() {
				llContainer := container.NewVBox()
				lessonList := mc.NewLessonList(date, llContainer)
				for _, lesson := range dailySchedule.Lessons {
					lessonHolder := cc.NewCalLessonHolder(date, lesson)
					lessonList.AddLesson(lessonHolder)

				}
				dateContainer.Add(llContainer)
				dateMap[date.Day()] = lessonList
			}
			dateGrid.Add(dateContainer)
		}
	}
	cal := container.NewBorder(widget.NewLabel(month.Format(time.DateOnly)), nil, nil, nil, dateGrid)
	return MonthCalendar{
		Calendar: cal,
		DateMap:  dateMap,
	}
}
