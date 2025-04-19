package routes

const (
	Courses           RoutePath = Term + "/courses"
	Course            RoutePath = Courses + ID
	CourseEdit        RoutePath = Course + "/edit"
	CourseFiles       RoutePath = Course + "/files"
	CourseCalendar    RoutePath = Course + "/calendar"
	CourseAssessments RoutePath = Course + "/assessments"
)

const (
	CourseFile        RoutePath = CourseFiles + "/*"
	CourseViewFile    RoutePath = Course + "/view-markdown/files/*"
	CourseEditFile    RoutePath = CourseFiles + "/edit/*"
	CourseAssessment  RoutePath = CourseAssessments + RoutePath(ID)
	CourseSlides      RoutePath = Course + "/slides"
	CourseEditSlides  RoutePath = CourseSlides + "/edit"
	CourseStandards   RoutePath = Course + "/standards"
	CourseStandard    RoutePath = CourseStandards + ID
)


// Course handler names
var (
	GetCourses     = NewHandlerName(GET, Course)
	PostCourse     = NewHandlerName(POST, Courses)
	GetEditCourse  = NewHandlerName(GET, CourseEdit)
	PostEditCourse = NewHandlerName(POST, CourseEdit)
	DeleteCourse   = NewHandlerName(DELETE, Course)
)

// Course file handler names
var (
	GetCourseFiles     = NewHandlerName(GET, CourseFiles)
	GetCourseFile      = NewHandlerName(GET, CourseFile)
	GetCourseEditFile  = NewHandlerName(GET, CourseFile)
	PostCourseEditFile = NewHandlerName(POST, CourseEditFile)
	ViewCourseFile     = NewHandlerName(GET, CourseViewFile)
)

// Course assessment handler names
var (
	GetCourseAssessments     = NewHandlerName(GET, CourseAssessments)
	PostCourseAssessment     = NewHandlerName(POST, CourseAssessments)
	GetCourseEditAssessment  = NewHandlerName(GET, CourseAssessment)
	PostCourseEditAssessment = NewHandlerName(POST, CourseAssessment)
	DeleteCourseAssessment   = NewHandlerName(DELETE, CourseAssessment)
)

// Course standard handler names
var (
	GetCourseStandards   = NewHandlerName(GET, CourseStandards)
	PostCourseStandard   = NewHandlerName(POST, CourseStandards)
	DeleteCourseStandard = NewHandlerName(DELETE, CourseStandard)
)
