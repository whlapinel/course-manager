package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/features/user"
)

type UserService struct {
	userService *user.Service
	termService *term.Service
}

func NewUserService(user *user.Service, term *term.Service) *UserService {
	return &UserService{
		userService: user,
		termService: term,
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

func (svc *UserService) ListTerms(userID string) (dto.User, error) {
	user, err := svc.userService.ByID(userID)
	if err != nil {
		return dto.User{}, err
	}
	terms, err := svc.termService.TermsByUser(userID)
	if err != nil {
		return dto.User{}, err
	}
	var termDTOs []dto.Term
	for _, term := range terms {
		termDTO := dto.Term{
			Term: term,
		}
		termDTOs = append(termDTOs, termDTO)
	}
	userDTO := dto.User{
		User:  user,
		Terms: termDTOs,
	}
	return userDTO, nil
}
