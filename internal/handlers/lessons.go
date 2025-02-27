package handlers

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/service"
	"io"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"

	mt "gh_static_portfolio/internal/templates/manager_templates"
	"log"

	"github.com/a-h/templ"
)

func LessonHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	nr := NewLessonRouter(svc, router)
	var routeHandlers []RouteHandler
	lr := nr.(*lessonRouter)
	lessonRouteHandlers := []RouteHandler{
		{EditSlides, ShowEditSlides, GET, lr.ShowEditSlides},
		{EditSlides, PostEditSlides, POST, lr.PostEditSlides},
	}
	routeHandlers = append(routeHandlers, lessonRouteHandlers...)
	var nodeHandlers = []RouteHandler{
		ShowFilesHandler(nr.ShowFiles, nr.GetRouter().emptyNodeSet...),
		ViewFilesHandler(nr.ViewFile, nr.GetRouter().emptyNodeSet...),
		ShowNodeDetailsHandler(nr.ShowDetails, nr.GetRouter().emptyNodeSet...),
		ShowEditHandler(nr.ShowEdit, nr.GetRouter().emptyNodeSet...),
		PostEditHandler(nr.PostEdit, nr.GetRouter().emptyNodeSet...),
		DeleteHandler(nr.Delete, nr.GetRouter().emptyNodeSet...),
	}
	routeHandlers = append(routeHandlers, nodeHandlers...)
	return routeHandlers
}

type lessonRouter struct {
	Router
}

// SetRouter implements NodeRouter.
func (r *lessonRouter) SetRouter(router Router) {
	r.Router = router
}

// Delete implements NodeRouter.
func (r *lessonRouter) Delete(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	return r.svc.DeleteLesson(r.params.LessonID)
}

// ListChildren implements NodeRouter.
// since lesson has no children, this will remain unimplemented.
func (l *lessonRouter) ListChildren(echo.Context) error {
	panic("unimplemented")
}
func (r *lessonRouter) PostEditSlides(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	content := c.FormValue(string(mt.EditSlidesTextAreaID))
	path := data.SlidesMarkdownFilePath(r.nodes.ToSlice()...)
	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(content))
	if err != nil {
		log.Println(err)
		return err
	}
	go func() error {
		slides, err := r.svc.GetSlides(r.params)
		if err != nil {
			return err
		}
		path := data.SlidesHTMLFilePath(r.nodes.ToSlice()...)
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		written, err := file.Write([]byte(slides))
		if err != nil {
			return err
		}
		log.Printf("%d written to file %s", written, path)
		return nil
	}()
	return c.Redirect(303, ShowDetailsURL(r))

}
func (r *lessonRouter) ShowEditSlides(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	markdownPath, err := r.svc.CreateSlidesIfNotExist(r.nodes.ToSlice()...)
	if err != nil {
		return err
	}
	markdownFile, err := os.Open(markdownPath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer markdownFile.Close()
	bytes, err := io.ReadAll(markdownFile)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(len(bytes), "bytes read")

	log.Println(string(bytes))
	template := mt.EditSlidesTemplate(r.params, string(bytes), string(PostEditSlides), r.app)
	return Respond(c, ShowDetailsURL(r), template, nil)
}

// PostEdit implements NodeRouter.
func (r *lessonRouter) PostEdit(c echo.Context) error {
	newParams, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = newParams
	nodes, err := r.svc.Nodes(r.params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "description":
			r.nodes.Lesson.Description = val[0]
		case "name":
			r.nodes.Lesson.Name = val[0]
		case "number":
			num, err := strconv.Atoi(val[0])
			if err != nil {
				return err
			}
			r.nodes.Lesson.Number = num
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}
		err := r.svc.UpdateLesson(r.nodes.Lesson)
		if err != nil {
			return err
		}

	}
	return c.Redirect(303, ShowDetailsURL(r))

}

// PostFile implements NodeRouter.
func (r *lessonRouter) PostFile(c echo.Context) error {
	return PostFile(c, r)
}

// PostNewChild implements NodeRouter.
// Lesson does not have child nodes
func (r *lessonRouter) PostNewChild(c echo.Context) error {
	panic("not implemented")

}

// Router implements NodeRouter.
func (r *lessonRouter) GetRouter() Router {
	return r.Router
}

