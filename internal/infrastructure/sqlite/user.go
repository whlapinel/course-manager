package sqlite

import (
	"context"
	"database/sql"
	"gh_static_portfolio/internal/features/user"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
)

type userRepo struct {
	queries *database.Queries
}

func NewUserRepo(queries *database.Queries) user.Repository {
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

func (repo *userRepo) Save(user user.User) (string, error) {
	dbUser, err := repo.queries.SaveUser(context.Background(), database.SaveUserParams{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Picture: sql.NullString{
			Valid:  user.Picture != "",
			String: user.Picture,
		},
	})
	if err != nil {
		return "", err
	}
	return dbUser.ID, nil
}

func (repo *userRepo) ByID(id string) (user.User, error) {
	dbUser, err := repo.queries.GetUser(context.Background(), id)
	if err != nil {
		return user.User{}, err
	}
	return user.User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Picture:   dbUser.Picture.String,
	}, nil

}
