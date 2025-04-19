package routes

const (
	Terms           RoutePath = User + "/terms"
	Term            RoutePath = Terms + RoutePath(ID)
	TermEdit        RoutePath = Term + "/edit"
	TermFiles       RoutePath = Term + "/files"
	TermFile        RoutePath = TermFiles + "/*"
	TermViewFile    RoutePath = Term + "/view-markdown/files/*"
	TermDates       RoutePath = Term + "/dates"
	TermDate        RoutePath = TermDates + "/:date"
	TermCalendar    RoutePath = Term + "/calendar"
	TermEditFile    RoutePath = TermFiles + "/edit/*"
	TermAssessments RoutePath = Term + "/assessments"
	TermAssessment  RoutePath = TermAssessments + RoutePath(ID)
)

// Term handler names
var (
	GetTerms     = NewHandlerName(GET, Terms)
	PostTerm     = NewHandlerName(POST, Terms)
	GetEditTerm  = NewHandlerName(GET, TermEdit)
	PostEditTerm = NewHandlerName(POST, TermEdit)
	DeleteTerm   = NewHandlerName(DELETE, Term)
)

// Term dates handler names (non-instructional dates)
var (
	GetTermDates   = NewHandlerName(GET, TermDates)
	PostTermDate   = NewHandlerName(POST, TermDates)
	DeleteTermDate = NewHandlerName(DELETE, TermDate)
)

// Term file handler names
var (
	GetTermFile      = NewHandlerName(GET, TermFile)
	PostTermFile     = NewHandlerName(POST, TermFile)
	GetTermFiles     = NewHandlerName(GET, TermFiles)
	GetTermEditFile  = NewHandlerName(GET, TermFile)
	PostTermEditFile = NewHandlerName(POST, TermEditFile)
	ViewTermFile     = NewHandlerName(GET, TermViewFile)
)
