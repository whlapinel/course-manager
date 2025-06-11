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
	service          Service[T, ID, ParentID]
	files            *services.FileService
	markdown         *services.MarkdownService
	nodes            *services.NodeService
	previewer        ports.SiteGenerator
	reverse          web.Reverse
	getNode          web.HandlerName
	viewNodeFile     web.HandlerName
	getNodeFile      web.HandlerName
	getNodeFiles     web.HandlerName
	getNodeEditFile  web.HandlerName
	postNodeFile     web.HandlerName
	postNodeEditFile web.HandlerName
	deleteNodeFile   web.HandlerName
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
	data := c.FormValue("content")
	err = h.markdown.Create([]byte(data), ".", nodes)
	if err != nil {
		return err
	}
	// Respond to the client
	return c.Redirect(200, h.reverse(h.getNodeFiles.String(), params.ToSlice()...))
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
		c.Attachment(fileInfo.AbsPath, filepath.Base(fileInfo.AbsPath))
	}
	parentPath := filepath.Dir(filePath)
	files, err := h.files.NodeFiles(filePath, nodes.ToSlice()...)
	if err != nil {
		return err
	}
	page := fileviews.FilesPage{
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
	content := c.FormValue("code-editor")
	log.Println("content", content)
	err = b.markdown.Update([]byte(content), filePath, nodes)
	if err != nil {
		return err
	}
	return c.Redirect(
		303,
		web.URLFunc(
			b.viewNodeFile,
			b.reverse,
			path.ToSlice()...,
		)(filePath),
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
		Contents:            string(content),
		PostEditFileURL:     web.URLFunc(h.postNodeEditFile, h.reverse, path.ToSlice()...)(filePath),
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
