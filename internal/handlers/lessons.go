package handlers

import (
	"bytes"
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/service"
	"gh_static_portfolio/internal/templates"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"

	mt "gh_static_portfolio/internal/templates/manager_templates"
	"io"
	"log"
	"os"

	"github.com/a-h/templ"
)

const (
	Lessons            RouteName = Unit + "/lessons"
	Lesson             RouteName = Lessons + RouteName(LessonID)
	NewLesson          RouteName = Lessons + "/new"
	EditLesson         RouteName = Lesson + "/edit"
	Slides             RouteName = Lesson + "/slides"
	LessonFiles        RouteName = Lesson + "/files/*"
	LessonViewMarkdown RouteName = Lesson + "/view-markdown/files/*"
	EditSlides         RouteName = Slides + "/edit"
	LessonStandards    RouteName = Lesson + "/standards"
	LessonStandard     RouteName = LessonStandards + RouteName(StandardID)
)

const (
	ListUnitLessons       = RouteHandlerName(GET + Lessons)
	LessonDetails         = RouteHandlerName(GET + Lesson)
	ShowNewLesson         = RouteHandlerName(GET + NewLesson)
	PostNewLesson         = RouteHandlerName(POST + NewLesson)
	ShowEditLesson        = RouteHandlerName(GET + EditLesson)
	PostEditLesson        = RouteHandlerName(POST + EditLesson)
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

func (h CourseHandler) LessonHandlers() []RouteHandler {
	return []RouteHandler{
		{Lessons, ListUnitLessons, GET, h.ListUnitLessons},
		{Lesson, LessonDetails, GET, h.LessonDetails},
		{NewLesson, ShowNewLesson, GET, h.ShowNewLesson},
		{NewLesson, PostNewLesson, POST, h.PostNewLesson},
		{EditLesson, ShowEditLesson, GET, h.ShowEditLesson},
		{EditLesson, PostEditLesson, POST, h.PostEditLesson},
		{Slides, ViewLessonSlides, GET, h.ViewLessonSlides},
		{LessonFiles, ShowLessonFiles, GET, h.ShowLessonFiles},
		{LessonViewMarkdown, GetLessonViewMarkdown, GET, h.GetLessonViewMarkdown},
		{LessonFiles, PostLessonFile, POST, h.PostLessonFile},
		{EditSlides, ShowEditSlides, GET, h.ShowEditSlides},
		{EditSlides, PostEditSlides, POST, h.PostEditSlides},
		{Lesson, DeleteLesson, DELETE, h.DeleteLesson},
		{LessonStandards, PostLessonStandard, POST, h.PostLessonStandard},
		{LessonStandard, DeleteLessonStandard, DELETE, h.DeleteLessonStandard},
	}
}

func (h CourseHandler) ListUnitLessons(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	unitID, err := ParseRouteParam(c, UnitID)
	if err != nil {
		log.Println(err)
		return err
	}
	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
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
		ParentNode:      unit,
		Children:        nodes,
		ChildDetailsRHN: LessonDetails.String(),
		CreateChildRHN:  ShowNewLesson.String(),
		DeleteChildRHN:  DeleteLesson.String(),
		UpNavURL:        h.e.Reverse(ListCourseUnits.String(), params.ToSlice()...),
		E:               h.e,
		BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit),
	}
	template := mt.LessonListTemplate(lessonList)
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) LessonDetails(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	page, err := h.lessonDetailsPage(params, "")
	if err != nil {
		return err
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) ShowEditLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	queryParam := c.QueryParam("field")
	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}

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
		Node:            lesson,
		GetEditNodeURL:  h.e.Reverse(ShowEditLesson.String(), params.ToSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditLesson.String(), params.ToSlice()...),
		UpNavURL:        h.e.Reverse(ListUnitLessons.String(), params.ToSlice()...),
		IsEdit:          true,
		BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit, lesson),
	}
	respond := func(component templ.Component) error {
		return Respond(c, h.e.Reverse(string(LessonDetails), params.ToSlice()...), component, nil)
	}
	if queryParam == templates.KebabCase(string(Description)) {
		return respond(mt.EditDescriptionComponent(details))
	} else if queryParam == templates.KebabCase(string(Name)) {
		return respond(mt.EditNameComponent(details))
	} else if queryParam == "number" {
		return respond(mt.EditNumberComponent(details))
	}
	errText := "field value is not expected"
	log.Println(errText)
	return fmt.Errorf("%s %s", errText, queryParam)

}

