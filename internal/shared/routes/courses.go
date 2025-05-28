package routes

import "gh_static_portfolio/internal/shared/web"

const (
	Courses           web.RoutePath = Term + "/courses"
	NewCourse         web.RoutePath = Courses + "/new"
	Course            web.RoutePath = Courses + web.RoutePath(CourseID)
	CourseEdit        web.RoutePath = Course + "/edit"
	CourseFiles       web.RoutePath = Course + "/files"
	CourseAssessments web.RoutePath = Course + "/assessments"
	CourseFile        web.RoutePath = CourseFiles + "/*"
	CourseViewFile    web.RoutePath = Course + "/view-markdown/files/*"
	CourseEditFile    web.RoutePath = CourseFiles + "/edit/*"
	CourseAssessment  web.RoutePath = CourseAssessments + web.RoutePath(AssessmentID)
	CourseSlides      web.RoutePath = Course + "/slides"
	CourseEditSlides  web.RoutePath = CourseSlides + "/edit"
	CourseStandards   web.RoutePath = Course + "/standards"
	CourseStandard    web.RoutePath = CourseStandards + web.RoutePath(StandardID)
	CourseOccasions   web.RoutePath = Course + "/occasions"
	CourseOccasion    web.RoutePath = CourseOccasions + web.RoutePath(OccasionID)
)

// Course calendar paths
const (
	CourseCalendar web.RoutePath = Course + "/calendar"
	ShiftLesson    web.RoutePath = Lesson + web.RoutePath(ShiftDirection) + web.RoutePath(Date)
	ExtendLesson   web.RoutePath = ShiftLesson + "/extend"
	DateUnits      web.RoutePath = CourseCalendar + web.RoutePath(Date)
	DateLessons    web.RoutePath = DateUnits + web.RoutePath(UnitID)
	LessonDates    web.RoutePath = Lesson + "/dates"
	LessonDate     web.RoutePath = LessonDates + web.RoutePath(Date)
)

// Course handler names
var (
	CreateCourseOccasion   = web.HandlerName(web.POST + CourseOccasions)
	DeleteCourseOccasion   = web.HandlerName(web.DELETE + CourseOccasion)
	ShowEditCourseOccasion = web.HandlerName(web.GET + CourseOccasion)
	PostEditCourseOccasion = web.HandlerName(web.POST + CourseOccasion)
	GetCourses             = web.NewHandlerName(web.GET, Courses)
	GetNewCourse           = web.NewHandlerName(web.GET, NewCourse)
	GetCourse              = web.NewHandlerName(web.GET, Course)
	PostNewCourse          = web.NewHandlerName(web.POST, NewCourse)
	GetEditCourse          = web.NewHandlerName(web.GET, CourseEdit)
	PostEditCourse         = web.NewHandlerName(web.POST, CourseEdit)
	DeleteCourse           = web.NewHandlerName(web.DELETE, Course)
)

// Course calendar handler names
var (
	GetCourseCalendar = web.NewHandlerName(web.GET, CourseCalendar)
	PostShiftLesson   = web.NewHandlerName(web.POST, ShiftLesson)
	PostExtendLesson  = web.NewHandlerName(web.POST, ExtendLesson)
	GetAddLessonDate  = web.NewHandlerName(web.GET, DateUnits)
	GetListLessons    = web.NewHandlerName(web.GET, DateLessons)
	DeleteLessonDate  = web.NewHandlerName(web.DELETE, LessonDate)
	PostAddLessonDate = web.NewHandlerName(web.POST, LessonDate)
)

// Course file handler names
var (
	GetCourseFiles     = web.NewHandlerName(web.GET, CourseFiles)
	PostCourseFile     = web.NewHandlerName(web.POST, CourseFiles)
	GetCourseFile      = web.NewHandlerName(web.GET, CourseFile)
	GetCourseEditFile  = web.NewHandlerName(web.GET, CourseFile)
	PostCourseEditFile = web.NewHandlerName(web.POST, CourseEditFile)
	ViewCourseFile     = web.NewHandlerName(web.GET, CourseViewFile)
)

// Course assessment handler names
var (
	GetCourseAssessments     = web.NewHandlerName(web.GET, CourseAssessments)
	PostCourseAssessment     = web.NewHandlerName(web.POST, CourseAssessments)
	GetCourseEditAssessment  = web.NewHandlerName(web.GET, CourseAssessment)
	PostCourseEditAssessment = web.NewHandlerName(web.POST, CourseAssessment)
	DeleteCourseAssessment   = web.NewHandlerName(web.DELETE, CourseAssessment)
)

// Course standard handler names
var (
	GetCourseStandards   = web.NewHandlerName(web.GET, CourseStandards)
	PostCourseStandard   = web.NewHandlerName(web.POST, CourseStandards)
	DeleteCourseStandard = web.NewHandlerName(web.DELETE, CourseStandard)
)