// ShowDetails implements NodeRouter.
func (r *lessonRouter) ShowDetails(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	nodeDetails := NodeDetailsPage(r, false)
	slides, err := r.svc.GetSlides(r.params)
	if err != nil {
		return err
	}
	page := mt.LessonDetailsPage{
		Slides:                  slides,
		E:                       r.app,
		Standards:               r.nodes.Course.StandardSet.Standards,
		NodeDetailsPage:         nodeDetails,
		PostLessonStandardURL:   r.app.Reverse(PostLessonStandard.String(), params.ToSlice()...),
		ViewMarkdownRHN:         string(ViewNodeFilesRHN(r.emptyNodeSet...)),
		FileRHN:                 string(ShowNodeFilesRHN(r.emptyNodeSet...)),
		DeleteLessonStandardRHN: DeleteLessonStandard.String(),
		GetEditAssessmentRHN:    GetEditAssessment.String(),
		DeleteAssessmentRHN:     DeleteAssessment.String(),
		PostAssessmentURL:       r.app.Reverse(PostAssessment.String(), params.ToSlice()...),
		EditSlidesURL:           r.app.Reverse(ShowEditSlides.String(), params.ToSlice()...),
	}
	component := page.Component()
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)
}

// ShowEdit implements NodeRouter.
func (r *lessonRouter) ShowEdit(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	r.params = params
	nodes, err := r.svc.Nodes(params)
	if err != nil {
		return err
	}
	r.nodes = nodes
	queryParam := c.QueryParam("field")
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := NodeDetailsPage(r, true)
	respond := func(component templ.Component) error {
		return Respond(c, r.app.Reverse(string(ShowNodeDetailsRHN(r.emptyNodeSet...)), params.ToSlice()...), component, nil)
	}
	if queryParam == "description" {
		return respond(mt.EditDescriptionComponent(details))
	} else if queryParam == "name" {
		return respond(mt.EditNameComponent(details))
	} else if queryParam == "number" {
		return respond(mt.EditNumberComponent(details))
	}
	errText := "field value is not expected"
	log.Println(errText)
	return fmt.Errorf("%s %s", errText, queryParam)

}

// ShowFiles implements NodeRouter.
func (r *lessonRouter) ShowFiles(c echo.Context) error {
	return ShowFiles(c, r)
}

// ShowNewChild implements NodeRouter.
func (l *lessonRouter) ShowNewChild(echo.Context) error {
	panic("unimplemented")
}

// ViewFile implements NodeRouter.
func (r *lessonRouter) ViewFile(c echo.Context) error {
	return ViewFile(c, r)
}

func NewLessonRouter(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &lessonRouter{
		Router: Router{
			svc:          svc,
			app:          app,
			emptyNodeSet: EmptyNodesLesson,
		},
	}
}

const (
	Lessons            RoutePath = Unit + "/lessons"
	Lesson             RoutePath = Lessons + RoutePath(LessonID)
	NewLesson          RoutePath = Lessons + "/new"
	EditLesson         RoutePath = Lesson + "/edit"
	Slides             RoutePath = Lesson + "/slides"
	LessonFiles        RoutePath = Lesson + "/files/*"
	LessonViewMarkdown RoutePath = Lesson + "/view-markdown/files/*"
	EditSlides         RoutePath = Slides + "/edit"
	LessonStandards    RoutePath = Lesson + "/standards"
	LessonStandard     RoutePath = LessonStandards + RoutePath(StandardID)
)

const (
	ViewLessonSlides      = RouteHandlerName(GET + Slides)
	ShowLessonFiles       = RouteHandlerName(GET + LessonFiles)
	GetLessonViewMarkdown = RouteHandlerName(GET + LessonViewMarkdown)
	PostLessonFile        = RouteHandlerName(POST + LessonFiles)
	ShowEditSlides        = RouteHandlerName(GET + EditSlides)
	PostEditSlides        = RouteHandlerName(POST + EditSlides)
	DeleteLesson          = RouteHandlerName(DELETE + Lesson)
	PostLessonStandard    = RouteHandlerName(POST + LessonStandards)
	DeleteLessonStandard  = RouteHandlerName(DELETE + LessonStandard)
)

