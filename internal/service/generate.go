package service

import (
	"gh_static_portfolio/internal/domain"
	sitegenerator "gh_static_portfolio/internal/gen_site"
	newgensite "gh_static_portfolio/internal/new_gen_site"
	"os/exec"
	"time"
)

func (svc CourseService) NewGenerateSite(userID string) error {
	user, err := svc.GetUser(userID)
	if err != nil {
		return err
	}
	terms, err := svc.GetTerms(userID)
	if err != nil {
		return err
	}
	var currTerm domain.Term
	for _, term := range terms {
		if term.Start.Before(time.Now()) && term.End.After(time.Now()) {
			currTerm = term
			break
		}
	}
	courses, err := svc.GetCourses(currTerm.ID)
	if err != nil {
		return err
	}
	for i, course := range courses {
		units, err := svc.GetUnits(course.ID)
		if err != nil {
			return err
		}
		for j, unit := range units {
			lessons, err := svc.GetLessons(unit.ID)
			if err != nil {
				return err
			}
			units[j].Lessons = lessons
		}
		courses[i].Units = units
	}
	currTerm.Courses = courses
	err = newgensite.Generate(user, currTerm)
	if err != nil {
		return err
	}
	return nil
}

func (svc CourseService) GenerateSite(userID string) error {
	user, err := svc.GetUser(userID)
	if err != nil {
		return err
	}
	err = sitegenerator.Generate(svc.repo, user)
	if err != nil {
		return err
	}
	return nil
}

func (svc CourseService) SyncSite() error {
	err := exec.Command("task", "sync-site").Run()
	if err != nil {
		return err
	}
	return nil
}
