package handlers

import (
	"bytes"
	"context"
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	fileviews "gh_static_portfolio/internal/app/views/files"
	markdownviews "gh_static_portfolio/internal/app/views/markdown"
	components "gh_static_portfolio/internal/basecomponents"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
	"gh_static_portfolio/internal/shared/web"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type intorstring interface {
	~int | ~string
}

type Service[T ports.Node, ID intorstring, ParentID intorstring] interface {
	Save(T) error
	Update(T) error
	Delete(ID) error
	ByID(ID) (T, error)
	ByParentID(ParentID) ([]T, error)
}

type baseHandler[T ports.Node, ID intorstring, ParentID intorstring] struct {
	service            Service[T, ID, ParentID]
	files              *services.FileService
	markdown           *services.MarkdownService
	nodes              *services.NodeService
	hasStatic          bool
	previewer          ports.SiteGenerator
	reverse            web.Reverse
	getCreateMarkdown  web.HandlerName
	postCreateMarkdown web.HandlerName
	getNode            web.HandlerName
	viewNodeFile       web.HandlerName
	getNodeFile        web.HandlerName
	getNodeFiles       web.HandlerName
	getNodeEditFile    web.HandlerName
	postNodeFile       web.HandlerName
	postNodeEditFile   web.HandlerName
	deleteNodeFile     web.HandlerName
}

func (h *baseHandler[T, ID, ParentID]) getCreateMarkdownFile(c echo.Context) error {
	info, err := parseAndFetchNodes(c, h.nodes)
	if err != nil {
		return err
	}
	editor := markdownviews.MarkdownEditor{
		Name:      "New Markdown File",
		SubmitURL: h.reverse(h.getCreateMarkdown.String(), info.NodePath.ToSlice()...),
		CancelURL: h.reverse(h.getNodeFiles.String(), info.NodePath.ToSlice()...),
	}
	return web.Respond(c, h.reverse(h.getNode.String(), info.NodePath.ToSlice()...), editor.Component(), nil)

}

func (h *baseHandler[T, ID, ParentID]) postCreateMarkdownFile(c echo.Context) error {
	params, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodes.Nodes(params)
	if err != nil {
		return err
	}
	name := c.FormValue("name")
	data := c.FormValue("code-editor")
	err = h.markdown.Create([]byte(data), name, nodes)
	if err != nil {
		return err
	}
	// Respond to the client
	return c.Redirect(303, h.reverse(h.viewNodeFile.String(), params.ToSlice(name)...))
}

func (h *baseHandler[T, ID, ParentID]) postUploadedFile(c echo.Context) error {
	filePath := c.Param("*")
	if filePath == "*" {
		filePath = "."
	}
	filePath, err := url.PathUnescape(filePath)
	if err != nil {
		return err
	}
	log.Println(filePath)
	params, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	nodes, err := h.nodes.Nodes(params)
	if err != nil {
		return err
	}
	// Parse the form to retrieve the file
	err = c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		return err
	}
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	filePath = filepath.Join(filePath, file.Filename)
	if filepath.Ext(filePath) == ".md" {
		err = h.markdown.Save(file, filePath, nodes)
		if err != nil {
			return err
		}
	} else {
		err = h.files.Save(file, filePath, nodes)
		if err != nil {
			return err
		}

	}
	// Respond to the client
	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", file.Filename))
}

func (h *baseHandler[T, ID, ParentID]) showFiles(c echo.Context) error {
	nodePath, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	log.Println("filePath", filePath)
	nodes, err := h.nodes.Nodes(nodePath)
	if err != nil {
		return err
	}
	if filePath == "" || filePath == "*" {
		filePath = "."
	}
	fileInfo, err := h.files.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir {
		c.File(fileInfo.AbsPath)
	}
	parentPath := filepath.Dir(filePath)
	files, err := h.files.NodeFiles(filePath, nodes.ToSlice()...)
	if err != nil {
		return err
	}
	page := fileviews.FilesPage{
		CreateMarkdownURL: h.reverse(h.getCreateMarkdown.String(), nodePath.ToSlice()...),
		UpNav: components.UpNav{
			Text: "Up",
			URL:  h.reverse(h.getNode.String(), nodePath.ToSlice()...),
		},
		Root: filePath == ".",
		ParentDirectory: fileviews.FilesPageItem{
			Name:  parentPath,
			URL:   h.reverse(h.getNodeFile.String(), nodes.ToSlice(filePath)),
			Path:  parentPath,
			IsDir: true,
		},
		CurrentDirectory: fileviews.FilesPageItem{
			Name:  filePath,
			URL:   h.reverse(h.getNodeFile.String(), nodes.ToSlice(filePath)),
			Path:  filePath,
			IsDir: filepath.Ext(filePath) == "",
		},
		PreviewURL: func(params ...any) string {
			if !h.hasStatic {
				return ""
			}
			return h.markdown.PreviewFileURL(nodes.User.(dto.User).LastName, params[0].(string), nodes)

		},
		OpenDirURL:          web.URLFunc(h.getNodeFiles, h.reverse, nodePath.ToSlice()...),
		ViewMarkdownURL:     web.URLFunc(h.viewNodeFile, h.reverse, nodePath.ToSlice()...),
		DeleteFileURL:       web.URLFunc(h.deleteNodeFile, h.reverse, nodePath.ToSlice()...),
		EditMarkdownFileURL: web.URLFunc(h.getNodeEditFile, h.reverse, nodePath.ToSlice()...),
		OpenFileURL:         web.URLFunc(h.getNodeFile, h.reverse, nodePath.ToSlice()...),
		UploadFileURL:       web.URLFunc(h.postNodeFile, h.reverse, nodePath.ToSlice()...),
		Node:                nodes.CurrentNode(),
		Files:               files,
		CourseManagerLayout: BaseLayout2(h.reverse, nodes.User.(dto.User)),
	}
	return Respond(c, page)
}

