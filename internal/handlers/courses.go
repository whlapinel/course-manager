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
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const (
	Courses          RouteName = Term + "/courses"
	Course           RouteName = Courses + RouteName(CourseID)
	CourseFiles      RouteName = Course + "/files/*"
	ViewMarkdown     RouteName = Course + "/view-markdown/files/*"
	CourseImage      RouteName = Course + "/image"
	NewCourse        RouteName = Courses + "/new"
	EditCourse       RouteName = Course + "/edit"
	CopyCourse       RouteName = Course + "/copy-to-term"
	CopyCourseToTerm RouteName = CopyCourse
	StandardSet      RouteName = Course + "/standard-set"
)
const (
	ListTermCourses       = RouteHandlerName(GET + Courses)
	CourseDetails         = RouteHandlerName(GET + Course)
	ShowCourseFiles       = RouteHandlerName(GET + CourseFiles)
	GetViewMarkdown       = RouteHandlerName(GET + ViewMarkdown)
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

func (h CourseHandler) CourseHandlers() []RouteHandler {
	return []RouteHandler{
		// Courses handlers
		// {Courses, ListTermCourses, GET, h.ListTermCourses},
		{Course, CourseDetails, GET, h.CourseDetails},
		{CourseFiles, ShowCourseFiles, GET, h.ShowCourseFiles},
		{ViewMarkdown, GetViewMarkdown, GET, h.GetViewMarkdown},
		{CourseFiles, PostCourseFiles, POST, h.PostCourseFile},
		{NewCourse, ShowNewCourse, GET, h.ShowNewCourse},
		{NewCourse, PostNewCourse, POST, h.PostNewCourse},
		{EditCourse, ShowEditCourse, GET, h.ShowEditCourse},
		{EditCourse, PostEditCourse, POST, h.PostEditCourse},
		{Course, DeleteCourse, DELETE, h.DeleteCourse},
		{CopyCourse, GetCopyCourse, GET, h.GetCopyCourse},
		{CopyCourseToTerm, PostCopyCourseToTerm, POST, h.PostCopyCourseToTerm},
		{StandardSet, PostSelectStandardSet, POST, h.PostSelectStandardSet},
	}
}

// func (h CourseHandler) ListTermCourses(c echo.Context) error {
// 	params := ParseCourseIDParams(c)
// 	user, err := h.svc.GetUser(params.UserID.Value.(string))
// 	if err != nil {
// 		return err
// 	}

// 	termID := params.TermID.Value.(int)
// 	term, err := h.svc.GetTerm(termID)
// 	if err != nil {
// 		return err
// 	}
// 	courses, err := h.svc.GetCourses(termID)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	term.Courses = courses
// 	page := mt.CourseListPage{
// 		ShowAssessmentsRHN: string(GetCourseAssessments),
// 		ShowCalendarRHN:    string(ShowCourseCalendar),
// 		NodeListPage: mt.NodeListPage{
// 			Params:           params,
// 			ParentNode:       term,
// 			Children:         term.Children(),
// 			ChildDetailsRHN:  CourseDetails.String(),
// 			CreateChildRHN:   ShowNewCourse.String(),
// 			ChildChildrenRHN: ListCourseUnits.String(),
// 			DeleteChildRHN:   DeleteCourse.String(),
// 			UpNavURL:         h.e.Reverse(ListTerms.String(), params.ToIntSlice()...),
// 			E:                h.e,
// 			BreadCrumbsData:  h.BreadCrumbs(params, user, term),
// 		},
// 	}

// 	component := page.Component()
// 	layout := h.CourseManagerLayout(component, user)
// 	return Respond(c, "", component, layout)
// }

func (h CourseHandler) CourseDetails(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	course, err := h.svc.GetCourse(params.CourseID.Value.(int))
	if err != nil {
		return err
	}
	sets, err := h.svc.GetStandardSets()
	if err != nil {
		return err
	}
	page := mt.CourseDetailsPage{
		GetCopyCourseURL:         h.e.Reverse(GetCopyCourse.String(), params.ToSlice()...),
		StandardSets:             sets,
		PostSelectStandardSetURL: h.e.Reverse(string(PostSelectStandardSet), params.ToSlice()...),
		NodeDetailsPage: mt.NodeDetailsPage{
			Params:          params,
			Node:            course,
			GetEditNodeURL:  h.e.Reverse(ShowEditCourse.String(), params.ToSlice()...),
			PostEditNodeURL: h.e.Reverse(PostEditCourse.String(), params.ToSlice()...),
			ListChildrenURL: h.e.Reverse(ListCourseUnits.String(), params.ToSlice()...),
			UpNavURL:        h.e.Reverse(ListTermCourses.String(), params.ToSlice()...),
			IsEdit:          false,
			BreadCrumbsData: h.BreadCrumbs(params, user, course.Term, course),
			GithubFilesURL:  string(templates.CourseFilesURL(course)),
			ServerFilesURL:  h.e.Reverse(ShowCourseFiles.String(), params.ToSlice("")...),
		},
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}
func (h CourseHandler) ShowCourseFiles(c echo.Context) error {
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
		Params:          params,
		CurrentPath:     path,
		OpenFileRHN:     ShowCourseFiles.String(),
		Files:           files,
		E:               h.e,
		BreadCrumbsData: h.BreadCrumbs(params, user, term, course),
		ViewMarkdownRHN: GetViewMarkdown.String(),
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) GetViewMarkdown(c echo.Context) error {
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
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) PostCourseFile(c echo.Context) error {
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
func (h CourseHandler) ShowNewCourse(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	termID, err := TermIDParam(params)
	if err != nil {
		return err
	}
	term, err := h.svc.GetTerm(termID)
	if err != nil {
		return err
	}
	nodeCreate := mt.NodeCreatePage{
		ParentNode:        term,
		NodeType:          domain.CourseTypeName,
		Params:            params,
		PostCreateNodeURL: h.e.Reverse(PostNewCourse.String(), params.ToSlice()...),
		CancelURL:         h.e.Reverse(ListTermCourses.String(), params.ToSlice()...),
		BreadCrumbsData:   h.BreadCrumbs(params, user, term),
	}
	template := mt.NodeCreateComponent(nodeCreate)
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) PostNewCourse(c echo.Context) error {
	params := ParseCourseIDParams(c)
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	form := c.Request().Form
	for key, val := range form {
		log.Println("key, val: ", key, val)
	}
	termID := ParseCourseIDParams(c).TermID
	name := c.FormValue("name")
	description := c.FormValue("description")
	course, err := h.svc.SaveCourse(service.SaveCourseParams{
		TermID:      termID.Value.(int),
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}
	page := mt.NodeDetailsPage{
		Node:            course,
		GetEditNodeURL:  h.e.Reverse(ShowEditCourse.String(), course.ID),
		PostEditNodeURL: h.e.Reverse(PostEditCourse.String(), course.ID),
		UpNavURL:        h.e.Reverse(ListTermCourses.String(), termID),
	}
	template := page.Component()
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) ShowEditCourse(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	queryParam := c.QueryParam("field")
	courseID, err := CourseIDParam(params)
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
		Params:          params,
		Node:            course,
		GetEditNodeURL:  h.e.Reverse(ShowEditCourse.String(), params.ToSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditCourse.String(), params.ToSlice()...),
		ListChildrenURL: h.e.Reverse(ListCourseUnits.String(), params.ToSlice()...),
		IsEdit:          true,
		BreadCrumbsData: h.BreadCrumbs(params, user, course.Term, course),
	}
	respond := func(component templ.Component) error {
		return Respond(c, h.e.Reverse(string(UnitDetails), params.ToSlice()...), component, nil)
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

func (h CourseHandler) PostEditCourse(c echo.Context) error {
	return nil
}

func (h CourseHandler) DeleteCourse(c echo.Context) error {
	params := ParseCourseIDParams(c)
	courseId := params.CourseID
	return h.svc.DeleteCourse(courseId.Value.(int))
}

func (h CourseHandler) GetCopyCourse(c echo.Context) error {
	params := ParseCourseIDParams(c)
	terms, err := h.svc.GetTerms(params.UserID.Value.(string))
	if err != nil {
		return err
	}
	data := mt.CopyCourseData{
		TermID:                  params.TermID.Value.(int),
		CourseID:                params.CourseID.Value.(int),
		Terms:                   terms,
		E:                       h.e,
		PostCopyCourseToTermRHN: string(PostCopyCourseToTerm),
	}
	component := data.Component()
	return Respond(c, h.e.Reverse(ListTermCourses.String(), params.ToSlice()...), component, nil)
}

func (h CourseHandler) PostCopyCourseToTerm(c echo.Context) error {
	params := ParseCourseIDParams(c)
	if params.CourseID.Valid && params.TermID.Valid {
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
		_, err = h.svc.CopyCourseToTerm(params.CourseID.Value.(int), termID)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("params not valid: courseID: %d and termID: %d", params.CourseID.Value, params.TermID.Value)
	}
	return c.Redirect(302, h.e.Reverse(ListTermCourses.String(), params.ToSlice()...))
}

func (h CourseHandler) PostSelectStandardSet(c echo.Context) error {
	params := ParseCourseIDParams(c)
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
	err = h.svc.SetStandardSet(params.CourseID.Value.(int), setID)
	if err != nil {
		return err
	}
	return c.Redirect(302, h.e.Reverse(CourseDetails.String(), params.ToSlice()...))
}
