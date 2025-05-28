package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/internal/features/course"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
	"gh_static_portfolio/internal/ports"
)

type courseRepo struct {
	queries *database.Queries
}

func NewCourseRepo(queries *database.Queries) course.Repository {
	return &courseRepo{
		queries: queries,
	}

}

func (repo *courseRepo) convertFromDB(dbCourse database.Course) course.Course {
	return course.Course{
		BaseNode: ports.BaseNode[int, int]{
			ID:          int(dbCourse.ID),
			ParentID:    int(dbCourse.TermID),
			Name:        dbCourse.Name,
			Description: dbCourse.Description.String,
		},
	}
}

// ByID implements course.Repository.
func (c *courseRepo) ByID(courseID int) (course.Course, error) {
	dbCourse, err := c.queries.GetCourseByCourseID(context.Background(), int64(courseID))
	if err != nil {
		return course.Course{}, err
	}
	return c.convertFromDB(dbCourse), err
}

// ByTermID implements course.Repository.
func (repo *courseRepo) ByTermID(termID int) ([]course.Course, error) {
	dbCourses, err := repo.queries.GetCoursesByTermID(context.Background(), int64(termID))
	if err != nil {
		return nil, err
	}
	var courses []course.Course
	for _, dbCourse := range dbCourses {
		courses = append(courses, repo.convertFromDB(dbCourse))
	}
	return courses, nil
}

// Delete implements course.Repository.
func (c *courseRepo) Delete(courseID int) error {
	_, err := c.queries.DeleteCourse(context.Background(), int64(courseID))
	if err != nil {
		return err
	}
	return nil
}

// Save implements course.Repository.
func (c *courseRepo) Save(newCourse course.Course) (int, error) {
	courseParams := database.SaveCourseParams{
		TermID: int64(newCourse.ParentID),
		Name:   newCourse.Name,

		Description: sql.NullString{
			Valid:  newCourse.Description != "",
			String: newCourse.Description,
		},
	}
	dbCourse, err := c.queries.SaveCourse(context.Background(), courseParams)
	if err != nil {
		return 0, fmt.Errorf("courseRepo.SaveCourse: %s", err)
	}
	return int(dbCourse.ID), nil
}

// Update implements course.Repository.
func (c *courseRepo) Update(updated course.Course) error {
	err := c.queries.UpdateCourse(context.Background(), database.UpdateCourseParams{
		ID:   int64(updated.ID),
		Name: updated.Name,
		Description: sql.NullString{
			Valid:  updated.Description != "",
			String: updated.Description,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
