package appcomponents

import (
	"gh_static_portfolio/internal/domain"
	templates "gh_static_portfolio/internal/templates/shared"
	"log"
	"net/url"

	"github.com/a-h/templ"
)

type Reverse func(name string, params ...any) string

const pageElementID ElementID = "page"

type ElementID string

// simply prefixes with '#'
func (i ElementID) Selector() string {
	return string("#" + i)
}

func (i ElementID) String() string {
	return string(i)
}

const (
	EditSlidesContainerID ElementID = "slides-editor-container"
	EditSlidesTextAreaID  ElementID = "slides-editor-text-area"
)

func AddQueryParam(path, key, value string) string {
	u, err := url.Parse(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	query := u.Query()
	query.Set(key, value)
	u.RawQuery = query.Encode()
	log.Println(u.String())
	return u.String()
}

type EditField struct {
	Params           domain.NodePath
	FieldName        string
	Content          string
	GetEditFieldURL  string
	PostEditFieldURL string
	InputComponent   templ.Component
	IsEdit           bool
}

func FieldContainerID(fieldName string) string {
	return templates.KebabCase(fieldName) + "-container"
}

func FieldInputID(fieldName string) string {
	return templates.KebabCase(fieldName)
}
