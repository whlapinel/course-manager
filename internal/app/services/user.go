package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/user"
)

type UserService struct {
	userService *user.Service
}

func NewUserService(user *user.Service) *UserService {
	return &UserService{
		userService: user,
	}
}

func (svc *UserService) ByID(userID string) (dto.User, error) {
	user, err := svc.userService.ByID(userID)
	if err != nil {
		return dto.User{}, err
	}
	return dto.User{
		User: user,
	}, nil
}
