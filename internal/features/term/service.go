package term

import "gh_static_portfolio/internal/core/term"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (svc *Service) Save(term term.Term) error {
	_, err := svc.repo.Save(term)
	return err
}
func (svc *Service) Update(term term.Term) error {
	return svc.repo.Update(term)
}

func (svc *Service) Delete(termID int) error {
	return svc.repo.Delete(termID)
}

func (svc *Service) ByID(termID int) (term.Term, error) {
	return svc.repo.ByID(termID)
}

func (svc *Service) TermsByUser(userID string) ([]term.Term, error) {
	return svc.repo.ByUserID(userID)
}
