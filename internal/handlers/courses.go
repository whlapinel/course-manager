package handlers

import (
	"bytes"
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/service"
	"gh_static_portfolio/internal/templates"
	mt "gh_static_portfolio/internal/templates/manager_templates"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func CourseHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	nodeRouter := NewCourseRouter(svc, router)
	var routeHandlers []RouteHandler
	cr := nodeRouter.(*courseRouter)
	courseRouteHandlers := []RouteHandler{
		{CopyCourse, GetCopyCourse, GET, cr.GetCopyCourse},
		{CopyCourseToTerm, PostCopyCourseToTerm, POST, cr.PostCopyCourseToTerm},
		{StandardSet, PostSelectStandardSet, POST, cr.PostSelectStandardSet},
	}
	routeHandlers = append(routeHandlers, courseRouteHandlers...)
	routeHandlers = append(routeHandlers, NodeHandlers(nodeRouter)...)
	return routeHandlers
}

type courseRouter struct {
	router Router
}

// PostFile implements NodeRouter.
func (r *courseRouter) PostFile(c echo.Context) error {
	h := r.router
	path := c.Param("*")
	log.Println("path: ", path)
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(h.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(h.params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	dirPath := data.NodeFilesDirPath(user, term, course)
	path = filepath.Join(dirPath, path)
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

func (r *courseRouter) GetRouter() Router {
	return r.router
}

// Delete implements NodeRouter.
func (r *courseRouter) Delete(c echo.Context) error {
	h := r.router
	h.params = ParseCourseIDParams(c)
	courseId := h.params.CourseID
	return h.svc.DeleteCourse(courseId.Value.(int))
}

// ListChildren implements NodeRouter. (implemented)
func (r *courseRouter) ListChildren(c echo.Context) error {
	h := r.router
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}

	course, err := h.svc.GetCourse(h.params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	units, err := h.svc.GetUnits(h.params.CourseID.Value.(int))
	if err != nil {
		log.Println(err)
		return err
	}
	course.Units = units
	h.node = course
	r.router = h
	unitList := NodeListPage(r)
	template := mt.NodeListComponent(unitList)
	layout := CourseManagerLayout(h.app, template, user)
	return Respond(c, "", template, layout)
}

// PostEdit implements NodeRouter.
func (c *courseRouter) PostEdit(echo.Context) error {
	panic("unimplemented")
}

// PostNewChild implements NodeRouter.
func (c *courseRouter) PostNewChild(echo.Context) error {
	panic("unimplemented")
}

// ShowDetails implements NodeRouter.
func (r *courseRouter) ShowDetails(c echo.Context) error {
	h := r.router
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}

	course, err := h.svc.GetCourse(h.params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	sets, err := h.svc.GetStandardSets()
	if err != nil {
		return err
	}
	page := mt.CourseDetailsPage{
		GetCopyCourseURL:         h.app.Reverse(GetCopyCourse.String(), h.params.ToSlice()...),
		StandardSets:             sets,
		PostSelectStandardSetURL: h.app.Reverse(string(PostSelectStandardSet), h.params.ToSlice()...),
		NodeDetailsPage: mt.NodeDetailsPage{
			Params:          h.params,
			Node:            course,
			GetEditNodeURL:  h.app.Reverse(ShowEditCourse.String(), h.params.ToSlice()...),
			PostEditNodeURL: h.app.Reverse(PostEditCourse.String(), h.params.ToSlice()...),
			ListChildrenURL: h.app.Reverse(ListCourseUnits.String(), h.params.ToSlice()...),
			UpNavURL:        h.app.Reverse(ListTermCourses.String(), h.params.ToSlice()...),
			IsEdit:          false,
			BreadCrumbsData: BreadCrumbs(h.app, h.params, user, course.Term, course),
			GithubFilesURL:  string(templates.CourseFilesURL(course)),
			ServerFilesURL:  h.app.Reverse(ShowCourseFiles.String(), h.params.ToSlice("")...),
		},
	}
	component := page.Component()
	layout := CourseManagerLayout(h.app, component, user)
	return Respond(c, "", component, layout)

}

// ShowEdit implements NodeRouter.
func (r *courseRouter) ShowEdit(c echo.Context) error {
	h := r.router
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	queryParam := c.QueryParam("field")
	courseID, err := CourseIDParam(h.params)
	if err != nil {
		log.Println(err)
		return err
	}
	course, err := h.svc.GetCourse(courseID)
	if err != nil {
		log.Println(err)
		return err
	}
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := mt.NodeDetailsPage{
		Params:          h.params,
		Node:            course,
		GetEditNodeURL:  h.app.Reverse(ShowEditCourse.String(), h.params.ToSlice()...),
		PostEditNodeURL: h.app.Reverse(PostEditCourse.String(), h.params.ToSlice()...),
		ListChildrenURL: h.app.Reverse(ListCourseUnits.String(), h.params.ToSlice()...),
		IsEdit:          true,
		BreadCrumbsData: BreadCrumbs(h.app, h.params, user, course.Term, course),
	}
	respond := func(component templ.Component) error {
		return Respond(c, h.app.Reverse(string(CourseDetails), h.params.ToSlice()...), component, nil)
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

// ShowFiles implements NodeRouter.
func (r *courseRouter) ShowFiles(c echo.Context) error {
	h := r.router
	path := c.Param("*")
	log.Println("path: ", path)
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}

	term, err := h.svc.GetTerm(h.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(h.params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	err = h.svc.CreateNodeFilesDir(user, term, course)
	if err != nil {
		return err
	}
	isDir, err := h.svc.IsDir(path, user, term, course)
	if err != nil {
		return err
	}
	if !isDir {
		c.Attachment(h.svc.NodeFilePath(path, user, term, course), filepath.Base(path))
	}
	files, err := h.svc.NodeFiles(path, user, term, course)
	for _, file := range files {
		log.Println(file.Path)
	}
	if err != nil {
		return err
	}
	log.Println(files)
	page := mt.FilesPage{
		Node:            course,
		Params:          h.params,
		CurrentPath:     path,
		OpenFileRHN:     ShowCourseFiles.String(),
		Files:           files,
		E:               h.app,
		BreadCrumbsData: BreadCrumbs(h.app, h.params, user, term, course),
		ViewMarkdownRHN: GetCourseViewMarkdown.String(),
	}
	component := page.Component()
	layout := CourseManagerLayout(h.app, component, user)
	return Respond(c, "", component, layout)

}

// ShowNewChild implements NodeRouter.
func (c *courseRouter) ShowNewChild(echo.Context) error {
	panic("unimplemented")
}

// ViewFile implements NodeRouter.
func (r *courseRouter) ViewFile(c echo.Context) error {
	h := r.router
	path := c.Param("*")
	log.Println("path: ", path)
	h.params = ParseCourseIDParams(c)
	user, err := h.svc.GetUser(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(h.params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(h.params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	err = h.svc.CreateNodeFilesDir(user, term, course)
	if err != nil {
		return err
	}
	pathRoot := data.NodeFilesDirPath(user, term, course)
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
	layout := CourseManagerLayout(h.app, component, user)
	return Respond(c, "", component, layout)
}

func NewCourseRouter(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &courseRouter{
		router: Router{
			svc:     svc,
			app:     app,
			nodeSet: EmptyNodesCourse,
		},
	}

}

const (
	Courses            RouteName = Term + "/courses"
	Course             RouteName = Courses + RouteName(CourseID)
	CourseFiles        RouteName = Course + "/files/*"
	CourseViewMarkdown RouteName = Course + "/view-markdown/files/*"
	CourseImage        RouteName = Course + "/image"
	NewCourse          RouteName = Courses + "/new"
	EditCourse         RouteName = Course + "/edit"
	CopyCourse         RouteName = Course + "/copy-to-term"
	CopyCourseToTerm   RouteName = CopyCourse
	StandardSet        RouteName = Course + "/standard-set"
)
const (
	ListTermCourses       = RouteHandlerName(GET + Courses)
	CourseDetails         = RouteHandlerName(GET + Course)
	ShowCourseFiles       = RouteHandlerName(GET + CourseFiles)
	GetCourseViewMarkdown = RouteHandlerName(GET + CourseViewMarkdown)
	PostCourseFiles       = RouteHandlerName(POST + CourseFiles)
	ShowEditCourse        = RouteHandlerName(GET + EditCourse)
	PostEditCourse        = RouteHandlerName(POST + EditCourse)
	ShowNewCourse         = RouteHandlerName(GET + NewCourse)
	PostNewCourse         = RouteHandlerName(POST + NewCourse)
	DeleteCourse          = RouteHandlerName(DELETE + Course)
	GetCopyCourse         = RouteHandlerName(GET + CopyCourse)
	PostCopyCourseToTerm  = RouteHandlerName(POST + CopyCourseToTerm)
	PostSelectStandardSet = RouteHandlerName(POST + StandardSet)
)

func (r *courseRouter) GetCopyCourse(c echo.Context) error {
	h := r.router
	h.params = ParseCourseIDParams(c)
	terms, err := h.svc.GetTerms(h.params.UserID.Value.(string))
	if err != nil {
		return err
	}
	data := mt.CopyCourseData{
		TermID:                  h.params.TermID.Value.(int),
		CourseID:                h.params.CourseID.Value.(int),
		Terms:                   terms,
		E:                       h.app,
		PostCopyCourseToTermRHN: string(PostCopyCourseToTerm),
	}
	component := data.Component()
	return Respond(c, h.app.Reverse(ListTermCourses.String(), h.params.ToSlice()...), component, nil)
}

func (r *courseRouter) PostCopyCourseToTerm(c echo.Context) error {
	log.Println("courseRouter.PostCopyCourseToTerm:")
	h := r.router
	h.params = ParseCourseIDParams(c)
	if h.params.CourseID.Valid && h.params.TermID.Valid {
		err := c.Request().ParseForm()
		if err != nil {
			return err
		}
		termIDParam := c.Request().Form.Get("term-id")
		log.Println("selected termID: ", termIDParam)
		termID, err := strconv.Atoi(termIDParam)
		if err != nil {
			return err
		}
		_, err = h.svc.CopyCourseToTerm(h.params.CourseID.Value.(int), termID)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("params not valid: courseID: %d and termID: %d", h.params.CourseID.Value, h.params.TermID.Value)
	}
	return c.Redirect(302, h.app.Reverse(ListTermCourses.String(), h.params.ToSlice()...))
}

func (r *courseRouter) PostSelectStandardSet(c echo.Context) error {
	log.Println("courseRouter.PostSelectStandardSet():")
	h := r.router
	h.params = ParseCourseIDParams(c)
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	standardSetParam := c.Request().Form.Get("standard-set")
	setID, err := strconv.Atoi(standardSetParam)
	if err != nil {
		return err
	}
	log.Println("selected set: ", standardSetParam)
	err = h.svc.SetStandardSet(h.params.CourseID.Value.(int), setID)
	if err != nil {
		return err
	}
	return c.Redirect(302, h.app.Reverse(CourseDetails.String(), h.params.ToSlice()...))
}
