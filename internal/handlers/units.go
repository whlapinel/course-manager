package handlers

import (
	"bytes"
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	"gh_static_portfolio/internal/templates"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func UnitHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	nodeRouter := NewUnitRouter(svc, router)
	var routeHandlers []RouteHandler
	// ur := nodeRouter.(*unitRouter)
	unitRouteHandlers := []RouteHandler{}
	routeHandlers = append(routeHandlers, unitRouteHandlers...)
	routeHandlers = append(routeHandlers, NodeHandlers(nodeRouter)...)
	return routeHandlers
}

type unitRouter struct {
	Router
}

// Delete implements NodeRouter.
func (r *unitRouter) Delete(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	return r.svc.DeleteUnit(r.params.UnitID.Value.(int))
}

// ListChildren implements NodeRouter.
func (r *unitRouter) ListChildren(c echo.Context) error {
	// 	params := ParseCourseIDParams(c)
	// 	user, err := r.svc.GetUser(params.UserID.Value.(string))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	unitID, err := ParseRouteParam(c, UnitID)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// 	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	unit, err := h.svc.GetUnit(unitID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	lessons, err := h.svc.GetLessons(unitID)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// 	var nodes []domain.CourseNode
	// 	for _, lesson := range lessons {
	// 		nodes = append(nodes, lesson)
	// 	}
	// 	lessonList := mt.NodeListPage{
	// 		Params:          params,
	// 		ParentNode:      unit,
	// 		Children:        nodes,
	// 		ChildDetailsRHN: LessonDetails.String(),
	// 		ShowNewChildURL: ShowNewLesson.String(),
	// 		DeleteChildRHN:  DeleteLesson.String(),
	// 		UpNavURL:        h.e.Reverse(ListCourseUnits.String(), params.ToSlice()...),
	// 		E:               h.e,
	// 		BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit),
	// 	}
	// 	template := mt.LessonListTemplate(lessonList)
	// 	layout := h.CourseManagerLayout(template, user)
	// 	return Respond(c, "", template, layout)

	panic("not implemented")
}

// PostEdit implements NodeRouter.
func (r *unitRouter) PostEdit(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := r.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	unitID, err := UnitIDParam(params)
	if err != nil {
		return err
	}
	term, err := r.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := r.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := r.svc.GetUnit(unitID)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	var updateUnit = func() error {
		err := r.svc.UpdateUnit(unit)
		if err != nil {
			return err
		}
		updatedUnit, err := r.svc.GetUnit(unitID)
		if err != nil {
			return err
		}
		unit = updatedUnit
		return nil
	}
	var pageData = func(unit domain.Unit) mt.NodeDetailsPage {
		return mt.NodeDetailsPage{
			Node:            unit,
			Params:          params,
			GetEditNodeURL:  r.app.Reverse(ShowEditUnit.String(), params.ToSlice()...),
			PostEditNodeURL: r.app.Reverse(PostEditUnit.String(), params.ToSlice()...),
			ListChildrenURL: r.app.Reverse(ListUnitLessons.String(), params.ToSlice()...),
			IsEdit:          false,
			BreadCrumbsData: BreadCrumbs(r.app, params, user, term, course, unit),
		}
	}
	var template templ.Component
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "description":
			unit.Description = val[0]
			err := updateUnit()
			if err != nil {
				return err
			}
			details := pageData(unit)
			template = mt.EditDescriptionComponent(details)
		case "name":
			unit.Name = val[0]
			err := updateUnit()
			if err != nil {
				return err
			}
			details := pageData(unit)
			template = mt.EditNameComponent(details)
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}

	}
	if template == nil {
		panic("template is nil!")
	}
	return Respond(c, r.app.Reverse(UnitDetails.String(), params.ToSlice()...), template, nil)

}

// PostFile implements NodeRouter.
func (r *unitRouter) PostFile(c echo.Context) error {
	path := c.Param("*")
	log.Println("path: ", path)
	params := ParseCourseIDParams(c)
	user, err := r.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := r.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := r.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := r.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	unitDirPath := data.NodeFilesDirPath(user, term, course, unit)
	path = filepath.Join(unitDirPath, path)
	// Parse the form to retrieve the file
	err = c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		return err
	}
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	// Open the file
	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to open file: %s", err))
	}
	defer src.Close()

	// Create a destination file
	dst, err := os.Create(filepath.Join(path, file.Filename))
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create destination file: %s", err))
	}
	defer dst.Close()

	// Copy the content of the uploaded file to the destination
	if _, err := io.Copy(dst, src); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save file")
	}

	// Respond to the client
	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", file.Filename))
}

// PostNewChild implements NodeRouter.
func (u *unitRouter) PostNewChild(echo.Context) error {
	panic("unimplemented")
}

// Router implements NodeRouter.
func (u *unitRouter) GetRouter() Router {
	return u.Router
}

