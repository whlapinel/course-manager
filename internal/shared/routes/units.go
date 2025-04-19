package routes

const (
	Units           RoutePath = Course + "/Units"
	Unit            RoutePath = Units + RoutePath(ID)
	UnitEdit        RoutePath = Unit + "/edit"
	UnitFiles       RoutePath = Unit + "/files"
	UnitFile        RoutePath = UnitFiles + "/*"
	UnitViewFile    RoutePath = Unit + "/view-markdown/files/*"
	UnitEditFile    RoutePath = UnitFiles + "/edit/*"
	UnitAssessments RoutePath = Unit + "/assessments"
	UnitAssessment  RoutePath = UnitAssessments + RoutePath(ID)
	UnitSlides      RoutePath = Unit + "/slides"
	UnitEditSlides  RoutePath = UnitSlides + "/edit"
	UnitStandards   RoutePath = Unit + "/standards"
	UnitStandard    RoutePath = UnitStandards + ID
)

// Unit handler names
var (
	GetUnits     = NewHandlerName(GET, Unit)
	PostUnit     = NewHandlerName(POST, Units)
	GetEditUnit  = NewHandlerName(GET, UnitEdit)
	PostEditUnit = NewHandlerName(POST, UnitEdit)
	DeleteUnit   = NewHandlerName(DELETE, Unit)
)

// Unit file handler names
var (
	GetUnitFiles     = NewHandlerName(GET, UnitFiles)
	GetUnitFile      = NewHandlerName(GET, UnitFile)
	GetUnitEditFile  = NewHandlerName(GET, UnitFile)
	PostUnitEditFile = NewHandlerName(POST, UnitEditFile)
	ViewUnitFile     = NewHandlerName(GET, UnitViewFile)
)

// Unit assessment handler names
var (
	GetUnitAssessments     = NewHandlerName(GET, UnitAssessments)
	PostUnitAssessment     = NewHandlerName(POST, UnitAssessments)
	GetUnitEditAssessment  = NewHandlerName(GET, UnitAssessment)
	PostUnitEditAssessment = NewHandlerName(POST, UnitAssessment)
	DeleteUnitAssessment   = NewHandlerName(DELETE, UnitAssessment)
)

// Unit standard handler names
var (
	GetUnitStandards   = NewHandlerName(GET, UnitStandards)
	PostUnitStandard   = NewHandlerName(POST, UnitStandards)
	DeleteUnitStandard = NewHandlerName(DELETE, UnitStandard)
)
