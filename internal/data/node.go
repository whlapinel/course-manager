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
		var dirName string
		if id, ok := node.GetID().(string); ok {
			dirName = fmt.Sprintf("%s_%s", strings.ToLower(node.TypeName()), id)
		} else if id, ok := node.GetID().(int); ok {
			dirName = fmt.Sprintf("%s_%d", strings.ToLower(node.TypeName()), id)
		}
		path = filepath.Join(
			path,
			dirName,
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
