package user

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) ByID(id string) (User, error) {
	return svc.repo.ByID(id)
}