// ShowDetails implements NodeRouter.
func (r *unitRouter) ShowDetails(c echo.Context) error {
	r.params = ParseCourseIDParams(c)
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := r.svc.GetTerm(r.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := r.svc.GetCourse(r.params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := r.svc.GetUnit(r.params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	r.node = unit
	r.ancestors = []domain.CourseNode{user, term, course, unit}
	page := NodeDetailsPage(r, false)
	// unitDetails := mt.NodeDetailsPage{
	// 	Params:          h.params,
	// 	Node:            unit,
	// 	GetEditNodeURL:  h.app.Reverse(ShowEditUnit.String(), h.params.ToSlice()...),
	// 	PostEditNodeURL: h.app.Reverse(PostEditUnit.String(), h.params.ToSlice()...),
	// 	ListChildrenURL: h.app.Reverse(ListUnitLessons.String(), h.params.ToSlice()...),
	// 	UpNavURL:        h.app.Reverse(ListCourseUnits.String(), h.params.ToSlice()...),
	// 	IsEdit:          false,
	// 	BreadCrumbsData: BreadCrumbs(h.app, h.params, user, term, course, unit),
	// 	GithubFilesURL:  string(templates.UnitFilesURL(unit, course)),
	// 	ServerFilesURL:  h.app.Reverse(ShowUnitFiles.String(), h.params.ToSlice("")...),
	// }
	component := mt.UnitDetailsComponent(page)
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)

}

// ShowEdit implements NodeRouter.
func (r *unitRouter) ShowEdit(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := r.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	queryParam := c.QueryParam("field")
	term, err := r.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := r.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := r.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		log.Println(err)
		return err
	}
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	r.node = unit
	r.ancestors = []domain.CourseNode{user, term, course, unit}
	details := NodeDetailsPage(r, true)
	respond := func(component templ.Component) error {
		return Respond(c, r.app.Reverse(string(UnitDetails), params.ToSlice()...), component, nil)
	}
	if queryParam == templates.KebabCase(string(Description)) {
		return respond(mt.EditDescriptionComponent(details))
	} else if queryParam == templates.KebabCase(string(Name)) {
		return respond(mt.EditNameComponent(details))
	}
	errText := "field value is not expected"
	log.Println(errText)
	return fmt.Errorf("%s %s", errText, queryParam)

}

// ShowFiles implements NodeRouter. (implemented)
func (r *unitRouter) ShowFiles(c echo.Context) error {
	path := c.Param("*")
	log.Println("path: ", path)
	r.params = ParseCourseIDParams(c)
	user, err := r.svc.GetUser(r.params.UserID.Value.(string))
	if err != nil {
		return err
	}

	term, err := r.svc.GetTerm(r.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := r.svc.GetCourse(r.params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := r.svc.GetUnit(r.params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	err = r.svc.CreateNodeFilesDir(user, term, course, unit)
	if err != nil {
		return err
	}
	isDir, err := r.svc.IsDir(path, user, term, course, unit)
	if err != nil {
		return err
	}
	if !isDir {
		c.Attachment(r.svc.NodeFilePath(path, user, term, course, unit), filepath.Base(path))
	}
	files, err := r.svc.NodeFiles(path, user, term, course, unit)
	for _, file := range files {
		log.Println(file.Path)
	}
	if err != nil {
		return err
	}
	log.Println(files)
	// page := mt.FilesPage{
	// 	Node:            unit,
	// 	Params:          r.params,
	// 	CurrentPath:     path,
	// 	OpenFileRHN:     ShowUnitFiles.String(),
	// 	ViewMarkdownRHN: GetUnitViewMarkdown.String(),
	// 	Files:           files,
	// 	E:               r.app,
	// 	BreadCrumbsData: BreadCrumbs(r.app, r.params, user, term, course, unit),
	// }
	page := NodeFilesPage(r, path, files)
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)

}

// ShowNewChild implements NodeRouter.
func (u *unitRouter) ShowNewChild(echo.Context) error {
	panic("unimplemented")
}

// ViewFile implements NodeRouter.
func (r *unitRouter) ViewFile(c echo.Context) error {
	path := c.Param("*")
	log.Println("path: ", path)
	params := ParseCourseIDParams(c)
	user, err := r.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := r.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := r.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := r.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	err = r.svc.CreateNodeFilesDir(user, term, course, unit)
	if err != nil {
		return err
	}
	pathRoot := data.NodeFilesDirPath(user, term, course, unit)
	path = filepath.Join(pathRoot, path)
	content, err := RenderMarkdownFile(path)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	data := mt.MarkdownDocument{
		Title:   filepath.Base(path),
		Content: string(content),
		Static:  false,
	}
	err = mt.DocLayout(data).Render(context.Background(), &buf)
	if err != nil {
		return err
	}
	data.Content = buf.String()
	component := mt.MarkdownIFrame(data)
	layout := CourseManagerLayout(r.app, component, user)
	return Respond(c, "", component, layout)

}

func NewUnitRouter(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &unitRouter{
		Router: Router{
			svc:     svc,
			app:     app,
			nodeSet: EmptyNodesUnit,
		},
	}

}

const (
	Units            RouteName = Course + "/units"
	Unit             RouteName = Units + RouteName(UnitID)
	NewUnit          RouteName = Units + "/new"
	EditUnit         RouteName = Unit + "/edit"
	UnitFiles        RouteName = Unit + "/files/*"
	UnitViewMarkdown RouteName = Unit + "/view-markdown/files/*"
)
const (
	ListCourseUnits     = RouteHandlerName(GET + Units)
	UnitDetails         = RouteHandlerName(GET + Unit)
	ShowUnitFiles       = RouteHandlerName(GET + UnitFiles)
	GetUnitViewMarkdown = RouteHandlerName(GET + UnitViewMarkdown)
	PostUnitFile        = RouteHandlerName(POST + UnitFiles)
	ShowEditUnit        = RouteHandlerName(GET + EditUnit)
	PostEditUnit        = RouteHandlerName(POST + EditUnit)
	ShowNewUnit         = RouteHandlerName(GET + NewUnit)
	PostNewUnit         = RouteHandlerName(POST + NewUnit)
	DeleteUnit          = RouteHandlerName(DELETE + Unit)
)
