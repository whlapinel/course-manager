package hugo

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func staticURLMaker(domain string) func(string, int) string {
	return func(lastName string, courseID int) string {
		lastName = strings.ReplaceAll(lastName, " ", "-")
		lastName = strings.ToLower(lastName)
		return fmt.Sprintf("https://%s-%d.%s", lastName, courseID, domain)
	}
}

type HugoConfig struct {
	BaseURL             string
	Title               string
	Subtitle            string
	Username            string
	UserDataPath        string
	ConfigPath          string
	SiteRoot            string // directory containing the hugo files e.g. data, public, content, etc.
	DestinationDataPath string // where JSON is written in site directory, used by Hugo to build site
}

type NewHugoConfigParams struct {
	dto.User
	SiteRoot            string // directory containing the hugo files e.g. data, public, content, etc.
	CourseID            int
	Domain              string
	Title               string
	Subtitle            string
	Username            string
	SourceDataPath      string // where data resides, used to write json to destination
	DestinationDataPath string // where JSON is written in site directory, used by Hugo to build site
	ConfigPath          string //
}

func NewConfig(user dto.User, params NewHugoConfigParams) HugoConfig {

	return HugoConfig{
		BaseURL: staticURLMaker(params.Domain)(user.LastName, params.CourseID),
		// BaseURL:      params.Domain,
		Title:               params.Title,
		Subtitle:            params.Subtitle,
		Username:            params.Username,
		UserDataPath:        params.SourceDataPath,
		ConfigPath:          params.ConfigPath,
		DestinationDataPath: params.DestinationDataPath,
		SiteRoot:            params.SiteRoot,
	}
}

func (c HugoConfig) Write() error {
	err := os.MkdirAll(filepath.Dir(c.ConfigPath), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(c.ConfigPath)
	if err != nil {
		return err
	}
	log.Println("working directory")
	log.Println(os.Getwd())
	tpl, err := template.ParseFiles("./internal/staticresources/config/config.toml")
	if err != nil {
		return err
	}
	return tpl.Execute(file, c)
}