func (b *baseHandler[T, ID, ParentID]) postEditMarkdown(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	if filePath == "" {
		return fmt.Errorf("path param is empty")
	}
	nodes, err := b.nodes.Nodes(path)
	if err != nil {
		return err
	}
	fileInfo, err := b.files.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if fileInfo.IsDir {
		return fmt.Errorf("%s is a directory", filePath)
	}
	name := c.FormValue("name")
	content := c.FormValue("code-editor")
	log.Println("content", content)
	err = b.markdown.Update([]byte(content), name, filePath, nodes)
	if err != nil {
		return err
	}
	return c.Redirect(
		303,
		web.URLFunc(
			b.viewNodeFile,
			b.reverse,
			path.ToSlice()...,
		)(name),
	)
}

func (h *baseHandler[T, ID, ParentID]) viewMarkdown(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	if filePath == "" {
		return fmt.Errorf("path param is empty")
	}
	nodes, err := h.nodes.Nodes(path)
	if err != nil {
		return err
	}
	html, err := h.markdown.ViewMarkdown(filePath, nodes.ToSlice()...)
	if err != nil {
		return err
	}
	doc := markdownviews.MarkdownDocument{
		Title:   filepath.Base(filePath),
		Content: string(html),
		Static:  false,
	}
	var buf bytes.Buffer
	err = markdownviews.DocLayout(doc).Render(context.Background(), &buf)
	if err != nil {
		return err
	}
	doc.Content = buf.String()
	component := markdownviews.MarkdownIFrame(doc)
	layout := BaseLayout3(h.reverse, nodes.User.(dto.User))
	return web.Respond(c, "", component, layout.WithPage(component))
}

func (h *baseHandler[T, ID, ParentID]) showEditFile(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	if filePath == "" {
		return fmt.Errorf("path param is empty")
	}
	nodes, err := h.nodes.Nodes(path)
	if err != nil {
		return err
	}
	fileInfo, err := h.files.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if fileInfo.IsDir {
		return fmt.Errorf("%s is a directory", filePath)
	}
	content, err := h.files.FileContent(filePath, nodes)
	if err != nil {
		return err
	}
	page := markdownviews.MarkdownEditor{
		Name:                filePath,
		Contents:            string(content),
		SubmitURL:           web.URLFunc(h.postNodeEditFile, h.reverse, path.ToSlice()...)(filePath),
		CancelURL:           h.reverse(h.getNodeFiles.String(), path.ToSlice()...),
		CourseManagerLayout: BaseLayout3(h.reverse, nodes.User.(dto.User)),
	}
	return web.Respond(
		c,
		h.reverse(
			h.getNode.String(),
			path.ToSlice()...,
		),
		page.Component(),
		nil,
	)
}

func (b *baseHandler[T, ID, ParentID]) deleteFile(c echo.Context) error {
	path, err := routes.ParseNodePath(c)
	if err != nil {
		return err
	}
	filePath := c.Param("*")
	if filePath == "" {
		return fmt.Errorf("path param is empty")
	}
	filePath, err = url.PathUnescape(filePath)
	if err != nil {
		return err
	}
	log.Println(filePath)
	nodes, err := b.nodes.Nodes(path)
	if err != nil {
		return err
	}
	fileInfo, err := b.files.FileInfo(filePath, nodes)
	if err != nil {
		return err
	}
	if fileInfo.IsDir {
		return fmt.Errorf("%s is a directory", filePath)
	}
	err = b.files.DeleteFile(filePath, nodes.ToSlice()...)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}
