package handlers

import (
	"gh_static_portfolio/internal/domain"
	"log"
	"testing"
)

func TestNodesRouteName(t *testing.T) {
	name := ChildNodesRouteName(
		domain.User{ID: "1"},
		domain.Term{ID: 2},
		domain.Course{ID: 3},
	)
	log.Println(name)

}

func TestNewNodeRouteName(t *testing.T) {
	name := NewNodeRouteName(
		domain.User{ID: "1"},
		domain.Term{ID: 2},
		domain.Course{ID: 3},
	)
	log.Println(name)

}
func TestEditNodeRouteName(t *testing.T) {
	name := EditNodeRouteName(
		domain.User{ID: "1"},
		domain.Term{ID: 2},
		domain.Course{ID: 3},
	)
	log.Println(name)

}
func TestNodeFilesRouteName(t *testing.T) {
	name := NodeFilesRouteName(
		domain.User{ID: "1"},
		domain.Term{ID: 2},
		domain.Course{ID: 3},
	)
	log.Println(name)

}
func TestListChildrenRHN(t *testing.T) {
	name := ListChildrenRHN(
		domain.User{ID: "1"},
		domain.Term{ID: 2},
		domain.Course{ID: 3},
	)
	log.Println(name)

}