// TODO: maybe the UpdateLesson functions should return an updated lesson instead of having to call GetLesson again
func (h CourseHandler) PostEditLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	lessonID, err := LessonIDParam(params)
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
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
		err := h.svc.UpdateLesson(lesson)
		if err != nil {
			return err
		}
		updatedLesson, err := h.svc.GetLesson(lessonID)
		if err != nil {
			return err
		}
		lesson = updatedLesson
		return nil
	}
	var lessonDetails = func(lesson domain.Lesson) mt.NodeDetailsPage {
		return mt.NodeDetailsPage{
			Node:            lesson,
			Params:          params,
			GetEditNodeURL:  h.e.Reverse(ShowEditLesson.String(), params.ToSlice()...),
			PostEditNodeURL: h.e.Reverse(PostEditLesson.String(), params.ToSlice()...),
			IsEdit:          false,
			BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit, lesson),
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
			details := lessonDetails(lesson)
			template = mt.EditDescriptionComponent(details)
		case "name":
			lesson.Name = val[0]
			err := updateLesson()
			if err != nil {
				return err
			}
			details := lessonDetails(lesson)
			template = mt.EditNameComponent(details)
		case "number":
			num, err := strconv.Atoi(val[0])
			if err != nil {
				return err
			}
			lesson.Number = num
			err = updateLesson()
			if err != nil {
				return err
			}
			details := lessonDetails(lesson)
			template = mt.EditNumberComponent(details)
		default:
			log.Println("form key:", key)
			panic("form key not expected!")
		}

	}
	if template == nil {
		panic("template is nil!")
	}
	return Respond(c, h.e.Reverse(string(LessonDetails), params.ToSlice()...), template, nil)
}

func (h CourseHandler) ViewLessonSlides(c echo.Context) error {
	params := ParseCourseIDParams(c)

	slidesContent, err := h.svc.GetSlides(params)
	if err != nil {
		return err
	}
	page, err := h.lessonDetailsPage(params, slidesContent)
	if err != nil {
		return err
	}
	component := page.Component()
	return Respond(c, h.e.Reverse(LessonDetails.String(), params.ToSlice()...), component, nil)
}

func (h CourseHandler) ShowLessonFiles(c echo.Context) error {
	path := c.Param("*")
	log.Println("path: ", path)
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	lesson, err := h.svc.GetLesson(params.LessonID.Value.(int))
	if err != nil {
		return err
	}
	err = h.svc.CreateNodeFilesDir(user, term, course, unit, lesson)
	if err != nil {
		return err
	}
	isDir, err := h.svc.IsDir(path, user, term, course, unit, lesson)
	if err != nil {
		return err
	}
	if !isDir {
		c.Attachment(h.svc.NodeFilePath(path, user, term, course, unit, lesson), filepath.Base(path))
	}
	files, err := h.svc.NodeFiles(path, user, term, course, unit, lesson)
	for _, file := range files {
		log.Println(file.Path)
	}
	if err != nil {
		return err
	}
	log.Println(files)
	page := mt.FilesPage{
		Node:            lesson,
		Params:          params,
		CurrentPath:     path,
		OpenFileRHN:     ShowLessonFiles.String(),
		ViewMarkdownRHN: GetLessonViewMarkdown.String(),
		Files:           files,
		E:               h.e,
		BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit, lesson),
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) GetLessonViewMarkdown(c echo.Context) error {
	params := ParseCourseIDParams(c)
	path := c.Param("*")
	log.Println("path: ", path)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	lesson, err := h.svc.GetLesson(params.LessonID.Value.(int))
	if err != nil {
		return err
	}
	err = h.svc.CreateNodeFilesDir(user, term, course, unit, lesson)
	if err != nil {
		return err
	}
	pathRoot := data.NodeFilesDirPath(user, term, course, unit, lesson)
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
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)

}

func (h CourseHandler) PostLessonFile(c echo.Context) error {
	path := c.Param("*")
	log.Println("path: ", path)
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	lesson, err := h.svc.GetLesson(params.LessonID.Value.(int))
	if err != nil {
		return err
	}
	lessonDirPath := data.NodeFilesDirPath(user, term, course, unit, lesson)
	path = filepath.Join(lessonDirPath, path)
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

func (h CourseHandler) ShowEditSlides(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	lesson, err := h.svc.GetLesson(params.LessonID.Value.(int))
	if err != nil {
		return err
	}
	markdownPath, err := h.svc.CreateSlidesIfNotExist(user, term, course, unit, lesson)
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
	return Respond(c, h.e.Reverse(LessonDetails.String(), params.ToSlice()...), template, nil)
}

func (h CourseHandler) PostEditSlides(c echo.Context) error {
	params := ParseCourseIDParams(c)
	log.Println(params)
	content := c.FormValue(string(mt.EditSlidesTextAreaID))
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	lesson, err := h.svc.GetLesson(params.LessonID.Value.(int))
	if err != nil {
		return err
	}

	path := data.SlidesMarkdownFilePath(user, term, course, unit, lesson)
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
	return c.Redirect(303, h.e.Reverse(ViewLessonSlides.String(), params.ToSlice()...))
}

func (h CourseHandler) ShowNewLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return err
	}
	page := mt.NodeCreatePage{
		ParentNode:        unit,
		NodeType:          domain.LessonTypeName,
		Params:            params,
		PostCreateNodeURL: h.e.Reverse(PostNewLesson.String(), params.ToSlice()...),
		CancelURL:         h.e.Reverse(ListUnitLessons.String(), params.ToSlice()...),
		BreadCrumbsData:   h.BreadCrumbs(params, user, term, course, unit),
	}
	template := mt.NodeCreateComponent(page)
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) PostNewLesson(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	for key, val := range form {
		log.Println("key, val: ", key, val)
	}
	params := ParseCourseIDParams(c)
	name := form.Get("name")
	numberParam := form.Get("number")
	number, err := strconv.Atoi(numberParam)
	if err != nil {
		return err
	}
	description := form.Get("description")
	lesson, err := h.svc.SaveLesson(service.SaveLessonParams{
		Lesson: domain.Lesson{
			UnitID:      params.UnitID.Value.(int),
			Number:      number,
			Name:        name,
			Description: description,
		},
	})

	if err != nil {
		return err
	}
	return c.Redirect(303, h.e.Reverse(LessonDetails.String(), params.ToSlice(lesson.ID)...))
}

