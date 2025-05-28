package pathing

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/ports"
	"path/filepath"
	"strings"
)

type nodePathService struct {
	segments []string
}

func NewNodePathService(root string) ports.PathingService {
	return &nodePathService{segments: []string{root}}
}

func (n *nodePathService) WithNewRoot(root string) ports.PathingService {
	return &nodePathService{segments: []string{root}}
}

func (n *nodePathService) WithSegment(segment string) ports.PathingService {
	return &nodePathService{segments: append(n.segments, segment)}
}

func (n *nodePathService) NodeDirPath(nodes ...ports.Node) string {
	var pathSegments = n.segments
	for _, node := range nodes {
		if node == nil {
			break
		}
		if _, ok := node.(dto.User); !ok {
			pathSegments = append(pathSegments, strings.ToLower(node.GetTypeName()+"s"))
		}
		var dirName string
		switch id := node.GetID().(type) {
		case string:
			dirName = fmt.Sprintf("%s_%s", strings.ToLower(node.GetTypeName()), id)
		case int:
			dirName = fmt.Sprintf("%s_%d", strings.ToLower(node.GetTypeName()), id)
		}
		pathSegments = append(pathSegments, dirName)
	}
	return filepath.Join(pathSegments...)
}

func (n *nodePathService) NodeChildrenDirPath(nodes ...ports.Node) string {
	path := n.NodeDirPath(nodes...)
	lastNode := nodes[len(nodes)-1]
	return filepath.Join(path, strings.ToLower(lastNode.GetChildTypeName())+"s")
}

func (n *nodePathService) NodeSlidesHTMLPath(nodes ...ports.Node) string {
	return filepath.Join(n.NodeDirPath(nodes...), "slides.html")
}

func (n *nodePathService) NodeSlidesMarkdownPath(nodes ...ports.Node) string {
	return filepath.Join(n.NodeDirPath(nodes...), "slides.md")
}

func (n *nodePathService) NodeFilesDirPath(nodes ...ports.Node) string {
	return filepath.Join(n.NodeDirPath(nodes...), "files")
}

func (n *nodePathService) NodeImagePath(nodes ...ports.Node) string {
	return filepath.Join(n.NodeDirPath(nodes...), "image.png")
}
