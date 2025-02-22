package handlers

import (
	"log"
	"testing"
)

func TestNodesRouteName(t *testing.T) {
	name := ChildNodesRouteName(EmptyNodesCourse...)
	log.Println(name)

}

func TestNewNodeRouteName(t *testing.T) {
	name := NewChildRouteName(EmptyNodesUnit...)
	log.Println(name)

}
func TestEditNodeRouteName(t *testing.T) {
	name := EditNodeRouteName(EmptyNodesCourse...)
	log.Println(name)

}
func TestNodeFilesRouteName(t *testing.T) {
	name := NodeFilesRouteName(EmptyNodesCourse...)
	log.Println(name)

}
func TestListChildrenRHN(t *testing.T) {
	name := ListChildrenRHN(EmptyNodesUser...)
	log.Println(name)

}

func TestEmptyNodes(t *testing.T) {
	log.Println(EmptyNodesTerm)
	log.Println(len(EmptyNodesTerm))

}

func TestShowChildDetails(t *testing.T) {
	courseSet := EmptyNodeSet(EmptyNodesCourse)
	log.Println(ShowChildDetailsRHN(courseSet...))
}

func TestShowParentDetailsRHN(t *testing.T) {
	courseSet := EmptyNodeSet(EmptyNodesCourse)
	log.Println(ShowParentDetailsRHN(courseSet...))

}

func TestListSiblingsURL(t *testing.T) {
	rhn := ListChildrenRHN(EmptyNodeSet(EmptyNodesTerm).ParentNodeSet()...)
	log.Println(rhn)
}

func TestListChildChildrenRHN(t *testing.T) {
	courseSet := EmptyNodeSet(EmptyNodesCourse)
	log.Println(ListChildChildrenRHN(courseSet...))
}
