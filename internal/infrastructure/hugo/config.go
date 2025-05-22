package hugo

import (
	"fmt"
	"gh_static_portfolio/internal/app/dto"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func staticURLMaker(domain string) func(string) string {
	return func(userID string) string {
		return fmt.Sprintf("https://%s.%s", userID, domain)
	}
}


type HugoConfig struct {
	BaseURL      string
	Title        string
	UserDataPath string
	ConfigPath   string
}

type NewHugoConfigParams struct {
	dto.User
	Domain       string
	Title        string
	UserDataPath string
	ConfigPath   string
}

func NewConfig(user dto.User, params NewHugoConfigParams) HugoConfig {
	return HugoConfig{
		BaseURL: staticURLMaker(params.Domain)(user.ID),
		// BaseURL:      params.Domain,
		Title:        params.Title,
		UserDataPath: params.UserDataPath,
		ConfigPath:   params.ConfigPath,
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
