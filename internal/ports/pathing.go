package ports

type PathingService interface {
	WithNewRoot(root string) PathingService
	WithSegment(segment string) PathingService
	NodeDirPath(...Node) string
	NodeChildrenDirPath(...Node) string
	NodeSlidesMarkdownPath(...Node) string
	NodeSlidesHTMLPath(...Node) string
	NodeFilesDirPath(...Node) string
	NodeImagePath(...Node) string
}
