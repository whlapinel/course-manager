package sqlite

import (
	"context"
	"gh_static_portfolio/internal/core/user"
	feature "gh_static_portfolio/internal/features/user"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
)

type userRepo struct {
	queries database.Queries
}

func NewUserRepo(queries database.Queries) feature.Repository {
	return &userRepo{
		queries: queries,
	}

}

func (repo *userRepo) All() ([]user.User, error) {
	dbUsers, err := repo.queries.AllUsers(context.Background())
	if err != nil {
		return nil, err
	}
	var users []user.User
	for _, dbUser := range dbUsers {
		user := user.User{
			ID:        dbUser.ID,
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
			Picture:   dbUser.Picture.String,
		}
		users = append(users, user)
	}
	return users, nil
}
