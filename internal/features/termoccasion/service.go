package termoccasion

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

func (svc *Service) ByParentID(termID int) ([]occasion.Occasion, error) {
	return svc.repo.ByParentID(termID)
}

func (svc *Service) ByID(occasionID int) (occasion.Occasion, error) {
	return svc.repo.ByID(occasionID)
}

func (svc *Service) Create(date time.Time, name string, termID int) (int, error) {
	occ := occasion.Occasion{
		Date:     date,
		Name:     name,
		ParentID: termID,
	}
	return svc.repo.Save(occ)
}

func (svc *Service) Update(id int, name string) error {
	var occ occasion.Occasion
	occ.ID = id
	occ.Name = name
	return svc.repo.Update(occ)
}
