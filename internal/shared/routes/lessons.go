package routes

const (
	Lessons           RoutePath = Unit + "/lessons"
	Lesson            RoutePath = Lessons + RoutePath(ID)
	LessonEdit        RoutePath = Lesson + "/edit"
	LessonFiles       RoutePath = Lesson + "/files"
	LessonFile        RoutePath = LessonFiles + "/*"
	LessonViewFile    RoutePath = Lesson + "/view-markdown/files/*"
	LessonEditFile    RoutePath = LessonFiles + "/edit/*"
	LessonAssessments RoutePath = Lesson + "/assessments"
	LessonAssessment  RoutePath = LessonAssessments + RoutePath(ID)
	LessonSlides      RoutePath = Lesson + "/slides"
	LessonEditSlides  RoutePath = LessonSlides + "/edit"
	LessonStandards   RoutePath = Lesson + "/standards"
	LessonStandard    RoutePath = LessonStandards + ID
)

// Lesson routes
var (
	GetLessons     = NewHandlerName(GET, Lesson)
	PostLesson     = NewHandlerName(POST, Lessons)
	GetEditLesson  = NewHandlerName(GET, LessonEdit)
	PostEditLesson = NewHandlerName(POST, LessonEdit)
	DeleteLesson   = NewHandlerName(DELETE, Lesson)
)

// Lesson file routes
var (
	GetLessonFiles     = NewHandlerName(GET, LessonFiles)
	GetLessonFile      = NewHandlerName(GET, LessonFile)
	GetLessonEditFile  = NewHandlerName(GET, LessonFile)
	PostLessonEditFile = NewHandlerName(POST, LessonEditFile)
	ViewLessonFile     = NewHandlerName(GET, LessonViewFile)
)

// Lesson slides routes
var (
	GetLessonSlides      = NewHandlerName(GET, LessonSlides)
	GetEditLessonSlides  = NewHandlerName(GET, LessonEditSlides)
	PostEditLessonSlides = NewHandlerName(POST, LessonEditSlides)
)

// Lesson assessment routes
var (
	GetLessonAssessments     = NewHandlerName(GET, LessonAssessments)
	PostLessonAssessment     = NewHandlerName(POST, LessonAssessments)
	GetLessonEditAssessment  = NewHandlerName(GET, LessonAssessment)
	PostLessonEditAssessment = NewHandlerName(POST, LessonAssessment)
	DeleteLessonAssessment   = NewHandlerName(DELETE, LessonAssessment)
)

// Lesson standard routes
var (
	GetLessonStandards   = NewHandlerName(GET, LessonStandards)
	PostLessonStandard   = NewHandlerName(POST, LessonStandards)
	DeleteLessonStandard = NewHandlerName(DELETE, LessonStandard)
)
