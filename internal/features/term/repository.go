package term

import (
	"time"
)

type Repository interface {
	ByUserID(userID string) ([]Term, error)
	ByID(termID int) (Term, error)
	Save(newTerm Term) (int, error)
	Update(updated Term) error
	Delete(termID int) error
	RemoveInstructionalDay(date time.Time, termID int) error
}
