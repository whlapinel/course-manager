package hugo

import (
	"path/filepath"
	"strings"
)

type BreadCrumbsPartial struct {
	Items []BreadCrumbsItem `json:"items"`
}

type BreadCrumbsItem struct {
	URL string `json:"url"`
}

func BreadCrumbs(url string) BreadCrumbsPartial {
	var bc BreadCrumbsPartial
	var items []BreadCrumbsItem
	segments := strings.Split(url, "/")
	for i := range segments {
		url := filepath.Join(segments[:i]...)
		items = append(items, BreadCrumbsItem{URL: url})
	}
	bc.Items = items
	return bc
}
