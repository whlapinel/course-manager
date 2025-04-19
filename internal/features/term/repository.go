package term

import (
	"gh_static_portfolio/internal/core/term"
)

type Repository interface {
	ByUserID(userID string) ([]term.Term, error)
	ByID(termID int) (term.Term, error)
	Save(newTerm term.Term) (int, error)
	Update(updated term.Term) error
	Delete(termID int) error
}
