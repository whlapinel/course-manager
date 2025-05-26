package unit

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) Save(unit Unit) error {
	_, err := svc.repo.Save(unit)
	if err != nil {
		return err
	}
	return nil
}
func (svc *Service) Update(unit Unit) error {
	return svc.repo.Update(unit)
}

func (svc *Service) ByID(unitID int) (Unit, error) {
	return svc.repo.ByID(unitID)
}

func (svc *Service) ByCourseID(courseID int) ([]Unit, error) {
	return svc.repo.ByCourseID(courseID)
}

func (svc *Service) Delete(unitID int) error {
	return svc.repo.Delete(unitID)
}
