package handlers

import (
	"bytes"
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
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

func (h CourseHandler) UnitHandlers() []RouteHandler {
	return []RouteHandler{
		// Units handlers
		{Units, ListCourseUnits, GET, h.ListCourseUnits},
		{Unit, UnitDetails, GET, h.UnitDetails},
		{UnitFiles, ShowUnitFiles, GET, h.ShowUnitFiles},
		{UnitViewMarkdown, GetUnitViewMarkdown, GET, h.GetUnitViewMarkdown},

		{UnitFiles, PostUnitFile, POST, h.PostUnitFile},
		{NewUnit, ShowNewUnit, GET, h.ShowNewUnit},
		{NewUnit, PostNewUnit, POST, h.PostNewUnit},
		{EditUnit, ShowEditUnit, GET, h.ShowEditUnit},
		{EditUnit, PostEditUnit, POST, h.PostEditUnit},
		{Unit, DeleteUnit, DELETE, h.DeleteUnit},
	}
}

func (h CourseHandler) ListCourseUnits(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	courseID, err := ParseRouteParam(c, CourseID)
	if err != nil {
		log.Println(err)
		return err
	}
	course, err := h.svc.GetCourse(courseID)
	if err != nil {
		return err
	}
	units, err := h.svc.GetUnits(courseID)
	if err != nil {
		log.Println(err)
		return err
	}
	course.Units = units
	unitList := mt.NodeListPage{
		Params:           params,
		ParentNode:       course,
		Children:         course.Children(),
		ChildDetailsRHN:  UnitDetails.String(),
		ChildChildrenRHN: ListUnitLessons.String(),
		CreateChildRHN:   ShowNewUnit.String(),
		UpNavURL:         h.e.Reverse(ListTermCourses.String(), params.ToSlice()...),
		E:                h.e,
		BreadCrumbsData:  h.BreadCrumbs(params, user, course.Term, course),
	}
	// old begins here
	// template := mt.UnitsListTemplate(termID, courseID, ListTermCourses.String(), ShowNewUnit.String(), units, ListUnitLessons.String(), UnitDetails.String(), h.e)
	template := mt.NodeListComponent(unitList)
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) UnitDetails(c echo.Context) error {
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
	unitDetails := mt.NodeDetailsPage{
		Params:          params,
		Node:            unit,
		GetEditNodeURL:  h.e.Reverse(ShowEditUnit.String(), params.ToSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditUnit.String(), params.ToSlice()...),
		ListChildrenURL: h.e.Reverse(ListUnitLessons.String(), params.ToSlice()...),
		UpNavURL:        h.e.Reverse(ListCourseUnits.String(), params.ToSlice()...),
		IsEdit:          false,
		BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit),
		GithubFilesURL:  string(templates.UnitFilesURL(unit, course)),
		ServerFilesURL:  h.e.Reverse(ShowUnitFiles.String(), params.ToSlice("")...),
	}
	template := mt.UnitDetailsComponent(unitDetails)
	layout := h.CourseManagerLayout(template, user)
	return Respond(c, "", template, layout)
}

func (h CourseHandler) ShowUnitFiles(c echo.Context) error {
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
	err = h.svc.CreateNodeFilesDir(user, term, course, unit)
	if err != nil {
		return err
	}
	isDir, err := h.svc.IsDir(path, user, term, course, unit)
	if err != nil {
		return err
	}
	if !isDir {
		c.Attachment(h.svc.NodeFilePath(path, user, term, course, unit), filepath.Base(path))
	}
	files, err := h.svc.NodeFiles(path, user, term, course, unit)
	for _, file := range files {
		log.Println(file.Path)
	}
	if err != nil {
		return err
	}
	log.Println(files)
	page := mt.FilesPage{
		Node:            unit,
		Params:          params,
		CurrentPath:     path,
		OpenFileRHN:     ShowUnitFiles.String(),
		ViewMarkdownRHN: GetUnitViewMarkdown.String(),
		Files:           files,
		E:               h.e,
		BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit),
	}
	component := page.Component()
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)
}

func (h CourseHandler) GetUnitViewMarkdown(c echo.Context) error {
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
	err = h.svc.CreateNodeFilesDir(user, term, course, unit)
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
	layout := h.CourseManagerLayout(component, user)
	return Respond(c, "", component, layout)

}

func (h CourseHandler) PostUnitFile(c echo.Context) error {
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

func (h CourseHandler) ShowNewUnit(c echo.Context) error {
	return nil
}

func (h CourseHandler) PostNewUnit(c echo.Context) error {
	return nil

}

func (h CourseHandler) ShowEditUnit(c echo.Context) error {
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
		log.Println(err)
		return err
	}
	if queryParam == "" {
		log.Println(err)
		return fmt.Errorf("field query param is missing")
	}
	details := mt.NodeDetailsPage{
		Params:          params,
		Node:            unit,
		GetEditNodeURL:  h.e.Reverse(ShowEditUnit.String(), params.ToSlice()...),
		PostEditNodeURL: h.e.Reverse(PostEditUnit.String(), params.ToSlice()...),
		ListChildrenURL: h.e.Reverse(ListUnitLessons.String(), params.ToSlice()...),
		UpNavURL:        h.e.Reverse(ListCourseUnits.String(), params.ToSlice()...),
		IsEdit:          true,
		BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit),
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

func (h CourseHandler) PostEditUnit(c echo.Context) error {
	params := ParseCourseIDParams(c)
	user, err := h.svc.GetUser(params.UserID.Value.(string))
	if err != nil {
		return err
	}

	unitID, err := UnitIDParam(params)
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
	unit, err := h.svc.GetUnit(unitID)
	if err != nil {
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		return err
	}
	form := c.Request().Form
	var updateUnit = func() error {
		err := h.svc.UpdateUnit(unit)
		if err != nil {
			return err
		}
		updatedUnit, err := h.svc.GetUnit(unitID)
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
			GetEditNodeURL:  h.e.Reverse(ShowEditUnit.String(), params.ToSlice()...),
			PostEditNodeURL: h.e.Reverse(PostEditUnit.String(), params.ToSlice()...),
			ListChildrenURL: h.e.Reverse(ListUnitLessons.String(), params.ToSlice()...),
			IsEdit:          false,
			BreadCrumbsData: h.BreadCrumbs(params, user, term, course, unit),
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
	return Respond(c, h.e.Reverse(UnitDetails.String(), params.ToSlice()...), template, nil)

}

func (h CourseHandler) DeleteUnit(c echo.Context) error {
	params := ParseCourseIDParams(c)
	return h.svc.DeleteUnit(params.UnitID.Value.(int))
}
