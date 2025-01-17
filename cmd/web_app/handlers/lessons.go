package handlers

import (
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"gh_static_portfolio/cmd/templates"

	"github.com/labstack/echo/v4"

	sitegenerator "gh_static_portfolio/cmd/gen_site"

	mt "gh_static_portfolio/cmd/templates/manager_templates"
	"io"
	"log"
	"os"

	"github.com/a-h/templ"
)

const (
	Lessons    RouteName = Unit + "/lessons"
	Lesson     RouteName = Lessons + RouteName(LessonID)
	NewLesson  RouteName = Lessons + "/new"
	EditLesson RouteName = Lesson + "/edit"
	Slides     RouteName = Lesson + "/slides"
	EditSlides RouteName = Slides + "/edit"
)

const (
	ListUnitLessons  = RouteHandlerName(GET + Lessons)
	LessonDetails    = RouteHandlerName(GET + Lesson)
	ShowNewLesson    = RouteHandlerName(GET + NewLesson)
	PostNewLesson    = RouteHandlerName(POST + NewLesson)
	ShowEditLesson   = RouteHandlerName(GET + EditLesson)
	PostEditLesson   = RouteHandlerName(POST + EditLesson)
	ViewLessonSlides = RouteHandlerName(GET + Slides)
	ShowEditSlides   = RouteHandlerName(GET + EditSlides)
	PostEditSlides   = RouteHandlerName(POST + EditSlides)
	DeleteLesson     = RouteHandlerName(DELETE + Lesson)
)

func (h CourseHandler) LessonHandlers() []RouteHandler {
	return []RouteHandler{
		{Lessons, ListUnitLessons, GET, h.ListUnitLessons},
		{Lesson, LessonDetails, GET, h.LessonDetails},
		{NewLesson, ShowNewLesson, GET, h.ShowNewLesson},
		{NewLesson, PostNewLesson, POST, h.PostNewLesson},
		{EditLesson, ShowEditLesson, GET, h.ShowEditLesson},
		{EditLesson, PostEditLesson, POST, h.PostEditLesson},
		{Slides, ViewLessonSlides, GET, h.ViewLessonSlides},
		{EditSlides, ShowEditSlides, GET, h.ShowEditSlides},
		{EditSlides, PostEditSlides, POST, h.PostEditSlides},
		{Lesson, DeleteLesson, DELETE, h.DeleteLesson},
	}
}

