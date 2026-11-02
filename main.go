package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("APP_NAME")
	version := os.Getenv("APP_VERSION")

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "%s %s\n", name, version)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
