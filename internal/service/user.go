package service

import "gh_static_portfolio/internal/domain"

type SaveUserParams struct {
	domain.User
}

func (svc CourseService) SaveUser(params SaveUserParams) (domain.User, error) {
	return svc.repo.SaveUser(params.User)

}

func (svc CourseService) GetUser(userID string) (domain.User, error) {
	return svc.repo.GetUser(userID)
}
