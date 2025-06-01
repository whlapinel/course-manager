package routes

import "gh_static_portfolio/internal/shared/web"

const (
	Users        web.RoutePath = "/users"
	User         web.RoutePath = Users + web.RoutePath(UserID)
	UserFiles    web.RoutePath = User + "/files"
	UserFile     web.RoutePath = UserFiles + "/*"
	UserViewFile web.RoutePath = User + "/view-markdown/files/*"
	UserEditFile web.RoutePath = UserFiles + "/edit/*"
)

var (

	GetUsers         = web.NewHandlerName(web.GET, Users)
	GetUser          = web.NewHandlerName(web.GET, User)
	GetUserFiles     = web.NewHandlerName(web.GET, UserFiles)
	PostUserFile     = web.NewHandlerName(web.POST, UserFiles)
	GetUserFile      = web.NewHandlerName(web.GET, UserFile)
	GetViewUserFile  = web.NewHandlerName(web.GET, UserViewFile)
	GetUserEditFile  = web.NewHandlerName(web.GET, UserEditFile)
	PostUserEditFile = web.NewHandlerName(web.POST, UserEditFile)
)
