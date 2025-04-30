package services

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/shared/node"
	"gh_static_portfolio/internal/shared/routes"
	"log"
	"path/filepath"
	"strings"
)

type GetUserDTO func(id string) (dto.User, error)
type GetTermDTO func(id int) (dto.Term, error)
type GetCourseDTO func(id int) (dto.Course, error)
type GetUnitDTO func(id int) (dto.Unit, error)
type GetLessonDTO func(id int) (dto.Lesson, error)

type NodeService struct {
	user   GetUserDTO
	term   GetTermDTO
	course GetCourseDTO
	unit   GetUnitDTO
	lesson GetLessonDTO
}

func NewNodeService(
	user GetUserDTO,
	term GetTermDTO,
	course GetCourseDTO,
	unit GetUnitDTO,
	lesson GetLessonDTO,
) *NodeService {

	return &NodeService{
		user:   user,
		term:   term,
		course: course,
		unit:   unit,
		lesson: lesson,
	}

}

func (svc *NodeService) Nodes(path routes.NodePath) (node.Nodes, error) {
	var nodes node.Nodes
	if path.UserID != "" {
		if svc == nil {
			return node.Nodes{}, fmt.Errorf("svc is nil")
		}
		if svc.user == nil {
			return node.Nodes{}, fmt.Errorf("svc.user function is nil")
		}
		user, err := svc.user(path.UserID)
		if err != nil {
			return node.Nodes{}, err
		}
		nodes.User = user
		if path.TermID != 0 {
			term, err := svc.term(path.TermID)
			if err != nil {
				return node.Nodes{}, err
			}
			nodes.Term = term
			if path.CourseID != 0 {
				course, err := svc.course(path.CourseID)
				if err != nil {
					return node.Nodes{}, err
				}
				nodes.Course = course
				if path.UnitID != 0 {
					unit, err := svc.unit(path.UnitID)
					if err != nil {
						return node.Nodes{}, err
					}
					nodes.Unit = unit
					if path.LessonID != 0 {
						lesson, err := svc.lesson(path.LessonID)
						if err != nil {
							return node.Nodes{}, err
						}
						nodes.Lesson = lesson
					}
				}
			}
		}
	}
	return nodes, nil
}

// nodes is path from root, where term is the root
func NodeDirPath(nodes ...node.Node) string {
	var path = "./internal/data"
	for _, node := range nodes {
		log.Println("node", node)
		if node == nil {
			break
		}
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

func NodeFilesDirPath(nodes ...node.Node) string {
	return filepath.Join(NodeDirPath(nodes...), "files")
}

func NodeImagePath(nodes ...node.Node) string {
	return filepath.Join(NodeDirPath(nodes...), "image.png")
}
