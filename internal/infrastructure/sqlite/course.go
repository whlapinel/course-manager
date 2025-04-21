package sqlite

import (
	"context"
	"gh_static_portfolio/internal/core/course"
	courseFeature "gh_static_portfolio/internal/features/course"
	database "gh_static_portfolio/internal/infrastructure/sqlite/sqlc"
)

type courseRepo struct {
	queries *database.Queries
}

func NewCourseRepo(queries *database.Queries) courseFeature.Repository {
	return &courseRepo{
		queries: queries,
	}

}

func (repo *courseRepo) convertFromDB(dbCourse database.Course) course.Course {
	return course.Course{
		ID:          int(dbCourse.ID),
		ParentID:    int(dbCourse.TermID),
		Name:        dbCourse.Name,
		Description: dbCourse.Description.String,
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
	panic("unimplemented")
}

// Save implements course.Repository.
func (c *courseRepo) Save(newTerm course.Course) (int, error) {
	panic("unimplemented")
}

// Update implements course.Repository.
func (c *courseRepo) Update(updated course.Course) error {
	panic("unimplemented")
}