// func (h CourseHandler) LessonHandlers() []RouteHandler {
// 	return []RouteHandler{
// 		{Lessons, ListUnitLessons, GET, h.ListUnitLessons},
// 		{Lesson, LessonDetails, GET, h.LessonDetails},
// 		{NewLesson, ShowNewLesson, GET, h.ShowNewLesson},
// 		{NewLesson, PostNewLesson, POST, h.PostNewLesson},
// 		{EditLesson, ShowEditLesson, GET, h.ShowEditLesson},
// 		{EditLesson, PostEditLesson, POST, h.PostEditLesson},
// 		{Slides, ViewLessonSlides, GET, h.ViewLessonSlides},
// 		{LessonFiles, ShowLessonFiles, GET, h.ShowLessonFiles},
// 		{LessonViewMarkdown, GetLessonViewMarkdown, GET, h.GetLessonViewMarkdown},
// 		{LessonFiles, PostLessonFile, POST, h.PostLessonFile},
// 		{EditSlides, ShowEditSlides, GET, h.ShowEditSlides},
// 		{EditSlides, PostEditSlides, POST, h.PostEditSlides},
// 		{Lesson, DeleteLesson, DELETE, h.DeleteLesson},
// 		{LessonStandards, PostLessonStandard, POST, h.PostLessonStandard},
// 		{LessonStandard, DeleteLessonStandard, DELETE, h.DeleteLessonStandard},
// 	}
// }

// func (h CourseHandler) ShowNewLesson(c echo.Context) error {
// 	params := ParseCourseIDParams(c)
// 	user, err := h.svc.GetUser(params.UserID.Value.(string))
// 	if err != nil {
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
// 	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
// 	if err != nil {
// 		return err
// 	}
// 	page := mt.NodeCreatePage{
// 		ParentNode:        unit,
// 		NodeType:          domain.LessonTypeName,
// 		Params:            params,
// 		PostCreateNodeURL: h.e.Reverse(PostNewLesson.String(), params.ToSlice()...),
// 		CancelURL:         h.e.Reverse(ListUnitLessons.String(), params.ToSlice()...),
// 		BreadCrumbsData:   h.BreadCrumbs(params, user, term, course, unit),
// 	}
// 	template := mt.NodeCreateComponent(page)
// 	layout := h.CourseManagerLayout(template, user)
// 	return Respond(c, "", template, layout)
// }

// func (h CourseHandler) PostNewLesson(c echo.Context) error {
// 	err := c.Request().ParseForm()
// 	if err != nil {
// 		return err
// 	}
// 	form := c.Request().Form
// 	for key, val := range form {
// 		log.Println("key, val: ", key, val)
// 	}
// 	params := ParseCourseIDParams(c)
// 	name := form.Get("name")
// 	numberParam := form.Get("number")
// 	number, err := strconv.Atoi(numberParam)
// 	if err != nil {
// 		return err
// 	}
// 	description := form.Get("description")
// 	lesson, err := h.svc.SaveLesson(service.SaveLessonParams{
// 		Lesson: domain.Lesson{
// 			UnitID:      params.UnitID.Value.(int),
// 			Number:      number,
// 			Name:        name,
// 			Description: description,
// 		},
// 	})

// 	if err != nil {
// 		return err
// 	}
// 	return c.Redirect(303, h.e.Reverse(LessonDetails.String(), params.ToSlice(lesson.ID)...))
// }

// func (h CourseHandler) PostLessonStandard(c echo.Context) error {
// 	params := ParseCourseIDParams(c)
// 	err := c.Request().ParseForm()
// 	if err != nil {
// 		return err
// 	}
// 	objectiveParam := c.Request().Form.Get("objective")
// 	objID, err := strconv.Atoi(objectiveParam)
// 	if err != nil {
// 		return err
// 	}
// 	err = h.svc.SetLessonObjective(params.LessonID.Value.(int), objID)
// 	if err != nil {
// 		return err
// 	}
// 	return c.Redirect(302, h.e.Reverse(LessonDetails.String(), params.ToSlice()...))
// }

// func (h CourseHandler) DeleteLessonStandard(c echo.Context) error {
// 	params := ParseCourseIDParams(c)
// 	user, err := h.svc.GetUser(params.UserID.Value.(string))
// 	if err != nil {
// 		return err
// 	}

// 	err = c.Request().ParseForm()
// 	if err != nil {
// 		return err
// 	}
// 	objectiveParam := ParseRouteStringParam(c, StandardID)
// 	objID, err := strconv.Atoi(objectiveParam)
// 	if err != nil {
// 		return err
// 	}
// 	log.Println(params.ToSlice()...)
// 	log.Println("deleting lesson standard ", objID)
// 	err = h.svc.DeleteLessonObjective(params.LessonID.Value.(int), objID)
// 	if err != nil {
// 		return err
// 	}
// 	page, err := h.lessonDetailsPage(params, "")
// 	if err != nil {
// 		return err
// 	}
// 	component := page.Component()
// 	layout := h.CourseManagerLayout(component, user)
// 	return Respond(c, "", component, layout)
// }
