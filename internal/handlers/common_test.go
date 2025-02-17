package handlers

import (
	"log"
	"testing"
)

func TestParseParam(t *testing.T) {
	param := RouteParam("/:user-id")
	name := param.Name()
	log.Println(name)
	
}