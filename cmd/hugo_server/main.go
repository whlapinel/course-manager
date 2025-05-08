package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("hugo", "--source", "/src", "--destination", "/public")
		err := cmd.Run()
		if err != nil {
			http.Error(w, "Hugo build failed", 500)
			return
		}
		w.Write([]byte("Build complete"))
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
