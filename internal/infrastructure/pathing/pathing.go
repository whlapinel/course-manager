package pathing

import (
	"fmt"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/node"
	"path/filepath"
	"strings"
)

type nodePathService struct {
	Root string
}

func NewNodePathService(root string) ports.PathingService {
	return &nodePathService{Root: root}
}

func (n *nodePathService) NodeDirPath(nodes ...node.Node) string {
	path := n.Root
	for _, node := range nodes {
		if node == nil {
			break
		}
		path = filepath.Join(path, strings.ToLower(node.TypeName()+"s"))
		var dirName string
		switch id := node.GetID().(type) {
		case string:
			dirName = fmt.Sprintf("%s_%s", strings.ToLower(node.TypeName()), id)
		case int:
			dirName = fmt.Sprintf("%s_%d", strings.ToLower(node.TypeName()), id)
		}
		path = filepath.Join(path, dirName)
	}
	return path
}

func (n *nodePathService) NodeFilesDirPath(nodes ...node.Node) string {
	return filepath.Join(n.NodeDirPath(nodes...), "files")
}

func (n *nodePathService) NodeImagePath(nodes ...node.Node) string {
	return filepath.Join(n.NodeDirPath(nodes...), "image.png")
}
