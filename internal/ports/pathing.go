package ports

type PathingService interface {
	NodeDirPath(...Node) string
	NodeSlidesMarkdownPath(...Node) string
	NodeSlidesHTMLPath(...Node) string
	NodeFilesDirPath(...Node) string
	NodeImagePath(...Node) string
}
