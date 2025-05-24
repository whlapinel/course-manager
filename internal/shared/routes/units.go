package routes

import "gh_static_portfolio/internal/shared/web"

const (
	Units           web.RoutePath = Course + "/units"
	NewUnit         web.RoutePath = Units + "/new"
	Unit            web.RoutePath = Units + web.RoutePath(UnitID)
	UnitEdit        web.RoutePath = Unit + "/edit"
	UnitFiles       web.RoutePath = Unit + "/files"
	UnitFile        web.RoutePath = UnitFiles + "/*"
	UnitViewFile    web.RoutePath = Unit + "/view-markdown/files/*"
	UnitEditFile    web.RoutePath = UnitFiles + "/edit/*"
	UnitAssessments web.RoutePath = Unit + "/assessments"
	UnitAssessment  web.RoutePath = UnitAssessments + web.RoutePath(AssessmentID)
	UnitSlides      web.RoutePath = Unit + "/slides"
	UnitEditSlides  web.RoutePath = UnitSlides + "/edit"
	UnitStandards   web.RoutePath = Unit + "/standards"
	UnitStandard    web.RoutePath = UnitStandards + web.RoutePath(StandardID)
)

// Unit handler names
var (
	GetUnit      = web.NewHandlerName(web.GET, Unit)
	GetUnits     = web.NewHandlerName(web.GET, Units)
	GetNewUnit   = web.NewHandlerName(web.GET, NewUnit)
	PostUnit     = web.NewHandlerName(web.POST, Units)
	GetEditUnit  = web.NewHandlerName(web.GET, UnitEdit)
	PostEditUnit = web.NewHandlerName(web.POST, UnitEdit)
	DeleteUnit   = web.NewHandlerName(web.DELETE, Unit)
)

// Unit file handler names
var (
	GetUnitFiles     = web.NewHandlerName(web.GET, UnitFiles)
	PostUnitFile     = web.NewHandlerName(web.POST, UnitFiles)
	GetUnitFile      = web.NewHandlerName(web.GET, UnitFile)
	GetUnitEditFile  = web.NewHandlerName(web.GET, UnitFile)
	PostUnitEditFile = web.NewHandlerName(web.POST, UnitEditFile)
	ViewUnitFile     = web.NewHandlerName(web.GET, UnitViewFile)
)

// Unit assessment handler names
var (
	GetUnitAssessments     = web.NewHandlerName(web.GET, UnitAssessments)
	PostUnitAssessment     = web.NewHandlerName(web.POST, UnitAssessments)
	GetUnitEditAssessment  = web.NewHandlerName(web.GET, UnitAssessment)
	PostUnitEditAssessment = web.NewHandlerName(web.POST, UnitAssessment)
	DeleteUnitAssessment   = web.NewHandlerName(web.DELETE, UnitAssessment)
)

// Unit standard handler names
var (
	GetUnitStandards   = web.NewHandlerName(web.GET, UnitStandards)
	PostUnitStandard   = web.NewHandlerName(web.POST, UnitStandards)
	DeleteUnitStandard = web.NewHandlerName(web.DELETE, UnitStandard)
)
