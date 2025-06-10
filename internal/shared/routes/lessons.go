package routes

import "gh_static_portfolio/internal/shared/web"

const (
	Lessons              web.RoutePath = Unit + "/lessons"
	NewLesson            web.RoutePath = Lessons + "/new"
	Lesson               web.RoutePath = Lessons + web.RoutePath(LessonID)
	LessonEdit           web.RoutePath = Lesson + "/edit"
	LessonFiles          web.RoutePath = Lesson + "/files"
	LessonFile           web.RoutePath = LessonFiles + "/*"
	LessonViewFile       web.RoutePath = Lesson + "/view-markdown/files/*"
	LessonCreateMarkdown web.RoutePath = Lesson + "/create-markdown"
	LessonEditFile       web.RoutePath = LessonFiles + "/edit/*"
	LessonAssessments    web.RoutePath = Lesson + "/assessments"
	LessonAssessment     web.RoutePath = LessonAssessments + web.RoutePath(AssessmentID)
	LessonSlides         web.RoutePath = Lesson + "/slides"
	LessonEditSlides     web.RoutePath = LessonSlides + "/edit"
	LessonStandards      web.RoutePath = Lesson + "/standards"
	LessonStandard       web.RoutePath = LessonStandards + web.RoutePath(StandardID)
)

// Lesson handler names
var (
	GetLessons     = web.NewHandlerName(web.GET, Lessons)
	GetNewLesson   = web.NewHandlerName(web.GET, NewLesson)
	PostNewLesson  = web.NewHandlerName(web.POST, NewLesson)
	GetLesson      = web.NewHandlerName(web.GET, Lesson)
	GetEditLesson  = web.NewHandlerName(web.GET, LessonEdit)
	PostEditLesson = web.NewHandlerName(web.POST, LessonEdit)
	DeleteLesson   = web.NewHandlerName(web.DELETE, Lesson)
)

// Lesson file handler names
var (
	GetLessonFiles           = web.NewHandlerName(web.GET, LessonFiles)
	PostLessonFile           = web.NewHandlerName(web.POST, LessonFiles)
	GetLessonFile            = web.NewHandlerName(web.GET, LessonFile)
	GetLessonEditFile        = web.NewHandlerName(web.GET, LessonFile)
	PostLessonEditFile       = web.NewHandlerName(web.POST, LessonEditFile)
	ViewLessonFile           = web.NewHandlerName(web.GET, LessonViewFile)
	PostLessonCreateMarkdown = web.NewHandlerName(web.POST, LessonCreateMarkdown)
	DeleteLessonFile = web.NewHandlerName(web.DELETE, LessonFile)
)

// Lesson slides handler names
var (
	GetLessonSlides      = web.NewHandlerName(web.GET, LessonSlides)
	GetEditLessonSlides  = web.NewHandlerName(web.GET, LessonEditSlides)
	PostEditLessonSlides = web.NewHandlerName(web.POST, LessonEditSlides)
)

// Lesson assessment handler names
var (
	GetLessonAssessments     = web.NewHandlerName(web.GET, LessonAssessments)
	PostLessonAssessment     = web.NewHandlerName(web.POST, LessonAssessments)
	GetLessonEditAssessment  = web.NewHandlerName(web.GET, LessonAssessment)
	PostLessonEditAssessment = web.NewHandlerName(web.POST, LessonAssessment)
	DeleteLessonAssessment   = web.NewHandlerName(web.DELETE, LessonAssessment)
)

// Lesson standard handler names
var (
	GetLessonStandards   = web.NewHandlerName(web.GET, LessonStandards)
	PostLessonStandard   = web.NewHandlerName(web.POST, LessonStandards)
	DeleteLessonStandard = web.NewHandlerName(web.DELETE, LessonStandard)
)
