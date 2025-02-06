package service

import (
	sitegenerator "gh_static_portfolio/internal/gen_site"
	"os/exec"
)

func (svc CourseService) GenerateSite() error {
	err := sitegenerator.Generate(svc.repo)
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
