package routes

import "gh_static_portfolio/internal/shared/web"

const (
	Terms           web.RoutePath = User + "/terms"
	NewTerm         web.RoutePath = Terms + "/new"
	Term            web.RoutePath = Terms + web.RoutePath(TermID)
	TermEdit        web.RoutePath = Term + "/edit"
	TermFiles       web.RoutePath = Term + "/files"
	TermFile        web.RoutePath = TermFiles + "/*"
	TermViewFile    web.RoutePath = Term + "/view-markdown/files/*"
	TermDates       web.RoutePath = Term + "/dates"
	TermDate        web.RoutePath = TermDates + web.RoutePath(Date)
	TermCalendar    web.RoutePath = Term + "/calendar"
	TermOccasions   web.RoutePath = Term + "/occasions"
	TermOccasion    web.RoutePath = TermOccasions + web.RoutePath(OccasionID)
	TermEditFile    web.RoutePath = TermFiles + "/edit/*"
	TermAssessments web.RoutePath = Term + "/assessments"
	TermAssessment  web.RoutePath = TermAssessments + web.RoutePath(AssessmentID)
)

// Term handler names
var (
	CreateOccasion    = web.HandlerName(web.POST + TermOccasions)
	DeleteOccasion    = web.HandlerName(web.DELETE + TermOccasion)
	ShowEditOccasion  = web.HandlerName(web.GET + TermOccasion)
	PostEditOccasion  = web.HandlerName(web.POST + TermOccasion)
	ShowEditTermDates = web.HandlerName(web.GET + TermDates)
	PostEditTermDates = web.HandlerName(web.POST + TermDates)
	GetTerms          = web.NewHandlerName(web.GET, Terms)
	GetTerm           = web.NewHandlerName(web.GET, Term)
	GetNewTerm        = web.NewHandlerName(web.GET, NewTerm)
	GetTermCalendar   = web.NewHandlerName(web.GET, TermCalendar)
	PostTerm          = web.NewHandlerName(web.POST, Terms)
	GetEditTerm       = web.NewHandlerName(web.GET, TermEdit)
	PostEditTerm      = web.NewHandlerName(web.POST, TermEdit)
	DeleteTerm        = web.NewHandlerName(web.DELETE, Term)
)

// Term dates handler names (non-instructional dates)
var (
	GetTermDates   = web.NewHandlerName(web.GET, TermDates)
	PostTermDate   = web.NewHandlerName(web.POST, TermDates)
	DeleteTermDate = web.NewHandlerName(web.DELETE, TermDate)
)

// Term file handler names
var (
	GetTermFile      = web.NewHandlerName(web.GET, TermFile)
	PostTermFile     = web.NewHandlerName(web.POST, TermFile)
	GetTermFiles     = web.NewHandlerName(web.GET, TermFiles)
	GetTermEditFile  = web.NewHandlerName(web.GET, TermFile)
	PostTermEditFile = web.NewHandlerName(web.POST, TermEditFile)
	ViewTermFile     = web.NewHandlerName(web.GET, TermViewFile)
)
