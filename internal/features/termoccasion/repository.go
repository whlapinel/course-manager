package termoccasion

import "gh_static_portfolio/internal/core/occasion"

type Repository interface {
	Save(occasion.Occasion) (int, error)
	Update(occasion.Occasion) error
	Delete(id occasion.Occasion) error
	ByID(id int) (occasion.Occasion, error)
	ByTermID(termID int) ([]occasion.Occasion, error)
}
