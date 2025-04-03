package service

import (
	"gh_static_portfolio/internal/domain"
	"time"
)

func (svc CourseService) NewGenerateSite(userID string) error {
	user, err := svc.GetUser(userID)
	if err != nil {
		return err
	}
	term, err := svc.getDataForGenerate(user)
	if err != nil {
		return err
	}
	err = svc.generate(user, term)
	if err != nil {
		return err
	}
	return nil
}

// for new generator
func (svc CourseService) getDataForGenerate(user domain.User) (domain.Term, error) {
	terms, err := svc.GetTerms(user.ID)
	if err != nil {
		return domain.Term{}, err
	}
	var currTerm domain.Term
	for _, term := range terms {
		if term.Start.Before(time.Now()) && term.End.After(time.Now()) {
			currTerm = term
			break
		}
	}
	// GetTerm to get term occasions
	currTerm, err = svc.GetTerm(currTerm.ID)
	if err != nil {
		return domain.Term{}, err

	}
	courses, err := svc.GetCourses(currTerm.ID)
	if err != nil {
		return domain.Term{}, err
	}
	for i, course := range courses {
		units, err := svc.GetUnits(course.ID)
		if err != nil {
			return domain.Term{}, err
		}
		for j, unit := range units {
			lessons, err := svc.GetLessons(unit.ID)
			if err != nil {
				return domain.Term{}, err
			}
			units[j].Lessons = lessons
		}
		courses[i].Units = units
	}
	currTerm.Courses = courses
	return currTerm, nil
}
