package main

import (
	"gh_static_portfolio/cmd/data"
	"log"
	"time"
)

func main() {
	courseInstances, err := data.GenerateCourseInstancesFromCSV2(time.Now())
	if err != nil {
		log.Fatal(err)
	}
	err = data.WriteCourseInstancesToCSV(courseInstances)
	if err != nil {
		log.Fatalf("error writing schedule to csv, %v", err)
	}

}
