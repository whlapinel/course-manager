package courseoccasion

import (
	"gh_static_portfolio/internal/core/occasion"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) Delete(id int) error {
	return svc.repo.Delete(id)
}

func (svc *Service) ByCourseID(courseID int) ([]occasion.Occasion, error) {
	return svc.repo.ByParentID(courseID)
}

func (svc *Service) ByID(occasionID int) (occasion.Occasion, error) {
	return svc.repo.ByID(occasionID)
}

func (svc *Service) Create(date time.Time, name string, parentID int) (int, error) {
	occ := occasion.Occasion{
		Name:     name,
		Date:     date,
		ParentID: parentID,
	}
	return svc.repo.Save(occ)
}

func (svc *Service) Update(id int, name string) error {
	var occ occasion.Occasion
	occ.ID = id
	occ.Name = name
	return svc.repo.Update(occ)
}
