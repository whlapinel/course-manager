package data

import (
	"fmt"
	"gh_static_portfolio/internal/domain"
	"path/filepath"
	"strings"
)

// nodes is path from root, where term is the root
func NodeDirPath(nodes ...domain.CourseNode) string {
	var path = "./internal/data"
	for _, node := range nodes {
		path = filepath.Join(path, strings.ToLower(node.TypeName()+"s"))
		path = filepath.Join(
			path,
			fmt.Sprintf("%s_%d", strings.ToLower(node.TypeName()), node.GetID()),
		)
	}
	return path
}

func NodeFilesDirPath(nodes ...domain.CourseNode) string {
	return filepath.Join(NodeDirPath(nodes...), "files")
}

func NodeImagePath(nodes ...domain.CourseNode) string {
	return filepath.Join(NodeDirPath(nodes...), "image.png")
}