func (h CourseHandler) ListUnitLessons(c echo.Context) error {
	params := ParseCourseIDParams(c)
	unitID, err := ParseRouteParam(c, UnitID)
	if err != nil {
		log.Println(err)
		return err
	}
	unit, err := h.svc.GetUnit(unitID)
	if err != nil {
		return err
	}
	lessons, err := h.svc.GetLessons(unitID)
	if err != nil {
		log.Println(err)
		return err
	}
	var nodes []domain.CourseNode
	for _, lesson := range lessons {
		nodes = append(nodes, lesson)
	}
	lessonList := mt.NodeListPage{
		Params:          params,
		ParentNode:      *unit,
		Children:        nodes,
		ChildDetailsRHN: LessonDetails.String(),
		CreateChildRHN:  ShowNewLesson.String(),
		DeleteChildRHN:  DeleteLesson.String(),
		UpNavURL:        h.e.Reverse(ListCourseUnits.String(), params.ToIntSlice()...),
		E:               h.e,
	}
	template := mt.LessonListTemplate(lessonList)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) LessonDetails(c echo.Context) error {
	params := ParseCourseIDParams(c)
	termID, err := TermIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("termID:", termID)
	courseID, err := CourseIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	unitID, err := UnitIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	lessonID, err := LessonIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	lesson, err := h.svc.GetLesson(lessonID)
	if err != nil {
		log.Println("error getting lesson:", err)
		return err
	}
	unit, err := h.svc.GetUnit(unitID)
	if err != nil {
		log.Println("error getting unit:", err)
		return err
	}
	course, err := h.svc.GetCourse(courseID)
	if err != nil {
		log.Println("error getting course:", err)
		return err
	}
	nodeDetails := mt.NodeDetailsPage{
		Params:          params,
		Node:            *lesson,
		GetEditNodeURL:  h.e.Reverse(ShowEditLesson.String(), params.ToIntSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditLesson.String(), params.ToIntSlice()...),
		UpNavURL:        h.e.Reverse(ListUnitLessons.String(), params.ToIntSlice()...),
		IsEdit:          false,
	}
	lessonDetails := mt.LessonDetailsPage{
		NodeDetailsPage: nodeDetails,
		GetSlidesURL:    h.e.Reverse(ViewLessonSlides.String(), params.ToIntSlice()...),
		EditSlidesURL:   h.e.Reverse(ShowEditSlides.String(), params.ToIntSlice()...),
		FilesURL:        string(templates.LessonFilesURL(*lesson, *unit, *course)),
	}
	template := mt.LessonDetailsComponent(lessonDetails)
	//old stuff begins here
	// editor := mt.NewLessonEditor(params, string(ShowEditLesson), string(PostEditLesson), h.e, *lesson)
	// template := mt.LessonDetailsTemplate(
	// 	params,
	// 	lesson,
	// 	unit,
	// 	course,
	// 	ViewLessonSlides.String(),
	// 	ShowEditSlides.String(),
	// 	ShowEditLesson.String(),
	// 	PostEditLesson.String(),
	// 	ListUnitLessons.String(),
	// 	editor,
	// 	h.e,
	// )
	// old stuff ends here
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) ShowEditLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	queryParam := c.QueryParam("field")
	lessonID, err := LessonIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	lesson, err := h.svc.GetLesson(lessonID)
	if err != nil {
		log.Println(err)
		return err
	}
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := mt.NodeDetailsPage{
		Params:          params,
		Node:            *lesson,
		GetEditNodeURL:  h.e.Reverse(ShowEditLesson.String(), params.ToIntSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditLesson.String(), params.ToIntSlice()...),
		UpNavURL:        h.e.Reverse(ListUnitLessons.String(), params.ToIntSlice()...),
		IsEdit:          true,
	}
	respond := func(component templ.Component) error {
		return Respond(c, h.e.Reverse(string(LessonDetails), params.ToIntSlice()...), component, nil)
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

// TODO: maybe the UpdateLesson functions should return an updated lesson instead of having to call GetLesson again
func (h CourseHandler) PostEditLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	lessonID, err := LessonIDParam(params)
	if err != nil {
		return err
	}
	lesson, err := h.svc.GetLesson(lessonID)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	var updateLesson = func() error {
		err := h.svc.UpdateLesson(*lesson)
		if err != nil {
			return err
		}
		updatedLesson, err := h.svc.GetLesson(lessonID)
		if err != nil {
			return err
		}
		*lesson = *updatedLesson
		return nil
	}
	var lessonDetails = func(lesson domain.Lesson) mt.NodeDetailsPage {
		return mt.NodeDetailsPage{
			Node:            lesson,
			Params:          params,
			GetEditNodeURL:  h.e.Reverse(ShowEditLesson.String(), params.ToIntSlice()...),
			PostEditNodeURL: h.e.Reverse(PostEditLesson.String(), params.ToIntSlice()...),
			IsEdit:          false,
		}
	}
	var template templ.Component
	for key, val := range form {
		log.Println(key, val)
		switch key {
		case "description":
			lesson.Description = val[0]
			err := updateLesson()
			if err != nil {
				return err
			}
			details := lessonDetails(*lesson)
			template = mt.EditDescriptionComponent(details)
		case "name":
			lesson.Name = val[0]
			err := updateLesson()
			if err != nil {
				return err
			}
			details := lessonDetails(*lesson)
			template = mt.EditNameComponent(details)
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}

	}
	if template == nil {
		panic("template is nil!")
	}
	return Respond(c, h.e.Reverse(string(LessonDetails), params.ToIntSlice()...), template, nil)
}

func (h CourseHandler) ViewLessonSlides(c echo.Context) error {
	params := ParseCourseIDParams(c)
	lessonID, err := ParseRouteParam(c, LessonID)
	if err != nil {
		log.Println(err)
		return err
	}
	sitegenerator.GenerateSlides(lessonID)
	slidesPath := data.NewSlidesHTMLFilePath(lessonID)
	log.Println(slidesPath)
	slidesContent, err := os.ReadFile(slidesPath)
	if err != nil {
		return err
	}
	template := mt.Slides(string(slidesContent))
	return Respond(c, h.e.Reverse(LessonDetails.String(), params.ToIntSlice()...), template, nil)
}

func (h CourseHandler) ShowEditSlides(c echo.Context) error {
	params := ParseCourseIDParams(c)
	lessonID, err := LessonIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	markdownPath, err := h.svc.CreateSlidesIfNotExist(lessonID, h.svc.GetLesson)
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
	template := mt.EditSlidesTemplate(params, string(bytes), string(PostEditSlides), h.e)
	return Respond(c, h.e.Reverse(LessonDetails.String(), params.ToIntSlice()...), template, nil)
}

func (h CourseHandler) PostEditSlides(c echo.Context) error {
	params := ParseCourseIDParams(c)
	log.Println(params)
	content := c.FormValue(string(mt.EditSlidesTextAreaID))
	lessonID, err := LessonIDParam(params)
	if err != nil {
		log.Println(err)
		return err
	}
	path := data.NewSlidesMarkdownFilePath(lessonID)
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
	return nil
}

func (h CourseHandler) ShowNewLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	unitID, err := UnitIDParam(params)
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(unitID)
	if err != nil {
		return err
	}
	page := mt.NodeCreatePage{
		ParentNode:        *unit,
		NodeType:          domain.LessonTypeName,
		Params:            params,
		PostCreateNodeURL: h.e.Reverse(PostNewLesson.String(), params.ToIntSlice()...),
		CancelURL:         h.e.Reverse(ListUnitLessons.String(), params.ToIntSlice()...),
	}
	template := mt.NodeCreateComponent(page)
	layout := h.CourseManagerLayout(template)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) PostNewLesson(c echo.Context) error {
	return nil
}

func (h CourseHandler) DeleteLesson(c echo.Context) error {
	panic("not implemented")
}
