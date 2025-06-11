package fileviews

import (
	ac "gh_static_portfolio/internal/app/components"
	cmp "gh_static_portfolio/internal/basecomponents"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"strings"

	"github.com/a-h/templ"
)

type FilesPage struct {
	cmp.UpNav
	Root                bool
	ParentDirectory     FilesPageItem
	CurrentDirectory    FilesPageItem
	Params              routes.NodePath
	CurrentPath         string
	PreviewURL          web.AddParams
	OpenDirURL          web.AddParams
	OpenFileURL         web.AddParams
	UploadFileURL       web.AddParams
	Node                ports.Node
	Files               []FilesPageItem
	PopRouteSegmentRHN  string
	BreadCrumbsData     ac.BreadCrumbs
	ViewMarkdownURL     web.AddParams
	EditMarkdownFileURL web.AddParams
	DeleteFileURL       web.AddParams
	ac.CourseManagerLayout
}

func (p FilesPage) HTMXResponse() templ.Component {
	return p.Component()
}

func (p FilesPage) NonHTMXResponse() templ.Component {
	return p.CourseManagerLayout.WithPage(p.Component())

}

type FilesPageItem struct {
	Name       string
	URL        string
	Path       string
	IsMarkdown bool
	IsDir      bool
}

func (data FilesPage) BreadCrumbs() ac.BreadCrumbs {
	return data.BreadCrumbsData
}

func AddParams(params routes.NodePath, additionalParams ...any) []any {
	pathSlice := params.ToSlice()
	pathSlice = append(pathSlice, additionalParams...)
	return pathSlice
}

func (data FilesPage) PageLayout() cmp.PageLayout {
	return cmp.PageLayout{
		PageTitle: "Files for " + data.Node.GetName(),
		UpNav:     data.UpNav,
		Crumbs:    data.BreadCrumbs().BreadCrumbs(),
	}
}

func (data FilesPage) DeleteFileButton(file FilesPageItem) templ.Component {
	return cmp.Button{
		Text:      "Delete",
		Method:    cmp.HxDelete,
		URL:       data.DeleteFileURL(file.Path),
		Image:     cmp.DeleteImage(),
		HxConfirm: "Are you sure you want to delete this file? " + file.Name,
		HxTarget:  "closest li",
	}.Component()
}

func (data FilesPage) EditMarkdownButton(file FilesPageItem) templ.Component {
	return cmp.Button{
		Text:     "Edit",
		Method:   cmp.HxGet,
		HxTarget: "#markdown",
		PushURL:  true,
		URL:      data.EditMarkdownFileURL(file.Path),
		Image:    cmp.EditImage(),
	}.Component()
}

// view as HTML, rendered
func (data FilesPage) ViewMarkdownButton(file FilesPageItem) templ.Component {
	return cmp.Button{
		Text:     "View As HTML",
		Method:   cmp.HxGet,
		HxTarget: "#markdown",
		URL:      data.ViewMarkdownURL(file.Name),
		PushURL:  true,
	}.Component()
}

// view as HTML
func (data FilesPage) PreviewPage(file FilesPageItem) templ.Component {
	return cmp.Link{
		Element: cmp.Element{
			ID: "preview-button",
		},
		Text:   "Preview Static Site Page",
		Target: "preview",
		URL:    data.PreviewURL(file.Name),
	}.Component()
}

func (data FilesPage) FileListItemElementID(file FilesPageItem) string {
	return strings.ReplaceAll(strings.ToLower(file.Name), " ", "-")
}

func (data FilesPage) Component() templ.Component {
	return NewFilesComponent(data)
}
