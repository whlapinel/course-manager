package routes

import "gh_static_portfolio/internal/shared/web"

const (
	Courses           web.RoutePath = Term + "/courses"
	NewCourse         web.RoutePath = Courses + "/new"
	Course            web.RoutePath = Courses + web.RoutePath(CourseID)
	CourseEdit        web.RoutePath = Course + "/edit"
	CourseFiles       web.RoutePath = Course + "/files"
	CourseCalendar    web.RoutePath = Course + "/calendar"
	CourseAssessments web.RoutePath = Course + "/assessments"
)

const (
	CourseFile       web.RoutePath = CourseFiles + "/*"
	CourseViewFile   web.RoutePath = Course + "/view-markdown/files/*"
	CourseEditFile   web.RoutePath = CourseFiles + "/edit/*"
	CourseAssessment web.RoutePath = CourseAssessments + web.RoutePath(AssessmentID)
	CourseSlides     web.RoutePath = Course + "/slides"
	CourseEditSlides web.RoutePath = CourseSlides + "/edit"
	CourseStandards  web.RoutePath = Course + "/standards"
	CourseStandard   web.RoutePath = CourseStandards + web.RoutePath(StandardID)
)

// Course handler names
var (
	GetCourses        = web.NewHandlerName(web.GET, Courses)
	GetNewCourse      = web.NewHandlerName(web.GET, NewCourse)
	GetCourseCalendar = web.NewHandlerName(web.GET, CourseCalendar)
	GetCourse         = web.NewHandlerName(web.GET, Course)
	PostCourse        = web.NewHandlerName(web.POST, Courses)
	GetEditCourse     = web.NewHandlerName(web.GET, CourseEdit)
	PostEditCourse    = web.NewHandlerName(web.POST, CourseEdit)
	DeleteCourse      = web.NewHandlerName(web.DELETE, Course)
)

// Course file handler names
var (
	GetCourseFiles     = web.NewHandlerName(web.GET, CourseFiles)
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
