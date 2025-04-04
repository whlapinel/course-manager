package handlers

import (
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/service"
	"io"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"

	mt "gh_static_portfolio/internal/templates/app"
	"log"
)

func LessonHandlers(svc service.CourseService, router *echo.Echo) []RouteHandler {
	nr := NewLessonRouter(svc, router)
	var routeHandlers []RouteHandler
	lr := nr.(*lessonRouter)
	lessonRouteHandlers := []RouteHandler{
		{Slides, ViewLessonSlides, GET, lr.ShowLessonSlides},
		{EditSlides, ShowEditSlides, GET, lr.ShowEditSlides},
		{EditSlides, PostEditSlides, POST, lr.PostEditSlides},
		{LessonAssessments, GetAssessments, GET, lr.GetAssessments},
		{NewLessonAssessment, ShowNewAsssessmentForm, GET, lr.ShowNewAsssessmentForm},

		// TODO: implement
		{LessonStandards, GetStandards, GET, lr.GetStandards},
		{LessonStandards, PostLessonStandard, POST, lr.PostLessonStandard},
		{LessonStandard, DeleteLessonStandard, DELETE, lr.DeleteLessonStandard},
	}
	routeHandlers = append(routeHandlers, lessonRouteHandlers...)
	var nodeHandlers = []RouteHandler{
		ShowFilesHandler(nr.ShowFiles, nr.GetRouter().emptyNodeSet...),
		PostFileHandler(nr.PostFile, nr.GetRouter().emptyNodeSet...),
		ViewFilesHandler(nr.ViewFile, nr.GetRouter().emptyNodeSet...),
		ShowNodeDetailsHandler(nr.ShowDetails, nr.GetRouter().emptyNodeSet...),
		ShowEditHandler(nr.ShowEdit, nr.GetRouter().emptyNodeSet...),
		ShowEditFileHandler(nr.ShowEditFile, nr.GetRouter().emptyNodeSet...),
		PostEditFileHandler(nr.PostEditFile, nr.GetRouter().emptyNodeSet...),
		PostEditHandler(nr.PostEdit, nr.GetRouter().emptyNodeSet...),
		DeleteHandler(nr.Delete, nr.GetRouter().emptyNodeSet...),
	}
	routeHandlers = append(routeHandlers, nodeHandlers...)
	return routeHandlers
}

type lessonRouter struct {
	router
}

func (r *lessonRouter) GetStandards(c echo.Context) error {
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
	nodeStandards, err := r.svc.GetLessonStandards(nodes.Lesson)
	if err != nil {
		return err
	}
	courseStandards, err := r.svc.GetCourseStandardsWithObjectives(nodes.Course)
	if err != nil {
		return err
	}
	standardFrag := mt.StandardsFragment{
		NodeStandards:   nodeStandards,
		CourseStandards: courseStandards,
		PostStandardURL: URL(r, PostLessonStandard),
		DeleteStandardURL: func(id any) string {
			return URL(r, DeleteLessonStandard, id)
		},
	}.Component()
	return Respond(c, ShowDetailsURL(r), standardFrag, nil)

}

func (r *lessonRouter) GetAssessments(c echo.Context) error {
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
	assessments, err := r.svc.GetLessonAssessments(nodes.Lesson.ID)
	if err != nil {
		return err
	}
	assFrag := mt.AssessmentsFragment{
		NewAssessmentURL: URL(r, ShowNewAsssessmentForm),
		GetEditAssessmentURL: func(id any) string {
			altURL := URL(r, GetEditAssessment, id)
			return altURL
		},
		ViewFileURL: func(relPath string) string {
			return URL(r, GetLessonViewMarkdown, relPath)
		},
		Assessments: assessments,
	}.Component()
	return Respond(c, ShowDetailsURL(r), assFrag, nil)

}
func (r *lessonRouter) ShowNewAsssessmentForm(c echo.Context) error {
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
	component := mt.NewAssessmentForm{
		PostAssessmentURL: URL(r, PostAssessment),
		CancelURL:         ShowDetailsURL(r),
		NodeID:            params.LessonID,
	}.Component()
	return Respond(c, ShowDetailsURL(r), component, nil)

}

// ShowEditFile implements NodeRouter.
func (r *lessonRouter) ShowEditFile(c echo.Context) error {
	return ShowEditFile(c, r, "/")
}

func (r *lessonRouter) PostEditFile(c echo.Context) error {
	return PostEditFile(c, r, ShowDetailsURL(r))
}

// SetRouter implements NodeRouter.
func (r *lessonRouter) SetRouter(router router) {
	r.router = router
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

func (r *lessonRouter) ShowLessonSlides(c echo.Context) error {
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
	slides, err := r.svc.GetSlides(r.params)
	if err != nil {
		return err
	}
	slideFrag := mt.Slides{
		HTML:          slides,
		EditSlidesURL: URL(r, ShowEditSlides),
	}.Component()
	return Respond(c, ShowDetailsURL(r), slideFrag, nil)
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
func (r *lessonRouter) GetRouter() router {
	return r.router
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
	page, err := r.DetailsPage(false)
	if err != nil {
		return err
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
	page, err := r.DetailsPage(true)
	if err != nil {
		return err
	}
	component := page.DetailsFormComponent(true)
	layout := CourseManagerLayout(r.app, component, r.nodes.User)
	return Respond(c, "", component, layout)

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
	return ViewFile(c, r, ShowDetailsURL(r))
}

func NewLessonRouter(svc service.CourseService, app *echo.Echo) NodeRouter {
	return &lessonRouter{
		router: router{
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
	GetAssessments        = RouteHandlerName(GET + LessonAssessments)
	GetStandards          = RouteHandlerName(GET + LessonStandards)
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

func (r *lessonRouter) PostLessonStandard(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	objectiveParam := c.Request().Form.Get("objective")
	objID, err := strconv.Atoi(objectiveParam)
	if err != nil {
		return err
	}
	err = r.svc.SetLessonObjective(params.LessonID, objID)
	if err != nil {
		return err
	}
	return c.Redirect(303, ShowDetailsURL(r))
}

func (r *lessonRouter) DeleteLessonStandard(c echo.Context) error {
	params, err := ParseNodePath(c)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	objectiveParam := ParseRouteStringParam(c, StandardID)
	objID, err := strconv.Atoi(objectiveParam)
	if err != nil {
		return err
	}
	log.Println(params.ToSlice()...)
	log.Println("deleting lesson standard ", objID)
	err = r.svc.DeleteLessonObjective(params.LessonID, objID)
	if err != nil {
		return err
	}
	return c.Redirect(303, ShowDetailsURL(r))
}

func (r *lessonRouter) DetailsPage(isEdit bool) (mt.LessonDetailsPage, error) {
	page := mt.LessonDetailsPage{
		E:                 r.app,
		NodeDetailsPage:   NodeDetailsPage(r, isEdit),
		ViewMarkdownRHN:   string(ViewNodeFilesRHN(r.emptyNodeSet...)),
		FileRHN:           string(ShowNodeFilesRHN(r.emptyNodeSet...)),
		AssetsURLFunc:     AssetsURLFunc,
		GetAssessmentsURL: URL(r, GetAssessments),
		GetStandardsURL:   URL(r, GetStandards),
		GetSlidesURL:      URL(r, ViewLessonSlides),
		EditSlidesURL:     r.app.Reverse(ShowEditSlides.String(), r.params.ToSlice()...),
	}

	return page, nil
}
