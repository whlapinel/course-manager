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
