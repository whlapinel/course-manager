package term

import (
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
func (svc *Service) Save(term Term) error {
	term = New(term)
	_, err := svc.repo.Save(term)
	return err
}
func (svc *Service) Update(term Term) error {
	return svc.repo.Update(term)
}

func (svc *Service) Delete(termID int) error {
	return svc.repo.Delete(termID)
}

func (svc *Service) ByID(termID int) (Term, error) {
	return svc.repo.ByID(termID)
}

func (svc *Service) TermsByUser(userID string) ([]Term, error) {
	return svc.repo.ByUserID(userID)
}

func (svc *Service) RemoveInstructionalDay(date time.Time, termID int) error {
	return svc.repo.RemoveInstructionalDay(date, termID)
}
