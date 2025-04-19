package user

import "gh_static_portfolio/internal/core/term"

type TermReader interface {
	ByUserID(userID string) ([]term.Term, error)
}
type Service struct {
	repo      Repository
	termQuery TermReader
}

func NewService(repo Repository, termQuery TermReader) *Service {
	return &Service{
		repo:      repo,
		termQuery: termQuery,
	}
}