func (h CourseHandler) DeleteLesson(c echo.Context) error {
	params := ParseCourseIDParams(c)
	err := h.svc.DeleteLesson(params.LessonID.Value.(int))
	if err != nil {
		return err
	}
	return nil
}

func (h CourseHandler) PostLessonStandard(c echo.Context) error {
	params := ParseCourseIDParams(c)
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	objectiveParam := c.Request().Form.Get("objective")
	objID, err := strconv.Atoi(objectiveParam)
	if err != nil {
		return err
	}
	err = h.svc.SetLessonObjective(params.LessonID.Value.(int), objID)
	if err != nil {
		return err
	}
	return c.Redirect(302, h.e.Reverse(LessonDetails.String(), params.ToSlice()...))
}

func (h CourseHandler) DeleteLessonStandard(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
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
	err = h.svc.DeleteLessonObjective(params.LessonID.Value.(int), objID)
	if err != nil {
		return err
	}
	page, err := h.lessonDetailsPage(params, "")
	if err != nil {
		return err
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) lessonDetailsPage(params mt.CourseIDParams, slides string) (mt.LessonDetailsPage, error) {
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return mt.LessonDetailsPage{}, err
	}

	term, err := h.svc.GetTerm(params.TermID.Value.(int))
	if err != nil {
		return mt.LessonDetailsPage{}, err
	}
	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return mt.LessonDetailsPage{}, err
	}
	unit, err := h.svc.GetUnit(params.UnitID.Value.(int))
	if err != nil {
		return mt.LessonDetailsPage{}, err
	}
	lesson, err := h.svc.GetLesson(params.LessonID.Value.(int))
	if err != nil {
		log.Println("error getting lesson:", err)
		return mt.LessonDetailsPage{}, err
	}
	nodeDetails := mt.NodeDetailsPage{
		Params:            params,
		ParentNode:        unit,
		Node:              lesson,
		GetEditNodeURL:    h.e.Reverse(ShowEditLesson.String(), params.ToSlice()...),
		PostEditNodeURL:   h.e.Reverse(PostEditLesson.String(), params.ToSlice()...),
		UpNavURL:          h.e.Reverse(ListUnitLessons.String(), params.ToSlice()...),
		IsEdit:            false,
		BreadCrumbsData:   h.BreadCrumbs(params, user, term, course, unit, lesson),
		Slides:            slides,
		CourseCalendarURL: h.e.Reverse(ShowCourseCalendar.String(), params.ToSlice()...),
		GithubFilesURL:    string(templates.LessonFilesURL(lesson, unit, course)),
		ServerFilesURL:    h.e.Reverse(ShowLessonFiles.String(), params.ToSlice("")...),
	}

	// var idParams = params.ToIntSlice()
	// var pathParam interface{} = ""
	// var withPath = append(idParams, pathParam)

	lessonDetails := mt.LessonDetailsPage{
		E:                       h.e,
		Standards:               course.StandardSet.Standards,
		NodeDetailsPage:         nodeDetails,
		PostLessonStandardURL:   h.e.Reverse(PostLessonStandard.String(), params.ToSlice()...),
		DeleteLessonStandardRHN: DeleteLessonStandard.String(),
		GetEditAssessmentRHN:    GetEditAssessment.String(),
		DeleteAssessmentRHN:     DeleteAssessment.String(),
		PostAssessmentURL:       h.e.Reverse(PostAssessment.String(), params.ToSlice()...),
		GetSlidesURL:            h.e.Reverse(ViewLessonSlides.String(), params.ToSlice()...),
		EditSlidesURL:           h.e.Reverse(ShowEditSlides.String(), params.ToSlice()...),
		GithubFilesURL:          string(templates.LessonFilesURL(lesson, unit, course)),
		ServerFilesURL:          h.e.Reverse(ShowLessonFiles.String(), params.ToSlice("")...),
	}
	return lessonDetails, nil
}
