package ports

import (
	"gh_static_portfolio/internal/shared/node"
)

type PathingService interface {
	NodeDirPath(...node.Node) string
	NodeFilesDirPath(...node.Node) string
	NodeImagePath(...node.Node) string
}
