package standard

import "gh_static_portfolio/internal/core/standard"

type Repository interface {
	LessonObjectives(id int) ([]standard.Objective, error)
	StandardByID(id int) (standard.Standard, error)
	ObjectiveByID(id int) (standard.Objective, error)
	SaveStandard(standard standard.Standard) (standard.Standard, error)
	SaveObjective(objective standard.Objective) (standard.Objective, error)
	UpdateStandard(updated standard.Standard) error
	UpdateObjective(updated standard.Standard) error
	Delete(courseID int) error
}
