package hugo

import (
	"gh_static_portfolio/internal/ports"
	"path/filepath"
	"strings"
)

func (h *hugoGenerator) BreadCrumbs(url string) ports.BreadCrumbsPartial {
	var bc ports.BreadCrumbsPartial
	var items []ports.BreadCrumbsItem
	// remove trailing slash
	url = "/" + url
	segments := strings.SplitAfter(url, "/")
	for i := range segments {
		url := filepath.Join(segments[:i+1]...)
		items = append(items, ports.BreadCrumbsItem{URL: url})
	}
	bc.Items = items
	return bc
}
