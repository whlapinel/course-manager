package managertemplates

import (
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	cmp "gh_static_portfolio/internal/templates/components/base"

	"github.com/a-h/templ"
)

type FilesPage struct {
	Root                bool
	ParentDirectory     FilesPageItem
	CurrentDirectory    FilesPageItem
	Params              routes.NodePath
	CurrentPath         string
	OpenDirURL          web.AddParams
	OpenFileURL         web.AddParams
	UploadFileURL       web.AddParams
	Node                node.Node
	Files               []FilesPageItem
	PopRouteSegmentRHN  string
	BreadCrumbsData     BreadCrumbs
	ViewMarkdownURL     web.AddParams
	EditMarkdownFileURL web.AddParams
	CourseManagerLayout
}

func (p FilesPage) HTMXResponse() templ.Component {
	return p.Component()
}

func (p FilesPage) NonHTMXResponse() templ.Component {
	return p.CourseManagerLayout.Component2(p.Component())

}

type FilesPageItem struct {
	Name       string
	URL        string
	Path       string
	IsMarkdown bool
	IsDir      bool
}

type MarkdownEditor struct {
	Name            string
	Contents        string
	PostEditFileURL string
}

func (data MarkdownEditor) Component() templ.Component {
	return MarkdownEditorComponent(data)
}

func (data FilesPage) BreadCrumbs() BreadCrumbs {
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
		UpNav: cmp.UpNav{
			Text: "Up one level",
		},
		Crumbs: data.BreadCrumbs().BreadCrumbs(),
	}
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

// view as HTML
func (data FilesPage) ViewMarkdownButton(file FilesPageItem) templ.Component {
	return cmp.Button{
		Text:     "View As HTML",
		Method:   cmp.HxGet,
		HxTarget: "#markdown",
		URL:      data.ViewMarkdownURL(file.Name),
		PushURL:  true,
	}.Component()

}
func (data FilesPage) Component() templ.Component {
	return NewFilesComponent(data)
}
