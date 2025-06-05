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
	// remove trailing slash
	url = "/" + url
	segments := strings.SplitAfter(url, "/")
	for i := range segments {
		url := filepath.Join(segments[:i+1]...)
		items = append(items, BreadCrumbsItem{URL: url})
	}
	bc.Items = items
	return bc
}
