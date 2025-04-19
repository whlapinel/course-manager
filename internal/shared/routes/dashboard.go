package routes

const (
	Users        RoutePath = "/users"
	User         RoutePath = Users + RoutePath(ID)
	UserFiles    RoutePath = User + "/files"
	UserFile     RoutePath = UserFiles + "/*"
	UserViewFile RoutePath = User + "/view-markdown/files/*"
	UserEditFile RoutePath = UserFiles + "/edit/*"
)

var (
	GetUsers         = NewHandlerName(GET, Users)
	GetUser          = NewHandlerName(GET, User)
	GetUserFiles     = NewHandlerName(GET, UserFiles)
	PostUserFile     = NewHandlerName(POST, UserFiles)
	GetUserFile      = NewHandlerName(GET, UserFile)
	GetViewUserFile  = NewHandlerName(GET, UserViewFile)
	GetUserEditFile  = NewHandlerName(GET, UserEditFile)
	PostUserEditFile = NewHandlerName(POST, UserEditFile)
)
