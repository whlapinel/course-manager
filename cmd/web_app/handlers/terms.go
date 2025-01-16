package handlers

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	mt "gh_static_portfolio/cmd/templates/manager_templates"

	"github.com/labstack/echo/v4"
)

const (
	Terms    RouteName = "/terms"
	Term     RouteName = Terms + RouteName(TermID)
	EditTerm RouteName = Term + "/edit"
	NewTerm  RouteName = Terms + "/new"
)

const (
	ListTerms    = RouteHandlerName(GET + Terms)
	TermDetails  = RouteHandlerName(GET + Term)
	ShowEditTerm = RouteHandlerName(GET + EditTerm)
	PostEditTerm = RouteHandlerName(POST + EditTerm)
	ShowNewTerm  = RouteHandlerName(GET + NewTerm)
	PostNewTerm  = RouteHandlerName(POST + NewTerm)
	DeleteTerm   = RouteHandlerName(POST + Term)
)

func (h CourseHandler) TermHandlers() []RouteHandler {
	return []RouteHandler{
		// Terms handlers
		{Terms, ListTerms, GET, h.ListTerms},
		{Term, TermDetails, GET, h.TermDetails},
		{NewTerm, ShowNewTerm, GET, h.ShowNewTerm},
		{NewTerm, PostNewTerm, POST, h.PostNewTerm},
		{EditTerm, ShowEditTerm, GET, h.ShowEditTerm},
		{EditTerm, PostEditTerm, POST, h.PostEditTerm},
		{Term, DeleteTerm, DELETE, h.DeleteTerm},
	}
}

func (h CourseHandler) ListTerms(c echo.Context) error {
	params := ParseCourseIDParams(c)
	terms, err := h.svc.GetTerms()
	if err != nil {
		return fmt.Errorf("error in CourseHandler.ListTerms: %s", err)
	}
	var termNodes []domain.CourseNode
	for _, term := range terms {
		termNodes = append(termNodes, term)
	}
	termsList := mt.NodeListPage{
		Params:           params,
		ParentNode:       domain.RootCourseNode{},
		Children:         termNodes,
		ChildDetailsRHN:  TermDetails.String(),
		ChildChildrenRHN: ListTermCourses.String(),
		CreateChildRHN:   ShowNewTerm.String(),
		DeleteChildRHN:   DeleteTerm.String(),
		UpNavURL:         h.e.Reverse(string(ShowHome)),
		E:                h.e,
	}
	template := mt.NodeListComponent(termsList)
	layout := h.CourseManagerLayout(template, "Terms")
	return Respond(c, "", template, layout)
}

func (h CourseHandler) TermDetails(c echo.Context) error {
	return nil
}

func (h CourseHandler) ShowNewTerm(c echo.Context) error {
	params := ParseCourseIDParams(c)
	nodeCreate := mt.NodeCreatePage{
		NodeType:          domain.TermTypeName,
		Params:            params,
		PostCreateNodeURL: h.e.Reverse(PostNewTerm.String(), params.ToIntSlice()...),
		CancelURL:         h.e.Reverse(ListTerms.String(), params.ToIntSlice()...),
	}
	template := mt.NodeCreateComponent(nodeCreate)
	layout := h.CourseManagerLayout(template, "New Term")
	return Respond(c, "", template, layout)

}

func (h CourseHandler) PostNewTerm(c echo.Context) error {
	return nil
}

func (h CourseHandler) ShowEditTerm(c echo.Context) error {
	return nil
}

func (h CourseHandler) PostEditTerm(c echo.Context) error {
	return nil
}

func (h CourseHandler) DeleteTerm(c echo.Context) error {
	return nil
}
