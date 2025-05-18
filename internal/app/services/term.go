package services

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/term"
	"gh_static_portfolio/internal/features/termoccasion"
)

type TermService struct {
	termService *term.Service
	occasionSvc *termoccasion.Service
}

func NewTermService(
	termSvc *term.Service,
	occasionSvc *termoccasion.Service,
) *TermService {
	return &TermService{
		termService: termSvc,
		occasionSvc: occasionSvc,
	}
}

func (svc *TermService) ByID(termID int) (dto.Term, error) {
	term, err := svc.termService.ByID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	termDTO := dto.Term{
		Term: term,
	}
	return termDTO, nil
}

func (svc *TermService) Save(term dto.Term) error {
	return svc.termService.Save(term.Term)
}

func (svc *TermService) Update(term dto.Term) error {
	return svc.termService.Update(term.Term)
}

func (svc *TermService) Delete(termID int) error {
	return svc.termService.Delete(termID)
}

func (svc *TermService) ByUserID(userID string) ([]dto.Term, error) {
	var termDTOs []dto.Term
	terms, err := svc.termService.TermsByUser(userID)
	if err != nil {
		return termDTOs, err
	}
	for _, term := range terms {
		termDTO := dto.Term{
			Term: term,
		}
		termDTOs = append(termDTOs, termDTO)
	}
	return termDTOs, nil

}

func (svc *TermService) WithOccasions(termID int) (dto.Term, error) {
	term, err := svc.termService.ByID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	occasions, err := svc.occasionSvc.ByTermID(termID)
	if err != nil {
		return dto.Term{}, err
	}
	return dto.Term{
		Term:      term,
		Occasions: occasions,
	}, nil
}
