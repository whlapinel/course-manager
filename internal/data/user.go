package data

import (
	"context"
	"database/sql"
	"gh_static_portfolio/internal/data/database"
	"gh_static_portfolio/internal/domain"
)

func (cr CourseRepo) SaveUser(user domain.User) (domain.User, error) {
	dbUser, err := cr.queries.SaveUser(context.Background(), database.SaveUserParams{
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
		return domain.User{}, err
	}
	return domain.User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Picture:   dbUser.Picture.String,
	}, nil

}

func (cr CourseRepo) GetUser(id string) (domain.User, error) {
	dbUser, err := cr.queries.GetUser(context.Background(), id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Picture:   dbUser.Picture.String,
	}, nil

}
