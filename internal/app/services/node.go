package services

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/ports"
	"gh_static_portfolio/internal/shared/routes"
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

func (svc *NodeService) Nodes(path routes.NodePath) (ports.Nodes, error) {
	var nodes ports.Nodes
	if path.UserID != "" {
		if svc == nil {
			return nodes, fmt.Errorf("svc is nil")
		}
		if svc.user == nil {
			return nodes, fmt.Errorf("svc.user function is nil")
		}
		user, err := svc.user(path.UserID)
		if err != nil {
			return nodes, err
		}
		nodes.User = user
		if path.TermID != 0 {
			term, err := svc.term(path.TermID)
			if err != nil {
				return nodes, err
			}
			nodes.Term = term
			if path.CourseID != 0 {
				course, err := svc.course(path.CourseID)
				if err != nil {
					return nodes, err
				}
				nodes.Course = course
				if path.UnitID != 0 {
					unit, err := svc.unit(path.UnitID)
					if err != nil {
						return nodes, err
					}
					nodes.Unit = unit
					if path.LessonID != 0 {
						lesson, err := svc.lesson(path.LessonID)
						if err != nil {
							return nodes, err
						}
						nodes.Lesson = lesson
					}
				}
			}
		}
	}
	return nodes, nil
}
